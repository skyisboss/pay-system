package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Notify struct {
	ent.Schema
}

func (Notify) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
		field.Uint64("chain_id"),
		field.Uint64("product_id"),
		field.Uint64("item_from").Comment("id关联来自"),
		field.Int64("item_type"),
		field.String("nonce"),
		field.String("notify_type"),
		field.String("send_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)",
		}),
		field.String("send_body").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)",
		}),
		field.Int64("send_retry"),
		field.Int64("handle_status").Default(0),
		field.String("handle_msg").Default("init"),
		field.Time("handle_time").Optional(),
	}
}
func (Notify) Edges() []ent.Edge {
	return nil
}
func (Notify) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "notify"}}
}
