#!/bin/bash

for f in pid/*.pid; do
  docker stop `cat $f`
  rm $f
done
