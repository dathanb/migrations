#!/bin/bash

mkdir -p pid

for f in pid/*; do
  if [ -f "$f" ]; then
    echo "Fakestack may already be running; please run stop.sh first"
    exit 1
  fi
done

./start-postgres.sh
./start-envoy.sh
