// Code generated by ent, DO NOT EDIT.

package apprun

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the apprun type in the database.
	Label = "apprun"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldHandler holds the string denoting the handler field in the database.
	FieldHandler = "handler"
	// FieldRuning holds the string denoting the runing field in the database.
	FieldRuning = "runing"
	// FieldTotal holds the string denoting the total field in the database.
	FieldTotal = "total"
	// Table holds the table name of the apprun in the database.
	Table = "apprun"
)

// Columns holds all SQL columns for apprun fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldHandler,
	FieldRuning,
	FieldTotal,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultRuning holds the default value on creation for the "runing" field.
	DefaultRuning uint64
	// DefaultTotal holds the default value on creation for the "total" field.
	DefaultTotal uint64
)

// OrderOption defines the ordering options for the Apprun queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByHandler orders the results by the handler field.
func ByHandler(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHandler, opts...).ToFunc()
}

// ByRuning orders the results by the runing field.
func ByRuning(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRuning, opts...).ToFunc()
}

// ByTotal orders the results by the total field.
func ByTotal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotal, opts...).ToFunc()
}
