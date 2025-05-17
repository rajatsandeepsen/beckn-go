#!/bin/bash

docker compose run --rm generator-transaction
./edit.sh transaction

docker compose run --rm generator-meta
./edit.sh meta

docker compose run --rm generator-registry
./edit.sh registry

