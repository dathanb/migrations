#!/bin/bash

PROJECT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$PROJECT_DIR"

printf "starting fakestack: "

docker run \
  --name fakestack \
  --rm \
  --cidfile=pid/fakestack.pid \
  -e FAKESTACK_DB_USERNAME=postgres \
  -e FAKESTACK_DB_PASSWORD=password \
  -e FAKESTACK_DB_PORT=5432 \
  -e FAKESTACK_DB_HOST=localhost \
  -e FAKESTACK_DB_DBNAME=fakestack \
  -e FAKESTACK_DB_SSL_MODE=disable \
  -e FAKESTACK_SERVER_PORT=8080 \
  fakestack

