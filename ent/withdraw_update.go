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
	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent/predicate"
	"github.com/skyisboss/pay-system/ent/withdraw"
)

// WithdrawUpdate is the builder for updating Withdraw entities.
type WithdrawUpdate struct {
	config
	hooks    []Hook
	mutation *WithdrawMutation
}

// Where appends a list predicates to the WithdrawUpdate builder.
func (wu *WithdrawUpdate) Where(ps ...predicate.Withdraw) *WithdrawUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetUpdatedAt sets the "updated_at" field.
func (wu *WithdrawUpdate) SetUpdatedAt(t time.Time) *WithdrawUpdate {
	wu.mutation.SetUpdatedAt(t)
	return wu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableUpdatedAt(t *time.Time) *WithdrawUpdate {
	if t != nil {
		wu.SetUpdatedAt(*t)
	}
	return wu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (wu *WithdrawUpdate) ClearUpdatedAt() *WithdrawUpdate {
	wu.mutation.ClearUpdatedAt()
	return wu
}

// SetDeletedAt sets the "deleted_at" field.
func (wu *WithdrawUpdate) SetDeletedAt(t time.Time) *WithdrawUpdate {
	wu.mutation.SetDeletedAt(t)
	return wu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableDeletedAt(t *time.Time) *WithdrawUpdate {
	if t != nil {
		wu.SetDeletedAt(*t)
	}
	return wu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (wu *WithdrawUpdate) ClearDeletedAt() *WithdrawUpdate {
	wu.mutation.ClearDeletedAt()
	return wu
}

// SetProductID sets the "product_id" field.
func (wu *WithdrawUpdate) SetProductID(i int64) *WithdrawUpdate {
	wu.mutation.ResetProductID()
	wu.mutation.SetProductID(i)
	return wu
}

// AddProductID adds i to the "product_id" field.
func (wu *WithdrawUpdate) AddProductID(i int64) *WithdrawUpdate {
	wu.mutation.AddProductID(i)
	return wu
}

// SetChainID sets the "chain_id" field.
func (wu *WithdrawUpdate) SetChainID(u uint64) *WithdrawUpdate {
	wu.mutation.ResetChainID()
	wu.mutation.SetChainID(u)
	return wu
}

// AddChainID adds u to the "chain_id" field.
func (wu *WithdrawUpdate) AddChainID(u int64) *WithdrawUpdate {
	wu.mutation.AddChainID(u)
	return wu
}

// SetToAddress sets the "to_address" field.
func (wu *WithdrawUpdate) SetToAddress(s string) *WithdrawUpdate {
	wu.mutation.SetToAddress(s)
	return wu
}

// SetAmountStr sets the "amount_str" field.
func (wu *WithdrawUpdate) SetAmountStr(s string) *WithdrawUpdate {
	wu.mutation.SetAmountStr(s)
	return wu
}

// SetAmountRaw sets the "amount_raw" field.
func (wu *WithdrawUpdate) SetAmountRaw(d decimal.Decimal) *WithdrawUpdate {
	wu.mutation.ResetAmountRaw()
	wu.mutation.SetAmountRaw(d)
	return wu
}

// AddAmountRaw adds d to the "amount_raw" field.
func (wu *WithdrawUpdate) AddAmountRaw(d decimal.Decimal) *WithdrawUpdate {
	wu.mutation.AddAmountRaw(d)
	return wu
}

// SetSerialID sets the "serial_id" field.
func (wu *WithdrawUpdate) SetSerialID(s string) *WithdrawUpdate {
	wu.mutation.SetSerialID(s)
	return wu
}

// SetTxHash sets the "tx_hash" field.
func (wu *WithdrawUpdate) SetTxHash(s string) *WithdrawUpdate {
	wu.mutation.SetTxHash(s)
	return wu
}

// SetNillableTxHash sets the "tx_hash" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableTxHash(s *string) *WithdrawUpdate {
	if s != nil {
		wu.SetTxHash(*s)
	}
	return wu
}

// ClearTxHash clears the value of the "tx_hash" field.
func (wu *WithdrawUpdate) ClearTxHash() *WithdrawUpdate {
	wu.mutation.ClearTxHash()
	return wu
}

// SetHandleStatus sets the "handle_status" field.
func (wu *WithdrawUpdate) SetHandleStatus(i int64) *WithdrawUpdate {
	wu.mutation.ResetHandleStatus()
	wu.mutation.SetHandleStatus(i)
	return wu
}

// SetNillableHandleStatus sets the "handle_status" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableHandleStatus(i *int64) *WithdrawUpdate {
	if i != nil {
		wu.SetHandleStatus(*i)
	}
	return wu
}

// AddHandleStatus adds i to the "handle_status" field.
func (wu *WithdrawUpdate) AddHandleStatus(i int64) *WithdrawUpdate {
	wu.mutation.AddHandleStatus(i)
	return wu
}

// SetHandleMsg sets the "handle_msg" field.
func (wu *WithdrawUpdate) SetHandleMsg(s string) *WithdrawUpdate {
	wu.mutation.SetHandleMsg(s)
	return wu
}

// SetNillableHandleMsg sets the "handle_msg" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableHandleMsg(s *string) *WithdrawUpdate {
	if s != nil {
		wu.SetHandleMsg(*s)
	}
	return wu
}

// ClearHandleMsg clears the value of the "handle_msg" field.
func (wu *WithdrawUpdate) ClearHandleMsg() *WithdrawUpdate {
	wu.mutation.ClearHandleMsg()
	return wu
}

// SetHandleTime sets the "handle_time" field.
func (wu *WithdrawUpdate) SetHandleTime(t time.Time) *WithdrawUpdate {
	wu.mutation.SetHandleTime(t)
	return wu
}

// SetNillableHandleTime sets the "handle_time" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableHandleTime(t *time.Time) *WithdrawUpdate {
	if t != nil {
		wu.SetHandleTime(*t)
	}
	return wu
}

