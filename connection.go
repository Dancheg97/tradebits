package main

import (
	pb "sync_tree/api"
)

func (s *server) Connect(
	in *pb.ConnectionRequests_In,
	stream pb.Connection_ConnectServer,
) error {
	// 1 - Check if connecoted is worth it
	// 2 - Start 'key-value queue' in database, to send changes that will occure while the most data is being sent
	// 3 - 
	return nil
}
