# Data model

Here is a brief description of data types, that may be stored in dgraph. Graphql (the variation, that is used in dgraph) is used to describe stored data. 


```graphql

type User {
    id: ID!
    pubkey: String! @id @search(by: [hash])
    balance: Int!
    messages: [String]
    orders: [Orders] @hasInverse(field: "user")
    operations: [Operation] @hasInverse(field: "user")
}

type Market {
    id: ID!
    name: String!
    pubkey: String!
    descr: String!
    img: String!
    worktime: String!
    fee: String!
}

type Order {
    id: ID!
    offer: Int!
    recieve: Int!
    ratio: Float!
    user: [User] @hasInverse(field: "orders")
    market: String!
}

type Net {
    id: ID!
    links: [String]
}

type Operation {
    id: ID!
    desciption: String!
    user: [User] @hasInverse(field: "operations")
}

```