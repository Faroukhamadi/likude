package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Faroukhamadi/likude/ent"
	"github.com/Faroukhamadi/likude/graphql/gql/generated"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CommentOrder, where *ent.CommentWhereInput) (*ent.CommentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Communities is the resolver for the communities field.
func (r *queryResolver) Communities(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CommunityOrder, where *ent.CommunityWhereInput) (*ent.CommunityConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PostOrder, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	return r.client.Post.Query().
		Paginate(ctx, after, first, before, last, ent.WithPostOrder(orderBy))
}

// Replies is the resolver for the replies field.
func (r *queryResolver) Replies(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ReplyOrder, where *ent.ReplyWhereInput) (*ent.ReplyConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Topics is the resolver for the topics field.
func (r *queryResolver) Topics(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.TopicWhereInput) (*ent.TopicConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Topicrelateds is the resolver for the topicrelateds field.
func (r *queryResolver) Topicrelateds(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.TopicRelatedWhereInput) (*ent.TopicRelatedConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	// if auth.ForContext(ctx) != nil {
	// 	return r.client.User.Query().
	// 		Paginate(ctx, after, first, before, last, ent.WithUserOrder(orderBy))
	// }
	// return nil, auth.ErrNotAuthenticated

	return r.client.User.Query().
		Paginate(ctx, after, first, before, last, ent.WithUserOrder(orderBy))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
