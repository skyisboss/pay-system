// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/skyisboss/pay-system/ent/apprun"
)

// Apprun is the model entity for the Apprun schema.
type Apprun struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Handler holds the value of the "handler" field.
	Handler string `json:"handler,omitempty"`
	// Runing holds the value of the "runing" field.
	Runing uint64 `json:"runing,omitempty"`
	// Total holds the value of the "total" field.
	Total        uint64 `json:"total,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Apprun) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case apprun.FieldID, apprun.FieldRuning, apprun.FieldTotal:
			values[i] = new(sql.NullInt64)
		case apprun.FieldHandler:
			values[i] = new(sql.NullString)
		case apprun.FieldCreatedAt, apprun.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Apprun fields.
func (a *Apprun) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apprun.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint64(value.Int64)
		case apprun.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case apprun.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case apprun.FieldHandler:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field handler", values[i])
			} else if value.Valid {
				a.Handler = value.String
			}
		case apprun.FieldRuning:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field runing", values[i])
			} else if value.Valid {
				a.Runing = uint64(value.Int64)
			}
		case apprun.FieldTotal:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total", values[i])
			} else if value.Valid {
				a.Total = uint64(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Apprun.
// This includes values selected through modifiers, order, etc.
func (a *Apprun) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Apprun.
// Note that you need to call Apprun.Unwrap() before calling this method if this Apprun
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Apprun) Update() *ApprunUpdateOne {
	return NewApprunClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Apprun entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Apprun) Unwrap() *Apprun {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Apprun is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Apprun) String() string {
	var builder strings.Builder
	builder.WriteString("Apprun(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("handler=")
	builder.WriteString(a.Handler)
	builder.WriteString(", ")
	builder.WriteString("runing=")
	builder.WriteString(fmt.Sprintf("%v", a.Runing))
	builder.WriteString(", ")
	builder.WriteString("total=")
	builder.WriteString(fmt.Sprintf("%v", a.Total))
	builder.WriteByte(')')
	return builder.String()
}

// Appruns is a parsable slice of Apprun.
type Appruns []*Apprun
