package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Apprun struct {
	ent.Schema
}

func (Apprun) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Time("created_at").Optional(),
		field.Time("updated_at").Optional(),
		field.String("handler").Unique(),
		field.Uint64("runing").Default(0),
		field.Uint64("total").Default(0),
	}
}
func (Apprun) Edges() []ent.Edge {
	return nil
}
func (Apprun) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "apprun"}}
}
