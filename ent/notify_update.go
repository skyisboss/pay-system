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
	"github.com/skyisboss/pay-system/ent/notify"
	"github.com/skyisboss/pay-system/ent/predicate"
)

// NotifyUpdate is the builder for updating Notify entities.
type NotifyUpdate struct {
	config
	hooks    []Hook
	mutation *NotifyMutation
}

// Where appends a list predicates to the NotifyUpdate builder.
func (nu *NotifyUpdate) Where(ps ...predicate.Notify) *NotifyUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NotifyUpdate) SetUpdatedAt(t time.Time) *NotifyUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nu *NotifyUpdate) SetNillableUpdatedAt(t *time.Time) *NotifyUpdate {
	if t != nil {
		nu.SetUpdatedAt(*t)
	}
	return nu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (nu *NotifyUpdate) ClearUpdatedAt() *NotifyUpdate {
	nu.mutation.ClearUpdatedAt()
	return nu
}

// SetDeletedAt sets the "deleted_at" field.
func (nu *NotifyUpdate) SetDeletedAt(t time.Time) *NotifyUpdate {
	nu.mutation.SetDeletedAt(t)
	return nu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nu *NotifyUpdate) SetNillableDeletedAt(t *time.Time) *NotifyUpdate {
	if t != nil {
		nu.SetDeletedAt(*t)
	}
	return nu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (nu *NotifyUpdate) ClearDeletedAt() *NotifyUpdate {
	nu.mutation.ClearDeletedAt()
	return nu
}

// SetChainID sets the "chain_id" field.
func (nu *NotifyUpdate) SetChainID(u uint64) *NotifyUpdate {
	nu.mutation.ResetChainID()
	nu.mutation.SetChainID(u)
	return nu
}

// AddChainID adds u to the "chain_id" field.
func (nu *NotifyUpdate) AddChainID(u int64) *NotifyUpdate {
	nu.mutation.AddChainID(u)
	return nu
}

// SetProductID sets the "product_id" field.
func (nu *NotifyUpdate) SetProductID(u uint64) *NotifyUpdate {
	nu.mutation.ResetProductID()
	nu.mutation.SetProductID(u)
	return nu
}

// AddProductID adds u to the "product_id" field.
func (nu *NotifyUpdate) AddProductID(u int64) *NotifyUpdate {
	nu.mutation.AddProductID(u)
	return nu
}

// SetItemFrom sets the "item_from" field.
func (nu *NotifyUpdate) SetItemFrom(u uint64) *NotifyUpdate {
	nu.mutation.ResetItemFrom()
	nu.mutation.SetItemFrom(u)
	return nu
}

// AddItemFrom adds u to the "item_from" field.
func (nu *NotifyUpdate) AddItemFrom(u int64) *NotifyUpdate {
	nu.mutation.AddItemFrom(u)
	return nu
}

// SetItemType sets the "item_type" field.
func (nu *NotifyUpdate) SetItemType(i int64) *NotifyUpdate {
	nu.mutation.ResetItemType()
	nu.mutation.SetItemType(i)
	return nu
}

// AddItemType adds i to the "item_type" field.
func (nu *NotifyUpdate) AddItemType(i int64) *NotifyUpdate {
	nu.mutation.AddItemType(i)
	return nu
}

// SetNonce sets the "nonce" field.
func (nu *NotifyUpdate) SetNonce(s string) *NotifyUpdate {
	nu.mutation.SetNonce(s)
	return nu
}

// SetNotifyType sets the "notify_type" field.
func (nu *NotifyUpdate) SetNotifyType(s string) *NotifyUpdate {
	nu.mutation.SetNotifyType(s)
	return nu
}

// SetSendURL sets the "send_url" field.
func (nu *NotifyUpdate) SetSendURL(s string) *NotifyUpdate {
	nu.mutation.SetSendURL(s)
	return nu
}

// SetSendBody sets the "send_body" field.
func (nu *NotifyUpdate) SetSendBody(s string) *NotifyUpdate {
	nu.mutation.SetSendBody(s)
	return nu
}

// SetSendRetry sets the "send_retry" field.
func (nu *NotifyUpdate) SetSendRetry(i int64) *NotifyUpdate {
	nu.mutation.ResetSendRetry()
	nu.mutation.SetSendRetry(i)
	return nu
}

// AddSendRetry adds i to the "send_retry" field.
func (nu *NotifyUpdate) AddSendRetry(i int64) *NotifyUpdate {
	nu.mutation.AddSendRetry(i)
	return nu
}

// SetHandleStatus sets the "handle_status" field.
func (nu *NotifyUpdate) SetHandleStatus(i int64) *NotifyUpdate {
	nu.mutation.ResetHandleStatus()
	nu.mutation.SetHandleStatus(i)
	return nu
}

// SetNillableHandleStatus sets the "handle_status" field if the given value is not nil.
func (nu *NotifyUpdate) SetNillableHandleStatus(i *int64) *NotifyUpdate {
	if i != nil {
		nu.SetHandleStatus(*i)
	}
	return nu
}

// AddHandleStatus adds i to the "handle_status" field.
func (nu *NotifyUpdate) AddHandleStatus(i int64) *NotifyUpdate {
	nu.mutation.AddHandleStatus(i)
	return nu
}

// SetHandleMsg sets the "handle_msg" field.
func (nu *NotifyUpdate) SetHandleMsg(s string) *NotifyUpdate {
	nu.mutation.SetHandleMsg(s)
	return nu
}

// SetNillableHandleMsg sets the "handle_msg" field if the given value is not nil.
func (nu *NotifyUpdate) SetNillableHandleMsg(s *string) *NotifyUpdate {
	if s != nil {
		nu.SetHandleMsg(*s)
	}
	return nu
}

// SetHandleTime sets the "handle_time" field.
func (nu *NotifyUpdate) SetHandleTime(t time.Time) *NotifyUpdate {
	nu.mutation.SetHandleTime(t)
	return nu
}

// SetNillableHandleTime sets the "handle_time" field if the given value is not nil.
func (nu *NotifyUpdate) SetNillableHandleTime(t *time.Time) *NotifyUpdate {
	if t != nil {
		nu.SetHandleTime(*t)
	}
	return nu
}

// ClearHandleTime clears the value of the "handle_time" field.
func (nu *NotifyUpdate) ClearHandleTime() *NotifyUpdate {
	nu.mutation.ClearHandleTime()
	return nu
}

// Mutation returns the NotifyMutation object of the builder.
func (nu *NotifyUpdate) Mutation() *NotifyMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NotifyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NotifyUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NotifyUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NotifyUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NotifyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(notify.Table, notify.Columns, sqlgraph.NewFieldSpec(notify.FieldID, field.TypeUint64))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.SetField(notify.FieldUpdatedAt, field.TypeTime, value)
	}
	if nu.mutation.UpdatedAtCleared() {
		_spec.ClearField(notify.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := nu.mutation.DeletedAt(); ok {
		_spec.SetField(notify.FieldDeletedAt, field.TypeTime, value)
	}
	if nu.mutation.DeletedAtCleared() {
		_spec.ClearField(notify.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := nu.mutation.ChainID(); ok {
		_spec.SetField(notify.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := nu.mutation.AddedChainID(); ok {
		_spec.AddField(notify.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := nu.mutation.ProductID(); ok {
		_spec.SetField(notify.FieldProductID, field.TypeUint64, value)
	}
	if value, ok := nu.mutation.AddedProductID(); ok {
		_spec.AddField(notify.FieldProductID, field.TypeUint64, value)
	}
	if value, ok := nu.mutation.ItemFrom(); ok {
		_spec.SetField(notify.FieldItemFrom, field.TypeUint64, value)
	}
	if value, ok := nu.mutation.AddedItemFrom(); ok {
		_spec.AddField(notify.FieldItemFrom, field.TypeUint64, value)
	}
	if value, ok := nu.mutation.ItemType(); ok {
		_spec.SetField(notify.FieldItemType, field.TypeInt64, value)
	}
	if value, ok := nu.mutation.AddedItemType(); ok {
		_spec.AddField(notify.FieldItemType, field.TypeInt64, value)
	}
	if value, ok := nu.mutation.Nonce(); ok {
		_spec.SetField(notify.FieldNonce, field.TypeString, value)
	}
	if value, ok := nu.mutation.NotifyType(); ok {
		_spec.SetField(notify.FieldNotifyType, field.TypeString, value)
	}
	if value, ok := nu.mutation.SendURL(); ok {
		_spec.SetField(notify.FieldSendURL, field.TypeString, value)
	}
	if value, ok := nu.mutation.SendBody(); ok {
		_spec.SetField(notify.FieldSendBody, field.TypeString, value)
	}
	if value, ok := nu.mutation.SendRetry(); ok {
		_spec.SetField(notify.FieldSendRetry, field.TypeInt64, value)
	}
	if value, ok := nu.mutation.AddedSendRetry(); ok {
		_spec.AddField(notify.FieldSendRetry, field.TypeInt64, value)
	}
	if value, ok := nu.mutation.HandleStatus(); ok {
		_spec.SetField(notify.FieldHandleStatus, field.TypeInt64, value)
	}
	if value, ok := nu.mutation.AddedHandleStatus(); ok {
		_spec.AddField(notify.FieldHandleStatus, field.TypeInt64, value)
	}
	if value, ok := nu.mutation.HandleMsg(); ok {
		_spec.SetField(notify.FieldHandleMsg, field.TypeString, value)
	}
	if value, ok := nu.mutation.HandleTime(); ok {
		_spec.SetField(notify.FieldHandleTime, field.TypeTime, value)
	}
	if nu.mutation.HandleTimeCleared() {
		_spec.ClearField(notify.FieldHandleTime, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notify.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NotifyUpdateOne is the builder for updating a single Notify entity.
type NotifyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotifyMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NotifyUpdateOne) SetUpdatedAt(t time.Time) *NotifyUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nuo *NotifyUpdateOne) SetNillableUpdatedAt(t *time.Time) *NotifyUpdateOne {
	if t != nil {
		nuo.SetUpdatedAt(*t)
	}
	return nuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (nuo *NotifyUpdateOne) ClearUpdatedAt() *NotifyUpdateOne {
	nuo.mutation.ClearUpdatedAt()
	return nuo
}

// SetDeletedAt sets the "deleted_at" field.
func (nuo *NotifyUpdateOne) SetDeletedAt(t time.Time) *NotifyUpdateOne {
	nuo.mutation.SetDeletedAt(t)
	return nuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nuo *NotifyUpdateOne) SetNillableDeletedAt(t *time.Time) *NotifyUpdateOne {
	if t != nil {
		nuo.SetDeletedAt(*t)
	}
	return nuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (nuo *NotifyUpdateOne) ClearDeletedAt() *NotifyUpdateOne {
	nuo.mutation.ClearDeletedAt()
	return nuo
}

// SetChainID sets the "chain_id" field.
func (nuo *NotifyUpdateOne) SetChainID(u uint64) *NotifyUpdateOne {
	nuo.mutation.ResetChainID()
	nuo.mutation.SetChainID(u)
	return nuo
}

// AddChainID adds u to the "chain_id" field.
func (nuo *NotifyUpdateOne) AddChainID(u int64) *NotifyUpdateOne {
	nuo.mutation.AddChainID(u)
	return nuo
}

// SetProductID sets the "product_id" field.
func (nuo *NotifyUpdateOne) SetProductID(u uint64) *NotifyUpdateOne {
	nuo.mutation.ResetProductID()
	nuo.mutation.SetProductID(u)
	return nuo
}

// AddProductID adds u to the "product_id" field.
func (nuo *NotifyUpdateOne) AddProductID(u int64) *NotifyUpdateOne {
	nuo.mutation.AddProductID(u)
	return nuo
}

// SetItemFrom sets the "item_from" field.
func (nuo *NotifyUpdateOne) SetItemFrom(u uint64) *NotifyUpdateOne {
	nuo.mutation.ResetItemFrom()
	nuo.mutation.SetItemFrom(u)
	return nuo
}

// AddItemFrom adds u to the "item_from" field.
func (nuo *NotifyUpdateOne) AddItemFrom(u int64) *NotifyUpdateOne {
	nuo.mutation.AddItemFrom(u)
	return nuo
}

// SetItemType sets the "item_type" field.
func (nuo *NotifyUpdateOne) SetItemType(i int64) *NotifyUpdateOne {
	nuo.mutation.ResetItemType()
	nuo.mutation.SetItemType(i)
	return nuo
}

// AddItemType adds i to the "item_type" field.
func (nuo *NotifyUpdateOne) AddItemType(i int64) *NotifyUpdateOne {
	nuo.mutation.AddItemType(i)
	return nuo
}

// SetNonce sets the "nonce" field.
func (nuo *NotifyUpdateOne) SetNonce(s string) *NotifyUpdateOne {
	nuo.mutation.SetNonce(s)
	return nuo
}

// SetNotifyType sets the "notify_type" field.
func (nuo *NotifyUpdateOne) SetNotifyType(s string) *NotifyUpdateOne {
	nuo.mutation.SetNotifyType(s)
	return nuo
}

// SetSendURL sets the "send_url" field.
func (nuo *NotifyUpdateOne) SetSendURL(s string) *NotifyUpdateOne {
	nuo.mutation.SetSendURL(s)
	return nuo
}

// SetSendBody sets the "send_body" field.
func (nuo *NotifyUpdateOne) SetSendBody(s string) *NotifyUpdateOne {
	nuo.mutation.SetSendBody(s)
	return nuo
}

// SetSendRetry sets the "send_retry" field.
func (nuo *NotifyUpdateOne) SetSendRetry(i int64) *NotifyUpdateOne {
	nuo.mutation.ResetSendRetry()
	nuo.mutation.SetSendRetry(i)
	return nuo
}

// AddSendRetry adds i to the "send_retry" field.
func (nuo *NotifyUpdateOne) AddSendRetry(i int64) *NotifyUpdateOne {
	nuo.mutation.AddSendRetry(i)
	return nuo
}

// SetHandleStatus sets the "handle_status" field.
func (nuo *NotifyUpdateOne) SetHandleStatus(i int64) *NotifyUpdateOne {
	nuo.mutation.ResetHandleStatus()
	nuo.mutation.SetHandleStatus(i)
	return nuo
}

// SetNillableHandleStatus sets the "handle_status" field if the given value is not nil.
func (nuo *NotifyUpdateOne) SetNillableHandleStatus(i *int64) *NotifyUpdateOne {
	if i != nil {
		nuo.SetHandleStatus(*i)
	}
	return nuo
}

// AddHandleStatus adds i to the "handle_status" field.
func (nuo *NotifyUpdateOne) AddHandleStatus(i int64) *NotifyUpdateOne {
	nuo.mutation.AddHandleStatus(i)
	return nuo
}

// SetHandleMsg sets the "handle_msg" field.
func (nuo *NotifyUpdateOne) SetHandleMsg(s string) *NotifyUpdateOne {
	nuo.mutation.SetHandleMsg(s)
	return nuo
}

// SetNillableHandleMsg sets the "handle_msg" field if the given value is not nil.
func (nuo *NotifyUpdateOne) SetNillableHandleMsg(s *string) *NotifyUpdateOne {
	if s != nil {
		nuo.SetHandleMsg(*s)
	}
	return nuo
}

// SetHandleTime sets the "handle_time" field.
func (nuo *NotifyUpdateOne) SetHandleTime(t time.Time) *NotifyUpdateOne {
	nuo.mutation.SetHandleTime(t)
	return nuo
}

// SetNillableHandleTime sets the "handle_time" field if the given value is not nil.
func (nuo *NotifyUpdateOne) SetNillableHandleTime(t *time.Time) *NotifyUpdateOne {
	if t != nil {
		nuo.SetHandleTime(*t)
	}
	return nuo
}

// ClearHandleTime clears the value of the "handle_time" field.
func (nuo *NotifyUpdateOne) ClearHandleTime() *NotifyUpdateOne {
	nuo.mutation.ClearHandleTime()
	return nuo
}

// Mutation returns the NotifyMutation object of the builder.
func (nuo *NotifyUpdateOne) Mutation() *NotifyMutation {
	return nuo.mutation
}

// Where appends a list predicates to the NotifyUpdate builder.
func (nuo *NotifyUpdateOne) Where(ps ...predicate.Notify) *NotifyUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NotifyUpdateOne) Select(field string, fields ...string) *NotifyUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Notify entity.
func (nuo *NotifyUpdateOne) Save(ctx context.Context) (*Notify, error) {
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NotifyUpdateOne) SaveX(ctx context.Context) *Notify {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NotifyUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NotifyUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NotifyUpdateOne) sqlSave(ctx context.Context) (_node *Notify, err error) {
	_spec := sqlgraph.NewUpdateSpec(notify.Table, notify.Columns, sqlgraph.NewFieldSpec(notify.FieldID, field.TypeUint64))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Notify.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notify.FieldID)
		for _, f := range fields {
			if !notify.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notify.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.SetField(notify.FieldUpdatedAt, field.TypeTime, value)
	}
	if nuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(notify.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := nuo.mutation.DeletedAt(); ok {
		_spec.SetField(notify.FieldDeletedAt, field.TypeTime, value)
	}
	if nuo.mutation.DeletedAtCleared() {
		_spec.ClearField(notify.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := nuo.mutation.ChainID(); ok {
		_spec.SetField(notify.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := nuo.mutation.AddedChainID(); ok {
		_spec.AddField(notify.FieldChainID, field.TypeUint64, value)
	}
	if value, ok := nuo.mutation.ProductID(); ok {
		_spec.SetField(notify.FieldProductID, field.TypeUint64, value)
	}
	if value, ok := nuo.mutation.AddedProductID(); ok {
		_spec.AddField(notify.FieldProductID, field.TypeUint64, value)
	}
	if value, ok := nuo.mutation.ItemFrom(); ok {
		_spec.SetField(notify.FieldItemFrom, field.TypeUint64, value)
	}
	if value, ok := nuo.mutation.AddedItemFrom(); ok {
		_spec.AddField(notify.FieldItemFrom, field.TypeUint64, value)
	}
	if value, ok := nuo.mutation.ItemType(); ok {
		_spec.SetField(notify.FieldItemType, field.TypeInt64, value)
	}
	if value, ok := nuo.mutation.AddedItemType(); ok {
		_spec.AddField(notify.FieldItemType, field.TypeInt64, value)
	}
	if value, ok := nuo.mutation.Nonce(); ok {
		_spec.SetField(notify.FieldNonce, field.TypeString, value)
	}
	if value, ok := nuo.mutation.NotifyType(); ok {
		_spec.SetField(notify.FieldNotifyType, field.TypeString, value)
	}
	if value, ok := nuo.mutation.SendURL(); ok {
		_spec.SetField(notify.FieldSendURL, field.TypeString, value)
	}
	if value, ok := nuo.mutation.SendBody(); ok {
		_spec.SetField(notify.FieldSendBody, field.TypeString, value)
	}
	if value, ok := nuo.mutation.SendRetry(); ok {
		_spec.SetField(notify.FieldSendRetry, field.TypeInt64, value)
	}
	if value, ok := nuo.mutation.AddedSendRetry(); ok {
		_spec.AddField(notify.FieldSendRetry, field.TypeInt64, value)
	}
	if value, ok := nuo.mutation.HandleStatus(); ok {
		_spec.SetField(notify.FieldHandleStatus, field.TypeInt64, value)
	}
	if value, ok := nuo.mutation.AddedHandleStatus(); ok {
		_spec.AddField(notify.FieldHandleStatus, field.TypeInt64, value)
	}
	if value, ok := nuo.mutation.HandleMsg(); ok {
		_spec.SetField(notify.FieldHandleMsg, field.TypeString, value)
	}
	if value, ok := nuo.mutation.HandleTime(); ok {
		_spec.SetField(notify.FieldHandleTime, field.TypeTime, value)
	}
	if nuo.mutation.HandleTimeCleared() {
		_spec.ClearField(notify.FieldHandleTime, field.TypeTime)
	}
	_node = &Notify{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notify.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}