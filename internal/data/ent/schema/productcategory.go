package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductCategory holds the schema definition for the ProductCategory entity.
type ProductCategory struct {
	ent.Schema
}

func (ProductCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_category"},
	}
}

// Fields of the ProductCategory.
func (ProductCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("status").Default(0),
		field.Int32("id").Default(0),
		field.Int32("level").Default(0),
		field.Int32("parent_id").Default(0).Optional(),
		field.Int32("sort").Default(0),
		field.String("name").Default(""),
	}
}

// Edges of the ProductCategory.
func (ProductCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProductCategory.Type).
			From("parent").
			Unique().
			Field("parent_id"),
		edge.To("products", Product.Type),
	}
}
