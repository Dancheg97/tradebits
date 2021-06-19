# Sync tree server

Server software to run sync tree server model.

Every dependency can go only from bottom level to top (as they are listed in folder), so that `main` package can import evrything, packages on above can not import anything and depend on each other.

Here is current listing (arrow represents import ability):

- `api`
- `calc`
- `data`
- `lock`
- `net`

  ↑
- `market`
- `node`
- `user`

  ↑
- `main`

## Logs

Methods that are used according to level of importance of some message.

- `Info`
- `Warning`
- `Error`
- `Fatal`

## Data

All of the methods has checks and logging. Every attemt to do some unexpected action (put instead of change/ get value that does not exist) or any external database issues are gonna be logged with CRITICAL tag.
In current implementation leveldb is used to store data, function list:
- `Get` - method to get bytes from database
- `Put` - method to put element to database
- `Change` - method to change element in database
- `Check` - method is cheking some value is written by some key
- `Trns` - method used to write new transaction


## Calc

This package implements methods used to calculate hashes, verify and sign messages.

- `Hash` - take a hash from value
- `Sign` - sign a byte message using private key bytes
- `Verify` - verify byte message by sign using public key bytes
- `Rand` - Create random 64 bytes for some adress

## Lock

This package provides functions to lock and unlock byte slices (with length of 64), to prevent double spending. Functions:

- `Lock` - lock byte slice (len 64, according to blake2b hash length)
- `Unlcok` - unlock byte slice

## Api

That package contains interfaces that are used to communicate to outer space. In current implementation gRPC is used for compact messages and fast serialization.

This package is automatically generated, by `api.proto` file, which is in `net` package

protoc api/api.proto --go-grpc_out=api
