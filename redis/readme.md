# Lock package
Lock is the main abstraction that is helping to prevent any form of double spending in this system.

Before any transaction is processed in system, this package enures that changed object is not being locked yet.

Each lock operation is putting value to redis via SETNX, by setting value if that one is not existing, whilst each unlock operation removes that value from redis.
