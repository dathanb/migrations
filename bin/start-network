#!/bin/bash

PROJECT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$PROJECT_DIR"

printf "Removing network..."
docker network rm fakestack-network 2> /dev/null 1>&2
printf " done\n"

printf "Creating network: "
docker network create --driver bridge fakestack-network

