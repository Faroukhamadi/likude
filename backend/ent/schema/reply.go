package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Faroukhamadi/likude/ent/schema/mixin/timestamps"
)

// Reply holds the schema definition for the Reply entity.
type Reply struct {
	ent.Schema
}

func (Reply) Mixin() []ent.Mixin {
	return []ent.Mixin{
		timestamps.Mixin{},
	}
}

// Fields of the Reply.
func (Reply) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
		field.Float32("points"),
	}
}

// Edges of the Reply.
func (Reply) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("comment", Comment.Type).
			Ref("replies").
			Annotations(entgql.RelayConnection()),
	}
}

// Annotations of the User
func (Reply) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
