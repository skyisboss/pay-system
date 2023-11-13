// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/skyisboss/pay-system/ent/notify"
)

// NotifyCreate is the builder for creating a Notify entity.
type NotifyCreate struct {
	config
	mutation *NotifyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (nc *NotifyCreate) SetCreatedAt(t time.Time) *NotifyCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NotifyCreate) SetNillableCreatedAt(t *time.Time) *NotifyCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the "updated_at" field.
func (nc *NotifyCreate) SetUpdatedAt(t time.Time) *NotifyCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nc *NotifyCreate) SetNillableUpdatedAt(t *time.Time) *NotifyCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetDeletedAt sets the "deleted_at" field.
func (nc *NotifyCreate) SetDeletedAt(t time.Time) *NotifyCreate {
	nc.mutation.SetDeletedAt(t)
	return nc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nc *NotifyCreate) SetNillableDeletedAt(t *time.Time) *NotifyCreate {
	if t != nil {
		nc.SetDeletedAt(*t)
	}
	return nc
}

// SetChainID sets the "chain_id" field.
func (nc *NotifyCreate) SetChainID(u uint64) *NotifyCreate {
	nc.mutation.SetChainID(u)
	return nc
}

// SetProductID sets the "product_id" field.
func (nc *NotifyCreate) SetProductID(u uint64) *NotifyCreate {
	nc.mutation.SetProductID(u)
	return nc
}

// SetItemFrom sets the "item_from" field.
func (nc *NotifyCreate) SetItemFrom(u uint64) *NotifyCreate {
	nc.mutation.SetItemFrom(u)
	return nc
}

// SetItemType sets the "item_type" field.
func (nc *NotifyCreate) SetItemType(i int64) *NotifyCreate {
	nc.mutation.SetItemType(i)
	return nc
}

// SetNonce sets the "nonce" field.
func (nc *NotifyCreate) SetNonce(s string) *NotifyCreate {
	nc.mutation.SetNonce(s)
	return nc
}

// SetNotifyType sets the "notify_type" field.
func (nc *NotifyCreate) SetNotifyType(s string) *NotifyCreate {
	nc.mutation.SetNotifyType(s)
	return nc
}

// SetSendURL sets the "send_url" field.
func (nc *NotifyCreate) SetSendURL(s string) *NotifyCreate {
	nc.mutation.SetSendURL(s)
	return nc
}

// SetSendBody sets the "send_body" field.
func (nc *NotifyCreate) SetSendBody(s string) *NotifyCreate {
	nc.mutation.SetSendBody(s)
	return nc
}

// SetSendRetry sets the "send_retry" field.
func (nc *NotifyCreate) SetSendRetry(i int64) *NotifyCreate {
	nc.mutation.SetSendRetry(i)
	return nc
}

// SetHandleStatus sets the "handle_status" field.
func (nc *NotifyCreate) SetHandleStatus(i int64) *NotifyCreate {
	nc.mutation.SetHandleStatus(i)
	return nc
}

// SetNillableHandleStatus sets the "handle_status" field if the given value is not nil.
func (nc *NotifyCreate) SetNillableHandleStatus(i *int64) *NotifyCreate {
	if i != nil {
		nc.SetHandleStatus(*i)
	}
	return nc
}

// SetHandleMsg sets the "handle_msg" field.
func (nc *NotifyCreate) SetHandleMsg(s string) *NotifyCreate {
	nc.mutation.SetHandleMsg(s)
	return nc
}

// SetNillableHandleMsg sets the "handle_msg" field if the given value is not nil.
func (nc *NotifyCreate) SetNillableHandleMsg(s *string) *NotifyCreate {
	if s != nil {
		nc.SetHandleMsg(*s)
	}
	return nc
}

// SetHandleTime sets the "handle_time" field.
func (nc *NotifyCreate) SetHandleTime(t time.Time) *NotifyCreate {
	nc.mutation.SetHandleTime(t)
	return nc
}

// SetNillableHandleTime sets the "handle_time" field if the given value is not nil.
func (nc *NotifyCreate) SetNillableHandleTime(t *time.Time) *NotifyCreate {
	if t != nil {
		nc.SetHandleTime(*t)
	}
	return nc
}

// SetID sets the "id" field.
func (nc *NotifyCreate) SetID(u uint64) *NotifyCreate {
	nc.mutation.SetID(u)
	return nc
}

// Mutation returns the NotifyMutation object of the builder.
func (nc *NotifyCreate) Mutation() *NotifyMutation {
	return nc.mutation
}

