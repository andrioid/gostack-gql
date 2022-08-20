package main

import (
	"os"

	sqlc "github.com/kyleconroy/sqlc/pkg/cli"
)

func main() {
	args := []string{"generate"}
	os.Exit(sqlc.Run(args))
}
