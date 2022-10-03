#!/bin/sh
## Usage:
##   . ./set-env.sh ; $COMMAND

export IMG_NAME="postgres:14-alpine"
export CONTAINER_NAME="postgres14"
export POSTGRES_USER="root"
export POSTGRES_PWD="secret"
export POSTGRES_DB="simple_bank"