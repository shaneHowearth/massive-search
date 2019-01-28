#!/bin/bash

# Build/run the images/containers
docker-compose up -d

# Run the integration tests
./integration-tests.sh

# Power the  containers down
docker-compose stop
