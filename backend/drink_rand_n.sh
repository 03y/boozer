#!/bin/bash

# This script sends n random drink consumption events to the boozer API.
# It requires httpie to be installed (https://httpie.io/docs/cli/installation)
#
# Usage:
# ./drink_rand_n.sh <auth_cookie_value> [n]
#
# If n is not provided, it defaults to 10.

if [ -z "$1" ]; then
  echo "Usage: $0 <auth_cookie> [n]"
  exit 1
fi

AUTH_COOKIE=$1
COUNT=${2:-10}

for i in $(seq 1 $COUNT)
do
  ITEM_ID=$(( (RANDOM % 200) + 1 ))
  echo "Consuming item #$ITEM_ID"
  http --verify=no POST https://localhost/api/v2/consumptions "item_id:=${ITEM_ID}" "Cookie:token=${AUTH_COOKIE}"
  sleep 0.1
done
