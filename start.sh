#!/bin/bash

mkdir -p pid

for f in pid/*; do
  if [ -f "$f" ]; then
    echo "Fakestack may already be running; please run stop.sh first"
    exit 1
  fi
done

printf "Removing network..."
docker network rm fakestack-network 2> /dev/null 1>&2
printf " done\n"

printf "Creating network: "
docker network create --driver bridge fakestack-network

./start-postgres.sh
./start-envoy.sh
