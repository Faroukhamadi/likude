package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Faroukhamadi/likude/ent"
)

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input ent.CreateCommentInput) (*ent.Comment, error) {
	fmt.Println("inside create comment mutation")
	return ent.FromContext(ctx).Comment.
		Create().
		SetInput(input).
		Save(ctx)
}
