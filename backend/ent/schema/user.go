package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"

	"entgo.io/ent/schema"

	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Faroukhamadi/likude/ent/schema/mixin/timestamps"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		timestamps.Mixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Immutable().
			Unique().
			Annotations(
				entgql.OrderField("USERNAME"),
			),
		field.String("password"),
		field.Int("karma").
			Default(0).
			Nillable().
			Optional().
			Annotations(
				entgql.OrderField("KARMA"),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.From("communities", Community.Type).
			Ref("users"),
	}
}

// Annotations of the User
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
