# goauth - OAuth 2.0 Authorization Code Flow with PKCE 

Trivial implementation (for now), of the OAuth 2.0 Authorization Code Flow with PKCE
(backend side).

## Run

```bash
go build ./cmd/goauth
./goauth
```

## Build
For creating a container image the script __build-container-image.sh__ can be used. This script
supports building the image with podman and docker. Podman is the default configuration of this
script. For using docker the flag "D" has to be set.

Build with podman:
```bash
./build-container-image.sh <tag-name>
```
Build with docker:
```bash
./build-container-image.sh -D <tag-name>
```

## Tools

### keygen
**Dependencies:** crypto
 
_Tool for generating the secret key for creating a jwt_

**Running**
```bash
gcc -o myKeygen keygen.c -lcrypto
```

### sqlitecli
**Dependencies:** sqlite3

_Toy tool for listing tables and executing queries from cli_
