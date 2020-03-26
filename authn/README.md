# Authn

our lightweight authentication service designed to scale

## Abstract

A dedicated service combining session and user management designed for small projects with minimal resources.

## Outline

Authn retains a store of users and their hashed passwords. And we utilize jwt-like Session Tokens and CSRF Tokens to verify user sessions.

It's an ancillary service designed to separate the authentication of a webapp's users from the webapp itself.

The goal is to have a reliable and isolated authentication service that can (almost) be copied and pasted for new projects and reduce development time for a ubiquitous service found in every modern webapp.

## Code Patterns

We use a semantic pattern to help scale:

Interfaces:

Code tying the current server to another physical device.
For virtual devices on server's physical device, we use the pattern: \<name\>x
For distant devices, we use the pattern: \<name\>xd
Interfaces should not change based on a service's state of scale. It could be a monolith, a small microservice, or a global POP, and interfaces will not change.

Controllers:

Code representing CRUD operations on a table through an interface.
Controllers should not change based on a service's state of scale.

Metas:

Code that relies on multiple interfaces and controllers. We call them metas because they consist of both interfaces and controllers and are the implementation of a blueprint required by the state of scale.

As such, Meta Interfaces are subject to change based on the state of scale. Metas for POPs will be different than Metas for a monolith because the utility of a POP and the utility of a monolith are fundamentally different.
