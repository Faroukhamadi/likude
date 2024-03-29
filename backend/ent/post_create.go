// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Faroukhamadi/likude/ent/comment"
	"github.com/Faroukhamadi/likude/ent/post"
	"github.com/Faroukhamadi/likude/ent/user"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (pc *PostCreate) SetCreatedAt(t time.Time) *PostCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableCreatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PostCreate) SetUpdatedAt(t time.Time) *PostCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableUpdatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetTitle sets the "title" field.
func (pc *PostCreate) SetTitle(s string) *PostCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetContent sets the "content" field.
func (pc *PostCreate) SetContent(s string) *PostCreate {
	pc.mutation.SetContent(s)
	return pc
}

// SetPoints sets the "points" field.
func (pc *PostCreate) SetPoints(f float64) *PostCreate {
	pc.mutation.SetPoints(f)
	return pc
}

// SetWriterID sets the "writer" edge to the User entity by ID.
func (pc *PostCreate) SetWriterID(id int) *PostCreate {
	pc.mutation.SetWriterID(id)
	return pc
}

// SetNillableWriterID sets the "writer" edge to the User entity by ID if the given value is not nil.
func (pc *PostCreate) SetNillableWriterID(id *int) *PostCreate {
	if id != nil {
		pc = pc.SetWriterID(*id)
	}
	return pc
}

// SetWriter sets the "writer" edge to the User entity.
func (pc *PostCreate) SetWriter(u *User) *PostCreate {
	return pc.SetWriterID(u.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (pc *PostCreate) AddCommentIDs(ids ...int) *PostCreate {
	pc.mutation.AddCommentIDs(ids...)
	return pc
}

// AddComments adds the "comments" edges to the Comment entity.
func (pc *PostCreate) AddComments(c ...*Comment) *PostCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCommentIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	var (
		err  error
		node *Post
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Post)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PostMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PostCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PostCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PostCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := post.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := post.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Post.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Post.updated_at"`)}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Post.title"`)}
	}
	if _, ok := pc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Post.content"`)}
	}
	if _, ok := pc.mutation.Points(); !ok {
		return &ValidationError{Name: "points", err: errors.New(`ent: missing required field "Post.points"`)}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: post.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: post.FieldID,
			},
		}
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: post.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: post.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := pc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := pc.mutation.Points(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: post.FieldPoints,
		})
		_node.Points = value
	}
	if nodes := pc.mutation.WriterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.WriterTable,
			Columns: []string{post.WriterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_posts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Post.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PostUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (pc *PostCreate) OnConflict(opts ...sql.ConflictOption) *PostUpsertOne {
	pc.conflict = opts
	return &PostUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pc *PostCreate) OnConflictColumns(columns ...string) *PostUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PostUpsertOne{
		create: pc,
	}
}

type (
	// PostUpsertOne is the builder for "upsert"-ing
	//  one Post node.
	PostUpsertOne struct {
		create *PostCreate
	}

	// PostUpsert is the "OnConflict" setter.
	PostUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *PostUpsert) SetCreatedAt(v time.Time) *PostUpsert {
	u.Set(post.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PostUpsert) UpdateCreatedAt() *PostUpsert {
	u.SetExcluded(post.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PostUpsert) SetUpdatedAt(v time.Time) *PostUpsert {
	u.Set(post.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PostUpsert) UpdateUpdatedAt() *PostUpsert {
	u.SetExcluded(post.FieldUpdatedAt)
	return u
}

// SetTitle sets the "title" field.
func (u *PostUpsert) SetTitle(v string) *PostUpsert {
	u.Set(post.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PostUpsert) UpdateTitle() *PostUpsert {
	u.SetExcluded(post.FieldTitle)
	return u
}

// SetContent sets the "content" field.
func (u *PostUpsert) SetContent(v string) *PostUpsert {
	u.Set(post.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *PostUpsert) UpdateContent() *PostUpsert {
	u.SetExcluded(post.FieldContent)
	return u
}

// SetPoints sets the "points" field.
func (u *PostUpsert) SetPoints(v float64) *PostUpsert {
	u.Set(post.FieldPoints, v)
	return u
}

// UpdatePoints sets the "points" field to the value that was provided on create.
func (u *PostUpsert) UpdatePoints() *PostUpsert {
	u.SetExcluded(post.FieldPoints)
	return u
}

// AddPoints adds v to the "points" field.
func (u *PostUpsert) AddPoints(v float64) *PostUpsert {
	u.Add(post.FieldPoints, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *PostUpsertOne) UpdateNewValues() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(post.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Post.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *PostUpsertOne) Ignore() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PostUpsertOne) DoNothing() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PostCreate.OnConflict
// documentation for more info.
func (u *PostUpsertOne) Update(set func(*PostUpsert)) *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PostUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PostUpsertOne) SetCreatedAt(v time.Time) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateCreatedAt() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PostUpsertOne) SetUpdatedAt(v time.Time) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateUpdatedAt() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetTitle sets the "title" field.
func (u *PostUpsertOne) SetTitle(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateTitle() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateTitle()
	})
}

// SetContent sets the "content" field.
func (u *PostUpsertOne) SetContent(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateContent() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateContent()
	})
}

// SetPoints sets the "points" field.
func (u *PostUpsertOne) SetPoints(v float64) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetPoints(v)
	})
}

// AddPoints adds v to the "points" field.
func (u *PostUpsertOne) AddPoints(v float64) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.AddPoints(v)
	})
}

// UpdatePoints sets the "points" field to the value that was provided on create.
func (u *PostUpsertOne) UpdatePoints() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdatePoints()
	})
}

// Exec executes the query.
func (u *PostUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PostCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PostUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PostUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PostUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	builders []*PostCreate
	conflict []sql.ConflictOption
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PostCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PostCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Post.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PostUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (pcb *PostCreateBulk) OnConflict(opts ...sql.ConflictOption) *PostUpsertBulk {
	pcb.conflict = opts
	return &PostUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pcb *PostCreateBulk) OnConflictColumns(columns ...string) *PostUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PostUpsertBulk{
		create: pcb,
	}
}

// PostUpsertBulk is the builder for "upsert"-ing
// a bulk of Post nodes.
type PostUpsertBulk struct {
	create *PostCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *PostUpsertBulk) UpdateNewValues() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(post.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *PostUpsertBulk) Ignore() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PostUpsertBulk) DoNothing() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PostCreateBulk.OnConflict
// documentation for more info.
func (u *PostUpsertBulk) Update(set func(*PostUpsert)) *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PostUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PostUpsertBulk) SetCreatedAt(v time.Time) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateCreatedAt() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PostUpsertBulk) SetUpdatedAt(v time.Time) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateUpdatedAt() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetTitle sets the "title" field.
func (u *PostUpsertBulk) SetTitle(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateTitle() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateTitle()
	})
}

// SetContent sets the "content" field.
func (u *PostUpsertBulk) SetContent(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateContent() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateContent()
	})
}

// SetPoints sets the "points" field.
func (u *PostUpsertBulk) SetPoints(v float64) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetPoints(v)
	})
}

// AddPoints adds v to the "points" field.
func (u *PostUpsertBulk) AddPoints(v float64) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.AddPoints(v)
	})
}

// UpdatePoints sets the "points" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdatePoints() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdatePoints()
	})
}

// Exec executes the query.
func (u *PostUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PostCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PostCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PostUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
