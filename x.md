# Turn statefull go application to stateless
In this article i will provide a complete example of clean stateless go server call which can be further extended for various implementations.

To start with, truely stateless clients are harder to write, because we need to communicate with separated statefull agents, like queues(RabbitMQ, Kafka) or databases(PostgreSQL, Redis, Dgraph, etc).

There is no point in attempt to make stateless application, if your application runs good enough on a single machine, but if you run out of resources, than you need to redesign your application to increase scalability.

In this article we will be looking at application for transactions between different users. We will start with statefull application, than can only be runned as a single instance, and than scale it at the expense of moving statefull components outside of the app.

## Step - 1 writing statefull app

First, we need to write simple for understanding statefull application, to get material to work with:
```go
package main

type state_holder_example struct {
	sync.Mutex
	lockmap map[string]struct{}
}

var holder = state_holder_example{}

func state(senderID string, recieverID string, amount int) error {
	state_holder_example.mutex.Lock()
    // say transaction between sender and reciever
    state_holder_example.mutex.Unlock()
	return nil
}
```

This code is used to illustrate, that we can have shared variable with mutex, to isolate operations with same ID's from different sources. This is an approach that is suitable for small applications that do not require much of a scale.

## Step 2 - face the problem

The problem with that code, is that you can not have multiple copies of the same app running on mutiple machines, because if you try to launch 2 nodes with that application, intruders may try to send same operations to different instances to get 2 outputs. So if they catch correct timings, both of spends are gonna be executed simoltaneously on different machines.

<p align="center">
  <img src="https://elf11.github.io/images/doubleSpending-blockchain.png" alt="logo"/>
</p>

How do we solve this problem?

## Step 3 - move statefull components outside of the app

What can we do - is move statefull components, outside of actual application, and say use [RedSync](https://github.com/go-redsync/redsync) or other syncronization patterns.

<p align="center">
<img go align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="ﬁ" height="68px" src="https://juststickers.in/wp-content/uploads/2016/07/go-programming-language.png" /> 
<img go align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="68px" height="68px" src="https://juststickers.in/wp-content/uploads/2016/07/go-programming-language.png" /> 
<img plus align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="68px"  height="68px" src="https://www.picng.com/upload/plus/png_plus_52208.png" />
<img redis align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="68px"  height="68px" src="https://camo.githubusercontent.com/4050472d0036e02ed3805e8329474f062eac6ae847ca0ac107d4889fa778711a/68747470733a2f2f6973332d73736c2e6d7a7374617469632e636f6d2f696d6167652f7468756d622f507572706c653132342f76342f31372f63642f61322f31376364613261302d623634312d633364302d336432322d3134313730346134306565662f49636f6e2e706e672f313230307836333062622e706e67" />
</p>

## Step 4 - think twice, it's gonna break!

Ok, than we need to think, what is gonna happen if we lock access to shared resource, but the executed code is gonna break...
That may cause infinite access lock for some resources, thus very strange unexpected behaviour.

```go
package main

type redlocker struct {
	coolpackage.someCoolDistributedMutex
	lockmap map[string]struct{}
}

func state(senderID string, recieverID string, amount int) error {
	state_holder_example.mutex.Lock()
    // <<===== AND HERE IT BREAKS
    // <<===== And we are getting infinite lock untill the sun burns out...
    state_holder_example.mutex.Unlock()
	return nil
}
```

## Step 5 - Make a timer

We have context library for such purpose. 
So what can we do - is set context with timeout just pass it for required operations:

```go
package main

type redlocker struct {
	coolpackage.someCoolDistributedMutex
	lockmap map[string]struct{}
}

func state(senderID string, recieverID string, amount int) error {
    ctx := context.WithTimeout( ... )
	state_holder_example.mutex.Lock(ctx)
    // passing ctx here aswell, in such a way, that if context is running out, we are backing evrything up to previous condition
	return nil
}
```

# Things to consider...

As you have seen in this simple example, it is way harder to design stateless applications, even when it comes to tiny simple examples. 

There are few importand things, that i recommend to do, when trying to write truely stateless app:

- Don't forget to move all statefull components outside
- Try to use power of ACID databases, for transaciton isolation
- Learn and use [distributed locking patterns](https://martin.kleppmann.com/2016/02/08/how-to-do-distributed-locking.html)
- Think what is gonna happend, if node breaks on each single line of code

