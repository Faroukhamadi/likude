package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
)

// Hello is the resolver for the hello field.
func (r *queryResolver) Hello(ctx context.Context) (*string, error) {
	var greeting = "Hello world!"
	log.Println("query is getting executed on the backend, greeting: ", greeting)
	return &greeting, nil
}

// Bye is the resolver for the bye field.
func (r *queryResolver) Bye(ctx context.Context) (*string, error) {
	var greeting = "Bye world!"
	return &greeting, nil
}
