package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/Faroukhamadi/likude/ent"
	"github.com/Faroukhamadi/likude/ent/user"
	"github.com/Faroukhamadi/likude/jwt"
)

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context, tokenString string) (*ent.User, error) {
	username, err := jwt.ParseToken(tokenString)
	if err != nil {
		log.Println("this is error", err)
		return nil, err
	}
	log.Println("this is the username after parsing the token, how cool is this!!", username)
	user, err := r.client.User.Query().Where(user.Username(username)).Only(ctx)
	return user, err
}
