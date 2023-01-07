package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Site holds the schema definition for the Site entity.
type Site struct {
	ent.Schema
}

func (Site) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "site"},
	}
}


// Fields of the Site.
func (Site) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.Int32("sort"),
		field.Int("status"),
		field.String("name"),
		field.String("link"),
		field.String("category"),
	}
}

// Edges of the Site.
func (Site) Edges() []ent.Edge {
	return nil
}
