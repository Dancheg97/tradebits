
# <p  align="center" style="font-family:courier;font-size:530%" size=210px> SETGRAPH - distributed trading system </p> 

<p align="center">
  <img height="330px" src="https://www.pngkey.com/png/full/437-4379380_networking-networking-png-portable-network-graphics.png" alt="logo"/>
</p>

# What is setgraph?

Setgraph - is a software, that allows markets to process operations in a distributed network and communicate with each other. Each node in system is it's market, together they provide system's security and stability.

Users can connect directly to setgraph network, to process input, output and trading operations. Users and markets are allowed to process operations using their cryptographic keys.

# How does it work?

To understand how it works, lets start with covering local terms:
- Node - single instance of working market, unary element in setgraph network
- User - individual, that can process transaction in network
- Transaction - single operation in network, that is changing actual data

Each node in network is a unique market, and has to store all actual data and verify transactions passing throw network. All nodes can accept, check, and distribute transactions across the network.

You can find complete listing of transaciton types below.

# Our tools:


<img align="left" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="82px" height="82px" src="https://juststickers.in/wp-content/uploads/2016/07/go-programming-language.png" />
<img align="left" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="82px" height="82px" src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/17/GraphQL_Logo.svg/2048px-GraphQL_Logo.svg.png"/>
<img align="left" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="82px" height="82px" src="https://camo.githubusercontent.com/e6c89a3654756437bd520290bdbe8062bea43e97d38ef2a95d1873d0edd0e014/68747470733a2f2f63646e2e66726565626965737570706c792e636f6d2f6c6f676f732f6c617267652f32782f677261796c6f672d6c6f676f2d706e672d7472616e73706172656e742e706e67" />
<img align="left" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="82px" height="82px" src="https://camo.githubusercontent.com/2c530b38cb14e74d785ebe8d7bf1a649fb44d3e9f43a8dbc103dc01d1fbfce0e/68747470733a2f2f7777772e646f636b65722e636f6d2f73697465732f64656661756c742f66696c65732f64382f323031392d30372f766572746963616c2d6c6f676f2d6d6f6e6f6368726f6d617469632e706e67" />
<img align="left" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="82px" height="82px" src="https://camo.githubusercontent.com/5d442673be6109d82be8dd19f0a2ed6844044bbb58d3e938e9fce7cd346a7946/68747470733a2f2f69312e77702e636f6d2f7465636878706f7365722e636f6d2f77702d636f6e74656e742f75706c6f6164732f323031392f31322f677270632d69636f6e2e706e673f6669743d363236253243363634" />
<img align="left" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="82px" height="82px" src="https://camo.githubusercontent.com/296247907281a8a54eebff1e3af9c89d6d28b6cc531c83befb810c57181d51d8/68747470733a2f2f75706c6f61642e77696b696d656469612e6f72672f77696b6970656469612f636f6d6d6f6e732f302f30302f4b756265726e657465735f253238636f6e7461696e65725f656e67696e652532392e706e67" />
<img align="left" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="82px" height="82px" src="https://is3-ssl.mzstatic.com/image/thumb/Purple124/v4/17/cd/a2/17cda2a0-b641-c3d0-3d22-141704a40eef/Icon.png/1200x630bb.png" />

<br/><br/><br/><br/><br/><br/><br/>


# Tech Stack

Here is listing of technologies used for backend side of token implementation:
- Go
- Dgraph
- gRPC
- Redis
- GrayLog

Go - is language we are using achieve blazing speed. Due to it's concurrent support, relyability and speed it is the best choice for building complex systems as ORB.

Dgraph is used to store persistant data. Due to native support of graph structure of dgraph database, it is the most suitable solution for this currency ecosystem.

Different implementations of dgraph are supported for data storage. In development process classic dgraph/standalone docker image may be the best solution. In production the most appropriate choice - is a full dgraph cluster.

gRPC is used as a method of communication between different nodes and client side. gRPC is fast in transportation and serialization speed, which is making it best option for web communication.

Redis is used for locking operations, to store information about

# API description

API is description of ways different markets communicate with each other and users.

```protobuf
syntax = "proto3";

option go_package = "./api";

package api;

service Info {
    rpc UserBalance(PublicKey) returns (Balance);
    rpc HasTrades(PublicKey) returns (Bool);
    rpc MarketInfo(Empty) returns (InfoMarket);
    rpc NetInfo(Offset) returns (Links);
    rpc GetMessages(Offset) returns (Messages);
    rpc OrderInfo(PublicKey) returns (Orders);
}

message PublicKey {
    string key = 1;
}

message Balance {
    int32 balance = 1;
}

message Bool {
    bool bool = 1;
}

message Empty { }

message InfoMarket {
    string name = 1;
    string pubkey = 2;
    string descr = 3;
    string img = 4;
    string worktime = 5;
    int32 fee = 6;
}

message Offset {
    int32 offset = 1;
}

message Links {
    repeated string links = 1;
}

message Messages {
    repeated string messages = 1;
}

message Orders {
    repeated Order orders = 1;
}

message Order {
    string id = 1;
    int32 offer = 2;
    int32 recieve = 3;
}

service User {
    rpc Message(MessageRequest) returns (Bool);
    rpc Remmittance(RemmittanceRequest) returns (Bool);
    rpc PlaceOrder(PlaceOrderRequest) returns (Bool);
    rpc CancelOrder(CancelOrderRequest) returns (Bool);
}

message MessageRequest {
    string operationid = 1;
    string userkey = 2;
    string message = 3;
    string sign = 4;
}

message RemmittanceRequest {
    string operationid = 1;
    string senderkey = 2;
    string recieverkey = 3;
    int32 amount = 4;
    string sign = 5;
}

message PlaceOrderRequest {
    string operationid = 1;
    string userkey = 2;
    string marketkey = 3;
    int32 offer = 4;
    int32 recieve = 5;
    string sign = 6;
}

message CancelOrderRequest {
    string operationid = 1;
    string userkey = 2;
    string sign = 6;   
}

service Market {
    rpc DecreaseOrder(DecreaseOrderRequest) returns (Bool);
    rpc CloseOrder(CloseOrderRequest) returns (Bool);
}

message DecreaseOrderRequest {
    string operationid = 1;
    string orderid = 2;
    string marketkey = 3;
    int32 newoffer = 4;
    int32 newrecieve = 5;
    string sign  = 6;
}

message CloseOrderRequest {
    string operationid = 1;
    string orderid = 2;
    string marketkey = 3;
    string sign  = 4;
}
```


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

And about transaction types, as a branches:

```graphql

```

### Send

- From - string
- To - string
- Amount - int
- Type - string

### Trade-open

- User - string
- Market - string
- Offer - int
- Recieve - int
- Buy - bool

### Trade-close

- User - bytes
- Market - bytes
- Recieved - int
- Left - int

### Cancel

- User - bytes
- Market - bytes
- Recieved - int
- Base - bool

### Update user

- User - bytes
- Name - string

### Update market

- User - 

### Deposit

- User - bytes
- Market - bytes

### Withdrawal

- User - bytes
- Info - string
