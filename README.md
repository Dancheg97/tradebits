# Sync tree server

Server software to run sync tree server model.

Every dependency can go only from bottom level to top (as they are listed in folder), so that `main` package can import evrything, while package `logs` is unabale to import anything.

Here is current listing (arrow represents import ability):


- `logs` (low utils)

  ↑
- `calc` (utils)
- `data` (utils)
- `lock` (utils)
- `net` (utils)

  ↑
- `market` (entities)
- `node` (entities)
- `user` (entities)

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

## Net

That package contains interfaces that are used to communicate to outer space. In current implementation gRPC is used for compact messages and fast serialization.

This package is automatically generated, by `api.proto` file, which is in `net` package

# Tests

To run all tests just run `tests.cmd` command