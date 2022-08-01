package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/Faroukhamadi/likude/ent"
	"github.com/Faroukhamadi/likude/ent/user"
	"github.com/Faroukhamadi/likude/graphql/gql/generated"
	"github.com/Faroukhamadi/likude/graphql/gql/model"
	"github.com/Faroukhamadi/likude/jwt"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*ent.User, error) {
	user, err := r.client.User.Query().Where(user.Username(input.Username)).Only(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if user.Password == input.Password {
		log.Println("password matches database")
	} else {
		log.Println("why the fuck doesnt password")
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return nil, err
	}
	log.Println("this is the generated token", token)
	log.Println("this is the context value", ctx.Value(userCtxKey.name))
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}
