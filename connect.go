package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dgraph-io/dgraph/client"
	protos "github.com/dgraph-io/dgraph/protos/api"
	x "github.com/dgraph-io/dgraph/x"
	"google.golang.org/grpc"
)

// Testing connect and query to specified node on cluster
func query() {
	conn1, err := grpc.Dial("127.0.0.1:9081", grpc.WithInsecure())
	x.Checkf(err, "While trying to dial gRPC")
	defer conn1.Close()
	dc1 := protos.NewDgraphClient(conn1)

	conn2, err := grpc.Dial("127.0.0.1:9082", grpc.WithInsecure())
	x.Checkf(err, "While trying to dial gRPC")
	defer conn2.Close()
	dc2 := protos.NewDgraphClient(conn2)

	conn3, err := grpc.Dial("127.0.0.1:9083", grpc.WithInsecure())
	x.Checkf(err, "While trying to dial gRPC")
	defer conn2.Close()
	dc3 := protos.NewDgraphClient(conn3)

	dg := client.NewDgraphClient(dc1, dc2, dc3)

	resp, err := dg.NewTxn().Query(context.Background(), `{
		me(func: eq(iduser,"5510345")) {
		  favorite{
				doing @filter(le(idactivity,"6")){
				idactivity
			  }
		  }
		}
	  }`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response: %s\n", resp.Json)
}
