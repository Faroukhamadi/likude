package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Faroukhamadi/likude/ent"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input ent.CreatePostInput) (*ent.Post, error) {
	fmt.Println("this is the user id: ", input.WriterID)
	return ent.FromContext(ctx).Post.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input ent.UpdatePostInput) (*ent.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id int) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpvotePost is the resolver for the upvotePost field.
func (r *mutationResolver) UpvotePost(ctx context.Context, id int) (*ent.Post, error) {
	return r.client.Post.UpdateOneID(id).AddPoints(1).Save(ctx)
}

// DownvotePost is the resolver for the downvotePost field.
func (r *mutationResolver) DownvotePost(ctx context.Context, id int) (*ent.Post, error) {
	return r.client.Post.UpdateOneID(id).AddPoints(-1).Save(ctx)
}
