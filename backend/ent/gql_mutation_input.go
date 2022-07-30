// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateCommentInput represents a mutation input for creating comments.
type CreateCommentInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Content   string
	Points    float64
	PostIDs   []int
	ReplyIDs  []int
}

// Mutate applies the CreateCommentInput on the CommentMutation builder.
func (i *CreateCommentInput) Mutate(m *CommentMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetContent(i.Content)
	m.SetPoints(i.Points)
	if v := i.PostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.ReplyIDs; len(v) > 0 {
		m.AddReplyIDs(v...)
	}
}

// SetInput applies the change-set in the CreateCommentInput on the CommentCreate builder.
func (c *CommentCreate) SetInput(i CreateCommentInput) *CommentCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCommentInput represents a mutation input for updating comments.
type UpdateCommentInput struct {
	UpdatedAt      *time.Time
	Content        *string
	Points         *float64
	AddPostIDs     []int
	RemovePostIDs  []int
	AddReplyIDs    []int
	RemoveReplyIDs []int
}

// Mutate applies the UpdateCommentInput on the CommentMutation builder.
func (i *UpdateCommentInput) Mutate(m *CommentMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Content; v != nil {
		m.SetContent(*v)
	}
	if v := i.Points; v != nil {
		m.SetPoints(*v)
	}
	if v := i.AddPostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.RemovePostIDs; len(v) > 0 {
		m.RemovePostIDs(v...)
	}
	if v := i.AddReplyIDs; len(v) > 0 {
		m.AddReplyIDs(v...)
	}
	if v := i.RemoveReplyIDs; len(v) > 0 {
		m.RemoveReplyIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateCommentInput on the CommentUpdate builder.
func (c *CommentUpdate) SetInput(i UpdateCommentInput) *CommentUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCommentInput on the CommentUpdateOne builder.
func (c *CommentUpdateOne) SetInput(i UpdateCommentInput) *CommentUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateCommunityInput represents a mutation input for creating communities.
type CreateCommunityInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Name      string
	About     string
	UserIDs   []int
}

// Mutate applies the CreateCommunityInput on the CommunityMutation builder.
func (i *CreateCommunityInput) Mutate(m *CommunityMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
	m.SetAbout(i.About)
	if v := i.UserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
}

// SetInput applies the change-set in the CreateCommunityInput on the CommunityCreate builder.
func (c *CommunityCreate) SetInput(i CreateCommunityInput) *CommunityCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCommunityInput represents a mutation input for updating communities.
type UpdateCommunityInput struct {
	UpdatedAt     *time.Time
	Name          *string
	About         *string
	AddUserIDs    []int
	RemoveUserIDs []int
}

// Mutate applies the UpdateCommunityInput on the CommunityMutation builder.
func (i *UpdateCommunityInput) Mutate(m *CommunityMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.About; v != nil {
		m.SetAbout(*v)
	}
	if v := i.AddUserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
	if v := i.RemoveUserIDs; len(v) > 0 {
		m.RemoveUserIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateCommunityInput on the CommunityUpdate builder.
func (c *CommunityUpdate) SetInput(i UpdateCommunityInput) *CommunityUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCommunityInput on the CommunityUpdateOne builder.
func (c *CommunityUpdateOne) SetInput(i UpdateCommunityInput) *CommunityUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreatePostInput represents a mutation input for creating posts.
type CreatePostInput struct {
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	Title      string
	Content    string
	Points     float64
	WriterID  int
	CommentIDs []int
}

// Mutate applies the CreatePostInput on the PostMutation builder.
func (i *CreatePostInput) Mutate(m *PostMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetTitle(i.Title)
	m.SetContent(i.Content)
	m.SetPoints(i.Points)
	m.SetWriterID(i.WriterID)
	if v := i.CommentIDs; len(v) > 0 {
		m.AddCommentIDs(v...)
	}
}

// SetInput applies the change-set in the CreatePostInput on the PostCreate builder.
func (c *PostCreate) SetInput(i CreatePostInput) *PostCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdatePostInput represents a mutation input for updating posts.
type UpdatePostInput struct {
	UpdatedAt        *time.Time
	Title            *string
	Content          *string
	Points           *float64
	AddWriterIDs     []int
	RemoveWriterIDs  []int
	AddCommentIDs    []int
	RemoveCommentIDs []int
}

// Mutate applies the UpdatePostInput on the PostMutation builder.
func (i *UpdatePostInput) Mutate(m *PostMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if v := i.Content; v != nil {
		m.SetContent(*v)
	}
	if v := i.Points; v != nil {
		m.SetPoints(*v)
	}
	if v := i.AddCommentIDs; len(v) > 0 {
		m.AddCommentIDs(v...)
	}
	if v := i.RemoveCommentIDs; len(v) > 0 {
		m.RemoveCommentIDs(v...)
	}
}

// SetInput applies the change-set in the UpdatePostInput on the PostUpdate builder.
func (c *PostUpdate) SetInput(i UpdatePostInput) *PostUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdatePostInput on the PostUpdateOne builder.
func (c *PostUpdateOne) SetInput(i UpdatePostInput) *PostUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateReplyInput represents a mutation input for creating replies.
type CreateReplyInput struct {
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	Content    string
	Points     float64
	CommentIDs []int
}

// Mutate applies the CreateReplyInput on the ReplyMutation builder.
func (i *CreateReplyInput) Mutate(m *ReplyMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetContent(i.Content)
	m.SetPoints(i.Points)
	if v := i.CommentIDs; len(v) > 0 {
		m.AddCommentIDs(v...)
	}
}

// SetInput applies the change-set in the CreateReplyInput on the ReplyCreate builder.
func (c *ReplyCreate) SetInput(i CreateReplyInput) *ReplyCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateReplyInput represents a mutation input for updating replies.
type UpdateReplyInput struct {
	UpdatedAt        *time.Time
	Content          *string
	Points           *float64
	AddCommentIDs    []int
	RemoveCommentIDs []int
}

// Mutate applies the UpdateReplyInput on the ReplyMutation builder.
func (i *UpdateReplyInput) Mutate(m *ReplyMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Content; v != nil {
		m.SetContent(*v)
	}
	if v := i.Points; v != nil {
		m.SetPoints(*v)
	}
	if v := i.AddCommentIDs; len(v) > 0 {
		m.AddCommentIDs(v...)
	}
	if v := i.RemoveCommentIDs; len(v) > 0 {
		m.RemoveCommentIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateReplyInput on the ReplyUpdate builder.
func (c *ReplyUpdate) SetInput(i UpdateReplyInput) *ReplyUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateReplyInput on the ReplyUpdateOne builder.
func (c *ReplyUpdateOne) SetInput(i UpdateReplyInput) *ReplyUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateTopicInput represents a mutation input for creating topics.
type CreateTopicInput struct {
	Name            string
	RelatedTopicIDs []int
}

// Mutate applies the CreateTopicInput on the TopicMutation builder.
func (i *CreateTopicInput) Mutate(m *TopicMutation) {
	m.SetName(i.Name)
	if v := i.RelatedTopicIDs; len(v) > 0 {
		m.AddRelatedTopicIDs(v...)
	}
}

// SetInput applies the change-set in the CreateTopicInput on the TopicCreate builder.
func (c *TopicCreate) SetInput(i CreateTopicInput) *TopicCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTopicInput represents a mutation input for updating topics.
type UpdateTopicInput struct {
	Name                  *string
	AddRelatedTopicIDs    []int
	RemoveRelatedTopicIDs []int
}

// Mutate applies the UpdateTopicInput on the TopicMutation builder.
func (i *UpdateTopicInput) Mutate(m *TopicMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.AddRelatedTopicIDs; len(v) > 0 {
		m.AddRelatedTopicIDs(v...)
	}
	if v := i.RemoveRelatedTopicIDs; len(v) > 0 {
		m.RemoveRelatedTopicIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateTopicInput on the TopicUpdate builder.
func (c *TopicUpdate) SetInput(i UpdateTopicInput) *TopicUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTopicInput on the TopicUpdateOne builder.
func (c *TopicUpdateOne) SetInput(i UpdateTopicInput) *TopicUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateTopicRelatedInput represents a mutation input for creating topicrelateds.
type CreateTopicRelatedInput struct {
	TopicID        int
	RelatedTopicID int
}

// Mutate applies the CreateTopicRelatedInput on the TopicRelatedMutation builder.
func (i *CreateTopicRelatedInput) Mutate(m *TopicRelatedMutation) {
	m.SetTopicID(i.TopicID)
	m.SetRelatedTopicID(i.RelatedTopicID)
}

// SetInput applies the change-set in the CreateTopicRelatedInput on the TopicRelatedCreate builder.
func (c *TopicRelatedCreate) SetInput(i CreateTopicRelatedInput) *TopicRelatedCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTopicRelatedInput represents a mutation input for updating topicrelateds.
type UpdateTopicRelatedInput struct {
	ClearTopic        bool
	TopicID           *int
	ClearRelatedTopic bool
	RelatedTopicID    *int
}

// Mutate applies the UpdateTopicRelatedInput on the TopicRelatedMutation builder.
func (i *UpdateTopicRelatedInput) Mutate(m *TopicRelatedMutation) {
	if i.ClearTopic {
		m.ClearTopic()
	}
	if v := i.TopicID; v != nil {
		m.SetTopicID(*v)
	}
	if i.ClearRelatedTopic {
		m.ClearRelatedTopic()
	}
	if v := i.RelatedTopicID; v != nil {
		m.SetRelatedTopicID(*v)
	}
}

// SetInput applies the change-set in the UpdateTopicRelatedInput on the TopicRelatedUpdate builder.
func (c *TopicRelatedUpdate) SetInput(i UpdateTopicRelatedInput) *TopicRelatedUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTopicRelatedInput on the TopicRelatedUpdateOne builder.
func (c *TopicRelatedUpdateOne) SetInput(i UpdateTopicRelatedInput) *TopicRelatedUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	Username     string
	Password     string
	Karma        *int
	PostIDs      []int
	CommunityIDs []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetUsername(i.Username)
	m.SetPassword(i.Password)
	if v := i.Karma; v != nil {
		m.SetKarma(*v)
	}
	if v := i.PostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.CommunityIDs; len(v) > 0 {
		m.AddCommunityIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	UpdatedAt          *time.Time
	Password           *string
	ClearKarma         bool
	Karma              *int
	AddPostIDs         []int
	RemovePostIDs      []int
	AddCommunityIDs    []int
	RemoveCommunityIDs []int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if i.ClearKarma {
		m.ClearKarma()
	}
	if v := i.Karma; v != nil {
		m.SetKarma(*v)
	}
	if v := i.AddPostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.RemovePostIDs; len(v) > 0 {
		m.RemovePostIDs(v...)
	}
	if v := i.AddCommunityIDs; len(v) > 0 {
		m.AddCommunityIDs(v...)
	}
	if v := i.RemoveCommunityIDs; len(v) > 0 {
		m.RemoveCommunityIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
