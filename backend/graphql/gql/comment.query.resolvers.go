package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Faroukhamadi/likude/ent"
	"github.com/Faroukhamadi/likude/ent/comment"
	"github.com/Faroukhamadi/likude/ent/post"
)

// PostComments is the resolver for the PostComments field.
func (r *queryResolver) PostComments(ctx context.Context, postID int) ([]*ent.Comment, error) {
	return r.client.Comment.Query().Where(comment.HasPostWith(post.ID(postID))).All(ctx)
}
