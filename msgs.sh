#!/bin/bash
set -uex
# Load shell variables
SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/variables.sh

$BINARY  tx maki mint-token 1000ubtc  --from=$VAL1 $conf
 