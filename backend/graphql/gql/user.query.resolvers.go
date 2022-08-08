package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Faroukhamadi/likude/auth"
	"github.com/Faroukhamadi/likude/ent"
)

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	if auth.ForContext(ctx) != nil {
		return r.client.User.Get(ctx, auth.ForContext(ctx).ID)
	}
	return nil, auth.ErrNotAuthenticated
}
