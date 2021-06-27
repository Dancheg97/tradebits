package main

import (
	"context"
	"fmt"
	pb "sync_tree/api"
	"sync_tree/data"
)

func (s *server) UserSearch(
	ctx context.Context,
	in *pb.UserSearchRequest,
) (*pb.Markets, error) {
	fmt.Println("user made a search request on: ", in.Name)
	results := data.Search(in.Name)
	return &pb.Markets{Adresses: results}, nil
}
