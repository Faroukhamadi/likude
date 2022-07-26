//go:build ignore
// +build ignore

package main

import (
	"log"

	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithSchemaPath("../graphql/schema/ent.graphql"),
		entgql.WithConfigPath("../graphql/gqlgen.yml"),
		entgql.WithWhereInputs(true),
		entgql.WithRelaySpec(true),
		entgql.WithSchemaGenerator(),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(ex),
	}
	fmt.Println("hello from entc")
	if err := entc.Generate("../ent/schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureEntQL,
			gen.FeaturePrivacy,
			gen.FeatureUpsert,
		},
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
