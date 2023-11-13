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
	"github.com/skyisboss/pay-system/ent/txn"
)

// TxnCreate is the builder for creating a Txn entity.
type TxnCreate struct {
	config
	mutation *TxnMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TxnCreate) SetCreatedAt(t time.Time) *TxnCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TxnCreate) SetNillableCreatedAt(t *time.Time) *TxnCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TxnCreate) SetUpdatedAt(t time.Time) *TxnCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TxnCreate) SetNillableUpdatedAt(t *time.Time) *TxnCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetDeletedAt sets the "deleted_at" field.
func (tc *TxnCreate) SetDeletedAt(t time.Time) *TxnCreate {
	tc.mutation.SetDeletedAt(t)
	return tc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tc *TxnCreate) SetNillableDeletedAt(t *time.Time) *TxnCreate {
	if t != nil {
		tc.SetDeletedAt(*t)
	}
	return tc
}

// SetTxID sets the "tx_id" field.
func (tc *TxnCreate) SetTxID(s string) *TxnCreate {
	tc.mutation.SetTxID(s)
	return tc
}

// SetChainID sets the "chain_id" field.
func (tc *TxnCreate) SetChainID(u uint64) *TxnCreate {
	tc.mutation.SetChainID(u)
	return tc
}

// SetProductID sets the "product_id" field.
func (tc *TxnCreate) SetProductID(i int64) *TxnCreate {
	tc.mutation.SetProductID(i)
	return tc
}

// SetFromAddress sets the "from_address" field.
func (tc *TxnCreate) SetFromAddress(s string) *TxnCreate {
	tc.mutation.SetFromAddress(s)
	return tc
}

// SetToAddress sets the "to_address" field.
func (tc *TxnCreate) SetToAddress(s string) *TxnCreate {
	tc.mutation.SetToAddress(s)
	return tc
}

// SetAmountStr sets the "amount_str" field.
func (tc *TxnCreate) SetAmountStr(s string) *TxnCreate {
	tc.mutation.SetAmountStr(s)
	return tc
}

// SetAmountRaw sets the "amount_raw" field.
func (tc *TxnCreate) SetAmountRaw(d decimal.Decimal) *TxnCreate {
	tc.mutation.SetAmountRaw(d)
	return tc
}

// SetHandleStatus sets the "handle_status" field.
func (tc *TxnCreate) SetHandleStatus(i int64) *TxnCreate {
	tc.mutation.SetHandleStatus(i)
	return tc
}

// SetNillableHandleStatus sets the "handle_status" field if the given value is not nil.
func (tc *TxnCreate) SetNillableHandleStatus(i *int64) *TxnCreate {
	if i != nil {
		tc.SetHandleStatus(*i)
	}
	return tc
}

// SetHandleMsg sets the "handle_msg" field.
func (tc *TxnCreate) SetHandleMsg(s string) *TxnCreate {
	tc.mutation.SetHandleMsg(s)
	return tc
}

// SetNillableHandleMsg sets the "handle_msg" field if the given value is not nil.
func (tc *TxnCreate) SetNillableHandleMsg(s *string) *TxnCreate {
	if s != nil {
		tc.SetHandleMsg(*s)
	}
	return tc
}

// SetHandleTime sets the "handle_time" field.
func (tc *TxnCreate) SetHandleTime(t time.Time) *TxnCreate {
	tc.mutation.SetHandleTime(t)
	return tc
}

// SetNillableHandleTime sets the "handle_time" field if the given value is not nil.
func (tc *TxnCreate) SetNillableHandleTime(t *time.Time) *TxnCreate {
	if t != nil {
		tc.SetHandleTime(*t)
	}
	return tc
}

// SetCollectStatus sets the "collect_status" field.
func (tc *TxnCreate) SetCollectStatus(i int64) *TxnCreate {
	tc.mutation.SetCollectStatus(i)
	return tc
}

// SetNillableCollectStatus sets the "collect_status" field if the given value is not nil.
func (tc *TxnCreate) SetNillableCollectStatus(i *int64) *TxnCreate {
	if i != nil {
		tc.SetCollectStatus(*i)
	}
	return tc
}

// SetCollectMsg sets the "collect_msg" field.
func (tc *TxnCreate) SetCollectMsg(s string) *TxnCreate {
	tc.mutation.SetCollectMsg(s)
	return tc
}

// SetNillableCollectMsg sets the "collect_msg" field if the given value is not nil.
func (tc *TxnCreate) SetNillableCollectMsg(s *string) *TxnCreate {
	if s != nil {
		tc.SetCollectMsg(*s)
	}
	return tc
}

// SetCollectTime sets the "collect_time" field.
func (tc *TxnCreate) SetCollectTime(t time.Time) *TxnCreate {
	tc.mutation.SetCollectTime(t)
	return tc
}

