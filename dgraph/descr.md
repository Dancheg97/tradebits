# Data model

Here is a brief description of data types, that may be stored in dgraph. Graphql (the variation, that is used in dgraph) is used to describe stored data. 

There are two main types of data in the system - branch and leaf. Leafs - are those pieces of data that describe current condition of the network, while branches describe the operations, that led to final condition.

You can think about those entities (they might change with time) as a leafs:

```graphql
type User {
    id: ID!
    name: String! @id @search(by: [fulltext])
    pubkey: String! @id @search(by: [hash])
    balance: Int!
    messages: [String]
    buys: [Buy] @hasInverse(field: "user")
    sells: [Sell] @hasInverse(field: "user")
}

type Market {
    id: ID!
    name: String! @id @search(by: [fulltext])
    pubkey: String! @id @search(by: [hash])
    descr: String!
    img: String!
    inputfee: Int!
    outputfee: Int!
    worktime: String!
    buys: [Buy] @hasInverse(field: "market")
    sells: [Sell] @hasInverse(field: "market")
}

type Buy {
    offer: Int!
    recieve: Int!
    user: [User] @hasInverse(field: "buys")
    market: [Market] @hasInverse(field: "buys")
}

type Sell {
    offer: Int!
    recieve: Int!
    user: [User] @hasInverse(field: "sells")
    market: [Market] @hasInverse(field: "sells")
}
```