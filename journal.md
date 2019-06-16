# 2019-06-11

It seems a little cumbersome to manage the docker network via shell scripts. That's kinda what docker-compose is
supposed to do. But docker compose doesn't appear to enable dynamic scaling of instances.

Looks like docker swarm does support that, though: <https://stackoverflow.com/a/38027260/114583>

So I could create two services: fakestack-blue and fakestack-green.

When I want to deploy, I could just `docker service create --name fakestack-green --replicas 2 -p 8080:8080/tcp fakestack:latest`.

That doesn't solve healthchecks, though, or adding the new instances to envoy.

So maybe, for starters, I'll just script the thing. It'll be a great learning experience anyway, and that's really what
this is all about.

But monitoring healthchecks, terminating instances, restarting them, etc. is a lot of work, and kind of distracts from
what else I'm trying to accomplish here.

But I guess we don't really need to monitor healthchecks persistently -- we just need to wait for healthchecks from the
new service to go green. We can block on that easily enough without implementing persistent monitoring of them.

So yeah, let's start with the primitive version where we manage everything through docker directly. And we'll move to
docker swarm or kubernetes later to simplify management. But in the meantime I'll get some experience building the
blue/green deployment myself.

Instead of managing that myself, it looks like I could use Traefik to get a lot of the routing for free. Not entirely
clear how to do blue/green in a docker setting with it, but it looks to handle the docker scaling and round-robin
routing just fine.

# 2019-06-12

