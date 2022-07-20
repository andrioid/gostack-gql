package main

import (
	"fmt"

	"github.com/eventpuffin/eventpuffin/pkg/gql"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("woot")
	gql.Server()
}
