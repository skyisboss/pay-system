// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent/withdraw"
)

// WithdrawCreate is the builder for creating a Withdraw entity.
type WithdrawCreate struct {
	config
	mutation *WithdrawMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (wc *WithdrawCreate) SetCreatedAt(t time.Time) *WithdrawCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WithdrawCreate) SetNillableCreatedAt(t *time.Time) *WithdrawCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetUpdatedAt sets the "updated_at" field.
func (wc *WithdrawCreate) SetUpdatedAt(t time.Time) *WithdrawCreate {
	wc.mutation.SetUpdatedAt(t)
	return wc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wc *WithdrawCreate) SetNillableUpdatedAt(t *time.Time) *WithdrawCreate {
	if t != nil {
		wc.SetUpdatedAt(*t)
	}
	return wc
}

// SetDeletedAt sets the "deleted_at" field.
func (wc *WithdrawCreate) SetDeletedAt(t time.Time) *WithdrawCreate {
	wc.mutation.SetDeletedAt(t)
	return wc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wc *WithdrawCreate) SetNillableDeletedAt(t *time.Time) *WithdrawCreate {
	if t != nil {
		wc.SetDeletedAt(*t)
	}
	return wc
}

// SetProductID sets the "product_id" field.
func (wc *WithdrawCreate) SetProductID(i int64) *WithdrawCreate {
	wc.mutation.SetProductID(i)
	return wc
}

// SetChainID sets the "chain_id" field.
func (wc *WithdrawCreate) SetChainID(u uint64) *WithdrawCreate {
	wc.mutation.SetChainID(u)
	return wc
}

// SetToAddress sets the "to_address" field.
func (wc *WithdrawCreate) SetToAddress(s string) *WithdrawCreate {
	wc.mutation.SetToAddress(s)
	return wc
}

// SetAmountStr sets the "amount_str" field.
func (wc *WithdrawCreate) SetAmountStr(s string) *WithdrawCreate {
	wc.mutation.SetAmountStr(s)
	return wc
}

// SetAmountRaw sets the "amount_raw" field.
func (wc *WithdrawCreate) SetAmountRaw(d decimal.Decimal) *WithdrawCreate {
	wc.mutation.SetAmountRaw(d)
	return wc
}

// SetSerialID sets the "serial_id" field.
func (wc *WithdrawCreate) SetSerialID(s string) *WithdrawCreate {
	wc.mutation.SetSerialID(s)
	return wc
}

// SetTxHash sets the "tx_hash" field.
func (wc *WithdrawCreate) SetTxHash(s string) *WithdrawCreate {
	wc.mutation.SetTxHash(s)
	return wc
}

// SetNillableTxHash sets the "tx_hash" field if the given value is not nil.
func (wc *WithdrawCreate) SetNillableTxHash(s *string) *WithdrawCreate {
	if s != nil {
		wc.SetTxHash(*s)
	}
	return wc
}

// SetHandleStatus sets the "handle_status" field.
func (wc *WithdrawCreate) SetHandleStatus(i int64) *WithdrawCreate {
	wc.mutation.SetHandleStatus(i)
	return wc
}

// SetNillableHandleStatus sets the "handle_status" field if the given value is not nil.
func (wc *WithdrawCreate) SetNillableHandleStatus(i *int64) *WithdrawCreate {
	if i != nil {
		wc.SetHandleStatus(*i)
	}
	return wc
}

// SetHandleMsg sets the "handle_msg" field.
func (wc *WithdrawCreate) SetHandleMsg(s string) *WithdrawCreate {
	wc.mutation.SetHandleMsg(s)
	return wc
}

// SetNillableHandleMsg sets the "handle_msg" field if the given value is not nil.
func (wc *WithdrawCreate) SetNillableHandleMsg(s *string) *WithdrawCreate {
	if s != nil {
		wc.SetHandleMsg(*s)
	}
	return wc
}

// SetHandleTime sets the "handle_time" field.
func (wc *WithdrawCreate) SetHandleTime(t time.Time) *WithdrawCreate {
	wc.mutation.SetHandleTime(t)
	return wc
}

// SetNillableHandleTime sets the "handle_time" field if the given value is not nil.
func (wc *WithdrawCreate) SetNillableHandleTime(t *time.Time) *WithdrawCreate {
	if t != nil {
		wc.SetHandleTime(*t)
	}
	return wc
}

// SetID sets the "id" field.
func (wc *WithdrawCreate) SetID(u uint64) *WithdrawCreate {
	wc.mutation.SetID(u)
	return wc
}

// Mutation returns the WithdrawMutation object of the builder.
func (wc *WithdrawCreate) Mutation() *WithdrawMutation {
	return wc.mutation
}

