#!/bin/bash

set +e

for f in `find pid -type f`; do
  docker stop `cat $f`
  rm $f
done
