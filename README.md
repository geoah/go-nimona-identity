# Nimona Identity POC

This is a just an experimental repo, please move on. :D

## TL;DR

Using CRDTs to create user and peer identities.  
An identity should be able to create a sub-identity will limited scope.  
All identities can expire, be revoked, or have their keys updated.  

### IdentityCreate

```json
{
    "key": {
        "id": "4A4B64DF",
        "fingerprint": "D4C3 7CF9 90D6 F506 9F0B  5B7B A341 09A9 4A4B 64DF",
        "public": "...",
        "algorithm": "rsa",
        "length": 2048,
        "capabilities": "esca",
        "created": "28 May 2015 at 19:42",
        "expires": "28 May 2025 at 19:42"
    }
}
```

## Development

Protobuf3 generation is done using `https://github.com/src-d/proteus`.  
Follow their instructions and run `go generate` when modifying Block structs.