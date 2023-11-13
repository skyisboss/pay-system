package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

type Withdraw struct {
	ent.Schema
}

func (Withdraw) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
		field.Int64("product_id"),
		field.Uint64("chain_id"),
		field.String("to_address"),
		// field.String("amount"),
		field.String("amount_str"),
		field.Float("amount_raw").GoType(decimal.Decimal{}).Comment("wei单位无小数点").SchemaType(map[string]string{
			dialect.MySQL: "decimal(32,0)",
		}),
		field.String("serial_id"),
		field.String("tx_hash").Optional(),
		field.Int64("handle_status").Default(0),
		field.String("handle_msg").Optional(),
		field.Time("handle_time").Optional(),
	}
}
func (Withdraw) Edges() []ent.Edge {
	return nil
}
func (Withdraw) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "withdraw"}}
}
