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
	"github.com/skyisboss/pay-system/ent/apprun"
	"github.com/skyisboss/pay-system/ent/predicate"
)

// ApprunUpdate is the builder for updating Apprun entities.
type ApprunUpdate struct {
	config
	hooks    []Hook
	mutation *ApprunMutation
}

// Where appends a list predicates to the ApprunUpdate builder.
func (au *ApprunUpdate) Where(ps ...predicate.Apprun) *ApprunUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *ApprunUpdate) SetCreatedAt(t time.Time) *ApprunUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *ApprunUpdate) SetNillableCreatedAt(t *time.Time) *ApprunUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// ClearCreatedAt clears the value of the "created_at" field.
func (au *ApprunUpdate) ClearCreatedAt() *ApprunUpdate {
	au.mutation.ClearCreatedAt()
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *ApprunUpdate) SetUpdatedAt(t time.Time) *ApprunUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (au *ApprunUpdate) SetNillableUpdatedAt(t *time.Time) *ApprunUpdate {
	if t != nil {
		au.SetUpdatedAt(*t)
	}
	return au
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (au *ApprunUpdate) ClearUpdatedAt() *ApprunUpdate {
	au.mutation.ClearUpdatedAt()
	return au
}

// SetHandler sets the "handler" field.
func (au *ApprunUpdate) SetHandler(s string) *ApprunUpdate {
	au.mutation.SetHandler(s)
	return au
}

// SetRuning sets the "runing" field.
func (au *ApprunUpdate) SetRuning(u uint64) *ApprunUpdate {
	au.mutation.ResetRuning()
	au.mutation.SetRuning(u)
	return au
}

// SetNillableRuning sets the "runing" field if the given value is not nil.
func (au *ApprunUpdate) SetNillableRuning(u *uint64) *ApprunUpdate {
	if u != nil {
		au.SetRuning(*u)
	}
	return au
}

// AddRuning adds u to the "runing" field.
func (au *ApprunUpdate) AddRuning(u int64) *ApprunUpdate {
	au.mutation.AddRuning(u)
	return au
}

// SetTotal sets the "total" field.
func (au *ApprunUpdate) SetTotal(u uint64) *ApprunUpdate {
	au.mutation.ResetTotal()
	au.mutation.SetTotal(u)
	return au
}

// SetNillableTotal sets the "total" field if the given value is not nil.
func (au *ApprunUpdate) SetNillableTotal(u *uint64) *ApprunUpdate {
	if u != nil {
		au.SetTotal(*u)
	}
	return au
}

// AddTotal adds u to the "total" field.
func (au *ApprunUpdate) AddTotal(u int64) *ApprunUpdate {
	au.mutation.AddTotal(u)
	return au
}

// Mutation returns the ApprunMutation object of the builder.
func (au *ApprunUpdate) Mutation() *ApprunMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ApprunUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ApprunUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ApprunUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ApprunUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ApprunUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(apprun.Table, apprun.Columns, sqlgraph.NewFieldSpec(apprun.FieldID, field.TypeUint64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(apprun.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.CreatedAtCleared() {
		_spec.ClearField(apprun.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(apprun.FieldUpdatedAt, field.TypeTime, value)
	}
	if au.mutation.UpdatedAtCleared() {
		_spec.ClearField(apprun.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := au.mutation.Handler(); ok {
		_spec.SetField(apprun.FieldHandler, field.TypeString, value)
	}
	if value, ok := au.mutation.Runing(); ok {
		_spec.SetField(apprun.FieldRuning, field.TypeUint64, value)
	}
	if value, ok := au.mutation.AddedRuning(); ok {
		_spec.AddField(apprun.FieldRuning, field.TypeUint64, value)
	}
	if value, ok := au.mutation.Total(); ok {
		_spec.SetField(apprun.FieldTotal, field.TypeUint64, value)
	}
	if value, ok := au.mutation.AddedTotal(); ok {
		_spec.AddField(apprun.FieldTotal, field.TypeUint64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apprun.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ApprunUpdateOne is the builder for updating a single Apprun entity.
type ApprunUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApprunMutation
}

// SetCreatedAt sets the "created_at" field.
func (auo *ApprunUpdateOne) SetCreatedAt(t time.Time) *ApprunUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *ApprunUpdateOne) SetNillableCreatedAt(t *time.Time) *ApprunUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (auo *ApprunUpdateOne) ClearCreatedAt() *ApprunUpdateOne {
	auo.mutation.ClearCreatedAt()
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *ApprunUpdateOne) SetUpdatedAt(t time.Time) *ApprunUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auo *ApprunUpdateOne) SetNillableUpdatedAt(t *time.Time) *ApprunUpdateOne {
	if t != nil {
		auo.SetUpdatedAt(*t)
	}
	return auo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (auo *ApprunUpdateOne) ClearUpdatedAt() *ApprunUpdateOne {
	auo.mutation.ClearUpdatedAt()
	return auo
}

// SetHandler sets the "handler" field.
func (auo *ApprunUpdateOne) SetHandler(s string) *ApprunUpdateOne {
	auo.mutation.SetHandler(s)
	return auo
}

// SetRuning sets the "runing" field.
func (auo *ApprunUpdateOne) SetRuning(u uint64) *ApprunUpdateOne {
	auo.mutation.ResetRuning()
	auo.mutation.SetRuning(u)
	return auo
}

// SetNillableRuning sets the "runing" field if the given value is not nil.
func (auo *ApprunUpdateOne) SetNillableRuning(u *uint64) *ApprunUpdateOne {
	if u != nil {
		auo.SetRuning(*u)
	}
	return auo
}

// AddRuning adds u to the "runing" field.
func (auo *ApprunUpdateOne) AddRuning(u int64) *ApprunUpdateOne {
	auo.mutation.AddRuning(u)
	return auo
}

// SetTotal sets the "total" field.
func (auo *ApprunUpdateOne) SetTotal(u uint64) *ApprunUpdateOne {
	auo.mutation.ResetTotal()
	auo.mutation.SetTotal(u)
	return auo
}

// SetNillableTotal sets the "total" field if the given value is not nil.
func (auo *ApprunUpdateOne) SetNillableTotal(u *uint64) *ApprunUpdateOne {
	if u != nil {
		auo.SetTotal(*u)
	}
	return auo
}

// AddTotal adds u to the "total" field.
func (auo *ApprunUpdateOne) AddTotal(u int64) *ApprunUpdateOne {
	auo.mutation.AddTotal(u)
	return auo
}

// Mutation returns the ApprunMutation object of the builder.
func (auo *ApprunUpdateOne) Mutation() *ApprunMutation {
	return auo.mutation
}

// Where appends a list predicates to the ApprunUpdate builder.
func (auo *ApprunUpdateOne) Where(ps ...predicate.Apprun) *ApprunUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ApprunUpdateOne) Select(field string, fields ...string) *ApprunUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Apprun entity.
func (auo *ApprunUpdateOne) Save(ctx context.Context) (*Apprun, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ApprunUpdateOne) SaveX(ctx context.Context) *Apprun {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ApprunUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ApprunUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ApprunUpdateOne) sqlSave(ctx context.Context) (_node *Apprun, err error) {
	_spec := sqlgraph.NewUpdateSpec(apprun.Table, apprun.Columns, sqlgraph.NewFieldSpec(apprun.FieldID, field.TypeUint64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Apprun.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apprun.FieldID)
		for _, f := range fields {
			if !apprun.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apprun.FieldID {
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
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(apprun.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.CreatedAtCleared() {
		_spec.ClearField(apprun.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(apprun.FieldUpdatedAt, field.TypeTime, value)
	}
	if auo.mutation.UpdatedAtCleared() {
		_spec.ClearField(apprun.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.Handler(); ok {
		_spec.SetField(apprun.FieldHandler, field.TypeString, value)
	}
	if value, ok := auo.mutation.Runing(); ok {
		_spec.SetField(apprun.FieldRuning, field.TypeUint64, value)
	}
	if value, ok := auo.mutation.AddedRuning(); ok {
		_spec.AddField(apprun.FieldRuning, field.TypeUint64, value)
	}
	if value, ok := auo.mutation.Total(); ok {
		_spec.SetField(apprun.FieldTotal, field.TypeUint64, value)
	}
	if value, ok := auo.mutation.AddedTotal(); ok {
		_spec.AddField(apprun.FieldTotal, field.TypeUint64, value)
	}
	_node = &Apprun{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apprun.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
