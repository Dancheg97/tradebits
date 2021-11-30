package main

import (
	"errors"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/calc"
	"sync_tree/data"
	"time"
)

func (s *server) Connect(
	in *pb.ConnectionRequests_In,
	stream pb.Connection_ConnectServer,
) error {
	concMes := [][]byte{
		in.PublicKey,
		[]byte(in.WebAdress),
	}
	signCheckErr := calc.Verify(concMes, in.PublicKey, in.Sign)
	if signCheckErr == nil {
		fmt.Println("[ConnectMessage] - Sign error")
		return errors.New("sign error")
	}
	// verify input data
	iterator, base := data.GetIterator()
	for iterator.Next() {
		kvPair := pb.ConnectionRequests_Out{
			Key:   iterator.Key(),
			Value: iterator.Value(),
		}
		err := stream.Send(&kvPair)
		if err != nil {
			data.CloseConnection()
			fmt.Println("[ConnectMessage] - Connection error")
			return errors.New("connection refused by client")
		}
	}
	data.ReOpenBase(base)
	for {
		time.Sleep(time.Millisecond)
		
	}
	return nil
}
