package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Faroukhamadi/likude/ent"
	"github.com/Faroukhamadi/likude/ent/post"
	"github.com/Faroukhamadi/likude/ent/user"
)

// UserPosts is the resolver for the UserPosts field.
func (r *queryResolver) UserPosts(ctx context.Context, userID int) ([]*ent.Post, error) {
	return r.client.Post.Query().Where(post.HasWriterWith(user.ID(userID))).All(ctx)
}
