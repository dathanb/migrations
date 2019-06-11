#!/bin/bash

mkdir -p pid

for f in pid/*; do
  if [ -f "$f" ]; then
    echo "Fakestack may already be running; please run stop.sh first"
    exit 1
  fi
done

docker run \
  --name fakestack_db \
  --rm \
  -d \
  --cidfile=pid/db.pid \
  postgres
