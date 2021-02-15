// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/cleaningroom"
	"github.com/team15/app/ent/lengthtime"
	"github.com/team15/app/ent/predicate"
)

// LengthtimeQuery is the builder for querying Lengthtime entities.
type LengthtimeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Lengthtime
	// eager-loading edges.
	withCleaningrooms *CleaningroomQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (lq *LengthtimeQuery) Where(ps ...predicate.Lengthtime) *LengthtimeQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit adds a limit step to the query.
func (lq *LengthtimeQuery) Limit(limit int) *LengthtimeQuery {
	lq.limit = &limit
	return lq
}

// Offset adds an offset step to the query.
func (lq *LengthtimeQuery) Offset(offset int) *LengthtimeQuery {
	lq.offset = &offset
	return lq
}

// Order adds an order step to the query.
func (lq *LengthtimeQuery) Order(o ...OrderFunc) *LengthtimeQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QueryCleaningrooms chains the current query on the cleaningrooms edge.
func (lq *LengthtimeQuery) QueryCleaningrooms() *CleaningroomQuery {
	query := &CleaningroomQuery{config: lq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lengthtime.Table, lengthtime.FieldID, lq.sqlQuery()),
			sqlgraph.To(cleaningroom.Table, cleaningroom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, lengthtime.CleaningroomsTable, lengthtime.CleaningroomsColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Lengthtime entity in the query. Returns *NotFoundError when no lengthtime was found.
func (lq *LengthtimeQuery) First(ctx context.Context) (*Lengthtime, error) {
	ls, err := lq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ls) == 0 {
		return nil, &NotFoundError{lengthtime.Label}
	}
	return ls[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LengthtimeQuery) FirstX(ctx context.Context) *Lengthtime {
	l, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return l
}

// FirstID returns the first Lengthtime id in the query. Returns *NotFoundError when no id was found.
func (lq *LengthtimeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lengthtime.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (lq *LengthtimeQuery) FirstXID(ctx context.Context) int {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Lengthtime entity in the query, returns an error if not exactly one entity was returned.
func (lq *LengthtimeQuery) Only(ctx context.Context) (*Lengthtime, error) {
	ls, err := lq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ls) {
	case 1:
		return ls[0], nil
	case 0:
		return nil, &NotFoundError{lengthtime.Label}
	default:
		return nil, &NotSingularError{lengthtime.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LengthtimeQuery) OnlyX(ctx context.Context) *Lengthtime {
	l, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return l
}

// OnlyID returns the only Lengthtime id in the query, returns an error if not exactly one id was returned.
func (lq *LengthtimeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{lengthtime.Label}
	default:
		err = &NotSingularError{lengthtime.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LengthtimeQuery) OnlyIDX(ctx context.Context) int {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Lengthtimes.
func (lq *LengthtimeQuery) All(ctx context.Context) ([]*Lengthtime, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return lq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (lq *LengthtimeQuery) AllX(ctx context.Context) []*Lengthtime {
	ls, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ls
}

// IDs executes the query and returns a list of Lengthtime ids.
func (lq *LengthtimeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := lq.Select(lengthtime.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LengthtimeQuery) IDsX(ctx context.Context) []int {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LengthtimeQuery) Count(ctx context.Context) (int, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return lq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LengthtimeQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LengthtimeQuery) Exist(ctx context.Context) (bool, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return lq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LengthtimeQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LengthtimeQuery) Clone() *LengthtimeQuery {
	return &LengthtimeQuery{
		config:     lq.config,
		limit:      lq.limit,
		offset:     lq.offset,
		order:      append([]OrderFunc{}, lq.order...),
		unique:     append([]string{}, lq.unique...),
		predicates: append([]predicate.Lengthtime{}, lq.predicates...),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
}

//  WithCleaningrooms tells the query-builder to eager-loads the nodes that are connected to
// the "cleaningrooms" edge. The optional arguments used to configure the query builder of the edge.
func (lq *LengthtimeQuery) WithCleaningrooms(opts ...func(*CleaningroomQuery)) *LengthtimeQuery {
	query := &CleaningroomQuery{config: lq.config}
	for _, opt := range opts {
		opt(query)
	}
	lq.withCleaningrooms = query
	return lq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Lengthtime string `json:"lengthtime,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Lengthtime.Query().
//		GroupBy(lengthtime.FieldLengthtime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (lq *LengthtimeQuery) GroupBy(field string, fields ...string) *LengthtimeGroupBy {
	group := &LengthtimeGroupBy{config: lq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Lengthtime string `json:"lengthtime,omitempty"`
//	}
//
//	client.Lengthtime.Query().
//		Select(lengthtime.FieldLengthtime).
//		Scan(ctx, &v)
//
func (lq *LengthtimeQuery) Select(field string, fields ...string) *LengthtimeSelect {
	selector := &LengthtimeSelect{config: lq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lq.sqlQuery(), nil
	}
	return selector
}

func (lq *LengthtimeQuery) prepareQuery(ctx context.Context) error {
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LengthtimeQuery) sqlAll(ctx context.Context) ([]*Lengthtime, error) {
	var (
		nodes       = []*Lengthtime{}
		_spec       = lq.querySpec()
		loadedTypes = [1]bool{
			lq.withCleaningrooms != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Lengthtime{config: lq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := lq.withCleaningrooms; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Lengthtime)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Cleaningroom(func(s *sql.Selector) {
			s.Where(sql.InValues(lengthtime.CleaningroomsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.lengthtime_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "lengthtime_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "lengthtime_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Cleaningrooms = append(node.Edges.Cleaningrooms, n)
		}
	}

	return nodes, nil
}

func (lq *LengthtimeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LengthtimeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := lq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (lq *LengthtimeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lengthtime.Table,
			Columns: lengthtime.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lengthtime.FieldID,
			},
		},
		From:   lq.sql,
		Unique: true,
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LengthtimeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(lengthtime.Table)
	selector := builder.Select(t1.Columns(lengthtime.Columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(lengthtime.Columns...)...)
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LengthtimeGroupBy is the builder for group-by Lengthtime entities.
type LengthtimeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LengthtimeGroupBy) Aggregate(fns ...AggregateFunc) *LengthtimeGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the group-by query and scan the result into the given value.
func (lgb *LengthtimeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := lgb.path(ctx)
	if err != nil {
		return err
	}
	lgb.sql = query
	return lgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lgb *LengthtimeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := lgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (lgb *LengthtimeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(lgb.fields) > 1 {
		return nil, errors.New("ent: LengthtimeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := lgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lgb *LengthtimeGroupBy) StringsX(ctx context.Context) []string {
	v, err := lgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (lgb *LengthtimeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lengthtime.Label}
	default:
		err = fmt.Errorf("ent: LengthtimeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lgb *LengthtimeGroupBy) StringX(ctx context.Context) string {
	v, err := lgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (lgb *LengthtimeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(lgb.fields) > 1 {
		return nil, errors.New("ent: LengthtimeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := lgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lgb *LengthtimeGroupBy) IntsX(ctx context.Context) []int {
	v, err := lgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (lgb *LengthtimeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lengthtime.Label}
	default:
		err = fmt.Errorf("ent: LengthtimeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lgb *LengthtimeGroupBy) IntX(ctx context.Context) int {
	v, err := lgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (lgb *LengthtimeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(lgb.fields) > 1 {
		return nil, errors.New("ent: LengthtimeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := lgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lgb *LengthtimeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := lgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (lgb *LengthtimeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lengthtime.Label}
	default:
		err = fmt.Errorf("ent: LengthtimeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lgb *LengthtimeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := lgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (lgb *LengthtimeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(lgb.fields) > 1 {
		return nil, errors.New("ent: LengthtimeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := lgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lgb *LengthtimeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := lgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (lgb *LengthtimeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lengthtime.Label}
	default:
		err = fmt.Errorf("ent: LengthtimeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lgb *LengthtimeGroupBy) BoolX(ctx context.Context) bool {
	v, err := lgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lgb *LengthtimeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := lgb.sqlQuery().Query()
	if err := lgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lgb *LengthtimeGroupBy) sqlQuery() *sql.Selector {
	selector := lgb.sql
	columns := make([]string, 0, len(lgb.fields)+len(lgb.fns))
	columns = append(columns, lgb.fields...)
	for _, fn := range lgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(lgb.fields...)
}

// LengthtimeSelect is the builder for select fields of Lengthtime entities.
type LengthtimeSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ls *LengthtimeSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ls.path(ctx)
	if err != nil {
		return err
	}
	ls.sql = query
	return ls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ls *LengthtimeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ls *LengthtimeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ls.fields) > 1 {
		return nil, errors.New("ent: LengthtimeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ls *LengthtimeSelect) StringsX(ctx context.Context) []string {
	v, err := ls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ls *LengthtimeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lengthtime.Label}
	default:
		err = fmt.Errorf("ent: LengthtimeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ls *LengthtimeSelect) StringX(ctx context.Context) string {
	v, err := ls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ls *LengthtimeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ls.fields) > 1 {
		return nil, errors.New("ent: LengthtimeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ls *LengthtimeSelect) IntsX(ctx context.Context) []int {
	v, err := ls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ls *LengthtimeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lengthtime.Label}
	default:
		err = fmt.Errorf("ent: LengthtimeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ls *LengthtimeSelect) IntX(ctx context.Context) int {
	v, err := ls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ls *LengthtimeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ls.fields) > 1 {
		return nil, errors.New("ent: LengthtimeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ls *LengthtimeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ls *LengthtimeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lengthtime.Label}
	default:
		err = fmt.Errorf("ent: LengthtimeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ls *LengthtimeSelect) Float64X(ctx context.Context) float64 {
	v, err := ls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ls *LengthtimeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ls.fields) > 1 {
		return nil, errors.New("ent: LengthtimeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ls *LengthtimeSelect) BoolsX(ctx context.Context) []bool {
	v, err := ls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ls *LengthtimeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lengthtime.Label}
	default:
		err = fmt.Errorf("ent: LengthtimeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ls *LengthtimeSelect) BoolX(ctx context.Context) bool {
	v, err := ls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ls *LengthtimeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ls.sqlQuery().Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ls *LengthtimeSelect) sqlQuery() sql.Querier {
	selector := ls.sql
	selector.Select(selector.Columns(ls.fields...)...)
	return selector
}
