# API description

API is description of ways markets communicate with each other and its users.

Each sign of message user send to market is additionally reinforced by market name as indtent to ensure markets will not try to proceed same operations with users on other markets. That is not mention in API, but implemented on server and client sides. 

Each procedure on specific market need to be reinforced by market name in senders message, to prevent mimic behavior by markets collecting data. That will guarantee, that only owner of the private key is able to process transactions on specific market.


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
    rpc TradeInfo(PublicKey) returns (Trades);
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

message Trades {
    repeated Trade trades = 1;
}

message Trade {
    string id = 1;
    int32 offer = 2;
    int32 recieve = 3;
}

service User {
    rpc Message(MessageRequest) returns (Bool);
    rpc Remmittance(RemmittanceRequest) returns (Bool);
    rpc Trade(Order) returns (Bool);
    rpc CancelOrders(CancelOrdersRequest) returns (Bool);
}

message MessageRequest {
    string userkey = 1;
    string message = 2;
    string sign = 3;
}

message RemmittanceRequest {
    string senderkey = 1;
    string recieverkey = 2;
    int32 amount = 3;
    string sign = 4;
}

message Order {
    string userkey = 1;
    string buymarketkey = 2;
    int32 offer = 3;
    int32 recieve = 4;
    string sign = 5;
}

message CancelOrdersRequest {
    string userkey = 1;
    string sign = 2;
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

# Code gen

Commands to generate code from proto file:

Go:
```bash
protoc api/api.proto --go-grpc_out=. --go_out=.
```
Dart:
```bash
protoc lib/api/api.proto --dart_out=grpc:.
```