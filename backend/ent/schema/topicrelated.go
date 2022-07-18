package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TopicRelated holds the schema definition for the TopicRelated entity.
type TopicRelated struct {
	ent.Schema
}

// Fields of the TopicRelated.
func (TopicRelated) Fields() []ent.Field {
	return []ent.Field{
		field.Int("topic_id"),
		field.Int("related_topic_id"),
	}
}

// Edges of the TopicRelated.
func (TopicRelated) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("topic", Topic.Type).
			Required().
			Unique().
			Field("topic_id"),
		edge.To("related_topic", Topic.Type).
			Required().
			Unique().
			Field("related_topic_id"),
	}
}

// Annotations of the User
func (TopicRelated) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
