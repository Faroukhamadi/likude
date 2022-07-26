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
	"github.com/Faroukhamadi/likude/ent/reply"
)

// ReplyCreate is the builder for creating a Reply entity.
type ReplyCreate struct {
	config
	mutation *ReplyMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (rc *ReplyCreate) SetCreatedAt(t time.Time) *ReplyCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *ReplyCreate) SetNillableCreatedAt(t *time.Time) *ReplyCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *ReplyCreate) SetUpdatedAt(t time.Time) *ReplyCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *ReplyCreate) SetNillableUpdatedAt(t *time.Time) *ReplyCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetContent sets the "content" field.
func (rc *ReplyCreate) SetContent(s string) *ReplyCreate {
	rc.mutation.SetContent(s)
	return rc
}

// SetPoints sets the "points" field.
func (rc *ReplyCreate) SetPoints(f float64) *ReplyCreate {
	rc.mutation.SetPoints(f)
	return rc
}

// AddCommentIDs adds the "comment" edge to the Comment entity by IDs.
func (rc *ReplyCreate) AddCommentIDs(ids ...int) *ReplyCreate {
	rc.mutation.AddCommentIDs(ids...)
	return rc
}

// AddComment adds the "comment" edges to the Comment entity.
func (rc *ReplyCreate) AddComment(c ...*Comment) *ReplyCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return rc.AddCommentIDs(ids...)
}

// Mutation returns the ReplyMutation object of the builder.
func (rc *ReplyCreate) Mutation() *ReplyMutation {
	return rc.mutation
}

// Save creates the Reply in the database.
func (rc *ReplyCreate) Save(ctx context.Context) (*Reply, error) {
	var (
		err  error
		node *Reply
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReplyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Reply)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ReplyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReplyCreate) SaveX(ctx context.Context) *Reply {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReplyCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReplyCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReplyCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := reply.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := reply.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReplyCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Reply.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Reply.updated_at"`)}
	}
	if _, ok := rc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Reply.content"`)}
	}
	if _, ok := rc.mutation.Points(); !ok {
		return &ValidationError{Name: "points", err: errors.New(`ent: missing required field "Reply.points"`)}
	}
	return nil
}

func (rc *ReplyCreate) sqlSave(ctx context.Context) (*Reply, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *ReplyCreate) createSpec() (*Reply, *sqlgraph.CreateSpec) {
	var (
		_node = &Reply{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: reply.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: reply.FieldID,
			},
		}
	)
	_spec.OnConflict = rc.conflict
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: reply.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: reply.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reply.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := rc.mutation.Points(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: reply.FieldPoints,
		})
		_node.Points = value
	}
	if nodes := rc.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   reply.CommentTable,
			Columns: reply.CommentPrimaryKey,
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
//	client.Reply.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReplyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (rc *ReplyCreate) OnConflict(opts ...sql.ConflictOption) *ReplyUpsertOne {
	rc.conflict = opts
	return &ReplyUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Reply.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (rc *ReplyCreate) OnConflictColumns(columns ...string) *ReplyUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &ReplyUpsertOne{
		create: rc,
	}
}

