#!/usr/bin/env bash

# put relevant env var exports here
export FAKESTACK_DB_USERNAME=postgres
export FAKESTACK_DB_PASSWORD=password
export FAKESTACK_DB_PORT=5432
export FAKESTACK_DB_HOST=localhost
export FAKESTACK_DB_DBNAME=fakestack
export FAKESTACK_DB_SSL_MODE=disable

export FAKESTACK_SERVER_PORT=8080

case $1 in
*)
  "$@"
  ;;
esac

