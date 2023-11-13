// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/skyisboss/pay-system/ent/addres"
)

// AddresCreate is the builder for creating a Addres entity.
type AddresCreate struct {
	config
	mutation *AddresMutation
	hooks    []Hook
}

// SetChainID sets the "chain_id" field.
func (ac *AddresCreate) SetChainID(u uint64) *AddresCreate {
	ac.mutation.SetChainID(u)
	return ac
}

// SetAddress sets the "address" field.
func (ac *AddresCreate) SetAddress(s string) *AddresCreate {
	ac.mutation.SetAddress(s)
	return ac
}

// SetPassword sets the "password" field.
func (ac *AddresCreate) SetPassword(s string) *AddresCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetUUID sets the "uuid" field.
func (ac *AddresCreate) SetUUID(s string) *AddresCreate {
	ac.mutation.SetUUID(s)
	return ac
}

// SetUseTo sets the "use_to" field.
func (ac *AddresCreate) SetUseTo(i int64) *AddresCreate {
	ac.mutation.SetUseTo(i)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AddresCreate) SetCreatedAt(t time.Time) *AddresCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AddresCreate) SetNillableCreatedAt(t *time.Time) *AddresCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AddresCreate) SetUpdatedAt(t time.Time) *AddresCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AddresCreate) SetNillableUpdatedAt(t *time.Time) *AddresCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AddresCreate) SetDeletedAt(t time.Time) *AddresCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AddresCreate) SetNillableDeletedAt(t *time.Time) *AddresCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AddresCreate) SetID(u uint64) *AddresCreate {
	ac.mutation.SetID(u)
	return ac
}

// Mutation returns the AddresMutation object of the builder.
func (ac *AddresCreate) Mutation() *AddresMutation {
	return ac.mutation
}

// Save creates the Addres in the database.
func (ac *AddresCreate) Save(ctx context.Context) (*Addres, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AddresCreate) SaveX(ctx context.Context) *Addres {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AddresCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AddresCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AddresCreate) check() error {
	if _, ok := ac.mutation.ChainID(); !ok {
		return &ValidationError{Name: "chain_id", err: errors.New(`ent: missing required field "Addres.chain_id"`)}
	}
	if _, ok := ac.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Addres.address"`)}
	}
	if _, ok := ac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Addres.password"`)}
	}
	if _, ok := ac.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "Addres.uuid"`)}
	}
	if _, ok := ac.mutation.UseTo(); !ok {
		return &ValidationError{Name: "use_to", err: errors.New(`ent: missing required field "Addres.use_to"`)}
	}
	return nil
}

func (ac *AddresCreate) sqlSave(ctx context.Context) (*Addres, error) {
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

func (ac *AddresCreate) createSpec() (*Addres, *sqlgraph.CreateSpec) {
	var (
		_node = &Addres{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(addres.Table, sqlgraph.NewFieldSpec(addres.FieldID, field.TypeUint64))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.ChainID(); ok {
		_spec.SetField(addres.FieldChainID, field.TypeUint64, value)
		_node.ChainID = value
	}
	if value, ok := ac.mutation.Address(); ok {
		_spec.SetField(addres.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.SetField(addres.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := ac.mutation.UUID(); ok {
		_spec.SetField(addres.FieldUUID, field.TypeString, value)
		_node.UUID = value
	}
	if value, ok := ac.mutation.UseTo(); ok {
		_spec.SetField(addres.FieldUseTo, field.TypeInt64, value)
		_node.UseTo = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(addres.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(addres.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(addres.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	return _node, _spec
}

// AddresCreateBulk is the builder for creating many Addres entities in bulk.
type AddresCreateBulk struct {
	config
	err      error
	builders []*AddresCreate
}

// Save creates the Addres entities in the database.
func (acb *AddresCreateBulk) Save(ctx context.Context) ([]*Addres, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Addres, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AddresMutation)
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
func (acb *AddresCreateBulk) SaveX(ctx context.Context) []*Addres {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AddresCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AddresCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
