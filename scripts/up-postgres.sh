#! /bin/bash

docker run \
  --rm \
  --name some-postgres \
  -e POSTGRES_PASSWORD=mysecretpassword \
  -p 5432:5432 \
  -d \
  postgres
