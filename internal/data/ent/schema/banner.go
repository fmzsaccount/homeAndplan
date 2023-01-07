package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Banner holds the schema definition for the Banner entity.
type Banner struct {
	ent.Schema
}

func (Banner) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "banner"},
	}
}

// Fields of the Banner.
func (Banner) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").Default(""),
		field.Int32("id").Default(0),
		field.Int32("sort").Default(0),
		field.String("title").Default(""),
		field.String("type").Default(""),
		//field.String("bucket").Default(""),
		field.String("link_id").Default(""),
		field.Int("status").Default(0),
		//field.Bool("recommend").Default(false),
	}
}

// Edges of the Banner.
func (Banner) Edges() []ent.Edge {
	return nil
}
