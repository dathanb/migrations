FROM lyft/envoy:latest
RUN apt-get update

# expects envoy.yaml to be mouted from the host
CMD /usr/local/bin/envoy -c /etc/envoy.yaml
