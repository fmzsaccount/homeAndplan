package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"gy_home/internal/biz"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product"},
	}
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int("status").Default(0),
		field.Int("recommend").Default(0),
		field.Int32("id").Default(0),
		field.Int32("history").Default(0),
		field.Int32("category_id").Default(0).Optional(),
		field.String("category_name").Default(""),
		field.String("name").Default(""),
		field.String("content").Default(""),
		field.Time("utime"),
		field.String("service_link").Default(""),
		field.String("apply_link").Default(""),
		field.JSON("json",biz.Product{}),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent_id", ProductCategory.Type).
			Ref("products").
			Unique().
			Field("category_id"),
	}
}
