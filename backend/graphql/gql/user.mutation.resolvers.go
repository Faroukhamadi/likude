package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Faroukhamadi/likude/ent"
	"github.com/Faroukhamadi/likude/graphql/gql/generated"
	"github.com/Faroukhamadi/likude/graphql/gql/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	return r.client.User.Create().
		SetUsername(input.Username).
		SetPassword(input.Password).
		Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input model.UpdateUserInput) (*ent.User, error) {
	if input.Karma == nil && input.Password == nil {
		return r.client.User.Get(ctx, id)
	} else if input.Karma == nil {
		return r.client.User.UpdateOneID(id).
			SetPassword(*input.Password).
			Save(ctx)
	} else {
		return r.client.User.UpdateOneID(id).
			SetKarma(*input.Karma).
			Save(ctx)
	}
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (int, error) {
	return id, r.client.User.DeleteOneID(id).Exec(ctx)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