Reading up on Envoy, it seems like to dynamically update envoy we need to implement the EDS. There's a [reference
implementation](https://github.com/envoyproxy/java-control-plane), but I don't find an Envoy-supported docker container
for it, so might have to do that ourselves.

But it'd probably be much easier to just modify the envoy configuration in place and reload envoy. It seems that Envoy
[supports hot restart](https://blog.envoyproxy.io/envoy-hot-restart-1d16b14555b5), so that's an option. That just shifts
the dynamic component of our cluster management over to some local thing that has to modify Envoy's YAML file in place
instead of sending requests to a management API. I'll have to do some research to figure out which one's easier. We
don't have sophisticated requirements, so either one is probably tenable -- but it might be easier to build that
management service, package it in Docker, and then use curl to manipulate the endpoint list than to programmatically
change the YAML file in the same way.

Interesting thoughts:
1. Envoy can be configured with weights to route traffic non-uniformly to members of a cluster. We could incorporate
   that as part of the deployment management, and have new instances scale up slowly as long as their health checks
   remain green.
2. For implementing blue/green, I think we'd want a script that does
   - stand up new instances in docker, connected to the local network
   - block until their health checks are all green
   - Update envoy to route traffic to them.
   - Update envoy to route traffic *away* from the old instances.
   - Monitor the old instances' traffic (prolly through an HTTP endpoint, but maybe by listening to prometheus), and
     once their traffic goes to zero, tear down the docker containers.
     - HTTP endpoint is probably better
     - It's probably also more robust, since the connection to prometheus could be down while the service is still up
       and serving traffic (not that we care that much about that kind of robustness in this case, but it's worth
       noting)

Note that for #2, just removing an instance from discovery doesn't necessarily remove the instance from Envoy. It says
that in the case of the host being removed from discovery, but still having a green healthcheck, Envoy will continue to
route traffic to it (because it assumes that discovery is eventually consistent, not strongly consistent). That means
that the service probably needs some lifecycle awareness -- after we remove the instance from discovery, we also notify
the service in some way that it's obsolescent, at which point it continues to handle traffic normally but starts
returning a red healthcheck. Then Envoy will remove the instance from its routing table internally. After that, traffic
to the instance should die off, and we can terminate the container.

## Control plane

So I'm going to go the EDS route. So looking at the java-control-plane and go-control-plane implementations.

Unfortunately, it looks like they're not full-featured, they're libraries that you import into your own xDS
implementations. And while it might not be too hard, I don't feel like writing an EDS server, nor do I really feel like
writing gRPC client for interfacing with the control plane.

So instead I'll investigate YML editing and hot restarting envoy.

So what does our Envoy config look like anyway?

We're going to have a few things going:
1. Postgres (do we want this traffic going through Envoy, or is that overkill?)
2. Blue cluster
3. Green cluster
4. Some client app(s)

The idea is that the client app(s) will be furiously making requests to our servers as we're running migrations, and
Envoy is just making sure that traffic goes to the right hosts.

To start up a cluster, we'll
1. Start all the docker containers in parallel
2. Wait for all the healthchecks to go green
3. Update the envoy config to add the new cluster to the existing Fakestack cluster.
4. hot-restart envoy

To shut down a cluster, we'll
1. Remove the cluster members from the envoy config
2. Hot-restart envoy
3. Wait until the traffic on the cluster goes to zero
4. Shut down the containers

If we modify startup step 3 to replace the existing cluster instead of augmenting it with updated members, then shutdown
steps 1 and 2 are already taken care of, and we just do steps 3 and 4.

That actually sounds pretty easy to do. We can even generate it from sources instead of parsing a YAML file and editing
it in place.

So let's run with that plan for now. Since even that naive approach to generating YAML sounds painful to do in bash,
though, I think we'll mix some Python into this effort. Maybe it makes sense to have a Python script just do everything?
Or a Python service that keeps track of what's active right now and updates the Envoy YML on demand. But that sounds
like overkill -- since we can just regenerate the config file entirely, we can be mostly stateless, so for now a script
sounds like a winner.

There's one wrinkle, I think, which is that Envoy uses `listeners` that depend on its clusters. That works fine once we
have a cluster going, but I don't think we can configure a listener to direct traffic to a cluster that doesn't exist.
So we might have to start up fakestack first, then start Envoy pointing to fakestack, even though it'd be nice to be
able to start envoy first and then bootstrap the fakestack cluster through the same procedure we'll use to start an
updated fakestack cluster. I'll have to look into that -- maybe we'll get lucky and Envoy will support listeners with
empty clusters.

# 2019-06-13

I have some basic but nice scripts for starting and stopping the basic infrastructure (docker network, postgres, envoy).
I think the `tini` init running in Docker should forward `SIGHUP` to envoy, which should trigger a restart. Need to
double-check that, though.

Actually, it looks like we might have to run a Python script to handle the sighup and actually do the hot reloading. I
tested without the Python script and tini seems to translate SIGHUP to SIGTERM, so the whole container just exits.

So I'll download the python hot restart wrapper and use it instead of tini.

No luck yet -- and it appears that the hot restart wrapper actually swallows sigterm and prevents graceful shutdown. So
maybe back to square one.

# 2019-06-14

Turns out I wasn't rebuilding the Docker image. :facepalm: But once I did, and passed the `--restart-epoch` to envoy, I
start getting a failure due to inability to access a shared memory device.

Given that all I really want is round-robin load balancing and graceful restart, I think we'll just try nginx next.

Now that I've put all this time into scripts for starting and stopping services, I think it would actually be totally
possible to do this entirely with docker-compose.  I'd prefer to just use existing tooling rather than have to write all
these scripts that are probably going to be brittle. Plus, I'm backgrounding all the services right now, which makes
viewing the logs harder -- and I'd prefer to be able to watch the logs in realtime. So that's next on the agenda.

@ 2019-06-15

It was trivial to get the client streaming to fakestack in docker, so I'd say everything's working fine. I think the
next thing I want to do, though, is to experiment with docker-compose instead of these ad-hoc scripts.

And I'll need to do some work in the fakestack codebase to make the client a little more versatile.

And I'll want to dockerize the client so it can be started and restarted as part of the stack.

# 2019-06-16

If I migrate to docker compose, it's not entirely clear how to get the fakestack service to wait for postgres to be
ready.  We have a few options. e.g.,

1. Include wait-for-it in the fakestack image
2. Support in-container restart -- e.g., via s6-init
3. Have some auto-restart policy in docker?

Actually, it looks like docker compose supports restart policies: <https://docs.docker.com/compose/compose-file/#restart_policy> 


