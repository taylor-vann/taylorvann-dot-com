#!/bin/bash

podman-compose -f ./docker-compose.yml down
podman-compose -f ./docker-compose.yml up -d