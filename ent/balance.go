// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent/balance"
	"github.com/skyisboss/pay-system/ent/schema"
)

// Balance is the model entity for the Balance schema.
type Balance struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// ChainID holds the value of the "chain_id" field.
	ChainID uint64 `json:"chain_id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID uint64 `json:"product_id,omitempty"`
	// 当前余额
	BalanceAmount decimal.Decimal `json:"balance_amount,omitempty"`
	// 提款冻结金额
	BalanceFreeze decimal.Decimal `json:"balance_freeze,omitempty"`
	// 总存款金额
	TotalDeposit decimal.Decimal `json:"total_deposit,omitempty"`
	// 总取款金额
	TotalWithdraw decimal.Decimal `json:"total_withdraw,omitempty"`
	// 存款次数
	CountDeposit uint64 `json:"count_deposit,omitempty"`
	// 取款次数
	CountWithdraw uint64 `json:"count_withdraw,omitempty"`
	// 账变记录
	ChangeLogs []schema.ChangeLogs `json:"change_logs,omitempty"`
	// 数据版本
	Version      int64 `json:"version,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Balance) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case balance.FieldChangeLogs:
			values[i] = new([]byte)
		case balance.FieldBalanceAmount, balance.FieldBalanceFreeze, balance.FieldTotalDeposit, balance.FieldTotalWithdraw:
			values[i] = new(decimal.Decimal)
		case balance.FieldID, balance.FieldChainID, balance.FieldProductID, balance.FieldCountDeposit, balance.FieldCountWithdraw, balance.FieldVersion:
			values[i] = new(sql.NullInt64)
		case balance.FieldCreatedAt, balance.FieldUpdatedAt, balance.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Balance fields.
func (b *Balance) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case balance.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = uint64(value.Int64)
		case balance.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case balance.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case balance.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				b.DeletedAt = value.Time
			}
		case balance.FieldChainID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field chain_id", values[i])
			} else if value.Valid {
				b.ChainID = uint64(value.Int64)
			}
		case balance.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				b.ProductID = uint64(value.Int64)
			}
		case balance.FieldBalanceAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field balance_amount", values[i])
			} else if value != nil {
				b.BalanceAmount = *value
			}
		case balance.FieldBalanceFreeze:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field balance_freeze", values[i])
			} else if value != nil {
				b.BalanceFreeze = *value
			}
		case balance.FieldTotalDeposit:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field total_deposit", values[i])
			} else if value != nil {
				b.TotalDeposit = *value
			}
		case balance.FieldTotalWithdraw:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field total_withdraw", values[i])
			} else if value != nil {
				b.TotalWithdraw = *value
			}
		case balance.FieldCountDeposit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count_deposit", values[i])
			} else if value.Valid {
				b.CountDeposit = uint64(value.Int64)
			}
		case balance.FieldCountWithdraw:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count_withdraw", values[i])
			} else if value.Valid {
				b.CountWithdraw = uint64(value.Int64)
			}
		case balance.FieldChangeLogs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field change_logs", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &b.ChangeLogs); err != nil {
					return fmt.Errorf("unmarshal field change_logs: %w", err)
				}
			}
		case balance.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				b.Version = value.Int64
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Balance.
// This includes values selected through modifiers, order, etc.
func (b *Balance) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// Update returns a builder for updating this Balance.
// Note that you need to call Balance.Unwrap() before calling this method if this Balance
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Balance) Update() *BalanceUpdateOne {
	return NewBalanceClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Balance entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Balance) Unwrap() *Balance {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Balance is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Balance) String() string {
	var builder strings.Builder
	builder.WriteString("Balance(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(b.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("chain_id=")
	builder.WriteString(fmt.Sprintf("%v", b.ChainID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", b.ProductID))
	builder.WriteString(", ")
	builder.WriteString("balance_amount=")
	builder.WriteString(fmt.Sprintf("%v", b.BalanceAmount))
	builder.WriteString(", ")
	builder.WriteString("balance_freeze=")
	builder.WriteString(fmt.Sprintf("%v", b.BalanceFreeze))
	builder.WriteString(", ")
	builder.WriteString("total_deposit=")
	builder.WriteString(fmt.Sprintf("%v", b.TotalDeposit))
	builder.WriteString(", ")
	builder.WriteString("total_withdraw=")
	builder.WriteString(fmt.Sprintf("%v", b.TotalWithdraw))
	builder.WriteString(", ")
	builder.WriteString("count_deposit=")
	builder.WriteString(fmt.Sprintf("%v", b.CountDeposit))
	builder.WriteString(", ")
	builder.WriteString("count_withdraw=")
	builder.WriteString(fmt.Sprintf("%v", b.CountWithdraw))
	builder.WriteString(", ")
	builder.WriteString("change_logs=")
	builder.WriteString(fmt.Sprintf("%v", b.ChangeLogs))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", b.Version))
	builder.WriteByte(')')
	return builder.String()
}

// Balances is a parsable slice of Balance.
type Balances []*Balance