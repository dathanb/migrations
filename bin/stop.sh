#!/bin/bash

PROJECT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$PROJECT_DIR"

set +e

printf "Removing network..."
docker network rm fakestack-network 2> /dev/null 1>&2
printf " done\n"

for f in `find pid -type f`; do
  printf "Stopping container for $(basename "$f"): "
  docker stop `cat $f`
  rm $f
done
