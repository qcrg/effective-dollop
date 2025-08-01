#!/bin/bash
set -e
cd "$(dirname "$0")"

docker run \
  --rm \
  --name evdp_postgres \
  -e POSTGRES_USER=guest \
  -e POSTGRES_PASSWORD='asdf;lkj' \
  -e POSTGRES_DB=evdp \
  -v ./init:/docker-entrypoint-initdb.d \
  -p 5432:5432 \
  postgres:latest
