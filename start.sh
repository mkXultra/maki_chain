#!/bin/bash
set -uex
# Load shell variables
SCRIPT_DIR=$(cd $(dirname $0); pwd)
. $SCRIPT_DIR/variables.sh

rm -rf $NODE_HOME &> /dev/null

$BINARY init test --home $NODE_HOME --chain-id=$CHAINID_1
echo $VAL_MNEMONIC_1    | $BINARY keys add $VAL1 --home $NODE_HOME --recover --keyring-backend=test

$BINARY add-genesis-account $($BINARY --home $NODE_HOME keys show $VAL1 --keyring-backend test -a) 100000000000stake --home $NODE_HOME
$BINARY gentx $VAL1 7000000000$BINARY_MAIN_TOKEN --home $NODE_HOME --chain-id $CHAINID_1 --keyring-backend test
$BINARY collect-gentxs --home $NODE_HOME
$BINARY start --home=$NODE_HOME