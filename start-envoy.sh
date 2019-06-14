#!/bin/bash

docker run \
  --rm \
  --name fakestack-envoy \
  -v envoy.yaml:/etc/yaml \
  fakestack-envoy

