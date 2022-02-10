package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("names").NotEmpty(),
		field.String("lastnames").NotEmpty(),
		field.Time("birthday"),
		field.String("email").NotEmpty(),
		field.Bool("activate"),
		field.Time("createdAt"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("properties", Property.Type),
		edge.To("contracts", Contract.Type),
	}
}
