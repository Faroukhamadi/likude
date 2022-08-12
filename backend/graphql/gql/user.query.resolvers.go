package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Faroukhamadi/likude/auth"
	"github.com/Faroukhamadi/likude/ent"
	"github.com/Faroukhamadi/likude/ent/user"
)

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	if auth.ForContext(ctx) != nil {
		return r.client.User.Get(ctx, auth.ForContext(ctx).ID)
	}
	return nil, auth.ErrNotAuthenticated
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, username string) (int, error) {
	user, err := r.client.User.Query().Where(user.Username(username)).First(ctx)
	if err != nil {
		return -1, err
	}
	return user.ID, nil
}
