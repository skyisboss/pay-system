// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent/balance"
	"github.com/skyisboss/pay-system/ent/predicate"
	"github.com/skyisboss/pay-system/ent/schema"
)

// BalanceUpdate is the builder for updating Balance entities.
type BalanceUpdate struct {
	config
	hooks    []Hook
	mutation *BalanceMutation
}

// Where appends a list predicates to the BalanceUpdate builder.
func (bu *BalanceUpdate) Where(ps ...predicate.Balance) *BalanceUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BalanceUpdate) SetUpdatedAt(t time.Time) *BalanceUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bu *BalanceUpdate) SetNillableUpdatedAt(t *time.Time) *BalanceUpdate {
	if t != nil {
		bu.SetUpdatedAt(*t)
	}
	return bu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (bu *BalanceUpdate) ClearUpdatedAt() *BalanceUpdate {
	bu.mutation.ClearUpdatedAt()
	return bu
}

// SetDeletedAt sets the "deleted_at" field.
func (bu *BalanceUpdate) SetDeletedAt(t time.Time) *BalanceUpdate {
	bu.mutation.SetDeletedAt(t)
	return bu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bu *BalanceUpdate) SetNillableDeletedAt(t *time.Time) *BalanceUpdate {
	if t != nil {
		bu.SetDeletedAt(*t)
	}
	return bu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (bu *BalanceUpdate) ClearDeletedAt() *BalanceUpdate {
	bu.mutation.ClearDeletedAt()
	return bu
}

// SetChainID sets the "chain_id" field.
func (bu *BalanceUpdate) SetChainID(u uint64) *BalanceUpdate {
	bu.mutation.ResetChainID()
	bu.mutation.SetChainID(u)
	return bu
}

// AddChainID adds u to the "chain_id" field.
func (bu *BalanceUpdate) AddChainID(u int64) *BalanceUpdate {
	bu.mutation.AddChainID(u)
	return bu
}

// SetProductID sets the "product_id" field.
func (bu *BalanceUpdate) SetProductID(u uint64) *BalanceUpdate {
	bu.mutation.ResetProductID()
	bu.mutation.SetProductID(u)
	return bu
}

// AddProductID adds u to the "product_id" field.
func (bu *BalanceUpdate) AddProductID(u int64) *BalanceUpdate {
	bu.mutation.AddProductID(u)
	return bu
}

// SetBalanceAmount sets the "balance_amount" field.
func (bu *BalanceUpdate) SetBalanceAmount(d decimal.Decimal) *BalanceUpdate {
	bu.mutation.ResetBalanceAmount()
	bu.mutation.SetBalanceAmount(d)
	return bu
}

// AddBalanceAmount adds d to the "balance_amount" field.
func (bu *BalanceUpdate) AddBalanceAmount(d decimal.Decimal) *BalanceUpdate {
	bu.mutation.AddBalanceAmount(d)
	return bu
}

// SetBalanceFreeze sets the "balance_freeze" field.
func (bu *BalanceUpdate) SetBalanceFreeze(d decimal.Decimal) *BalanceUpdate {
	bu.mutation.ResetBalanceFreeze()
	bu.mutation.SetBalanceFreeze(d)
	return bu
}

// AddBalanceFreeze adds d to the "balance_freeze" field.
func (bu *BalanceUpdate) AddBalanceFreeze(d decimal.Decimal) *BalanceUpdate {
	bu.mutation.AddBalanceFreeze(d)
	return bu
}

// SetTotalDeposit sets the "total_deposit" field.
func (bu *BalanceUpdate) SetTotalDeposit(d decimal.Decimal) *BalanceUpdate {
	bu.mutation.ResetTotalDeposit()
	bu.mutation.SetTotalDeposit(d)
	return bu
}

// AddTotalDeposit adds d to the "total_deposit" field.
func (bu *BalanceUpdate) AddTotalDeposit(d decimal.Decimal) *BalanceUpdate {
	bu.mutation.AddTotalDeposit(d)
	return bu
}

// SetTotalWithdraw sets the "total_withdraw" field.
func (bu *BalanceUpdate) SetTotalWithdraw(d decimal.Decimal) *BalanceUpdate {
	bu.mutation.ResetTotalWithdraw()
	bu.mutation.SetTotalWithdraw(d)
	return bu
}

// AddTotalWithdraw adds d to the "total_withdraw" field.
func (bu *BalanceUpdate) AddTotalWithdraw(d decimal.Decimal) *BalanceUpdate {
	bu.mutation.AddTotalWithdraw(d)
	return bu
}

// SetCountDeposit sets the "count_deposit" field.
func (bu *BalanceUpdate) SetCountDeposit(u uint64) *BalanceUpdate {
	bu.mutation.ResetCountDeposit()
	bu.mutation.SetCountDeposit(u)
	return bu
}

// SetNillableCountDeposit sets the "count_deposit" field if the given value is not nil.
func (bu *BalanceUpdate) SetNillableCountDeposit(u *uint64) *BalanceUpdate {
	if u != nil {
		bu.SetCountDeposit(*u)
	}
	return bu
}

// AddCountDeposit adds u to the "count_deposit" field.
func (bu *BalanceUpdate) AddCountDeposit(u int64) *BalanceUpdate {
	bu.mutation.AddCountDeposit(u)
	return bu
}

// SetCountWithdraw sets the "count_withdraw" field.
func (bu *BalanceUpdate) SetCountWithdraw(u uint64) *BalanceUpdate {
	bu.mutation.ResetCountWithdraw()
	bu.mutation.SetCountWithdraw(u)
	return bu
}

// SetNillableCountWithdraw sets the "count_withdraw" field if the given value is not nil.
func (bu *BalanceUpdate) SetNillableCountWithdraw(u *uint64) *BalanceUpdate {
	if u != nil {
		bu.SetCountWithdraw(*u)
	}
	return bu
}

// AddCountWithdraw adds u to the "count_withdraw" field.
func (bu *BalanceUpdate) AddCountWithdraw(u int64) *BalanceUpdate {
	bu.mutation.AddCountWithdraw(u)
	return bu
}

// SetChangeLogs sets the "change_logs" field.
func (bu *BalanceUpdate) SetChangeLogs(sl []schema.ChangeLogs) *BalanceUpdate {
	bu.mutation.SetChangeLogs(sl)
	return bu
}

// AppendChangeLogs appends sl to the "change_logs" field.
func (bu *BalanceUpdate) AppendChangeLogs(sl []schema.ChangeLogs) *BalanceUpdate {
	bu.mutation.AppendChangeLogs(sl)
	return bu
}

// SetVersion sets the "version" field.
func (bu *BalanceUpdate) SetVersion(i int64) *BalanceUpdate {
	bu.mutation.ResetVersion()
	bu.mutation.SetVersion(i)
	return bu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (bu *BalanceUpdate) SetNillableVersion(i *int64) *BalanceUpdate {
	if i != nil {
		bu.SetVersion(*i)
	}
	return bu
}

// AddVersion adds i to the "version" field.
func (bu *BalanceUpdate) AddVersion(i int64) *BalanceUpdate {
	bu.mutation.AddVersion(i)
	return bu
}

// Mutation returns the BalanceMutation object of the builder.
func (bu *BalanceUpdate) Mutation() *BalanceMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BalanceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BalanceUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BalanceUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BalanceUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BalanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(balance.Table, balance.Columns, sqlgraph.NewFieldSpec(balance.FieldID, field.TypeUint64))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(balance.FieldUpdatedAt, field.TypeTime, value)
	}
	if bu.mutation.UpdatedAtCleared() {
		_spec.ClearField(balance.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := bu.mutation.DeletedAt(); ok {
		_spec.SetField(balance.FieldDeletedAt, field.TypeTime, value)
	}
	if bu.mutation.DeletedAtCleared() {
		_spec.ClearField(balance.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := bu.mutation.ChainID(); ok {
		_spec.SetField(balance.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := bu.mutation.AddedChainID(); ok {
		_spec.AddField(balance.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := bu.mutation.ProductID(); ok {
		_spec.SetField(balance.FieldProductID, field.TypeUint64, value)
	}
	if value, ok := bu.mutation.AddedProductID(); ok {
		_spec.AddField(balance.FieldProductID, field.TypeUint64, value)
	}
	if value, ok := bu.mutation.BalanceAmount(); ok {
		_spec.SetField(balance.FieldBalanceAmount, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedBalanceAmount(); ok {
		_spec.AddField(balance.FieldBalanceAmount, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.BalanceFreeze(); ok {
		_spec.SetField(balance.FieldBalanceFreeze, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedBalanceFreeze(); ok {
		_spec.AddField(balance.FieldBalanceFreeze, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.TotalDeposit(); ok {
		_spec.SetField(balance.FieldTotalDeposit, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedTotalDeposit(); ok {
		_spec.AddField(balance.FieldTotalDeposit, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.TotalWithdraw(); ok {
		_spec.SetField(balance.FieldTotalWithdraw, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedTotalWithdraw(); ok {
		_spec.AddField(balance.FieldTotalWithdraw, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.CountDeposit(); ok {
		_spec.SetField(balance.FieldCountDeposit, field.TypeUint64, value)
	}
	if value, ok := bu.mutation.AddedCountDeposit(); ok {
		_spec.AddField(balance.FieldCountDeposit, field.TypeUint64, value)
	}
	if value, ok := bu.mutation.CountWithdraw(); ok {
		_spec.SetField(balance.FieldCountWithdraw, field.TypeUint64, value)
	}
	if value, ok := bu.mutation.AddedCountWithdraw(); ok {
		_spec.AddField(balance.FieldCountWithdraw, field.TypeUint64, value)
	}
	if value, ok := bu.mutation.ChangeLogs(); ok {
		_spec.SetField(balance.FieldChangeLogs, field.TypeJSON, value)
	}
	if value, ok := bu.mutation.AppendedChangeLogs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, balance.FieldChangeLogs, value)
		})
	}
	if value, ok := bu.mutation.Version(); ok {
		_spec.SetField(balance.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.AddedVersion(); ok {
		_spec.AddField(balance.FieldVersion, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{balance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BalanceUpdateOne is the builder for updating a single Balance entity.
type BalanceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BalanceMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BalanceUpdateOne) SetUpdatedAt(t time.Time) *BalanceUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (buo *BalanceUpdateOne) SetNillableUpdatedAt(t *time.Time) *BalanceUpdateOne {
	if t != nil {
		buo.SetUpdatedAt(*t)
	}
	return buo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (buo *BalanceUpdateOne) ClearUpdatedAt() *BalanceUpdateOne {
	buo.mutation.ClearUpdatedAt()
	return buo
}

// SetDeletedAt sets the "deleted_at" field.
func (buo *BalanceUpdateOne) SetDeletedAt(t time.Time) *BalanceUpdateOne {
	buo.mutation.SetDeletedAt(t)
	return buo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (buo *BalanceUpdateOne) SetNillableDeletedAt(t *time.Time) *BalanceUpdateOne {
	if t != nil {
		buo.SetDeletedAt(*t)
	}
	return buo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (buo *BalanceUpdateOne) ClearDeletedAt() *BalanceUpdateOne {
	buo.mutation.ClearDeletedAt()
	return buo
}

// SetChainID sets the "chain_id" field.
func (buo *BalanceUpdateOne) SetChainID(u uint64) *BalanceUpdateOne {
	buo.mutation.ResetChainID()
	buo.mutation.SetChainID(u)
	return buo
}

// AddChainID adds u to the "chain_id" field.
func (buo *BalanceUpdateOne) AddChainID(u int64) *BalanceUpdateOne {
	buo.mutation.AddChainID(u)
	return buo
}

// SetProductID sets the "product_id" field.
func (buo *BalanceUpdateOne) SetProductID(u uint64) *BalanceUpdateOne {
	buo.mutation.ResetProductID()
	buo.mutation.SetProductID(u)
	return buo
}

// AddProductID adds u to the "product_id" field.
func (buo *BalanceUpdateOne) AddProductID(u int64) *BalanceUpdateOne {
	buo.mutation.AddProductID(u)
	return buo
}

// SetBalanceAmount sets the "balance_amount" field.
func (buo *BalanceUpdateOne) SetBalanceAmount(d decimal.Decimal) *BalanceUpdateOne {
	buo.mutation.ResetBalanceAmount()
	buo.mutation.SetBalanceAmount(d)
	return buo
}

// AddBalanceAmount adds d to the "balance_amount" field.
func (buo *BalanceUpdateOne) AddBalanceAmount(d decimal.Decimal) *BalanceUpdateOne {
	buo.mutation.AddBalanceAmount(d)
	return buo
}

// SetBalanceFreeze sets the "balance_freeze" field.
func (buo *BalanceUpdateOne) SetBalanceFreeze(d decimal.Decimal) *BalanceUpdateOne {
	buo.mutation.ResetBalanceFreeze()
	buo.mutation.SetBalanceFreeze(d)
	return buo
}

// AddBalanceFreeze adds d to the "balance_freeze" field.
func (buo *BalanceUpdateOne) AddBalanceFreeze(d decimal.Decimal) *BalanceUpdateOne {
	buo.mutation.AddBalanceFreeze(d)
	return buo
}

// SetTotalDeposit sets the "total_deposit" field.
func (buo *BalanceUpdateOne) SetTotalDeposit(d decimal.Decimal) *BalanceUpdateOne {
	buo.mutation.ResetTotalDeposit()
	buo.mutation.SetTotalDeposit(d)
	return buo
}

// AddTotalDeposit adds d to the "total_deposit" field.
func (buo *BalanceUpdateOne) AddTotalDeposit(d decimal.Decimal) *BalanceUpdateOne {
	buo.mutation.AddTotalDeposit(d)
	return buo
}

// SetTotalWithdraw sets the "total_withdraw" field.
func (buo *BalanceUpdateOne) SetTotalWithdraw(d decimal.Decimal) *BalanceUpdateOne {
	buo.mutation.ResetTotalWithdraw()
	buo.mutation.SetTotalWithdraw(d)
	return buo
}

// AddTotalWithdraw adds d to the "total_withdraw" field.
func (buo *BalanceUpdateOne) AddTotalWithdraw(d decimal.Decimal) *BalanceUpdateOne {
	buo.mutation.AddTotalWithdraw(d)
	return buo
}

// SetCountDeposit sets the "count_deposit" field.
func (buo *BalanceUpdateOne) SetCountDeposit(u uint64) *BalanceUpdateOne {
	buo.mutation.ResetCountDeposit()
	buo.mutation.SetCountDeposit(u)
	return buo
}

// SetNillableCountDeposit sets the "count_deposit" field if the given value is not nil.
func (buo *BalanceUpdateOne) SetNillableCountDeposit(u *uint64) *BalanceUpdateOne {
	if u != nil {
		buo.SetCountDeposit(*u)
	}
	return buo
}

// AddCountDeposit adds u to the "count_deposit" field.
func (buo *BalanceUpdateOne) AddCountDeposit(u int64) *BalanceUpdateOne {
	buo.mutation.AddCountDeposit(u)
	return buo
}

// SetCountWithdraw sets the "count_withdraw" field.
func (buo *BalanceUpdateOne) SetCountWithdraw(u uint64) *BalanceUpdateOne {
	buo.mutation.ResetCountWithdraw()
	buo.mutation.SetCountWithdraw(u)
	return buo
}

// SetNillableCountWithdraw sets the "count_withdraw" field if the given value is not nil.
func (buo *BalanceUpdateOne) SetNillableCountWithdraw(u *uint64) *BalanceUpdateOne {
	if u != nil {
		buo.SetCountWithdraw(*u)
	}
	return buo
}

// AddCountWithdraw adds u to the "count_withdraw" field.
func (buo *BalanceUpdateOne) AddCountWithdraw(u int64) *BalanceUpdateOne {
	buo.mutation.AddCountWithdraw(u)
	return buo
}

// SetChangeLogs sets the "change_logs" field.
func (buo *BalanceUpdateOne) SetChangeLogs(sl []schema.ChangeLogs) *BalanceUpdateOne {
	buo.mutation.SetChangeLogs(sl)
	return buo
}

// AppendChangeLogs appends sl to the "change_logs" field.
func (buo *BalanceUpdateOne) AppendChangeLogs(sl []schema.ChangeLogs) *BalanceUpdateOne {
	buo.mutation.AppendChangeLogs(sl)
	return buo
}

// SetVersion sets the "version" field.
func (buo *BalanceUpdateOne) SetVersion(i int64) *BalanceUpdateOne {
	buo.mutation.ResetVersion()
	buo.mutation.SetVersion(i)
	return buo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (buo *BalanceUpdateOne) SetNillableVersion(i *int64) *BalanceUpdateOne {
	if i != nil {
		buo.SetVersion(*i)
	}
	return buo
}

// AddVersion adds i to the "version" field.
func (buo *BalanceUpdateOne) AddVersion(i int64) *BalanceUpdateOne {
	buo.mutation.AddVersion(i)
	return buo
}

// Mutation returns the BalanceMutation object of the builder.
func (buo *BalanceUpdateOne) Mutation() *BalanceMutation {
	return buo.mutation
}

// Where appends a list predicates to the BalanceUpdate builder.
func (buo *BalanceUpdateOne) Where(ps ...predicate.Balance) *BalanceUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BalanceUpdateOne) Select(field string, fields ...string) *BalanceUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Balance entity.
func (buo *BalanceUpdateOne) Save(ctx context.Context) (*Balance, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BalanceUpdateOne) SaveX(ctx context.Context) *Balance {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BalanceUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BalanceUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BalanceUpdateOne) sqlSave(ctx context.Context) (_node *Balance, err error) {
	_spec := sqlgraph.NewUpdateSpec(balance.Table, balance.Columns, sqlgraph.NewFieldSpec(balance.FieldID, field.TypeUint64))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Balance.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, balance.FieldID)
		for _, f := range fields {
			if !balance.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != balance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(balance.FieldUpdatedAt, field.TypeTime, value)
	}
	if buo.mutation.UpdatedAtCleared() {
		_spec.ClearField(balance.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := buo.mutation.DeletedAt(); ok {
		_spec.SetField(balance.FieldDeletedAt, field.TypeTime, value)
	}
	if buo.mutation.DeletedAtCleared() {
		_spec.ClearField(balance.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := buo.mutation.ChainID(); ok {
		_spec.SetField(balance.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := buo.mutation.AddedChainID(); ok {
		_spec.AddField(balance.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := buo.mutation.ProductID(); ok {
		_spec.SetField(balance.FieldProductID, field.TypeUint64, value)
	}
	if value, ok := buo.mutation.AddedProductID(); ok {
		_spec.AddField(balance.FieldProductID, field.TypeUint64, value)
	}
	if value, ok := buo.mutation.BalanceAmount(); ok {
		_spec.SetField(balance.FieldBalanceAmount, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedBalanceAmount(); ok {
		_spec.AddField(balance.FieldBalanceAmount, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.BalanceFreeze(); ok {
		_spec.SetField(balance.FieldBalanceFreeze, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedBalanceFreeze(); ok {
		_spec.AddField(balance.FieldBalanceFreeze, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.TotalDeposit(); ok {
		_spec.SetField(balance.FieldTotalDeposit, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedTotalDeposit(); ok {
		_spec.AddField(balance.FieldTotalDeposit, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.TotalWithdraw(); ok {
		_spec.SetField(balance.FieldTotalWithdraw, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedTotalWithdraw(); ok {
		_spec.AddField(balance.FieldTotalWithdraw, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.CountDeposit(); ok {
		_spec.SetField(balance.FieldCountDeposit, field.TypeUint64, value)
	}
	if value, ok := buo.mutation.AddedCountDeposit(); ok {
		_spec.AddField(balance.FieldCountDeposit, field.TypeUint64, value)
	}
	if value, ok := buo.mutation.CountWithdraw(); ok {
		_spec.SetField(balance.FieldCountWithdraw, field.TypeUint64, value)
	}
	if value, ok := buo.mutation.AddedCountWithdraw(); ok {
		_spec.AddField(balance.FieldCountWithdraw, field.TypeUint64, value)
	}
	if value, ok := buo.mutation.ChangeLogs(); ok {
		_spec.SetField(balance.FieldChangeLogs, field.TypeJSON, value)
	}
	if value, ok := buo.mutation.AppendedChangeLogs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, balance.FieldChangeLogs, value)
		})
	}
	if value, ok := buo.mutation.Version(); ok {
		_spec.SetField(balance.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.AddedVersion(); ok {
		_spec.AddField(balance.FieldVersion, field.TypeInt64, value)
	}
	_node = &Balance{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{balance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
