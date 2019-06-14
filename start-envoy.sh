#!/bin/bash

rm -rf tmp
mkdir -p tmp

cp envoy/envoy.yaml tmp/

printf "starting envoy: "

docker run \
  -d \
  --rm \
  --cidfile=pid/envoy.pid \
  --name fakestack-envoy \
  --init \
  -v "$(pwd)/tmp/envoy.yaml:/etc/envoy.yaml" \
  fakestack-envoy

