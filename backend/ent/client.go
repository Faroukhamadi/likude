// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/Faroukhamadi/likude/ent/migrate"

	"github.com/Faroukhamadi/likude/ent/comment"
	"github.com/Faroukhamadi/likude/ent/community"
	"github.com/Faroukhamadi/likude/ent/post"
	"github.com/Faroukhamadi/likude/ent/reply"
	"github.com/Faroukhamadi/likude/ent/topic"
	"github.com/Faroukhamadi/likude/ent/topicrelated"
	"github.com/Faroukhamadi/likude/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Comment is the client for interacting with the Comment builders.
	Comment *CommentClient
	// Community is the client for interacting with the Community builders.
	Community *CommunityClient
	// Post is the client for interacting with the Post builders.
	Post *PostClient
	// Reply is the client for interacting with the Reply builders.
	Reply *ReplyClient
	// Topic is the client for interacting with the Topic builders.
	Topic *TopicClient
	// TopicRelated is the client for interacting with the TopicRelated builders.
	TopicRelated *TopicRelatedClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Comment = NewCommentClient(c.config)
	c.Community = NewCommunityClient(c.config)
	c.Post = NewPostClient(c.config)
	c.Reply = NewReplyClient(c.config)
	c.Topic = NewTopicClient(c.config)
	c.TopicRelated = NewTopicRelatedClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Comment:      NewCommentClient(cfg),
		Community:    NewCommunityClient(cfg),
		Post:         NewPostClient(cfg),
		Reply:        NewReplyClient(cfg),
		Topic:        NewTopicClient(cfg),
		TopicRelated: NewTopicRelatedClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Comment:      NewCommentClient(cfg),
		Community:    NewCommunityClient(cfg),
		Post:         NewPostClient(cfg),
		Reply:        NewReplyClient(cfg),
		Topic:        NewTopicClient(cfg),
		TopicRelated: NewTopicRelatedClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Comment.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Comment.Use(hooks...)
	c.Community.Use(hooks...)
	c.Post.Use(hooks...)
	c.Reply.Use(hooks...)
	c.Topic.Use(hooks...)
	c.TopicRelated.Use(hooks...)
	c.User.Use(hooks...)
}

// CommentClient is a client for the Comment schema.
type CommentClient struct {
	config
}

// NewCommentClient returns a client for the Comment from the given config.
func NewCommentClient(c config) *CommentClient {
	return &CommentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `comment.Hooks(f(g(h())))`.
func (c *CommentClient) Use(hooks ...Hook) {
	c.hooks.Comment = append(c.hooks.Comment, hooks...)
}

// Create returns a builder for creating a Comment entity.
func (c *CommentClient) Create() *CommentCreate {
	mutation := newCommentMutation(c.config, OpCreate)
	return &CommentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Comment entities.
func (c *CommentClient) CreateBulk(builders ...*CommentCreate) *CommentCreateBulk {
	return &CommentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Comment.
func (c *CommentClient) Update() *CommentUpdate {
	mutation := newCommentMutation(c.config, OpUpdate)
	return &CommentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommentClient) UpdateOne(co *Comment) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withComment(co))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommentClient) UpdateOneID(id int) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withCommentID(id))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Comment.
