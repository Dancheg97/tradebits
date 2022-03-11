package main

import (
	"context"
	"dbexample"
	"errors"
	"sync"
)

// struct to hold information about users
type user struct {
	id      string
	balance int
}

// struct to hold information about current locks
type locker struct {
	sync.Mutex
	lockmap map[string]struct{}
}

// var to hold state of locked users
var lock = locker{
	lockmap: map[string]struct{}{},
}

func statefullSend(senderID string, recieverID string, amount int) error {
	// lock users if possible
	lock.Mutex.Lock()
	ctx := context.WithTimeout( ... )
	_, err := lock.lockmap[senderID]
	if err {
		return errors.New("sender locked")
	}
	_, err = lock.lockmap[recieverID]
	if err {
		return errors.New("reciever locked")
	}
	lock.lockmap[senderID] = struct{}{}
	lock.lockmap[recieverID] = struct{}{}
	lock.Mutex.Unlock()
	// execute statefull transaction
	sender := user{}
	reciever := user{}
	dbexample.Get(senderID, &sender)
	dbexample.Get(recieverID, &reciever)
	// pls don't use this example \|/ in production :D, made for illustration
	sender.balance = sender.balance - amount
	reciever.balance = reciever.balance + amount
	// unlock statefull elements
	lock.Mutex.Lock()
	delete(lock.lockmap, senderID)
	delete(lock.lockmap, recieverID)
	return nil
}
