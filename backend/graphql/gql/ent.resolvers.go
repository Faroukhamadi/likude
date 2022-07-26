package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Faroukhamadi/likude/ent"
	"github.com/Faroukhamadi/likude/graphql/gql/generated"
)

// Points is the resolver for the points field.
func (r *commentResolver) Points(ctx context.Context, obj *ent.Comment) (float64, error) {
	panic(fmt.Errorf("points not implemented"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CommentOrder) (*ent.CommentConnection, error) {
	panic(fmt.Errorf("comments not implemented"))
}

// Communities is the resolver for the communities field.
func (r *queryResolver) Communities(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CommunityOrder) (*ent.CommunityConnection, error) {
	panic(fmt.Errorf("communities not implemented"))
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PostOrder) (*ent.PostConnection, error) {
	panic(fmt.Errorf("posts not implemented"))
}

// Replies is the resolver for the replies field.
func (r *queryResolver) Replies(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ReplyOrder) (*ent.ReplyConnection, error) {
	panic(fmt.Errorf("replies not implemented"))
}

// Topics is the resolver for the topics field.
func (r *queryResolver) Topics(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.TopicConnection, error) {
	panic(fmt.Errorf("topics not implemented"))
}

// Topicrelateds is the resolver for the topicrelateds field.
func (r *queryResolver) Topicrelateds(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.TopicRelatedConnection, error) {
	panic(fmt.Errorf("topicsRelated not implemented"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last, ent.WithUserOrder(orderBy))
}

// Points is the resolver for the points field.
func (r *replyResolver) Points(ctx context.Context, obj *ent.Reply) (float64, error) {
	panic(fmt.Errorf("points not implemented"))
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Reply returns generated.ReplyResolver implementation.
func (r *Resolver) Reply() generated.ReplyResolver { return &replyResolver{r} }

type commentResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type replyResolver struct{ *Resolver }
