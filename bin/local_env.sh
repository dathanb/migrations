#!/usr/bin/env bash

# put relevant env var exports here
export MIGRATION_DEMO_DB_USERNAME=postgres
export MIGRATION_DEMO_DB_PASSWORD=password
export MIGRATION_DEMO_DB_PORT=5432
export MIGRATION_DEMO_DB_HOST=localhost
export MIGRATION_DEMO_DB_DBNAME=migration-demo
export MIGRATION_DEMO_DB_SSL_MODE=disable

export MIGRATION_DEMO_SERVER_PORT=8080

case $1 in
*)
  "$@"
  ;;
esac

