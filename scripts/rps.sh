#!/usr/bin/env bash

# THIS IDEALLY SHOULD BE A E2E TEST WITH TESTCONTAINERS
# https://golang.testcontainers.org/

# YOU MUST HAVE A CHAIN RUNNING TO EXECUTE THIS SCRIPT
# make init && rpsd start --minimum-gas-prices 0rps

# Variables
NAME1="alice"
NAME2="bob"
BIN=$(which rpsd)

##################
# Create Student #
##################

echo "Create students for alice and bob"

# Get the addresses
ADDRESS1=$($BIN keys show $NAME1 -a --keyring-backend test)
ADDRESS2=$($BIN keys show $NAME2 -a --keyring-backend test)

# Create a new student for alice
$BIN tx rps create_student alice $ADDRESS1 20 --from $NAME1 -o json -y | jq

# Create a new student for bob
$BIN tx rps create_student bod $ADDRESS2 45 --from $NAME2 -o json -y | jq

sleep 5

####################
# Getting students #
####################

echo "Printing students"

# Get the student for alice
$BIN query rps student $ADDRESS1 -o json | jq

# Get the student for bob
$BIN query rps student $ADDRESS2 -o json | jq

#####################
# Deleting students #
#####################

echo "Deleting the bob student"
$BIN tx rps delete_student $ADDRESS2 --from $NAME2 -o json -y | jq

sleep 5

$BIN query rps student $ADDRESS2 -o json | jq

echo "This should be an error --^"