package main

import (
	pb "orb/api"
)

func (s *server) Connect(
	in *pb.ConnectionRequests_In,
	stream pb.Connection_ConnectServer,
) error {
	// 1 - Check if recipient has valid sign and balance
	// 2 - Start 'key-value queue' in database, to fix changes that will occure while the main part of data is streamed to recipient
	// 3 - Start iterating over database values and send all data through gRPC stream
	// 4 - Stop reciving new requests
	// 5 - Wait until there is 0 curretly active requests in stack
	// 6 - Start 'request queue' to resend currently active requests
	// 7 - Put 'request queue' in request queue pool to be fullfilled from all incoming requests
	// 8 - Start resending values from 'database queue'
	// 9 - Start resending values from 'request queue'
	return nil
}
