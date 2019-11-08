#!/bin/sh

# stop all required services

# apache starts by default sometimes
systemctl stop httpd.service

# nginx
systemctl stop nginx