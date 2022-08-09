package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Faroukhamadi/likude/ent/user"
	"github.com/Faroukhamadi/likude/graphql/gql/generated"
	"github.com/Faroukhamadi/likude/graphql/gql/model"
	"github.com/Faroukhamadi/likude/helpers"
	"github.com/Faroukhamadi/likude/jwt"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	user, err := r.client.User.Query().Where(user.Username(input.Username)).Only(ctx)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}
	correct := helpers.CheckPasswordHash(input.Password, user.Password)
	if correct {
		token, err := jwt.GenerateToken(user.Username)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		return "", fmt.Errorf("incorrect password")
	}
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