// Save creates the Notify in the database.
func (nc *NotifyCreate) Save(ctx context.Context) (*Notify, error) {
	nc.defaults()
	return withHooks(ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NotifyCreate) SaveX(ctx context.Context) *Notify {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NotifyCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NotifyCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NotifyCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := notify.DefaultCreatedAt
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.HandleStatus(); !ok {
		v := notify.DefaultHandleStatus
		nc.mutation.SetHandleStatus(v)
	}
	if _, ok := nc.mutation.HandleMsg(); !ok {
		v := notify.DefaultHandleMsg
		nc.mutation.SetHandleMsg(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NotifyCreate) check() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Notify.created_at"`)}
	}
	if _, ok := nc.mutation.ChainID(); !ok {
		return &ValidationError{Name: "chain_id", err: errors.New(`ent: missing required field "Notify.chain_id"`)}
	}
	if _, ok := nc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product_id", err: errors.New(`ent: missing required field "Notify.product_id"`)}
	}
	if _, ok := nc.mutation.ItemFrom(); !ok {
		return &ValidationError{Name: "item_from", err: errors.New(`ent: missing required field "Notify.item_from"`)}
	}
	if _, ok := nc.mutation.ItemType(); !ok {
		return &ValidationError{Name: "item_type", err: errors.New(`ent: missing required field "Notify.item_type"`)}
	}
	if _, ok := nc.mutation.Nonce(); !ok {
		return &ValidationError{Name: "nonce", err: errors.New(`ent: missing required field "Notify.nonce"`)}
	}
	if _, ok := nc.mutation.NotifyType(); !ok {
		return &ValidationError{Name: "notify_type", err: errors.New(`ent: missing required field "Notify.notify_type"`)}
	}
	if _, ok := nc.mutation.SendURL(); !ok {
		return &ValidationError{Name: "send_url", err: errors.New(`ent: missing required field "Notify.send_url"`)}
	}
	if _, ok := nc.mutation.SendBody(); !ok {
		return &ValidationError{Name: "send_body", err: errors.New(`ent: missing required field "Notify.send_body"`)}
	}
	if _, ok := nc.mutation.SendRetry(); !ok {
		return &ValidationError{Name: "send_retry", err: errors.New(`ent: missing required field "Notify.send_retry"`)}
	}
	if _, ok := nc.mutation.HandleStatus(); !ok {
		return &ValidationError{Name: "handle_status", err: errors.New(`ent: missing required field "Notify.handle_status"`)}
	}
	if _, ok := nc.mutation.HandleMsg(); !ok {
		return &ValidationError{Name: "handle_msg", err: errors.New(`ent: missing required field "Notify.handle_msg"`)}
	}
	return nil
}

func (nc *NotifyCreate) sqlSave(ctx context.Context) (*Notify, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NotifyCreate) createSpec() (*Notify, *sqlgraph.CreateSpec) {
	var (
		_node = &Notify{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(notify.Table, sqlgraph.NewFieldSpec(notify.FieldID, field.TypeUint64))
	)
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.SetField(notify.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.SetField(notify.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nc.mutation.DeletedAt(); ok {
		_spec.SetField(notify.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := nc.mutation.ChainID(); ok {
		_spec.SetField(notify.FieldChainID, field.TypeUint64, value)
		_node.ChainID = value
	}
	if value, ok := nc.mutation.ProductID(); ok {
		_spec.SetField(notify.FieldProductID, field.TypeUint64, value)
		_node.ProductID = value
	}
	if value, ok := nc.mutation.ItemFrom(); ok {
		_spec.SetField(notify.FieldItemFrom, field.TypeUint64, value)
		_node.ItemFrom = value
	}
	if value, ok := nc.mutation.ItemType(); ok {
		_spec.SetField(notify.FieldItemType, field.TypeInt64, value)
		_node.ItemType = value
	}
	if value, ok := nc.mutation.Nonce(); ok {
		_spec.SetField(notify.FieldNonce, field.TypeString, value)
		_node.Nonce = value
	}
	if value, ok := nc.mutation.NotifyType(); ok {
		_spec.SetField(notify.FieldNotifyType, field.TypeString, value)
		_node.NotifyType = value
	}
	if value, ok := nc.mutation.SendURL(); ok {
		_spec.SetField(notify.FieldSendURL, field.TypeString, value)
		_node.SendURL = value
	}
	if value, ok := nc.mutation.SendBody(); ok {
		_spec.SetField(notify.FieldSendBody, field.TypeString, value)
		_node.SendBody = value
	}
	if value, ok := nc.mutation.SendRetry(); ok {
		_spec.SetField(notify.FieldSendRetry, field.TypeInt64, value)
		_node.SendRetry = value
	}
	if value, ok := nc.mutation.HandleStatus(); ok {
		_spec.SetField(notify.FieldHandleStatus, field.TypeInt64, value)
		_node.HandleStatus = value
	}
	if value, ok := nc.mutation.HandleMsg(); ok {
		_spec.SetField(notify.FieldHandleMsg, field.TypeString, value)
		_node.HandleMsg = value
	}
	if value, ok := nc.mutation.HandleTime(); ok {
		_spec.SetField(notify.FieldHandleTime, field.TypeTime, value)
		_node.HandleTime = value
	}
	return _node, _spec
}

// NotifyCreateBulk is the builder for creating many Notify entities in bulk.
type NotifyCreateBulk struct {
	config
	err      error
	builders []*NotifyCreate
}

// Save creates the Notify entities in the database.
func (ncb *NotifyCreateBulk) Save(ctx context.Context) ([]*Notify, error) {
	if ncb.err != nil {
		return nil, ncb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Notify, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotifyMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NotifyCreateBulk) SaveX(ctx context.Context) []*Notify {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NotifyCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NotifyCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}