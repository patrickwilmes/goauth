# goauth - OAuth 2.0 Authorization Code Flow with PKCE 

Trivial implementation (for now), of the OAuth 2.0 Authorization Code Flow with PKCE
(backend side).

## Build And Run

```bash
go build ./cmd/goauth
./goauth
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
