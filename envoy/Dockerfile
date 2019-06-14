FROM lyft/envoy:latest
RUN apt-get update
RUN apt-get install -y python

RUN mkdir /app
COPY hot-restarter.py /app
COPY envoy.sh /app
WORKDIR /app

# expects envoy.yaml to be mounted from the host
CMD /usr/bin/env python /app/hot-restarter.py /app/envoy.sh
