package main

import (
	"context"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/data"
)

func (s *server) InfoFind(
	ctx context.Context,
	in *pb.InfoFindRequest,
) (*pb.InfoFindResponse, error) {
	fmt.Println("user made a search request on: ", in.Info)
	results := data.Search(in.Info)
	return &pb.InfoFindResponse{ConcMarkets: results}, nil
}
