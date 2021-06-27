package main

import (
	"context"
	pb "sync_tree/api"
	"sync_tree/data"
)

func (s *server) UserSearch(
	ctx context.Context,
	in *pb.UserSearchRequest,
) (*pb.Markets, error) {
	results := data.Search(in.Name)
	return &pb.Markets{Adresses: results}, nil
}
