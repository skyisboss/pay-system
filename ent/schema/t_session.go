package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type TSession struct {
	ent.Schema
}

func (TSession) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("key_name"),
		field.String("key_value"),
		field.String("ip"),
		field.Time("created_at").Optional(),
		field.Time("updated_at").Optional(),
	}
}
func (TSession) Edges() []ent.Edge {
	return nil
}
func (TSession) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "t_session"}}
}
