// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/skyisboss/pay-system/ent/addres"
	"github.com/skyisboss/pay-system/ent/predicate"
)

// AddresUpdate is the builder for updating Addres entities.
type AddresUpdate struct {
	config
	hooks    []Hook
	mutation *AddresMutation
}

// Where appends a list predicates to the AddresUpdate builder.
func (au *AddresUpdate) Where(ps ...predicate.Addres) *AddresUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetChainID sets the "chain_id" field.
func (au *AddresUpdate) SetChainID(u uint64) *AddresUpdate {
	au.mutation.ResetChainID()
	au.mutation.SetChainID(u)
	return au
}

// AddChainID adds u to the "chain_id" field.
func (au *AddresUpdate) AddChainID(u int64) *AddresUpdate {
	au.mutation.AddChainID(u)
	return au
}

// SetAddress sets the "address" field.
func (au *AddresUpdate) SetAddress(s string) *AddresUpdate {
	au.mutation.SetAddress(s)
	return au
}

// SetPassword sets the "password" field.
func (au *AddresUpdate) SetPassword(s string) *AddresUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetUUID sets the "uuid" field.
func (au *AddresUpdate) SetUUID(s string) *AddresUpdate {
	au.mutation.SetUUID(s)
	return au
}

// SetUseTo sets the "use_to" field.
func (au *AddresUpdate) SetUseTo(i int64) *AddresUpdate {
	au.mutation.ResetUseTo()
	au.mutation.SetUseTo(i)
	return au
}

// AddUseTo adds i to the "use_to" field.
func (au *AddresUpdate) AddUseTo(i int64) *AddresUpdate {
	au.mutation.AddUseTo(i)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AddresUpdate) SetCreatedAt(t time.Time) *AddresUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AddresUpdate) SetNillableCreatedAt(t *time.Time) *AddresUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// ClearCreatedAt clears the value of the "created_at" field.
