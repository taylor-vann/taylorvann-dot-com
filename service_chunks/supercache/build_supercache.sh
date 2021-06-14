#!/bin/bash

python3 create_supercache.py
podman-compose -f ./docker-compose.yml build