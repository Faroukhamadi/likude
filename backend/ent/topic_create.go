// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Faroukhamadi/likude/ent/topic"
	"github.com/Faroukhamadi/likude/ent/topicrelated"
)

// TopicCreate is the builder for creating a Topic entity.
type TopicCreate struct {
	config
	mutation *TopicMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TopicCreate) SetName(s string) *TopicCreate {
	tc.mutation.SetName(s)
	return tc
}

// AddRelatedTopicIDs adds the "related_topics" edge to the Topic entity by IDs.
func (tc *TopicCreate) AddRelatedTopicIDs(ids ...int) *TopicCreate {
	tc.mutation.AddRelatedTopicIDs(ids...)
	return tc
}

// AddRelatedTopics adds the "related_topics" edges to the Topic entity.
func (tc *TopicCreate) AddRelatedTopics(t ...*Topic) *TopicCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddRelatedTopicIDs(ids...)
}

// AddTopicRelationIDs adds the "topic_relations" edge to the TopicRelated entity by IDs.
func (tc *TopicCreate) AddTopicRelationIDs(ids ...int) *TopicCreate {
	tc.mutation.AddTopicRelationIDs(ids...)
	return tc
}

// AddTopicRelations adds the "topic_relations" edges to the TopicRelated entity.
func (tc *TopicCreate) AddTopicRelations(t ...*TopicRelated) *TopicCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTopicRelationIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tc *TopicCreate) Mutation() *TopicMutation {
	return tc.mutation
}

// Save creates the Topic in the database.
func (tc *TopicCreate) Save(ctx context.Context) (*Topic, error) {
	var (
		err  error
		node *Topic
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Topic)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TopicMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TopicCreate) SaveX(ctx context.Context) *Topic {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TopicCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TopicCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TopicCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Topic.name"`)}
	}
	return nil
}

func (tc *TopicCreate) sqlSave(ctx context.Context) (*Topic, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TopicCreate) createSpec() (*Topic, *sqlgraph.CreateSpec) {
	var (
		_node = &Topic{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: topic.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: topic.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldName,
		})
		_node.Name = value
	}
	if nodes := tc.mutation.RelatedTopicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.RelatedTopicsTable,
			Columns: topic.RelatedTopicsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TopicRelationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   topic.TopicRelationsTable,
			Columns: []string{topic.TopicRelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topicrelated.FieldID,
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

// TopicCreateBulk is the builder for creating many Topic entities in bulk.
type TopicCreateBulk struct {
	config
	builders []*TopicCreate
}

// Save creates the Topic entities in the database.
func (tcb *TopicCreateBulk) Save(ctx context.Context) ([]*Topic, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Topic, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TopicMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TopicCreateBulk) SaveX(ctx context.Context) []*Topic {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TopicCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TopicCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}