func (c *CommentClient) Delete() *CommentDelete {
	mutation := newCommentMutation(c.config, OpDelete)
	return &CommentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CommentClient) DeleteOne(co *Comment) *CommentDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CommentClient) DeleteOneID(id int) *CommentDeleteOne {
	builder := c.Delete().Where(comment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommentDeleteOne{builder}
}

// Query returns a query builder for Comment.
func (c *CommentClient) Query() *CommentQuery {
	return &CommentQuery{
		config: c.config,
	}
}

// Get returns a Comment entity by its id.
func (c *CommentClient) Get(ctx context.Context, id int) (*Comment, error) {
	return c.Query().Where(comment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommentClient) GetX(ctx context.Context, id int) *Comment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPost queries the post edge of a Comment.
func (c *CommentClient) QueryPost(co *Comment) *PostQuery {
	query := &PostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comment.PostTable, comment.PostColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReplies queries the replies edge of a Comment.
func (c *CommentClient) QueryReplies(co *Comment) *ReplyQuery {
	query := &ReplyQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(reply.Table, reply.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, comment.RepliesTable, comment.RepliesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CommentClient) Hooks() []Hook {
	return c.hooks.Comment
}

// CommunityClient is a client for the Community schema.
type CommunityClient struct {
	config
}

// NewCommunityClient returns a client for the Community from the given config.
func NewCommunityClient(c config) *CommunityClient {
	return &CommunityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `community.Hooks(f(g(h())))`.
func (c *CommunityClient) Use(hooks ...Hook) {
	c.hooks.Community = append(c.hooks.Community, hooks...)
}

// Create returns a builder for creating a Community entity.
func (c *CommunityClient) Create() *CommunityCreate {
	mutation := newCommunityMutation(c.config, OpCreate)
	return &CommunityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Community entities.
func (c *CommunityClient) CreateBulk(builders ...*CommunityCreate) *CommunityCreateBulk {
	return &CommunityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Community.
func (c *CommunityClient) Update() *CommunityUpdate {
	mutation := newCommunityMutation(c.config, OpUpdate)
	return &CommunityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommunityClient) UpdateOne(co *Community) *CommunityUpdateOne {
	mutation := newCommunityMutation(c.config, OpUpdateOne, withCommunity(co))
	return &CommunityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommunityClient) UpdateOneID(id int) *CommunityUpdateOne {
	mutation := newCommunityMutation(c.config, OpUpdateOne, withCommunityID(id))
	return &CommunityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Community.
func (c *CommunityClient) Delete() *CommunityDelete {
	mutation := newCommunityMutation(c.config, OpDelete)
	return &CommunityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CommunityClient) DeleteOne(co *Community) *CommunityDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CommunityClient) DeleteOneID(id int) *CommunityDeleteOne {
	builder := c.Delete().Where(community.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommunityDeleteOne{builder}
}

// Query returns a query builder for Community.
func (c *CommunityClient) Query() *CommunityQuery {
	return &CommunityQuery{
		config: c.config,
	}
}

// Get returns a Community entity by its id.
func (c *CommunityClient) Get(ctx context.Context, id int) (*Community, error) {
	return c.Query().Where(community.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommunityClient) GetX(ctx context.Context, id int) *Community {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Community.
func (c *CommunityClient) QueryUsers(co *Community) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(community.Table, community.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, community.UsersTable, community.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CommunityClient) Hooks() []Hook {
	return c.hooks.Community
}

// PostClient is a client for the Post schema.
type PostClient struct {
	config
}

// NewPostClient returns a client for the Post from the given config.
func NewPostClient(c config) *PostClient {
	return &PostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `post.Hooks(f(g(h())))`.
func (c *PostClient) Use(hooks ...Hook) {
	c.hooks.Post = append(c.hooks.Post, hooks...)
}

// Create returns a builder for creating a Post entity.
func (c *PostClient) Create() *PostCreate {
	mutation := newPostMutation(c.config, OpCreate)
	return &PostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Post entities.
func (c *PostClient) CreateBulk(builders ...*PostCreate) *PostCreateBulk {
	return &PostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Post.
func (c *PostClient) Update() *PostUpdate {
	mutation := newPostMutation(c.config, OpUpdate)
	return &PostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PostClient) UpdateOne(po *Post) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPost(po))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PostClient) UpdateOneID(id int) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPostID(id))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Post.
func (c *PostClient) Delete() *PostDelete {
	mutation := newPostMutation(c.config, OpDelete)
	return &PostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PostClient) DeleteOne(po *Post) *PostDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PostClient) DeleteOneID(id int) *PostDeleteOne {
	builder := c.Delete().Where(post.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PostDeleteOne{builder}
}

// Query returns a query builder for Post.
func (c *PostClient) Query() *PostQuery {
	return &PostQuery{
		config: c.config,
	}
}

// Get returns a Post entity by its id.
func (c *PostClient) Get(ctx context.Context, id int) (*Post, error) {
	return c.Query().Where(post.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PostClient) GetX(ctx context.Context, id int) *Post {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWriter queries the writer edge of a Post.
func (c *PostClient) QueryWriter(po *Post) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, post.WriterTable, post.WriterColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryComments queries the comments edge of a Post.
func (c *PostClient) QueryComments(po *Post) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, post.CommentsTable, post.CommentsColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PostClient) Hooks() []Hook {
	return c.hooks.Post
}

// ReplyClient is a client for the Reply schema.
type ReplyClient struct {
	config
}

// NewReplyClient returns a client for the Reply from the given config.
func NewReplyClient(c config) *ReplyClient {
	return &ReplyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `reply.Hooks(f(g(h())))`.
func (c *ReplyClient) Use(hooks ...Hook) {
	c.hooks.Reply = append(c.hooks.Reply, hooks...)
}

// Create returns a builder for creating a Reply entity.
func (c *ReplyClient) Create() *ReplyCreate {
	mutation := newReplyMutation(c.config, OpCreate)
	return &ReplyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Reply entities.
func (c *ReplyClient) CreateBulk(builders ...*ReplyCreate) *ReplyCreateBulk {
	return &ReplyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Reply.
func (c *ReplyClient) Update() *ReplyUpdate {
	mutation := newReplyMutation(c.config, OpUpdate)
	return &ReplyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReplyClient) UpdateOne(r *Reply) *ReplyUpdateOne {
	mutation := newReplyMutation(c.config, OpUpdateOne, withReply(r))
	return &ReplyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReplyClient) UpdateOneID(id int) *ReplyUpdateOne {
	mutation := newReplyMutation(c.config, OpUpdateOne, withReplyID(id))
	return &ReplyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Reply.
func (c *ReplyClient) Delete() *ReplyDelete {
	mutation := newReplyMutation(c.config, OpDelete)
	return &ReplyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ReplyClient) DeleteOne(r *Reply) *ReplyDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ReplyClient) DeleteOneID(id int) *ReplyDeleteOne {
	builder := c.Delete().Where(reply.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReplyDeleteOne{builder}
}

// Query returns a query builder for Reply.
func (c *ReplyClient) Query() *ReplyQuery {
	return &ReplyQuery{
		config: c.config,
	}
}

// Get returns a Reply entity by its id.
func (c *ReplyClient) Get(ctx context.Context, id int) (*Reply, error) {
	return c.Query().Where(reply.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReplyClient) GetX(ctx context.Context, id int) *Reply {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryComment queries the comment edge of a Reply.
func (c *ReplyClient) QueryComment(r *Reply) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(reply.Table, reply.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, reply.CommentTable, reply.CommentPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ReplyClient) Hooks() []Hook {
	return c.hooks.Reply
}

// TopicClient is a client for the Topic schema.
type TopicClient struct {
	config
}

// NewTopicClient returns a client for the Topic from the given config.
func NewTopicClient(c config) *TopicClient {
	return &TopicClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `topic.Hooks(f(g(h())))`.
func (c *TopicClient) Use(hooks ...Hook) {
	c.hooks.Topic = append(c.hooks.Topic, hooks...)
}

// Create returns a builder for creating a Topic entity.
func (c *TopicClient) Create() *TopicCreate {
	mutation := newTopicMutation(c.config, OpCreate)
	return &TopicCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Topic entities.
func (c *TopicClient) CreateBulk(builders ...*TopicCreate) *TopicCreateBulk {
	return &TopicCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Topic.
func (c *TopicClient) Update() *TopicUpdate {
	mutation := newTopicMutation(c.config, OpUpdate)
	return &TopicUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TopicClient) UpdateOne(t *Topic) *TopicUpdateOne {
	mutation := newTopicMutation(c.config, OpUpdateOne, withTopic(t))
	return &TopicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TopicClient) UpdateOneID(id int) *TopicUpdateOne {
	mutation := newTopicMutation(c.config, OpUpdateOne, withTopicID(id))
	return &TopicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Topic.
func (c *TopicClient) Delete() *TopicDelete {
	mutation := newTopicMutation(c.config, OpDelete)
	return &TopicDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TopicClient) DeleteOne(t *Topic) *TopicDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TopicClient) DeleteOneID(id int) *TopicDeleteOne {
	builder := c.Delete().Where(topic.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TopicDeleteOne{builder}
}

// Query returns a query builder for Topic.
func (c *TopicClient) Query() *TopicQuery {
	return &TopicQuery{
		config: c.config,
	}
}

// Get returns a Topic entity by its id.
func (c *TopicClient) Get(ctx context.Context, id int) (*Topic, error) {
	return c.Query().Where(topic.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TopicClient) GetX(ctx context.Context, id int) *Topic {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRelatedTopics queries the related_topics edge of a Topic.
func (c *TopicClient) QueryRelatedTopics(t *Topic) *TopicQuery {
	query := &TopicQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(topic.Table, topic.FieldID, id),
			sqlgraph.To(topic.Table, topic.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, topic.RelatedTopicsTable, topic.RelatedTopicsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTopicRelations queries the topic_relations edge of a Topic.
func (c *TopicClient) QueryTopicRelations(t *Topic) *TopicRelatedQuery {
	query := &TopicRelatedQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(topic.Table, topic.FieldID, id),
			sqlgraph.To(topicrelated.Table, topicrelated.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, topic.TopicRelationsTable, topic.TopicRelationsColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TopicClient) Hooks() []Hook {
	return c.hooks.Topic
}

// TopicRelatedClient is a client for the TopicRelated schema.
type TopicRelatedClient struct {
	config
}

// NewTopicRelatedClient returns a client for the TopicRelated from the given config.
func NewTopicRelatedClient(c config) *TopicRelatedClient {
	return &TopicRelatedClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `topicrelated.Hooks(f(g(h())))`.
func (c *TopicRelatedClient) Use(hooks ...Hook) {
	c.hooks.TopicRelated = append(c.hooks.TopicRelated, hooks...)
}

// Create returns a builder for creating a TopicRelated entity.
func (c *TopicRelatedClient) Create() *TopicRelatedCreate {
	mutation := newTopicRelatedMutation(c.config, OpCreate)
	return &TopicRelatedCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TopicRelated entities.
func (c *TopicRelatedClient) CreateBulk(builders ...*TopicRelatedCreate) *TopicRelatedCreateBulk {
	return &TopicRelatedCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TopicRelated.
func (c *TopicRelatedClient) Update() *TopicRelatedUpdate {
	mutation := newTopicRelatedMutation(c.config, OpUpdate)
	return &TopicRelatedUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TopicRelatedClient) UpdateOne(tr *TopicRelated) *TopicRelatedUpdateOne {
	mutation := newTopicRelatedMutation(c.config, OpUpdateOne, withTopicRelated(tr))
	return &TopicRelatedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TopicRelatedClient) UpdateOneID(id int) *TopicRelatedUpdateOne {
	mutation := newTopicRelatedMutation(c.config, OpUpdateOne, withTopicRelatedID(id))
	return &TopicRelatedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TopicRelated.
func (c *TopicRelatedClient) Delete() *TopicRelatedDelete {
	mutation := newTopicRelatedMutation(c.config, OpDelete)
	return &TopicRelatedDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TopicRelatedClient) DeleteOne(tr *TopicRelated) *TopicRelatedDeleteOne {
	return c.DeleteOneID(tr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TopicRelatedClient) DeleteOneID(id int) *TopicRelatedDeleteOne {
	builder := c.Delete().Where(topicrelated.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TopicRelatedDeleteOne{builder}
}

// Query returns a query builder for TopicRelated.
func (c *TopicRelatedClient) Query() *TopicRelatedQuery {
	return &TopicRelatedQuery{
		config: c.config,
	}
}

// Get returns a TopicRelated entity by its id.
func (c *TopicRelatedClient) Get(ctx context.Context, id int) (*TopicRelated, error) {
	return c.Query().Where(topicrelated.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TopicRelatedClient) GetX(ctx context.Context, id int) *TopicRelated {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTopic queries the topic edge of a TopicRelated.
func (c *TopicRelatedClient) QueryTopic(tr *TopicRelated) *TopicQuery {
	query := &TopicQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := tr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(topicrelated.Table, topicrelated.FieldID, id),
			sqlgraph.To(topic.Table, topic.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, topicrelated.TopicTable, topicrelated.TopicColumn),
		)
		fromV = sqlgraph.Neighbors(tr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRelatedTopic queries the related_topic edge of a TopicRelated.
func (c *TopicRelatedClient) QueryRelatedTopic(tr *TopicRelated) *TopicQuery {
	query := &TopicQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := tr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(topicrelated.Table, topicrelated.FieldID, id),
			sqlgraph.To(topic.Table, topic.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, topicrelated.RelatedTopicTable, topicrelated.RelatedTopicColumn),
		)
		fromV = sqlgraph.Neighbors(tr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TopicRelatedClient) Hooks() []Hook {
	return c.hooks.TopicRelated
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPosts queries the posts edge of a User.
func (c *UserClient) QueryPosts(u *User) *PostQuery {
	query := &PostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PostsTable, user.PostsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCommunities queries the communities edge of a User.
func (c *UserClient) QueryCommunities(u *User) *CommunityQuery {
	query := &CommunityQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(community.Table, community.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.CommunitiesTable, user.CommunitiesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
