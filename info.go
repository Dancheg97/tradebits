package main

import (
	"context"
	pb "setgraph/api"
	"setgraph/dgraph"
)

func (s *server) UserBalance(
	ctx context.Context, 
	in *pb.PublicKey,
) (*pb.Balance, error) {
	
}
