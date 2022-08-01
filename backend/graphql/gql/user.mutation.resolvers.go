package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Faroukhamadi/likude/ent"
	"github.com/Faroukhamadi/likude/jwt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	token, err := jwt.GenerateToken(input.Username)
	if err != nil {
		panic("token could not generated")
	}
	fmt.Println(token)
	return ent.FromContext(ctx).User.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (int, error) {
	return id, ent.FromContext(ctx).User.DeleteOneID(id).Exec(ctx)
}
