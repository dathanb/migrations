#!/bin/bash

PROJECT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$PROJECT_DIR"

printf "starting postgres: "

docker run \
  --name fakestack_db \
  --rm \
  -d \
  --cidfile=pid/db.pid \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  postgres

