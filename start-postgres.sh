#!/bin/bash

docker run \
  --name fakestack_db \
  --rm \
  -d \
  --cidfile=pid/db.pid \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  postgres

