Brian Taylor Vann
taylorvann dot com

Authentication

Simple JWT authentication in python

We want to issue JWTs to users with correct credentials.
Users can supply an email and password and get a JWT.

We want a white list of valid JWT signatures.
We want to store JWT signatures in memory to quickly verify validity.
We want JWT signatures to delete after a given time.

The process is a create, read, delete.

Create a webtoken, involves a username, password
@login
-> if valid credentials:
     token is created
     signature is stored in memory
     * here you can hook to perform logging or stats on user

Read a webtoken, involves checking existence in memory and expiration date.
@verify
-> if token signature exists:
     return boolean

Invalidate a webtoken, delete a jwt signature from memory
@logout
- if token signature exists:
     remove token from memeory
     return boolean


Create a user
@createUser
- if user credentials valid:
  -> save user & password
  -> create a webtoken
  -> save signature in memory
  -> return jwt


@updateUserCredentials
- if webtoken valid and password confirmation:
  -> update user in database (in this case just a password)
  -> (possibly invalidate all tokens)

@removeUser
- if webtoken valid and password confirmation:
  -> remove user from database
  -> invalidate all tokens



in redis memory, we have the chance to save jwts and a random encryption string.
this could help prevent "stolen" token strings and completely randomizes encryption.
At the worst, 