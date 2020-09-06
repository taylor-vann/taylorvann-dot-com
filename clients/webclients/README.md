## Clients

A simple file server.
A typescript development environment for web applications.

### Standards

Web Standards are used for static content like html, styles, and scripts.
No commonjs (except for testing ... for now) and no jsx.

### Tests

Mocha and Chai for unit tests.

It's almost worth building a library for typescript. (maybe later)

Currently, testing is handled in Typescript through a separate tsconfig.
Override TS-Node with an environment variable:
TS_NODE_PROJECT=tsconfig.test.json

It's not the best solution but it's simple.
