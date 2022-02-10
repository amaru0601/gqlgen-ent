package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Property holds the schema definition for the Property entity.
type Property struct {
	ent.Schema
}

// Fields of the Property.
func (Property) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("class").Values(
			"HOUSE",
			"APARTMENT",
			"PREMISES",
			"OFFICE",
			"VEHICLE",
		).Default("APARTMENT"),
		field.String("address"),
		field.String("city"),
		field.String("description"),
		field.Bool("deleted"),
	}
}

// Edges of the Property.
func (Property) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("properties").Unique(),
		edge.To("contract", Contract.Type),
	}
}
