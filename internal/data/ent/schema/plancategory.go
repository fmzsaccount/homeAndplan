package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PlanCategory holds the schema definition for the PlanCategory entity.
type PlanCategory struct {
	ent.Schema
}

func (PlanCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "plan_category"},
	}
}

// Fields of the PlanCategory.
func (PlanCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").Default(0),
		field.Int32("sort").Default(0),
		field.String("name").Default(""),
		field.String("content").Default(""),
		//field.String("bucket").Default(""),
		field.String("url").Default(""),
	}
}

// Edges of the PlanCategory.
func (PlanCategory) Edges() []ent.Edge {
	return nil
}
