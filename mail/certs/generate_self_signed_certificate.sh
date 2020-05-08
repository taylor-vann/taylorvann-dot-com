#!/bin/bash

# generate self-signed certificate from openssl
location=/usr/local/certs/mail/

openssl genrsa -out ${location}https-server.key 2048
openssl ecparam -genkey -name secp384r1 -out ${location}https-server.key
openssl req -new -x509 -sha256 -key ${location}https-server.key -out ${location}https-server.crt -days 365 -config ${location}cert.conf 