// ClearHandleTime clears the value of the "handle_time" field.
func (wu *WithdrawUpdate) ClearHandleTime() *WithdrawUpdate {
	wu.mutation.ClearHandleTime()
	return wu
}

// Mutation returns the WithdrawMutation object of the builder.
func (wu *WithdrawUpdate) Mutation() *WithdrawMutation {
	return wu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WithdrawUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WithdrawUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WithdrawUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WithdrawUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wu *WithdrawUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(withdraw.Table, withdraw.Columns, sqlgraph.NewFieldSpec(withdraw.FieldID, field.TypeUint64))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.UpdatedAt(); ok {
		_spec.SetField(withdraw.FieldUpdatedAt, field.TypeTime, value)
	}
	if wu.mutation.UpdatedAtCleared() {
		_spec.ClearField(withdraw.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := wu.mutation.DeletedAt(); ok {
		_spec.SetField(withdraw.FieldDeletedAt, field.TypeTime, value)
	}
	if wu.mutation.DeletedAtCleared() {
		_spec.ClearField(withdraw.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := wu.mutation.ProductID(); ok {
		_spec.SetField(withdraw.FieldProductID, field.TypeInt64, value)
	}
	if value, ok := wu.mutation.AddedProductID(); ok {
		_spec.AddField(withdraw.FieldProductID, field.TypeInt64, value)
	}
	if value, ok := wu.mutation.ChainID(); ok {
		_spec.SetField(withdraw.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := wu.mutation.AddedChainID(); ok {
		_spec.AddField(withdraw.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := wu.mutation.ToAddress(); ok {
		_spec.SetField(withdraw.FieldToAddress, field.TypeString, value)
	}
	if value, ok := wu.mutation.AmountStr(); ok {
		_spec.SetField(withdraw.FieldAmountStr, field.TypeString, value)
	}
	if value, ok := wu.mutation.AmountRaw(); ok {
		_spec.SetField(withdraw.FieldAmountRaw, field.TypeFloat64, value)
	}
	if value, ok := wu.mutation.AddedAmountRaw(); ok {
		_spec.AddField(withdraw.FieldAmountRaw, field.TypeFloat64, value)
	}
	if value, ok := wu.mutation.SerialID(); ok {
		_spec.SetField(withdraw.FieldSerialID, field.TypeString, value)
	}
	if value, ok := wu.mutation.TxHash(); ok {
		_spec.SetField(withdraw.FieldTxHash, field.TypeString, value)
	}
	if wu.mutation.TxHashCleared() {
		_spec.ClearField(withdraw.FieldTxHash, field.TypeString)
	}
	if value, ok := wu.mutation.HandleStatus(); ok {
		_spec.SetField(withdraw.FieldHandleStatus, field.TypeInt64, value)
	}
	if value, ok := wu.mutation.AddedHandleStatus(); ok {
		_spec.AddField(withdraw.FieldHandleStatus, field.TypeInt64, value)
	}
	if value, ok := wu.mutation.HandleMsg(); ok {
		_spec.SetField(withdraw.FieldHandleMsg, field.TypeString, value)
	}
	if wu.mutation.HandleMsgCleared() {
		_spec.ClearField(withdraw.FieldHandleMsg, field.TypeString)
	}
	if value, ok := wu.mutation.HandleTime(); ok {
		_spec.SetField(withdraw.FieldHandleTime, field.TypeTime, value)
	}
	if wu.mutation.HandleTimeCleared() {
		_spec.ClearField(withdraw.FieldHandleTime, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withdraw.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WithdrawUpdateOne is the builder for updating a single Withdraw entity.
type WithdrawUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WithdrawMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (wuo *WithdrawUpdateOne) SetUpdatedAt(t time.Time) *WithdrawUpdateOne {
	wuo.mutation.SetUpdatedAt(t)
	return wuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableUpdatedAt(t *time.Time) *WithdrawUpdateOne {
	if t != nil {
		wuo.SetUpdatedAt(*t)
	}
	return wuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (wuo *WithdrawUpdateOne) ClearUpdatedAt() *WithdrawUpdateOne {
	wuo.mutation.ClearUpdatedAt()
	return wuo
}

// SetDeletedAt sets the "deleted_at" field.
func (wuo *WithdrawUpdateOne) SetDeletedAt(t time.Time) *WithdrawUpdateOne {
	wuo.mutation.SetDeletedAt(t)
	return wuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableDeletedAt(t *time.Time) *WithdrawUpdateOne {
	if t != nil {
		wuo.SetDeletedAt(*t)
	}
	return wuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (wuo *WithdrawUpdateOne) ClearDeletedAt() *WithdrawUpdateOne {
	wuo.mutation.ClearDeletedAt()
	return wuo
}

// SetProductID sets the "product_id" field.
func (wuo *WithdrawUpdateOne) SetProductID(i int64) *WithdrawUpdateOne {
	wuo.mutation.ResetProductID()
	wuo.mutation.SetProductID(i)
	return wuo
}

// AddProductID adds i to the "product_id" field.
func (wuo *WithdrawUpdateOne) AddProductID(i int64) *WithdrawUpdateOne {
	wuo.mutation.AddProductID(i)
	return wuo
}

// SetChainID sets the "chain_id" field.
func (wuo *WithdrawUpdateOne) SetChainID(u uint64) *WithdrawUpdateOne {
	wuo.mutation.ResetChainID()
	wuo.mutation.SetChainID(u)
	return wuo
}

// AddChainID adds u to the "chain_id" field.
func (wuo *WithdrawUpdateOne) AddChainID(u int64) *WithdrawUpdateOne {
	wuo.mutation.AddChainID(u)
	return wuo
}

// SetToAddress sets the "to_address" field.
func (wuo *WithdrawUpdateOne) SetToAddress(s string) *WithdrawUpdateOne {
	wuo.mutation.SetToAddress(s)
	return wuo
}

// SetAmountStr sets the "amount_str" field.
func (wuo *WithdrawUpdateOne) SetAmountStr(s string) *WithdrawUpdateOne {
	wuo.mutation.SetAmountStr(s)
	return wuo
}

// SetAmountRaw sets the "amount_raw" field.
func (wuo *WithdrawUpdateOne) SetAmountRaw(d decimal.Decimal) *WithdrawUpdateOne {
	wuo.mutation.ResetAmountRaw()
	wuo.mutation.SetAmountRaw(d)
	return wuo
}

// AddAmountRaw adds d to the "amount_raw" field.
func (wuo *WithdrawUpdateOne) AddAmountRaw(d decimal.Decimal) *WithdrawUpdateOne {
	wuo.mutation.AddAmountRaw(d)
	return wuo
}

// SetSerialID sets the "serial_id" field.
func (wuo *WithdrawUpdateOne) SetSerialID(s string) *WithdrawUpdateOne {
	wuo.mutation.SetSerialID(s)
	return wuo
}

// SetTxHash sets the "tx_hash" field.
func (wuo *WithdrawUpdateOne) SetTxHash(s string) *WithdrawUpdateOne {
	wuo.mutation.SetTxHash(s)
	return wuo
}

// SetNillableTxHash sets the "tx_hash" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableTxHash(s *string) *WithdrawUpdateOne {
	if s != nil {
		wuo.SetTxHash(*s)
	}
	return wuo
}

// ClearTxHash clears the value of the "tx_hash" field.
func (wuo *WithdrawUpdateOne) ClearTxHash() *WithdrawUpdateOne {
	wuo.mutation.ClearTxHash()
	return wuo
}

// SetHandleStatus sets the "handle_status" field.
func (wuo *WithdrawUpdateOne) SetHandleStatus(i int64) *WithdrawUpdateOne {
	wuo.mutation.ResetHandleStatus()
	wuo.mutation.SetHandleStatus(i)
	return wuo
}

// SetNillableHandleStatus sets the "handle_status" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableHandleStatus(i *int64) *WithdrawUpdateOne {
	if i != nil {
		wuo.SetHandleStatus(*i)
	}
	return wuo
}

// AddHandleStatus adds i to the "handle_status" field.
func (wuo *WithdrawUpdateOne) AddHandleStatus(i int64) *WithdrawUpdateOne {
	wuo.mutation.AddHandleStatus(i)
	return wuo
}

// SetHandleMsg sets the "handle_msg" field.
func (wuo *WithdrawUpdateOne) SetHandleMsg(s string) *WithdrawUpdateOne {
	wuo.mutation.SetHandleMsg(s)
	return wuo
}

// SetNillableHandleMsg sets the "handle_msg" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableHandleMsg(s *string) *WithdrawUpdateOne {
	if s != nil {
		wuo.SetHandleMsg(*s)
	}
	return wuo
}

// ClearHandleMsg clears the value of the "handle_msg" field.
func (wuo *WithdrawUpdateOne) ClearHandleMsg() *WithdrawUpdateOne {
	wuo.mutation.ClearHandleMsg()
	return wuo
}

// SetHandleTime sets the "handle_time" field.
func (wuo *WithdrawUpdateOne) SetHandleTime(t time.Time) *WithdrawUpdateOne {
	wuo.mutation.SetHandleTime(t)
	return wuo
}

// SetNillableHandleTime sets the "handle_time" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableHandleTime(t *time.Time) *WithdrawUpdateOne {
	if t != nil {
		wuo.SetHandleTime(*t)
	}
	return wuo
}

// ClearHandleTime clears the value of the "handle_time" field.
func (wuo *WithdrawUpdateOne) ClearHandleTime() *WithdrawUpdateOne {
	wuo.mutation.ClearHandleTime()
	return wuo
}

// Mutation returns the WithdrawMutation object of the builder.
func (wuo *WithdrawUpdateOne) Mutation() *WithdrawMutation {
	return wuo.mutation
}

// Where appends a list predicates to the WithdrawUpdate builder.
func (wuo *WithdrawUpdateOne) Where(ps ...predicate.Withdraw) *WithdrawUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WithdrawUpdateOne) Select(field string, fields ...string) *WithdrawUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Withdraw entity.
func (wuo *WithdrawUpdateOne) Save(ctx context.Context) (*Withdraw, error) {
	return withHooks(ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WithdrawUpdateOne) SaveX(ctx context.Context) *Withdraw {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WithdrawUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WithdrawUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wuo *WithdrawUpdateOne) sqlSave(ctx context.Context) (_node *Withdraw, err error) {
	_spec := sqlgraph.NewUpdateSpec(withdraw.Table, withdraw.Columns, sqlgraph.NewFieldSpec(withdraw.FieldID, field.TypeUint64))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Withdraw.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withdraw.FieldID)
		for _, f := range fields {
			if !withdraw.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != withdraw.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.UpdatedAt(); ok {
		_spec.SetField(withdraw.FieldUpdatedAt, field.TypeTime, value)
	}
	if wuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(withdraw.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := wuo.mutation.DeletedAt(); ok {
		_spec.SetField(withdraw.FieldDeletedAt, field.TypeTime, value)
	}
	if wuo.mutation.DeletedAtCleared() {
		_spec.ClearField(withdraw.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := wuo.mutation.ProductID(); ok {
		_spec.SetField(withdraw.FieldProductID, field.TypeInt64, value)
	}
	if value, ok := wuo.mutation.AddedProductID(); ok {
		_spec.AddField(withdraw.FieldProductID, field.TypeInt64, value)
	}
	if value, ok := wuo.mutation.ChainID(); ok {
		_spec.SetField(withdraw.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := wuo.mutation.AddedChainID(); ok {
		_spec.AddField(withdraw.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := wuo.mutation.ToAddress(); ok {
		_spec.SetField(withdraw.FieldToAddress, field.TypeString, value)
	}
	if value, ok := wuo.mutation.AmountStr(); ok {
		_spec.SetField(withdraw.FieldAmountStr, field.TypeString, value)
	}
	if value, ok := wuo.mutation.AmountRaw(); ok {
		_spec.SetField(withdraw.FieldAmountRaw, field.TypeFloat64, value)
	}
	if value, ok := wuo.mutation.AddedAmountRaw(); ok {
		_spec.AddField(withdraw.FieldAmountRaw, field.TypeFloat64, value)
	}
	if value, ok := wuo.mutation.SerialID(); ok {
		_spec.SetField(withdraw.FieldSerialID, field.TypeString, value)
	}
	if value, ok := wuo.mutation.TxHash(); ok {
		_spec.SetField(withdraw.FieldTxHash, field.TypeString, value)
	}
	if wuo.mutation.TxHashCleared() {
		_spec.ClearField(withdraw.FieldTxHash, field.TypeString)
	}
	if value, ok := wuo.mutation.HandleStatus(); ok {
		_spec.SetField(withdraw.FieldHandleStatus, field.TypeInt64, value)
	}
	if value, ok := wuo.mutation.AddedHandleStatus(); ok {
		_spec.AddField(withdraw.FieldHandleStatus, field.TypeInt64, value)
	}
	if value, ok := wuo.mutation.HandleMsg(); ok {
		_spec.SetField(withdraw.FieldHandleMsg, field.TypeString, value)
	}
	if wuo.mutation.HandleMsgCleared() {
		_spec.ClearField(withdraw.FieldHandleMsg, field.TypeString)
	}
	if value, ok := wuo.mutation.HandleTime(); ok {
		_spec.SetField(withdraw.FieldHandleTime, field.TypeTime, value)
	}
	if wuo.mutation.HandleTimeCleared() {
		_spec.ClearField(withdraw.FieldHandleTime, field.TypeTime)
	}
	_node = &Withdraw{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withdraw.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}
