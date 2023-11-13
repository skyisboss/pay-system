// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/skyisboss/pay-system/ent/predicate"
	"github.com/skyisboss/pay-system/ent/withdraw"
)

// WithdrawQuery is the builder for querying Withdraw entities.
type WithdrawQuery struct {
	config
	ctx        *QueryContext
	order      []withdraw.OrderOption
	inters     []Interceptor
	predicates []predicate.Withdraw
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WithdrawQuery builder.
func (wq *WithdrawQuery) Where(ps ...predicate.Withdraw) *WithdrawQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit the number of records to be returned by this query.
func (wq *WithdrawQuery) Limit(limit int) *WithdrawQuery {
	wq.ctx.Limit = &limit
	return wq
}

// Offset to start from.
func (wq *WithdrawQuery) Offset(offset int) *WithdrawQuery {
	wq.ctx.Offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WithdrawQuery) Unique(unique bool) *WithdrawQuery {
	wq.ctx.Unique = &unique
	return wq
}

// Order specifies how the records should be ordered.
func (wq *WithdrawQuery) Order(o ...withdraw.OrderOption) *WithdrawQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// First returns the first Withdraw entity from the query.
// Returns a *NotFoundError when no Withdraw was found.
func (wq *WithdrawQuery) First(ctx context.Context) (*Withdraw, error) {
	nodes, err := wq.Limit(1).All(setContextOp(ctx, wq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{withdraw.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WithdrawQuery) FirstX(ctx context.Context) *Withdraw {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Withdraw ID from the query.
// Returns a *NotFoundError when no Withdraw ID was found.
func (wq *WithdrawQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = wq.Limit(1).IDs(setContextOp(ctx, wq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{withdraw.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WithdrawQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Withdraw entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Withdraw entity is found.
// Returns a *NotFoundError when no Withdraw entities are found.
func (wq *WithdrawQuery) Only(ctx context.Context) (*Withdraw, error) {
	nodes, err := wq.Limit(2).All(setContextOp(ctx, wq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{withdraw.Label}
	default:
		return nil, &NotSingularError{withdraw.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WithdrawQuery) OnlyX(ctx context.Context) *Withdraw {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Withdraw ID in the query.
// Returns a *NotSingularError when more than one Withdraw ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WithdrawQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = wq.Limit(2).IDs(setContextOp(ctx, wq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{withdraw.Label}
	default:
		err = &NotSingularError{withdraw.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WithdrawQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Withdraws.
func (wq *WithdrawQuery) All(ctx context.Context) ([]*Withdraw, error) {
	ctx = setContextOp(ctx, wq.ctx, "All")
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Withdraw, *WithdrawQuery]()
	return withInterceptors[[]*Withdraw](ctx, wq, qr, wq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wq *WithdrawQuery) AllX(ctx context.Context) []*Withdraw {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Withdraw IDs.
func (wq *WithdrawQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if wq.ctx.Unique == nil && wq.path != nil {
		wq.Unique(true)
	}
	ctx = setContextOp(ctx, wq.ctx, "IDs")
	if err = wq.Select(withdraw.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WithdrawQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WithdrawQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wq.ctx, "Count")
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wq, querierCount[*WithdrawQuery](), wq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WithdrawQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WithdrawQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wq.ctx, "Exist")
	switch _, err := wq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WithdrawQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WithdrawQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WithdrawQuery) Clone() *WithdrawQuery {
	if wq == nil {
		return nil
	}
	return &WithdrawQuery{
		config:     wq.config,
		ctx:        wq.ctx.Clone(),
		order:      append([]withdraw.OrderOption{}, wq.order...),
		inters:     append([]Interceptor{}, wq.inters...),
		predicates: append([]predicate.Withdraw{}, wq.predicates...),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Withdraw.Query().
//		GroupBy(withdraw.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wq *WithdrawQuery) GroupBy(field string, fields ...string) *WithdrawGroupBy {
	wq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WithdrawGroupBy{build: wq}
	grbuild.flds = &wq.ctx.Fields
	grbuild.label = withdraw.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Withdraw.Query().
//		Select(withdraw.FieldCreatedAt).
//		Scan(ctx, &v)
func (wq *WithdrawQuery) Select(fields ...string) *WithdrawSelect {
	wq.ctx.Fields = append(wq.ctx.Fields, fields...)
	sbuild := &WithdrawSelect{WithdrawQuery: wq}
	sbuild.label = withdraw.Label
	sbuild.flds, sbuild.scan = &wq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WithdrawSelect configured with the given aggregations.
func (wq *WithdrawQuery) Aggregate(fns ...AggregateFunc) *WithdrawSelect {
	return wq.Select().Aggregate(fns...)
}

func (wq *WithdrawQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wq); err != nil {
				return err
			}
		}
	}
	for _, f := range wq.ctx.Fields {
		if !withdraw.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WithdrawQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Withdraw, error) {
	var (
		nodes = []*Withdraw{}
		_spec = wq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Withdraw).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Withdraw{config: wq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wq *WithdrawQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	_spec.Node.Columns = wq.ctx.Fields
	if len(wq.ctx.Fields) > 0 {
		_spec.Unique = wq.ctx.Unique != nil && *wq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WithdrawQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(withdraw.Table, withdraw.Columns, sqlgraph.NewFieldSpec(withdraw.FieldID, field.TypeUint64))
	_spec.From = wq.sql
	if unique := wq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wq.path != nil {
		_spec.Unique = true
	}
	if fields := wq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withdraw.FieldID)
		for i := range fields {
			if fields[i] != withdraw.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WithdrawQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(withdraw.Table)
	columns := wq.ctx.Fields
	if len(columns) == 0 {
		columns = withdraw.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.ctx.Unique != nil && *wq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithdrawGroupBy is the group-by builder for Withdraw entities.
type WithdrawGroupBy struct {
	selector
	build *WithdrawQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WithdrawGroupBy) Aggregate(fns ...AggregateFunc) *WithdrawGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the selector query and scans the result into the given value.
func (wgb *WithdrawGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wgb.build.ctx, "GroupBy")
	if err := wgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WithdrawQuery, *WithdrawGroupBy](ctx, wgb.build, wgb, wgb.build.inters, v)
}

func (wgb *WithdrawGroupBy) sqlScan(ctx context.Context, root *WithdrawQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wgb.flds)+len(wgb.fns))
		for _, f := range *wgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WithdrawSelect is the builder for selecting fields of Withdraw entities.
type WithdrawSelect struct {
	*WithdrawQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ws *WithdrawSelect) Aggregate(fns ...AggregateFunc) *WithdrawSelect {
	ws.fns = append(ws.fns, fns...)
	return ws
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WithdrawSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ws.ctx, "Select")
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WithdrawQuery, *WithdrawSelect](ctx, ws.WithdrawQuery, ws, ws.inters, v)
}

func (ws *WithdrawSelect) sqlScan(ctx context.Context, root *WithdrawQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ws.fns))
	for _, fn := range ws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
