## Certificates

Place your certificates in this directory.

- cert.pem
- fullchain.pem
- privkey.pem


Make sure they match up with:

The environemnt variables required by our environment variables.

```
CERTS_CRT_FILEPATH=/usr/local/certs/live/cert.pem
CERTS_KEY_FILEPATH=/usr/local/certs/live/privkey.pem
```

And the Docker-Compose files.

```
...
volumes:
  - ./certs/:/usr/local/certs/live/
  - ./webapi/:/go/src/webapi/
```