// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/cleanername"
	"github.com/team15/app/ent/cleaningroom"
	"github.com/team15/app/ent/lengthtime"
	"github.com/team15/app/ent/predicate"
	"github.com/team15/app/ent/room"
)

// CleaningRoomQuery is the builder for querying CleaningRoom entities.
type CleaningRoomQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.CleaningRoom
	// eager-loading edges.
	withRoom        *RoomQuery
	withCleanerName *CleanerNameQuery
	withLengthTime  *LengthTimeQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (crq *CleaningRoomQuery) Where(ps ...predicate.CleaningRoom) *CleaningRoomQuery {
	crq.predicates = append(crq.predicates, ps...)
	return crq
}

// Limit adds a limit step to the query.
func (crq *CleaningRoomQuery) Limit(limit int) *CleaningRoomQuery {
	crq.limit = &limit
	return crq
}

// Offset adds an offset step to the query.
func (crq *CleaningRoomQuery) Offset(offset int) *CleaningRoomQuery {
	crq.offset = &offset
	return crq
}

// Order adds an order step to the query.
func (crq *CleaningRoomQuery) Order(o ...OrderFunc) *CleaningRoomQuery {
	crq.order = append(crq.order, o...)
	return crq
}

// QueryRoom chains the current query on the Room edge.
func (crq *CleaningRoomQuery) QueryRoom() *RoomQuery {
	query := &RoomQuery{config: crq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cleaningroom.Table, cleaningroom.FieldID, crq.sqlQuery()),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cleaningroom.RoomTable, cleaningroom.RoomColumn),
		)
		fromU = sqlgraph.SetNeighbors(crq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCleanerName chains the current query on the CleanerName edge.
func (crq *CleaningRoomQuery) QueryCleanerName() *CleanerNameQuery {
	query := &CleanerNameQuery{config: crq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cleaningroom.Table, cleaningroom.FieldID, crq.sqlQuery()),
			sqlgraph.To(cleanername.Table, cleanername.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cleaningroom.CleanerNameTable, cleaningroom.CleanerNameColumn),
		)
		fromU = sqlgraph.SetNeighbors(crq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLengthTime chains the current query on the LengthTime edge.
func (crq *CleaningRoomQuery) QueryLengthTime() *LengthTimeQuery {
	query := &LengthTimeQuery{config: crq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cleaningroom.Table, cleaningroom.FieldID, crq.sqlQuery()),
			sqlgraph.To(lengthtime.Table, lengthtime.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cleaningroom.LengthTimeTable, cleaningroom.LengthTimeColumn),
		)
		fromU = sqlgraph.SetNeighbors(crq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CleaningRoom entity in the query. Returns *NotFoundError when no cleaningroom was found.
func (crq *CleaningRoomQuery) First(ctx context.Context) (*CleaningRoom, error) {
	crs, err := crq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(crs) == 0 {
		return nil, &NotFoundError{cleaningroom.Label}
	}
	return crs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (crq *CleaningRoomQuery) FirstX(ctx context.Context) *CleaningRoom {
	cr, err := crq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return cr
}

// FirstID returns the first CleaningRoom id in the query. Returns *NotFoundError when no id was found.
func (crq *CleaningRoomQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cleaningroom.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (crq *CleaningRoomQuery) FirstXID(ctx context.Context) int {
	id, err := crq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only CleaningRoom entity in the query, returns an error if not exactly one entity was returned.
func (crq *CleaningRoomQuery) Only(ctx context.Context) (*CleaningRoom, error) {
	crs, err := crq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(crs) {
	case 1:
		return crs[0], nil
	case 0:
		return nil, &NotFoundError{cleaningroom.Label}
	default:
		return nil, &NotSingularError{cleaningroom.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (crq *CleaningRoomQuery) OnlyX(ctx context.Context) *CleaningRoom {
	cr, err := crq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return cr
}

// OnlyID returns the only CleaningRoom id in the query, returns an error if not exactly one id was returned.
func (crq *CleaningRoomQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cleaningroom.Label}
	default:
		err = &NotSingularError{cleaningroom.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (crq *CleaningRoomQuery) OnlyIDX(ctx context.Context) int {
	id, err := crq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CleaningRooms.
func (crq *CleaningRoomQuery) All(ctx context.Context) ([]*CleaningRoom, error) {
	if err := crq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return crq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (crq *CleaningRoomQuery) AllX(ctx context.Context) []*CleaningRoom {
	crs, err := crq.All(ctx)
	if err != nil {
		panic(err)
	}
	return crs
}

// IDs executes the query and returns a list of CleaningRoom ids.
func (crq *CleaningRoomQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := crq.Select(cleaningroom.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (crq *CleaningRoomQuery) IDsX(ctx context.Context) []int {
	ids, err := crq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (crq *CleaningRoomQuery) Count(ctx context.Context) (int, error) {
	if err := crq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return crq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (crq *CleaningRoomQuery) CountX(ctx context.Context) int {
	count, err := crq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (crq *CleaningRoomQuery) Exist(ctx context.Context) (bool, error) {
	if err := crq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return crq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (crq *CleaningRoomQuery) ExistX(ctx context.Context) bool {
	exist, err := crq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (crq *CleaningRoomQuery) Clone() *CleaningRoomQuery {
	return &CleaningRoomQuery{
		config:     crq.config,
		limit:      crq.limit,
		offset:     crq.offset,
		order:      append([]OrderFunc{}, crq.order...),
		unique:     append([]string{}, crq.unique...),
		predicates: append([]predicate.CleaningRoom{}, crq.predicates...),
		// clone intermediate query.
		sql:  crq.sql.Clone(),
		path: crq.path,
	}
}

//  WithRoom tells the query-builder to eager-loads the nodes that are connected to
// the "Room" edge. The optional arguments used to configure the query builder of the edge.
func (crq *CleaningRoomQuery) WithRoom(opts ...func(*RoomQuery)) *CleaningRoomQuery {
	query := &RoomQuery{config: crq.config}
	for _, opt := range opts {
		opt(query)
	}
	crq.withRoom = query
	return crq
}

//  WithCleanerName tells the query-builder to eager-loads the nodes that are connected to
// the "CleanerName" edge. The optional arguments used to configure the query builder of the edge.
func (crq *CleaningRoomQuery) WithCleanerName(opts ...func(*CleanerNameQuery)) *CleaningRoomQuery {
	query := &CleanerNameQuery{config: crq.config}
	for _, opt := range opts {
		opt(query)
	}
	crq.withCleanerName = query
	return crq
}

//  WithLengthTime tells the query-builder to eager-loads the nodes that are connected to
// the "LengthTime" edge. The optional arguments used to configure the query builder of the edge.
func (crq *CleaningRoomQuery) WithLengthTime(opts ...func(*LengthTimeQuery)) *CleaningRoomQuery {
	query := &LengthTimeQuery{config: crq.config}
	for _, opt := range opts {
		opt(query)
	}
	crq.withLengthTime = query
	return crq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Dateandstarttime time.Time `json:"dateandstarttime,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CleaningRoom.Query().
//		GroupBy(cleaningroom.FieldDateandstarttime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (crq *CleaningRoomQuery) GroupBy(field string, fields ...string) *CleaningRoomGroupBy {
	group := &CleaningRoomGroupBy{config: crq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return crq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Dateandstarttime time.Time `json:"dateandstarttime,omitempty"`
//	}
//
//	client.CleaningRoom.Query().
//		Select(cleaningroom.FieldDateandstarttime).
//		Scan(ctx, &v)
//
func (crq *CleaningRoomQuery) Select(field string, fields ...string) *CleaningRoomSelect {
	selector := &CleaningRoomSelect{config: crq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return crq.sqlQuery(), nil
	}
	return selector
}

func (crq *CleaningRoomQuery) prepareQuery(ctx context.Context) error {
	if crq.path != nil {
		prev, err := crq.path(ctx)
		if err != nil {
			return err
		}
		crq.sql = prev
	}
	return nil
}

func (crq *CleaningRoomQuery) sqlAll(ctx context.Context) ([]*CleaningRoom, error) {
	var (
		nodes       = []*CleaningRoom{}
		withFKs     = crq.withFKs
		_spec       = crq.querySpec()
		loadedTypes = [3]bool{
			crq.withRoom != nil,
			crq.withCleanerName != nil,
			crq.withLengthTime != nil,
		}
	)
	if crq.withRoom != nil || crq.withCleanerName != nil || crq.withLengthTime != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, cleaningroom.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &CleaningRoom{config: crq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
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
	if err := sqlgraph.QueryNodes(ctx, crq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := crq.withRoom; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CleaningRoom)
		for i := range nodes {
			if fk := nodes[i].room_cleaningrooms; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(room.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "room_cleaningrooms" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Room = n
			}
		}
	}

	if query := crq.withCleanerName; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CleaningRoom)
		for i := range nodes {
			if fk := nodes[i].cleanerroom_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(cleanername.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "cleanerroom_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.CleanerName = n
			}
		}
	}

	if query := crq.withLengthTime; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CleaningRoom)
		for i := range nodes {
			if fk := nodes[i].lengthtime_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(lengthtime.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "lengthtime_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.LengthTime = n
			}
		}
	}

	return nodes, nil
}

func (crq *CleaningRoomQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := crq.querySpec()
	return sqlgraph.CountNodes(ctx, crq.driver, _spec)
}

func (crq *CleaningRoomQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := crq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (crq *CleaningRoomQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cleaningroom.Table,
			Columns: cleaningroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cleaningroom.FieldID,
			},
		},
		From:   crq.sql,
		Unique: true,
	}
	if ps := crq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := crq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := crq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := crq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (crq *CleaningRoomQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(crq.driver.Dialect())
	t1 := builder.Table(cleaningroom.Table)
	selector := builder.Select(t1.Columns(cleaningroom.Columns...)...).From(t1)
	if crq.sql != nil {
		selector = crq.sql
		selector.Select(selector.Columns(cleaningroom.Columns...)...)
	}
	for _, p := range crq.predicates {
		p(selector)
	}
	for _, p := range crq.order {
		p(selector)
	}
	if offset := crq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := crq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CleaningRoomGroupBy is the builder for group-by CleaningRoom entities.
type CleaningRoomGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (crgb *CleaningRoomGroupBy) Aggregate(fns ...AggregateFunc) *CleaningRoomGroupBy {
	crgb.fns = append(crgb.fns, fns...)
	return crgb
}

// Scan applies the group-by query and scan the result into the given value.
func (crgb *CleaningRoomGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := crgb.path(ctx)
	if err != nil {
		return err
	}
	crgb.sql = query
	return crgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (crgb *CleaningRoomGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := crgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (crgb *CleaningRoomGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(crgb.fields) > 1 {
		return nil, errors.New("ent: CleaningRoomGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := crgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (crgb *CleaningRoomGroupBy) StringsX(ctx context.Context) []string {
	v, err := crgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (crgb *CleaningRoomGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = crgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cleaningroom.Label}
	default:
		err = fmt.Errorf("ent: CleaningRoomGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (crgb *CleaningRoomGroupBy) StringX(ctx context.Context) string {
	v, err := crgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (crgb *CleaningRoomGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(crgb.fields) > 1 {
		return nil, errors.New("ent: CleaningRoomGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := crgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (crgb *CleaningRoomGroupBy) IntsX(ctx context.Context) []int {
	v, err := crgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (crgb *CleaningRoomGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = crgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cleaningroom.Label}
	default:
		err = fmt.Errorf("ent: CleaningRoomGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (crgb *CleaningRoomGroupBy) IntX(ctx context.Context) int {
	v, err := crgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (crgb *CleaningRoomGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(crgb.fields) > 1 {
		return nil, errors.New("ent: CleaningRoomGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := crgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (crgb *CleaningRoomGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := crgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (crgb *CleaningRoomGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = crgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cleaningroom.Label}
	default:
		err = fmt.Errorf("ent: CleaningRoomGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (crgb *CleaningRoomGroupBy) Float64X(ctx context.Context) float64 {
	v, err := crgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (crgb *CleaningRoomGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(crgb.fields) > 1 {
		return nil, errors.New("ent: CleaningRoomGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := crgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (crgb *CleaningRoomGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := crgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (crgb *CleaningRoomGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = crgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cleaningroom.Label}
	default:
		err = fmt.Errorf("ent: CleaningRoomGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (crgb *CleaningRoomGroupBy) BoolX(ctx context.Context) bool {
	v, err := crgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (crgb *CleaningRoomGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := crgb.sqlQuery().Query()
	if err := crgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (crgb *CleaningRoomGroupBy) sqlQuery() *sql.Selector {
	selector := crgb.sql
	columns := make([]string, 0, len(crgb.fields)+len(crgb.fns))
	columns = append(columns, crgb.fields...)
	for _, fn := range crgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(crgb.fields...)
}

// CleaningRoomSelect is the builder for select fields of CleaningRoom entities.
type CleaningRoomSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (crs *CleaningRoomSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := crs.path(ctx)
	if err != nil {
		return err
	}
	crs.sql = query
	return crs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (crs *CleaningRoomSelect) ScanX(ctx context.Context, v interface{}) {
	if err := crs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (crs *CleaningRoomSelect) Strings(ctx context.Context) ([]string, error) {
	if len(crs.fields) > 1 {
		return nil, errors.New("ent: CleaningRoomSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := crs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (crs *CleaningRoomSelect) StringsX(ctx context.Context) []string {
	v, err := crs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (crs *CleaningRoomSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = crs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cleaningroom.Label}
	default:
		err = fmt.Errorf("ent: CleaningRoomSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (crs *CleaningRoomSelect) StringX(ctx context.Context) string {
	v, err := crs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (crs *CleaningRoomSelect) Ints(ctx context.Context) ([]int, error) {
	if len(crs.fields) > 1 {
		return nil, errors.New("ent: CleaningRoomSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := crs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (crs *CleaningRoomSelect) IntsX(ctx context.Context) []int {
	v, err := crs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (crs *CleaningRoomSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = crs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cleaningroom.Label}
	default:
		err = fmt.Errorf("ent: CleaningRoomSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (crs *CleaningRoomSelect) IntX(ctx context.Context) int {
	v, err := crs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (crs *CleaningRoomSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(crs.fields) > 1 {
		return nil, errors.New("ent: CleaningRoomSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := crs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (crs *CleaningRoomSelect) Float64sX(ctx context.Context) []float64 {
	v, err := crs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (crs *CleaningRoomSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = crs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cleaningroom.Label}
	default:
		err = fmt.Errorf("ent: CleaningRoomSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (crs *CleaningRoomSelect) Float64X(ctx context.Context) float64 {
	v, err := crs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (crs *CleaningRoomSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(crs.fields) > 1 {
		return nil, errors.New("ent: CleaningRoomSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := crs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (crs *CleaningRoomSelect) BoolsX(ctx context.Context) []bool {
	v, err := crs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (crs *CleaningRoomSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = crs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cleaningroom.Label}
	default:
		err = fmt.Errorf("ent: CleaningRoomSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (crs *CleaningRoomSelect) BoolX(ctx context.Context) bool {
	v, err := crs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (crs *CleaningRoomSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := crs.sqlQuery().Query()
	if err := crs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (crs *CleaningRoomSelect) sqlQuery() sql.Querier {
	selector := crs.sql
	selector.Select(selector.Columns(crs.fields...)...)
	return selector
}