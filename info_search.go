package main

import (
	"context"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/data"
)

func (s *server) InfoSearch(
	ctx context.Context,
	in *pb.InfoSearchRequest,
) (*pb.InfoSearchResponse, error) {
	fmt.Println("user made a search request on: ", in.Info)
	results := data.Search(in.Info)
	return &pb.InfoSearchResponse{ConcMarkets: results}, nil
}
