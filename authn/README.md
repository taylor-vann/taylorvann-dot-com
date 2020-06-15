# Authn

Validate Users and Sessions designed to separate and scale

## Abstract

Authn is the microservice that will issue sessions to users and microservices.

## Sessions

Control access to internal and public apis.

## Store

Emails and Passwords

Permissions for Users 

# Dependencies

Aside from the Golang standard library:

github.com/lib/pq/

github.com/gomodule/redigo/redis/

golang.org/x/crypto/argon2/

golang.org/x/net/publicsuffix/