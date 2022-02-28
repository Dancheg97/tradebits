
# <p  align="center" style="font-family:courier;font-size:130%" size=210px> TradeBits - interface for trading communications </p> 

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

# Our tools:


<p align="center">
<img align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="72px" height="72px" src="https://juststickers.in/wp-content/uploads/2016/07/go-programming-language.png" />
<img align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="72px" height="72px" src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/17/GraphQL_Logo.svg/2048px-GraphQL_Logo.svg.png"/>
<img align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="72px" height="72px" src="https://camo.githubusercontent.com/e6c89a3654756437bd520290bdbe8062bea43e97d38ef2a95d1873d0edd0e014/68747470733a2f2f63646e2e66726565626965737570706c792e636f6d2f6c6f676f732f6c617267652f32782f677261796c6f672d6c6f676f2d706e672d7472616e73706172656e742e706e67" />
<img align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="72px" height="72px" src="https://camo.githubusercontent.com/2c530b38cb14e74d785ebe8d7bf1a649fb44d3e9f43a8dbc103dc01d1fbfce0e/68747470733a2f2f7777772e646f636b65722e636f6d2f73697465732f64656661756c742f66696c65732f64382f323031392d30372f766572746963616c2d6c6f676f2d6d6f6e6f6368726f6d617469632e706e67" />
<img align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="72px" height="72px" src="https://camo.githubusercontent.com/5d442673be6109d82be8dd19f0a2ed6844044bbb58d3e938e9fce7cd346a7946/68747470733a2f2f69312e77702e636f6d2f7465636878706f7365722e636f6d2f77702d636f6e74656e742f75706c6f6164732f323031392f31322f677270632d69636f6e2e706e673f6669743d363236253243363634" />
<img align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="72px" height="72px" src="https://camo.githubusercontent.com/296247907281a8a54eebff1e3af9c89d6d28b6cc531c83befb810c57181d51d8/68747470733a2f2f75706c6f61642e77696b696d656469612e6f72672f77696b6970656469612f636f6d6d6f6e732f302f30302f4b756265726e657465735f253238636f6e7461696e65725f656e67696e652532392e706e67" />
<img align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="72px" height="72px" src="https://is3-ssl.mzstatic.com/image/thumb/Purple124/v4/17/cd/a2/17cda2a0-b641-c3d0-3d22-141704a40eef/Icon.png/1200x630bb.png" />
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

