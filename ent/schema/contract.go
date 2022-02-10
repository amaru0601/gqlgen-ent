package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Contract holds the schema definition for the Contract entity.
type Contract struct {
	ent.Schema
}

// Fields of the Contract.
func (Contract) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_date"),
		field.Time("end_date"),
		field.Float("pay_amount"),
		field.Time("pay_date"),
	}
}

// Edges of the Contract.
func (Contract) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("contracts"),
		edge.From("property", Property.Type).Ref("contract").Unique(),
	}
}
