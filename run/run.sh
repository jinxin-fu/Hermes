#!/bin/sh

step=3
# start hermes
./transform -f transform.yaml &

sleep $step

# start transform
./hermes -f hermes-api.yaml

