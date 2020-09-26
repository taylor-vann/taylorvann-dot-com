# Authn

Validate Users and Sessions

## Abstract

Authn is a microservice to create users, roles, and sessions.

## Setup

Provide the following files:

```
./environment-variables.env
./webapi/store_db.init.json
```

You can follow the patterns from the following files:

```
./environment-variables.example.dev
./webapi/store_db.init.json
```

Authn is self contained. You need to provide at least one user to authneticate and CRUD other users, roles, and sessions.

## Sessions

Control access to internal and public apis.

## Store

Users: Emails and Passwords

Roles: Permissions for Users

## Dependencies

Aside from the Golang standard library:

github.com/lib/pq/

github.com/gomodule/redigo/redis/

golang.org/x/crypto/argon2/

golang.org/x/net/publicsuffix/
