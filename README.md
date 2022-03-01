
# <p  align="center" style="font-family:courier;font-size:130%" size=212px> TradeBits - interface for trading communications </p> 

<p align="center">
  <img height="309px" src="https://www.downloadclipart.net/large/network-png-hd.png" alt="logo"/>
</p>

# What is TradeBits?

TradeBits - is a software, that allows markets to process operations in a distributed network and communicate with each other. Each node in system is it's market, together they provide system's security and stability.

Users can connect directly to TradeBits network markets, to process input, output and trading operations. Users and markets are allowed to process operations using their cryptographic keys.

# How does it work?

To understand how it works, lets start with covering local terms:
- Node - single instance of working market, unary element in TradeBits network
- User - individual, that can process transaction in network
- Transaction - single operation in network, that is changing actual data

Each node in network is a unique market, and has to store all actual data and verify related to that market.

Markets, own their private and public keys, and process trading operation placed by users with different markets using their keys. Markets do not use 3-rd party endpoints and connect directly to each other to process trading operations. Each market holds resposibility for it's users.

Users have to verify markets and process operations on their own risk.

# Our tools


<p align="center">
<img go align="center" style="padding-left: 12px; padding-right: 12px; padding-bottom: 12px;" width="74px" height="74px" src="https://raw.githubusercontent.com/golangci/awesome-go-linters/master/go.png" />
<img mongo align="center" style="padding-left: 12px; padding-right: 12px; padding-bottom: 12px;" width="74px" height="74px" src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/f9/Antu_mongodb.svg/2048px-Antu_mongodb.svg.png"/>
<img graylog align="center" style="padding-left: 12px; padding-right: 12px; padding-bottom: 12px;" width="74px" height="74px" src="https://camo.githubusercontent.com/e6c89a3654756437bd520290bdbe8062bea43e97d38ef2a95d1873d0edd0e014/68747470733a2f2f63646e2e66726565626965737570706c792e636f6d2f6c6f676f732f6c617267652f32782f677261796c6f672d6c6f676f2d706e672d7472616e73706172656e742e706e67" />
<img docker align="center" style="padding-left: 12px; padding-right: 12px; padding-bottom: 12px;" width="74px" height="74px" src="https://cdn-icons-png.flaticon.com/512/919/919853.png" />
<img swagger align="center" style="padding-left: 12px; padding-right: 12px; padding-bottom: 12px;" width="74px" height="74px" src="https://upload.wikimedia.org/wikipedia/commons/a/ab/Swagger-logo.png" />
<img redis align="center" style="padding-left: 12px; padding-right: 12px; padding-bottom: 12px;" width="74px" height="74px" src="https://is3-ssl.mzstatic.com/image/thumb/Purple124/v4/17/cd/a2/17cda2a0-b641-c3d0-3d22-141704a40eef/Icon.png/1200x630bb.png" />
</p>


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

# Links

- [API description](api/descr.md)
- [Data model descrtion](dgraph/descr.md)
- [Redis usage](redis/descr.md)
- [Setup guide]()
- [Mobile app repo]()

