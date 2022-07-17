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
	panic(fmt.Errorf("not implemented"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
}

// Points is the resolver for the points field.
func (r *replyResolver) Points(ctx context.Context, obj *ent.Reply) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}

// Points is the resolver for the points field.
func (r *commentWhereInputResolver) Points(ctx context.Context, obj *ent.CommentWhereInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsNeq is the resolver for the pointsNEQ field.
func (r *commentWhereInputResolver) PointsNeq(ctx context.Context, obj *ent.CommentWhereInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsIn is the resolver for the pointsIn field.
func (r *commentWhereInputResolver) PointsIn(ctx context.Context, obj *ent.CommentWhereInput, data []float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsNotIn is the resolver for the pointsNotIn field.
func (r *commentWhereInputResolver) PointsNotIn(ctx context.Context, obj *ent.CommentWhereInput, data []float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsGt is the resolver for the pointsGT field.
func (r *commentWhereInputResolver) PointsGt(ctx context.Context, obj *ent.CommentWhereInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsGte is the resolver for the pointsGTE field.
func (r *commentWhereInputResolver) PointsGte(ctx context.Context, obj *ent.CommentWhereInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsLt is the resolver for the pointsLT field.
func (r *commentWhereInputResolver) PointsLt(ctx context.Context, obj *ent.CommentWhereInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsLte is the resolver for the pointsLTE field.
func (r *commentWhereInputResolver) PointsLte(ctx context.Context, obj *ent.CommentWhereInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// Points is the resolver for the points field.
func (r *createCommentInputResolver) Points(ctx context.Context, obj *ent.CreateCommentInput, data float64) error {
	panic(fmt.Errorf("not implemented"))
}

// Points is the resolver for the points field.
func (r *createReplyInputResolver) Points(ctx context.Context, obj *ent.CreateReplyInput, data float64) error {
	panic(fmt.Errorf("not implemented"))
}

// Points is the resolver for the points field.
func (r *replyWhereInputResolver) Points(ctx context.Context, obj *ent.ReplyWhereInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsNeq is the resolver for the pointsNEQ field.
func (r *replyWhereInputResolver) PointsNeq(ctx context.Context, obj *ent.ReplyWhereInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsIn is the resolver for the pointsIn field.
func (r *replyWhereInputResolver) PointsIn(ctx context.Context, obj *ent.ReplyWhereInput, data []float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsNotIn is the resolver for the pointsNotIn field.
func (r *replyWhereInputResolver) PointsNotIn(ctx context.Context, obj *ent.ReplyWhereInput, data []float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsGt is the resolver for the pointsGT field.
func (r *replyWhereInputResolver) PointsGt(ctx context.Context, obj *ent.ReplyWhereInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsGte is the resolver for the pointsGTE field.
func (r *replyWhereInputResolver) PointsGte(ctx context.Context, obj *ent.ReplyWhereInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsLt is the resolver for the pointsLT field.
func (r *replyWhereInputResolver) PointsLt(ctx context.Context, obj *ent.ReplyWhereInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// PointsLte is the resolver for the pointsLTE field.
func (r *replyWhereInputResolver) PointsLte(ctx context.Context, obj *ent.ReplyWhereInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// Points is the resolver for the points field.
func (r *updateCommentInputResolver) Points(ctx context.Context, obj *ent.UpdateCommentInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// Points is the resolver for the points field.
func (r *updateReplyInputResolver) Points(ctx context.Context, obj *ent.UpdateReplyInput, data *float64) error {
	panic(fmt.Errorf("not implemented"))
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Reply returns generated.ReplyResolver implementation.
func (r *Resolver) Reply() generated.ReplyResolver { return &replyResolver{r} }

// CommentWhereInput returns generated.CommentWhereInputResolver implementation.
func (r *Resolver) CommentWhereInput() generated.CommentWhereInputResolver {
	return &commentWhereInputResolver{r}
}

// CreateCommentInput returns generated.CreateCommentInputResolver implementation.
func (r *Resolver) CreateCommentInput() generated.CreateCommentInputResolver {
	return &createCommentInputResolver{r}
}

// CreateReplyInput returns generated.CreateReplyInputResolver implementation.
func (r *Resolver) CreateReplyInput() generated.CreateReplyInputResolver {
	return &createReplyInputResolver{r}
}

// ReplyWhereInput returns generated.ReplyWhereInputResolver implementation.
func (r *Resolver) ReplyWhereInput() generated.ReplyWhereInputResolver {
	return &replyWhereInputResolver{r}
}

// UpdateCommentInput returns generated.UpdateCommentInputResolver implementation.
func (r *Resolver) UpdateCommentInput() generated.UpdateCommentInputResolver {
	return &updateCommentInputResolver{r}
}

// UpdateReplyInput returns generated.UpdateReplyInputResolver implementation.
func (r *Resolver) UpdateReplyInput() generated.UpdateReplyInputResolver {
	return &updateReplyInputResolver{r}
}

type commentResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type replyResolver struct{ *Resolver }
type commentWhereInputResolver struct{ *Resolver }
type createCommentInputResolver struct{ *Resolver }
type createReplyInputResolver struct{ *Resolver }
type replyWhereInputResolver struct{ *Resolver }
type updateCommentInputResolver struct{ *Resolver }
type updateReplyInputResolver struct{ *Resolver }