func (au *AddresUpdate) ClearCreatedAt() *AddresUpdate {
	au.mutation.ClearCreatedAt()
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AddresUpdate) SetUpdatedAt(t time.Time) *AddresUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (au *AddresUpdate) SetNillableUpdatedAt(t *time.Time) *AddresUpdate {
	if t != nil {
		au.SetUpdatedAt(*t)
	}
	return au
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (au *AddresUpdate) ClearUpdatedAt() *AddresUpdate {
	au.mutation.ClearUpdatedAt()
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *AddresUpdate) SetDeletedAt(t time.Time) *AddresUpdate {
	au.mutation.SetDeletedAt(t)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *AddresUpdate) SetNillableDeletedAt(t *time.Time) *AddresUpdate {
	if t != nil {
		au.SetDeletedAt(*t)
	}
	return au
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (au *AddresUpdate) ClearDeletedAt() *AddresUpdate {
	au.mutation.ClearDeletedAt()
	return au
}

// Mutation returns the AddresMutation object of the builder.
func (au *AddresUpdate) Mutation() *AddresMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AddresUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AddresUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AddresUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AddresUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AddresUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(addres.Table, addres.Columns, sqlgraph.NewFieldSpec(addres.FieldID, field.TypeUint64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.ChainID(); ok {
		_spec.SetField(addres.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := au.mutation.AddedChainID(); ok {
		_spec.AddField(addres.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := au.mutation.Address(); ok {
		_spec.SetField(addres.FieldAddress, field.TypeString, value)
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.SetField(addres.FieldPassword, field.TypeString, value)
	}
	if value, ok := au.mutation.UUID(); ok {
		_spec.SetField(addres.FieldUUID, field.TypeString, value)
	}
	if value, ok := au.mutation.UseTo(); ok {
		_spec.SetField(addres.FieldUseTo, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedUseTo(); ok {
		_spec.AddField(addres.FieldUseTo, field.TypeInt64, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(addres.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.CreatedAtCleared() {
		_spec.ClearField(addres.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(addres.FieldUpdatedAt, field.TypeTime, value)
	}
	if au.mutation.UpdatedAtCleared() {
		_spec.ClearField(addres.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.SetField(addres.FieldDeletedAt, field.TypeTime, value)
	}
	if au.mutation.DeletedAtCleared() {
		_spec.ClearField(addres.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{addres.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AddresUpdateOne is the builder for updating a single Addres entity.
type AddresUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AddresMutation
}

// SetChainID sets the "chain_id" field.
func (auo *AddresUpdateOne) SetChainID(u uint64) *AddresUpdateOne {
	auo.mutation.ResetChainID()
	auo.mutation.SetChainID(u)
	return auo
}

// AddChainID adds u to the "chain_id" field.
func (auo *AddresUpdateOne) AddChainID(u int64) *AddresUpdateOne {
	auo.mutation.AddChainID(u)
	return auo
}

// SetAddress sets the "address" field.
func (auo *AddresUpdateOne) SetAddress(s string) *AddresUpdateOne {
	auo.mutation.SetAddress(s)
	return auo
}

// SetPassword sets the "password" field.
func (auo *AddresUpdateOne) SetPassword(s string) *AddresUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetUUID sets the "uuid" field.
func (auo *AddresUpdateOne) SetUUID(s string) *AddresUpdateOne {
	auo.mutation.SetUUID(s)
	return auo
}

// SetUseTo sets the "use_to" field.
func (auo *AddresUpdateOne) SetUseTo(i int64) *AddresUpdateOne {
	auo.mutation.ResetUseTo()
	auo.mutation.SetUseTo(i)
	return auo
}

// AddUseTo adds i to the "use_to" field.
func (auo *AddresUpdateOne) AddUseTo(i int64) *AddresUpdateOne {
	auo.mutation.AddUseTo(i)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AddresUpdateOne) SetCreatedAt(t time.Time) *AddresUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AddresUpdateOne) SetNillableCreatedAt(t *time.Time) *AddresUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (auo *AddresUpdateOne) ClearCreatedAt() *AddresUpdateOne {
	auo.mutation.ClearCreatedAt()
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AddresUpdateOne) SetUpdatedAt(t time.Time) *AddresUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auo *AddresUpdateOne) SetNillableUpdatedAt(t *time.Time) *AddresUpdateOne {
	if t != nil {
		auo.SetUpdatedAt(*t)
	}
	return auo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (auo *AddresUpdateOne) ClearUpdatedAt() *AddresUpdateOne {
	auo.mutation.ClearUpdatedAt()
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *AddresUpdateOne) SetDeletedAt(t time.Time) *AddresUpdateOne {
	auo.mutation.SetDeletedAt(t)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *AddresUpdateOne) SetNillableDeletedAt(t *time.Time) *AddresUpdateOne {
	if t != nil {
		auo.SetDeletedAt(*t)
	}
	return auo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (auo *AddresUpdateOne) ClearDeletedAt() *AddresUpdateOne {
	auo.mutation.ClearDeletedAt()
	return auo
}

// Mutation returns the AddresMutation object of the builder.
func (auo *AddresUpdateOne) Mutation() *AddresMutation {
	return auo.mutation
}

// Where appends a list predicates to the AddresUpdate builder.
func (auo *AddresUpdateOne) Where(ps ...predicate.Addres) *AddresUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AddresUpdateOne) Select(field string, fields ...string) *AddresUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Addres entity.
func (auo *AddresUpdateOne) Save(ctx context.Context) (*Addres, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AddresUpdateOne) SaveX(ctx context.Context) *Addres {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AddresUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AddresUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AddresUpdateOne) sqlSave(ctx context.Context) (_node *Addres, err error) {
	_spec := sqlgraph.NewUpdateSpec(addres.Table, addres.Columns, sqlgraph.NewFieldSpec(addres.FieldID, field.TypeUint64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Addres.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, addres.FieldID)
		for _, f := range fields {
			if !addres.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != addres.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.ChainID(); ok {
		_spec.SetField(addres.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := auo.mutation.AddedChainID(); ok {
		_spec.AddField(addres.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := auo.mutation.Address(); ok {
		_spec.SetField(addres.FieldAddress, field.TypeString, value)
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.SetField(addres.FieldPassword, field.TypeString, value)
	}
	if value, ok := auo.mutation.UUID(); ok {
		_spec.SetField(addres.FieldUUID, field.TypeString, value)
	}
	if value, ok := auo.mutation.UseTo(); ok {
		_spec.SetField(addres.FieldUseTo, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedUseTo(); ok {
		_spec.AddField(addres.FieldUseTo, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(addres.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.CreatedAtCleared() {
		_spec.ClearField(addres.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(addres.FieldUpdatedAt, field.TypeTime, value)
	}
	if auo.mutation.UpdatedAtCleared() {
		_spec.ClearField(addres.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.SetField(addres.FieldDeletedAt, field.TypeTime, value)
	}
	if auo.mutation.DeletedAtCleared() {
		_spec.ClearField(addres.FieldDeletedAt, field.TypeTime)
	}
	_node = &Addres{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{addres.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
