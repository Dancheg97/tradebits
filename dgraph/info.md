# Description

Dgraph is used to store persistant data. Due to native support of graph structure of dgraph database, it is the most suitable solution for this currency ecosystem.

Different implementations of dgraph are supported for data storage. In development process classic dgraph/standalone docker image may be the best solution. In production the most appropriate choice - is a full dgraph cluster.

# Data model

Here is a brief description of data types, that may be stored in dgraph.

### User

- Adress - bytes
- MesKey - bytes
- Name - string
- Balances - bytes - uint64
- Chats - [links]

### Chat

- Adress1 - bytes
- Adress2 - bytes
- Messages - [strings]
- Count - int

### Market

- adress - bytes
- Name - string
- MesKey - bytes
- Descr - string
- Img - string
- InputFee - int
- OutputFee - int
- WorkTime - string
- Chats - [links]
- Buys - [links]
- Sells - [links]