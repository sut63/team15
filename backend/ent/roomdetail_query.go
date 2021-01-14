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
	"github.com/team15/app/ent/bedtype"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/lease"
	"github.com/team15/app/ent/petrule"
	"github.com/team15/app/ent/pledge"
	"github.com/team15/app/ent/predicate"
	"github.com/team15/app/ent/roomdetail"
	"github.com/team15/app/ent/staytype"
)

// RoomdetailQuery is the builder for querying Roomdetail entities.
type RoomdetailQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Roomdetail
	// eager-loading edges.
	withPledge      *PledgeQuery
	withPetrule     *PetruleQuery
	withBedtype     *BedtypeQuery
	withEmployee    *EmployeeQuery
	withStaytype    *StaytypeQuery
	withRoomdetails *LeaseQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (rq *RoomdetailQuery) Where(ps ...predicate.Roomdetail) *RoomdetailQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit adds a limit step to the query.
func (rq *RoomdetailQuery) Limit(limit int) *RoomdetailQuery {
	rq.limit = &limit
	return rq
}

// Offset adds an offset step to the query.
func (rq *RoomdetailQuery) Offset(offset int) *RoomdetailQuery {
	rq.offset = &offset
	return rq
}

// Order adds an order step to the query.
func (rq *RoomdetailQuery) Order(o ...OrderFunc) *RoomdetailQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryPledge chains the current query on the pledge edge.
func (rq *RoomdetailQuery) QueryPledge() *PledgeQuery {
	query := &PledgeQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roomdetail.Table, roomdetail.FieldID, rq.sqlQuery()),
			sqlgraph.To(pledge.Table, pledge.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, roomdetail.PledgeTable, roomdetail.PledgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPetrule chains the current query on the petrule edge.
func (rq *RoomdetailQuery) QueryPetrule() *PetruleQuery {
	query := &PetruleQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roomdetail.Table, roomdetail.FieldID, rq.sqlQuery()),
			sqlgraph.To(petrule.Table, petrule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, roomdetail.PetruleTable, roomdetail.PetruleColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBedtype chains the current query on the bedtype edge.
func (rq *RoomdetailQuery) QueryBedtype() *BedtypeQuery {
	query := &BedtypeQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roomdetail.Table, roomdetail.FieldID, rq.sqlQuery()),
			sqlgraph.To(bedtype.Table, bedtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, roomdetail.BedtypeTable, roomdetail.BedtypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEmployee chains the current query on the employee edge.
func (rq *RoomdetailQuery) QueryEmployee() *EmployeeQuery {
	query := &EmployeeQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roomdetail.Table, roomdetail.FieldID, rq.sqlQuery()),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, roomdetail.EmployeeTable, roomdetail.EmployeeColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStaytype chains the current query on the staytype edge.
func (rq *RoomdetailQuery) QueryStaytype() *StaytypeQuery {
	query := &StaytypeQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roomdetail.Table, roomdetail.FieldID, rq.sqlQuery()),
			sqlgraph.To(staytype.Table, staytype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, roomdetail.StaytypeTable, roomdetail.StaytypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoomdetails chains the current query on the roomdetails edge.
func (rq *RoomdetailQuery) QueryRoomdetails() *LeaseQuery {
	query := &LeaseQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roomdetail.Table, roomdetail.FieldID, rq.sqlQuery()),
			sqlgraph.To(lease.Table, lease.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, roomdetail.RoomdetailsTable, roomdetail.RoomdetailsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Roomdetail entity in the query. Returns *NotFoundError when no roomdetail was found.
func (rq *RoomdetailQuery) First(ctx context.Context) (*Roomdetail, error) {
	rs, err := rq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(rs) == 0 {
		return nil, &NotFoundError{roomdetail.Label}
	}
	return rs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RoomdetailQuery) FirstX(ctx context.Context) *Roomdetail {
	r, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return r
}

// FirstID returns the first Roomdetail id in the query. Returns *NotFoundError when no id was found.
func (rq *RoomdetailQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{roomdetail.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (rq *RoomdetailQuery) FirstXID(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Roomdetail entity in the query, returns an error if not exactly one entity was returned.
func (rq *RoomdetailQuery) Only(ctx context.Context) (*Roomdetail, error) {
	rs, err := rq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(rs) {
	case 1:
		return rs[0], nil
	case 0:
		return nil, &NotFoundError{roomdetail.Label}
	default:
		return nil, &NotSingularError{roomdetail.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RoomdetailQuery) OnlyX(ctx context.Context) *Roomdetail {
	r, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// OnlyID returns the only Roomdetail id in the query, returns an error if not exactly one id was returned.
func (rq *RoomdetailQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{roomdetail.Label}
	default:
		err = &NotSingularError{roomdetail.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RoomdetailQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Roomdetails.
func (rq *RoomdetailQuery) All(ctx context.Context) ([]*Roomdetail, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rq *RoomdetailQuery) AllX(ctx context.Context) []*Roomdetail {
	rs, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return rs
}

// IDs executes the query and returns a list of Roomdetail ids.
func (rq *RoomdetailQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rq.Select(roomdetail.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RoomdetailQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RoomdetailQuery) Count(ctx context.Context) (int, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RoomdetailQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RoomdetailQuery) Exist(ctx context.Context) (bool, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RoomdetailQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RoomdetailQuery) Clone() *RoomdetailQuery {
	return &RoomdetailQuery{
		config:     rq.config,
		limit:      rq.limit,
		offset:     rq.offset,
		order:      append([]OrderFunc{}, rq.order...),
		unique:     append([]string{}, rq.unique...),
		predicates: append([]predicate.Roomdetail{}, rq.predicates...),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

//  WithPledge tells the query-builder to eager-loads the nodes that are connected to
// the "pledge" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RoomdetailQuery) WithPledge(opts ...func(*PledgeQuery)) *RoomdetailQuery {
	query := &PledgeQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withPledge = query
	return rq
}

//  WithPetrule tells the query-builder to eager-loads the nodes that are connected to
// the "petrule" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RoomdetailQuery) WithPetrule(opts ...func(*PetruleQuery)) *RoomdetailQuery {
	query := &PetruleQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withPetrule = query
	return rq
}

//  WithBedtype tells the query-builder to eager-loads the nodes that are connected to
// the "bedtype" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RoomdetailQuery) WithBedtype(opts ...func(*BedtypeQuery)) *RoomdetailQuery {
	query := &BedtypeQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withBedtype = query
	return rq
}

//  WithEmployee tells the query-builder to eager-loads the nodes that are connected to
// the "employee" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RoomdetailQuery) WithEmployee(opts ...func(*EmployeeQuery)) *RoomdetailQuery {
	query := &EmployeeQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withEmployee = query
	return rq
}

//  WithStaytype tells the query-builder to eager-loads the nodes that are connected to
// the "staytype" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RoomdetailQuery) WithStaytype(opts ...func(*StaytypeQuery)) *RoomdetailQuery {
	query := &StaytypeQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withStaytype = query
	return rq
}

//  WithRoomdetails tells the query-builder to eager-loads the nodes that are connected to
// the "roomdetails" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RoomdetailQuery) WithRoomdetails(opts ...func(*LeaseQuery)) *RoomdetailQuery {
	query := &LeaseQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withRoomdetails = query
	return rq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Roomnumber string `json:"roomnumber,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Roomdetail.Query().
//		GroupBy(roomdetail.FieldRoomnumber).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rq *RoomdetailQuery) GroupBy(field string, fields ...string) *RoomdetailGroupBy {
	group := &RoomdetailGroupBy{config: rq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Roomnumber string `json:"roomnumber,omitempty"`
//	}
//
//	client.Roomdetail.Query().
//		Select(roomdetail.FieldRoomnumber).
//		Scan(ctx, &v)
//
func (rq *RoomdetailQuery) Select(field string, fields ...string) *RoomdetailSelect {
	selector := &RoomdetailSelect{config: rq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(), nil
	}
	return selector
}

func (rq *RoomdetailQuery) prepareQuery(ctx context.Context) error {
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RoomdetailQuery) sqlAll(ctx context.Context) ([]*Roomdetail, error) {
	var (
		nodes       = []*Roomdetail{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [6]bool{
			rq.withPledge != nil,
			rq.withPetrule != nil,
			rq.withBedtype != nil,
			rq.withEmployee != nil,
			rq.withStaytype != nil,
			rq.withRoomdetails != nil,
		}
	)
	if rq.withPledge != nil || rq.withPetrule != nil || rq.withBedtype != nil || rq.withEmployee != nil || rq.withStaytype != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, roomdetail.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Roomdetail{config: rq.config}
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
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rq.withPledge; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Roomdetail)
		for i := range nodes {
			if fk := nodes[i].pledge_roomdetails; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(pledge.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "pledge_roomdetails" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Pledge = n
			}
		}
	}

	if query := rq.withPetrule; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Roomdetail)
		for i := range nodes {
			if fk := nodes[i].petrule_roomdetails; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(petrule.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "petrule_roomdetails" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Petrule = n
			}
		}
	}

	if query := rq.withBedtype; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Roomdetail)
		for i := range nodes {
			if fk := nodes[i].bedtype_roomdetails; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(bedtype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "bedtype_roomdetails" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Bedtype = n
			}
		}
	}

	if query := rq.withEmployee; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Roomdetail)
		for i := range nodes {
			if fk := nodes[i].employee_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(employee.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "employee_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Employee = n
			}
		}
	}

	if query := rq.withStaytype; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Roomdetail)
		for i := range nodes {
			if fk := nodes[i].staytype_roomdetails; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(staytype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "staytype_roomdetails" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Staytype = n
			}
		}
	}

	if query := rq.withRoomdetails; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Roomdetail)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Lease(func(s *sql.Selector) {
			s.Where(sql.InValues(roomdetail.RoomdetailsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.room_num
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "room_num" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "room_num" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Roomdetails = n
		}
	}

	return nodes, nil
}

func (rq *RoomdetailQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RoomdetailQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rq *RoomdetailQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roomdetail.Table,
			Columns: roomdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomdetail.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RoomdetailQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(roomdetail.Table)
	selector := builder.Select(t1.Columns(roomdetail.Columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(roomdetail.Columns...)...)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoomdetailGroupBy is the builder for group-by Roomdetail entities.
type RoomdetailGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RoomdetailGroupBy) Aggregate(fns ...AggregateFunc) *RoomdetailGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the group-by query and scan the result into the given value.
func (rgb *RoomdetailGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rgb.path(ctx)
	if err != nil {
		return err
	}
	rgb.sql = query
	return rgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rgb *RoomdetailGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomdetailGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RoomdetailGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rgb *RoomdetailGroupBy) StringsX(ctx context.Context) []string {
	v, err := rgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomdetailGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomdetail.Label}
	default:
		err = fmt.Errorf("ent: RoomdetailGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rgb *RoomdetailGroupBy) StringX(ctx context.Context) string {
	v, err := rgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomdetailGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RoomdetailGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rgb *RoomdetailGroupBy) IntsX(ctx context.Context) []int {
	v, err := rgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomdetailGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomdetail.Label}
	default:
		err = fmt.Errorf("ent: RoomdetailGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rgb *RoomdetailGroupBy) IntX(ctx context.Context) int {
	v, err := rgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomdetailGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RoomdetailGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rgb *RoomdetailGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomdetailGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomdetail.Label}
	default:
		err = fmt.Errorf("ent: RoomdetailGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rgb *RoomdetailGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomdetailGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RoomdetailGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rgb *RoomdetailGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (rgb *RoomdetailGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomdetail.Label}
	default:
		err = fmt.Errorf("ent: RoomdetailGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rgb *RoomdetailGroupBy) BoolX(ctx context.Context) bool {
	v, err := rgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rgb *RoomdetailGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rgb.sqlQuery().Query()
	if err := rgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rgb *RoomdetailGroupBy) sqlQuery() *sql.Selector {
	selector := rgb.sql
	columns := make([]string, 0, len(rgb.fields)+len(rgb.fns))
	columns = append(columns, rgb.fields...)
	for _, fn := range rgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(rgb.fields...)
}

// RoomdetailSelect is the builder for select fields of Roomdetail entities.
type RoomdetailSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (rs *RoomdetailSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := rs.path(ctx)
	if err != nil {
		return err
	}
	rs.sql = query
	return rs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rs *RoomdetailSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (rs *RoomdetailSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RoomdetailSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rs *RoomdetailSelect) StringsX(ctx context.Context) []string {
	v, err := rs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (rs *RoomdetailSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomdetail.Label}
	default:
		err = fmt.Errorf("ent: RoomdetailSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rs *RoomdetailSelect) StringX(ctx context.Context) string {
	v, err := rs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (rs *RoomdetailSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RoomdetailSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rs *RoomdetailSelect) IntsX(ctx context.Context) []int {
	v, err := rs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (rs *RoomdetailSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomdetail.Label}
	default:
		err = fmt.Errorf("ent: RoomdetailSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rs *RoomdetailSelect) IntX(ctx context.Context) int {
	v, err := rs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (rs *RoomdetailSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RoomdetailSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rs *RoomdetailSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (rs *RoomdetailSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomdetail.Label}
	default:
		err = fmt.Errorf("ent: RoomdetailSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rs *RoomdetailSelect) Float64X(ctx context.Context) float64 {
	v, err := rs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (rs *RoomdetailSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RoomdetailSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rs *RoomdetailSelect) BoolsX(ctx context.Context) []bool {
	v, err := rs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (rs *RoomdetailSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomdetail.Label}
	default:
		err = fmt.Errorf("ent: RoomdetailSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rs *RoomdetailSelect) BoolX(ctx context.Context) bool {
	v, err := rs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rs *RoomdetailSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rs.sqlQuery().Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rs *RoomdetailSelect) sqlQuery() sql.Querier {
	selector := rs.sql
	selector.Select(selector.Columns(rs.fields...)...)
	return selector
}
