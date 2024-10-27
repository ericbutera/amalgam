//go:build ignore

package main

// TODO: move to tools

import (
	"context"
	"fmt"
	"os"

	generate "github.com/Khan/genqlient/generate"
	_ "github.com/alexflint/go-arg"
	gqlfetch "github.com/suessflorian/gqlfetch"
)

// Downloads the GraphQL schema from the locally running server.
// more info: https://github.com/Khan/genqlient/blob/main/docs/schema.md#fetching-your-schema
func main() {
	schema, err := gqlfetch.BuildClientSchema(context.Background(), "http://localhost:8082/query", false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = os.WriteFile("schema.graphql", []byte(schema), 0644); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	generate.Main()
}

//go:generate go run github.com/Khan/genqlient genqlient.yaml