// Save creates the Withdraw in the database.
func (wc *WithdrawCreate) Save(ctx context.Context) (*Withdraw, error) {
	wc.defaults()
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WithdrawCreate) SaveX(ctx context.Context) *Withdraw {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WithdrawCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WithdrawCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WithdrawCreate) defaults() {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := withdraw.DefaultCreatedAt
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.HandleStatus(); !ok {
		v := withdraw.DefaultHandleStatus
		wc.mutation.SetHandleStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WithdrawCreate) check() error {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Withdraw.created_at"`)}
	}
	if _, ok := wc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product_id", err: errors.New(`ent: missing required field "Withdraw.product_id"`)}
	}
	if _, ok := wc.mutation.ChainID(); !ok {
		return &ValidationError{Name: "chain_id", err: errors.New(`ent: missing required field "Withdraw.chain_id"`)}
	}
	if _, ok := wc.mutation.ToAddress(); !ok {
		return &ValidationError{Name: "to_address", err: errors.New(`ent: missing required field "Withdraw.to_address"`)}
	}
	if _, ok := wc.mutation.AmountStr(); !ok {
		return &ValidationError{Name: "amount_str", err: errors.New(`ent: missing required field "Withdraw.amount_str"`)}
	}
	if _, ok := wc.mutation.AmountRaw(); !ok {
		return &ValidationError{Name: "amount_raw", err: errors.New(`ent: missing required field "Withdraw.amount_raw"`)}
	}
	if _, ok := wc.mutation.SerialID(); !ok {
		return &ValidationError{Name: "serial_id", err: errors.New(`ent: missing required field "Withdraw.serial_id"`)}
	}
	if _, ok := wc.mutation.HandleStatus(); !ok {
		return &ValidationError{Name: "handle_status", err: errors.New(`ent: missing required field "Withdraw.handle_status"`)}
	}
	return nil
}

func (wc *WithdrawCreate) sqlSave(ctx context.Context) (*Withdraw, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WithdrawCreate) createSpec() (*Withdraw, *sqlgraph.CreateSpec) {
	var (
		_node = &Withdraw{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(withdraw.Table, sqlgraph.NewFieldSpec(withdraw.FieldID, field.TypeUint64))
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.SetField(withdraw.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.UpdatedAt(); ok {
		_spec.SetField(withdraw.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := wc.mutation.DeletedAt(); ok {
		_spec.SetField(withdraw.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := wc.mutation.ProductID(); ok {
		_spec.SetField(withdraw.FieldProductID, field.TypeInt64, value)
		_node.ProductID = value
	}
	if value, ok := wc.mutation.ChainID(); ok {
		_spec.SetField(withdraw.FieldChainID, field.TypeUint64, value)
		_node.ChainID = value
	}
	if value, ok := wc.mutation.ToAddress(); ok {
		_spec.SetField(withdraw.FieldToAddress, field.TypeString, value)
		_node.ToAddress = value
	}
	if value, ok := wc.mutation.AmountStr(); ok {
		_spec.SetField(withdraw.FieldAmountStr, field.TypeString, value)
		_node.AmountStr = value
	}
	if value, ok := wc.mutation.AmountRaw(); ok {
		_spec.SetField(withdraw.FieldAmountRaw, field.TypeFloat64, value)
		_node.AmountRaw = value
	}
	if value, ok := wc.mutation.SerialID(); ok {
		_spec.SetField(withdraw.FieldSerialID, field.TypeString, value)
		_node.SerialID = value
	}
	if value, ok := wc.mutation.TxHash(); ok {
		_spec.SetField(withdraw.FieldTxHash, field.TypeString, value)
		_node.TxHash = value
	}
	if value, ok := wc.mutation.HandleStatus(); ok {
		_spec.SetField(withdraw.FieldHandleStatus, field.TypeInt64, value)
		_node.HandleStatus = value
	}
	if value, ok := wc.mutation.HandleMsg(); ok {
		_spec.SetField(withdraw.FieldHandleMsg, field.TypeString, value)
		_node.HandleMsg = value
	}
	if value, ok := wc.mutation.HandleTime(); ok {
		_spec.SetField(withdraw.FieldHandleTime, field.TypeTime, value)
		_node.HandleTime = value
	}
	return _node, _spec
}

// WithdrawCreateBulk is the builder for creating many Withdraw entities in bulk.
type WithdrawCreateBulk struct {
	config
	err      error
	builders []*WithdrawCreate
}

// Save creates the Withdraw entities in the database.
func (wcb *WithdrawCreateBulk) Save(ctx context.Context) ([]*Withdraw, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Withdraw, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WithdrawMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WithdrawCreateBulk) SaveX(ctx context.Context) []*Withdraw {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WithdrawCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WithdrawCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
