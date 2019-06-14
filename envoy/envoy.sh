#!/bin/sh
/usr/local/bin/envoy \
  -c /etc/envoy.yaml \
  --restart-epoch 1 \
  -l debug

