package dgraph

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

var schema string
var dgraph *dgo.Dgraph

func Setup(adress string, schema string) {
	d, dialErr := grpc.Dial(adress, grpc.WithInsecure())
	if dialErr != nil {
		log.Panic("unable to open dgraph")
	}
	dgraph := dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
	bytes, err := ioutil.ReadFile(schema)
	if err != nil {
		log.Panic("unable to find schema.gql")
	}
	schemaErr := dgraph.Alter(
		context.Background(),
		&api.Operation{
			Schema: string(bytes),
		},
	)
	if schemaErr != nil {
		log.Panic("unable to set dgraph schema", schemaErr)
	}
}
