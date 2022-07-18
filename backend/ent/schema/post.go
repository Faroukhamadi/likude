package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Faroukhamadi/likude/ent/schema/mixin/timestamps"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		timestamps.Mixin{},
	}
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("content"),
		field.Float("points"),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("writer", User.Type).
			Ref("posts"),
		edge.To("comments", Comment.Type),
	}
}

// Annotations of the User
func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
