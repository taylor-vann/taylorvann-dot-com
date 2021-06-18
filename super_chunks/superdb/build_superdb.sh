#!/bin/bash

python3 create_supercache_templates.py
podman-compose -f ./docker-compose.yml build