#!/bin/bash

PROJECT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$PROJECT_DIR"

mkdir -p pid

for f in pid/*; do
  if [ -f "$f" ]; then
    echo "Fakestack may already be running; please run stop.sh first"
    exit 1
  fi
done

bin/start-network.sh
bin/start-postgres.sh
bin/start-envoy.sh