// SetNillableCollectTime sets the "collect_time" field if the given value is not nil.
func (tc *TxnCreate) SetNillableCollectTime(t *time.Time) *TxnCreate {
	if t != nil {
		tc.SetCollectTime(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TxnCreate) SetID(u uint64) *TxnCreate {
	tc.mutation.SetID(u)
	return tc
}

// Mutation returns the TxnMutation object of the builder.
func (tc *TxnCreate) Mutation() *TxnMutation {
	return tc.mutation
}

// Save creates the Txn in the database.
func (tc *TxnCreate) Save(ctx context.Context) (*Txn, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TxnCreate) SaveX(ctx context.Context) *Txn {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TxnCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TxnCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TxnCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := txn.DefaultCreatedAt
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.HandleStatus(); !ok {
		v := txn.DefaultHandleStatus
		tc.mutation.SetHandleStatus(v)
	}
	if _, ok := tc.mutation.CollectStatus(); !ok {
		v := txn.DefaultCollectStatus
		tc.mutation.SetCollectStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TxnCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Txn.created_at"`)}
	}
	if _, ok := tc.mutation.TxID(); !ok {
		return &ValidationError{Name: "tx_id", err: errors.New(`ent: missing required field "Txn.tx_id"`)}
	}
	if _, ok := tc.mutation.ChainID(); !ok {
		return &ValidationError{Name: "chain_id", err: errors.New(`ent: missing required field "Txn.chain_id"`)}
	}
	if _, ok := tc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product_id", err: errors.New(`ent: missing required field "Txn.product_id"`)}
	}
	if _, ok := tc.mutation.FromAddress(); !ok {
		return &ValidationError{Name: "from_address", err: errors.New(`ent: missing required field "Txn.from_address"`)}
	}
	if _, ok := tc.mutation.ToAddress(); !ok {
		return &ValidationError{Name: "to_address", err: errors.New(`ent: missing required field "Txn.to_address"`)}
	}
	if _, ok := tc.mutation.AmountStr(); !ok {
		return &ValidationError{Name: "amount_str", err: errors.New(`ent: missing required field "Txn.amount_str"`)}
	}
	if _, ok := tc.mutation.AmountRaw(); !ok {
		return &ValidationError{Name: "amount_raw", err: errors.New(`ent: missing required field "Txn.amount_raw"`)}
	}
	if _, ok := tc.mutation.HandleStatus(); !ok {
		return &ValidationError{Name: "handle_status", err: errors.New(`ent: missing required field "Txn.handle_status"`)}
	}
	if _, ok := tc.mutation.CollectStatus(); !ok {
		return &ValidationError{Name: "collect_status", err: errors.New(`ent: missing required field "Txn.collect_status"`)}
	}
	return nil
}

func (tc *TxnCreate) sqlSave(ctx context.Context) (*Txn, error) {
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

func (tc *TxnCreate) createSpec() (*Txn, *sqlgraph.CreateSpec) {
	var (
		_node = &Txn{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(txn.Table, sqlgraph.NewFieldSpec(txn.FieldID, field.TypeUint64))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(txn.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(txn.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.DeletedAt(); ok {
		_spec.SetField(txn.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := tc.mutation.TxID(); ok {
		_spec.SetField(txn.FieldTxID, field.TypeString, value)
		_node.TxID = value
	}
	if value, ok := tc.mutation.ChainID(); ok {
		_spec.SetField(txn.FieldChainID, field.TypeUint64, value)
		_node.ChainID = value
	}
	if value, ok := tc.mutation.ProductID(); ok {
		_spec.SetField(txn.FieldProductID, field.TypeInt64, value)
		_node.ProductID = value
	}
	if value, ok := tc.mutation.FromAddress(); ok {
		_spec.SetField(txn.FieldFromAddress, field.TypeString, value)
		_node.FromAddress = value
	}
	if value, ok := tc.mutation.ToAddress(); ok {
		_spec.SetField(txn.FieldToAddress, field.TypeString, value)
		_node.ToAddress = value
	}
	if value, ok := tc.mutation.AmountStr(); ok {
		_spec.SetField(txn.FieldAmountStr, field.TypeString, value)
		_node.AmountStr = value
	}
	if value, ok := tc.mutation.AmountRaw(); ok {
		_spec.SetField(txn.FieldAmountRaw, field.TypeFloat64, value)
		_node.AmountRaw = value
	}
	if value, ok := tc.mutation.HandleStatus(); ok {
		_spec.SetField(txn.FieldHandleStatus, field.TypeInt64, value)
		_node.HandleStatus = value
	}
	if value, ok := tc.mutation.HandleMsg(); ok {
		_spec.SetField(txn.FieldHandleMsg, field.TypeString, value)
		_node.HandleMsg = value
	}
	if value, ok := tc.mutation.HandleTime(); ok {
		_spec.SetField(txn.FieldHandleTime, field.TypeTime, value)
		_node.HandleTime = value
	}
	if value, ok := tc.mutation.CollectStatus(); ok {
		_spec.SetField(txn.FieldCollectStatus, field.TypeInt64, value)
		_node.CollectStatus = value
	}
	if value, ok := tc.mutation.CollectMsg(); ok {
		_spec.SetField(txn.FieldCollectMsg, field.TypeString, value)
		_node.CollectMsg = value
	}
	if value, ok := tc.mutation.CollectTime(); ok {
		_spec.SetField(txn.FieldCollectTime, field.TypeTime, value)
		_node.CollectTime = value
	}
	return _node, _spec
}

// TxnCreateBulk is the builder for creating many Txn entities in bulk.
type TxnCreateBulk struct {
	config
	err      error
	builders []*TxnCreate
}

// Save creates the Txn entities in the database.
func (tcb *TxnCreateBulk) Save(ctx context.Context) ([]*Txn, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Txn, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TxnMutation)
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
func (tcb *TxnCreateBulk) SaveX(ctx context.Context) []*Txn {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TxnCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TxnCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
