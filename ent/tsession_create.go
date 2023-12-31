// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/skyisboss/pay-system/ent/tsession"
)

// TSessionCreate is the builder for creating a TSession entity.
type TSessionCreate struct {
	config
	mutation *TSessionMutation
	hooks    []Hook
}

// SetKeyName sets the "key_name" field.
func (tc *TSessionCreate) SetKeyName(s string) *TSessionCreate {
	tc.mutation.SetKeyName(s)
	return tc
}

// SetKeyValue sets the "key_value" field.
func (tc *TSessionCreate) SetKeyValue(s string) *TSessionCreate {
	tc.mutation.SetKeyValue(s)
	return tc
}

// SetIP sets the "ip" field.
func (tc *TSessionCreate) SetIP(s string) *TSessionCreate {
	tc.mutation.SetIP(s)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TSessionCreate) SetCreatedAt(t time.Time) *TSessionCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TSessionCreate) SetNillableCreatedAt(t *time.Time) *TSessionCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TSessionCreate) SetUpdatedAt(t time.Time) *TSessionCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TSessionCreate) SetNillableUpdatedAt(t *time.Time) *TSessionCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TSessionCreate) SetID(u uint64) *TSessionCreate {
	tc.mutation.SetID(u)
	return tc
}

// Mutation returns the TSessionMutation object of the builder.
func (tc *TSessionCreate) Mutation() *TSessionMutation {
	return tc.mutation
}

// Save creates the TSession in the database.
func (tc *TSessionCreate) Save(ctx context.Context) (*TSession, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TSessionCreate) SaveX(ctx context.Context) *TSession {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TSessionCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TSessionCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TSessionCreate) check() error {
	if _, ok := tc.mutation.KeyName(); !ok {
		return &ValidationError{Name: "key_name", err: errors.New(`ent: missing required field "TSession.key_name"`)}
	}
	if _, ok := tc.mutation.KeyValue(); !ok {
		return &ValidationError{Name: "key_value", err: errors.New(`ent: missing required field "TSession.key_value"`)}
	}
	if _, ok := tc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`ent: missing required field "TSession.ip"`)}
	}
	return nil
}

func (tc *TSessionCreate) sqlSave(ctx context.Context) (*TSession, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TSessionCreate) createSpec() (*TSession, *sqlgraph.CreateSpec) {
	var (
		_node = &TSession{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tsession.Table, sqlgraph.NewFieldSpec(tsession.FieldID, field.TypeUint64))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.KeyName(); ok {
		_spec.SetField(tsession.FieldKeyName, field.TypeString, value)
		_node.KeyName = value
	}
	if value, ok := tc.mutation.KeyValue(); ok {
		_spec.SetField(tsession.FieldKeyValue, field.TypeString, value)
		_node.KeyValue = value
	}
	if value, ok := tc.mutation.IP(); ok {
		_spec.SetField(tsession.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(tsession.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(tsession.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// TSessionCreateBulk is the builder for creating many TSession entities in bulk.
type TSessionCreateBulk struct {
	config
	err      error
	builders []*TSessionCreate
}

// Save creates the TSession entities in the database.
func (tcb *TSessionCreateBulk) Save(ctx context.Context) ([]*TSession, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*TSession, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TSessionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TSessionCreateBulk) SaveX(ctx context.Context) []*TSession {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TSessionCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TSessionCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
