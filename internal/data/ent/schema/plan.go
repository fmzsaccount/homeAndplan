package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Plan holds the schema definition for the Plan entity.
type Plan struct {
	ent.Schema
}
func (Plan) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "plan"},
	}
}
// Fields of the Plan.
func (Plan) Fields() []ent.Field {
	return []ent.Field{
		field.Int("recommend").Default(0),
		field.Int32("id").Default(0),
		field.Int32("category_id").Default(0),
		field.String("category").Default(""),
		field.String("name").Default(""),
		field.String("url").Default(""),
		field.Time("utime"),

	}
}

// Edges of the Plan.
func (Plan) Edges() []ent.Edge {
	return nil
}
