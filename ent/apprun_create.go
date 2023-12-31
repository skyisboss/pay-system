// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/skyisboss/pay-system/ent/apprun"
)

// ApprunCreate is the builder for creating a Apprun entity.
type ApprunCreate struct {
	config
	mutation *ApprunMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *ApprunCreate) SetCreatedAt(t time.Time) *ApprunCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *ApprunCreate) SetNillableCreatedAt(t *time.Time) *ApprunCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *ApprunCreate) SetUpdatedAt(t time.Time) *ApprunCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *ApprunCreate) SetNillableUpdatedAt(t *time.Time) *ApprunCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetHandler sets the "handler" field.
func (ac *ApprunCreate) SetHandler(s string) *ApprunCreate {
	ac.mutation.SetHandler(s)
	return ac
}

// SetRuning sets the "runing" field.
func (ac *ApprunCreate) SetRuning(u uint64) *ApprunCreate {
	ac.mutation.SetRuning(u)
	return ac
}

// SetNillableRuning sets the "runing" field if the given value is not nil.
func (ac *ApprunCreate) SetNillableRuning(u *uint64) *ApprunCreate {
	if u != nil {
		ac.SetRuning(*u)
	}
	return ac
}

// SetTotal sets the "total" field.
func (ac *ApprunCreate) SetTotal(u uint64) *ApprunCreate {
	ac.mutation.SetTotal(u)
	return ac
}

// SetNillableTotal sets the "total" field if the given value is not nil.
func (ac *ApprunCreate) SetNillableTotal(u *uint64) *ApprunCreate {
	if u != nil {
		ac.SetTotal(*u)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *ApprunCreate) SetID(u uint64) *ApprunCreate {
	ac.mutation.SetID(u)
	return ac
}

// Mutation returns the ApprunMutation object of the builder.
func (ac *ApprunCreate) Mutation() *ApprunMutation {
	return ac.mutation
}

// Save creates the Apprun in the database.
func (ac *ApprunCreate) Save(ctx context.Context) (*Apprun, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ApprunCreate) SaveX(ctx context.Context) *Apprun {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ApprunCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ApprunCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ApprunCreate) defaults() {
	if _, ok := ac.mutation.Runing(); !ok {
		v := apprun.DefaultRuning
		ac.mutation.SetRuning(v)
	}
	if _, ok := ac.mutation.Total(); !ok {
		v := apprun.DefaultTotal
		ac.mutation.SetTotal(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ApprunCreate) check() error {
	if _, ok := ac.mutation.Handler(); !ok {
		return &ValidationError{Name: "handler", err: errors.New(`ent: missing required field "Apprun.handler"`)}
	}
	if _, ok := ac.mutation.Runing(); !ok {
		return &ValidationError{Name: "runing", err: errors.New(`ent: missing required field "Apprun.runing"`)}
	}
	if _, ok := ac.mutation.Total(); !ok {
		return &ValidationError{Name: "total", err: errors.New(`ent: missing required field "Apprun.total"`)}
	}
	return nil
}

func (ac *ApprunCreate) sqlSave(ctx context.Context) (*Apprun, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ApprunCreate) createSpec() (*Apprun, *sqlgraph.CreateSpec) {
	var (
		_node = &Apprun{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(apprun.Table, sqlgraph.NewFieldSpec(apprun.FieldID, field.TypeUint64))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(apprun.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(apprun.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.Handler(); ok {
		_spec.SetField(apprun.FieldHandler, field.TypeString, value)
		_node.Handler = value
	}
	if value, ok := ac.mutation.Runing(); ok {
		_spec.SetField(apprun.FieldRuning, field.TypeUint64, value)
		_node.Runing = value
	}
	if value, ok := ac.mutation.Total(); ok {
		_spec.SetField(apprun.FieldTotal, field.TypeUint64, value)
		_node.Total = value
	}
	return _node, _spec
}

// ApprunCreateBulk is the builder for creating many Apprun entities in bulk.
type ApprunCreateBulk struct {
	config
	err      error
	builders []*ApprunCreate
}

// Save creates the Apprun entities in the database.
func (acb *ApprunCreateBulk) Save(ctx context.Context) ([]*Apprun, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Apprun, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApprunMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ApprunCreateBulk) SaveX(ctx context.Context) []*Apprun {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ApprunCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ApprunCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
