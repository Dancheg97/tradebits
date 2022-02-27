package dgraph

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

func GetClient(adress string) *dgo.Dgraph {
	d, err := grpc.Dial(adress, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	dgraph := dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
	dgraph.Alter(
		context.Background(),
		&api.Operation{
			Schema: `
				name: string @index(term) .
				balance: int .
			`,
		},
	)
	return dgraph
}
