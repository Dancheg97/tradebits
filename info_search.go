package main

import (
	"context"
	pb "sync_tree/api"
	"sync_tree/search"
)

func (s *server) InfoSearch(
	ctx context.Context,
	in *pb.InfoSearchRequest,
) (*pb.InfoSearchResponse, error) {
	//fmt.Println("user made a search request on: ", in.Info)
	results := search.Search(in.Info)
	return &pb.InfoSearchResponse{ConcMarkets: results}, nil
}
