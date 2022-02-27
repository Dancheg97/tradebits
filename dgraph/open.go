package dgraph

import (
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

var dgraph = newClient("0.0.0.0:9080")

func newClient(adress string) *dgo.Dgraph {
	d, err := grpc.Dial(adress, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}