type (
	// ReplyUpsertOne is the builder for "upsert"-ing
	//  one Reply node.
	ReplyUpsertOne struct {
		create *ReplyCreate
	}

	// ReplyUpsert is the "OnConflict" setter.
	ReplyUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *ReplyUpsert) SetCreatedAt(v time.Time) *ReplyUpsert {
	u.Set(reply.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ReplyUpsert) UpdateCreatedAt() *ReplyUpsert {
	u.SetExcluded(reply.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ReplyUpsert) SetUpdatedAt(v time.Time) *ReplyUpsert {
	u.Set(reply.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ReplyUpsert) UpdateUpdatedAt() *ReplyUpsert {
	u.SetExcluded(reply.FieldUpdatedAt)
	return u
}

// SetContent sets the "content" field.
func (u *ReplyUpsert) SetContent(v string) *ReplyUpsert {
	u.Set(reply.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *ReplyUpsert) UpdateContent() *ReplyUpsert {
	u.SetExcluded(reply.FieldContent)
	return u
}

// SetPoints sets the "points" field.
func (u *ReplyUpsert) SetPoints(v float64) *ReplyUpsert {
	u.Set(reply.FieldPoints, v)
	return u
}

// UpdatePoints sets the "points" field to the value that was provided on create.
func (u *ReplyUpsert) UpdatePoints() *ReplyUpsert {
	u.SetExcluded(reply.FieldPoints)
	return u
}

// AddPoints adds v to the "points" field.
func (u *ReplyUpsert) AddPoints(v float64) *ReplyUpsert {
	u.Add(reply.FieldPoints, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Reply.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ReplyUpsertOne) UpdateNewValues() *ReplyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(reply.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Reply.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ReplyUpsertOne) Ignore() *ReplyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReplyUpsertOne) DoNothing() *ReplyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReplyCreate.OnConflict
// documentation for more info.
func (u *ReplyUpsertOne) Update(set func(*ReplyUpsert)) *ReplyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReplyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ReplyUpsertOne) SetCreatedAt(v time.Time) *ReplyUpsertOne {
	return u.Update(func(s *ReplyUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ReplyUpsertOne) UpdateCreatedAt() *ReplyUpsertOne {
	return u.Update(func(s *ReplyUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ReplyUpsertOne) SetUpdatedAt(v time.Time) *ReplyUpsertOne {
	return u.Update(func(s *ReplyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ReplyUpsertOne) UpdateUpdatedAt() *ReplyUpsertOne {
	return u.Update(func(s *ReplyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetContent sets the "content" field.
func (u *ReplyUpsertOne) SetContent(v string) *ReplyUpsertOne {
	return u.Update(func(s *ReplyUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *ReplyUpsertOne) UpdateContent() *ReplyUpsertOne {
	return u.Update(func(s *ReplyUpsert) {
		s.UpdateContent()
	})
}

// SetPoints sets the "points" field.
func (u *ReplyUpsertOne) SetPoints(v float64) *ReplyUpsertOne {
	return u.Update(func(s *ReplyUpsert) {
		s.SetPoints(v)
	})
}

// AddPoints adds v to the "points" field.
func (u *ReplyUpsertOne) AddPoints(v float64) *ReplyUpsertOne {
	return u.Update(func(s *ReplyUpsert) {
		s.AddPoints(v)
	})
}

// UpdatePoints sets the "points" field to the value that was provided on create.
func (u *ReplyUpsertOne) UpdatePoints() *ReplyUpsertOne {
	return u.Update(func(s *ReplyUpsert) {
		s.UpdatePoints()
	})
}

// Exec executes the query.
func (u *ReplyUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReplyCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReplyUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ReplyUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ReplyUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ReplyCreateBulk is the builder for creating many Reply entities in bulk.
type ReplyCreateBulk struct {
	config
	builders []*ReplyCreate
	conflict []sql.ConflictOption
}

// Save creates the Reply entities in the database.
func (rcb *ReplyCreateBulk) Save(ctx context.Context) ([]*Reply, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Reply, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReplyMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReplyCreateBulk) SaveX(ctx context.Context) []*Reply {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReplyCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReplyCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Reply.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReplyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (rcb *ReplyCreateBulk) OnConflict(opts ...sql.ConflictOption) *ReplyUpsertBulk {
	rcb.conflict = opts
	return &ReplyUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Reply.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (rcb *ReplyCreateBulk) OnConflictColumns(columns ...string) *ReplyUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &ReplyUpsertBulk{
		create: rcb,
	}
}

// ReplyUpsertBulk is the builder for "upsert"-ing
// a bulk of Reply nodes.
type ReplyUpsertBulk struct {
	create *ReplyCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Reply.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ReplyUpsertBulk) UpdateNewValues() *ReplyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(reply.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Reply.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ReplyUpsertBulk) Ignore() *ReplyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReplyUpsertBulk) DoNothing() *ReplyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReplyCreateBulk.OnConflict
// documentation for more info.
func (u *ReplyUpsertBulk) Update(set func(*ReplyUpsert)) *ReplyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReplyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ReplyUpsertBulk) SetCreatedAt(v time.Time) *ReplyUpsertBulk {
	return u.Update(func(s *ReplyUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ReplyUpsertBulk) UpdateCreatedAt() *ReplyUpsertBulk {
	return u.Update(func(s *ReplyUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ReplyUpsertBulk) SetUpdatedAt(v time.Time) *ReplyUpsertBulk {
	return u.Update(func(s *ReplyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ReplyUpsertBulk) UpdateUpdatedAt() *ReplyUpsertBulk {
	return u.Update(func(s *ReplyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetContent sets the "content" field.
func (u *ReplyUpsertBulk) SetContent(v string) *ReplyUpsertBulk {
	return u.Update(func(s *ReplyUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *ReplyUpsertBulk) UpdateContent() *ReplyUpsertBulk {
	return u.Update(func(s *ReplyUpsert) {
		s.UpdateContent()
	})
}

// SetPoints sets the "points" field.
func (u *ReplyUpsertBulk) SetPoints(v float64) *ReplyUpsertBulk {
	return u.Update(func(s *ReplyUpsert) {
		s.SetPoints(v)
	})
}

// AddPoints adds v to the "points" field.
func (u *ReplyUpsertBulk) AddPoints(v float64) *ReplyUpsertBulk {
	return u.Update(func(s *ReplyUpsert) {
		s.AddPoints(v)
	})
}

// UpdatePoints sets the "points" field to the value that was provided on create.
func (u *ReplyUpsertBulk) UpdatePoints() *ReplyUpsertBulk {
	return u.Update(func(s *ReplyUpsert) {
		s.UpdatePoints()
	})
}

// Exec executes the query.
func (u *ReplyUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ReplyCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReplyCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReplyUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
