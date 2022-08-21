package main

import (
	"github.com/andrioid/gostack-gql/pkg/graph"
	sqlc "github.com/kyleconroy/sqlc/pkg/cli"
)

func main() {
	args := []string{"generate"}
	sqlc.Run(args)
	graph.Generate()
}
