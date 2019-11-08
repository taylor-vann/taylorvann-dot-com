#!/bin/sh

# default installation instruction for a centos7 mulitpurpose server

# EPEL release
yum install epel-release -y
yum install dnf -y

# Web Server
dnf install nginx

