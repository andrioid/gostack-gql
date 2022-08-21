package graph

import (
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

func Generate() {
	// GraphQL // TODO: Move into graphql pkg
	cfg, err := config.LoadConfig("gqlgen.yml")
	if err != nil {
		panic(err)
	}
	err = api.Generate((cfg))

	if err != nil {
		panic(err)
	}
}
