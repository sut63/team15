// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/deposit"
	"github.com/team15/app/ent/predicate"
	"github.com/team15/app/ent/statusd"
)

// StatusdUpdate is the builder for updating Statusd entities.
type StatusdUpdate struct {
	config
	hooks      []Hook
	mutation   *StatusdMutation
	predicates []predicate.Statusd
}

// Where adds a new predicate for the builder.
func (su *StatusdUpdate) Where(ps ...predicate.Statusd) *StatusdUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetStatusdname sets the statusdname field.
func (su *StatusdUpdate) SetStatusdname(s string) *StatusdUpdate {
	su.mutation.SetStatusdname(s)
	return su
}

// AddStatusdIDs adds the statusds edge to Deposit by ids.
func (su *StatusdUpdate) AddStatusdIDs(ids ...int) *StatusdUpdate {
	su.mutation.AddStatusdIDs(ids...)
	return su
}

// AddStatusds adds the statusds edges to Deposit.
func (su *StatusdUpdate) AddStatusds(d ...*Deposit) *StatusdUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return su.AddStatusdIDs(ids...)
}

// Mutation returns the StatusdMutation object of the builder.
func (su *StatusdUpdate) Mutation() *StatusdMutation {
	return su.mutation
}

// RemoveStatusdIDs removes the statusds edge to Deposit by ids.
func (su *StatusdUpdate) RemoveStatusdIDs(ids ...int) *StatusdUpdate {
	su.mutation.RemoveStatusdIDs(ids...)
	return su
}

// RemoveStatusds removes statusds edges to Deposit.
func (su *StatusdUpdate) RemoveStatusds(d ...*Deposit) *StatusdUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return su.RemoveStatusdIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *StatusdUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := su.mutation.Statusdname(); ok {
		if err := statusd.StatusdnameValidator(v); err != nil {
			return 0, &ValidationError{Name: "statusdname", err: fmt.Errorf("ent: validator failed for field \"statusdname\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusdMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *StatusdUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StatusdUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StatusdUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StatusdUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statusd.Table,
			Columns: statusd.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statusd.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Statusdname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: statusd.FieldStatusdname,
		})
	}
	if nodes := su.mutation.RemovedStatusdsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statusd.StatusdsTable,
			Columns: []string{statusd.StatusdsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deposit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.StatusdsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statusd.StatusdsTable,
			Columns: []string{statusd.StatusdsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deposit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statusd.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// StatusdUpdateOne is the builder for updating a single Statusd entity.
type StatusdUpdateOne struct {
	config
	hooks    []Hook
	mutation *StatusdMutation
}

// SetStatusdname sets the statusdname field.
func (suo *StatusdUpdateOne) SetStatusdname(s string) *StatusdUpdateOne {
	suo.mutation.SetStatusdname(s)
	return suo
}

// AddStatusdIDs adds the statusds edge to Deposit by ids.
func (suo *StatusdUpdateOne) AddStatusdIDs(ids ...int) *StatusdUpdateOne {
	suo.mutation.AddStatusdIDs(ids...)
	return suo
}

// AddStatusds adds the statusds edges to Deposit.
func (suo *StatusdUpdateOne) AddStatusds(d ...*Deposit) *StatusdUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return suo.AddStatusdIDs(ids...)
}

// Mutation returns the StatusdMutation object of the builder.
func (suo *StatusdUpdateOne) Mutation() *StatusdMutation {
	return suo.mutation
}

// RemoveStatusdIDs removes the statusds edge to Deposit by ids.
func (suo *StatusdUpdateOne) RemoveStatusdIDs(ids ...int) *StatusdUpdateOne {
	suo.mutation.RemoveStatusdIDs(ids...)
	return suo
}

// RemoveStatusds removes statusds edges to Deposit.
func (suo *StatusdUpdateOne) RemoveStatusds(d ...*Deposit) *StatusdUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return suo.RemoveStatusdIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *StatusdUpdateOne) Save(ctx context.Context) (*Statusd, error) {
	if v, ok := suo.mutation.Statusdname(); ok {
		if err := statusd.StatusdnameValidator(v); err != nil {
			return nil, &ValidationError{Name: "statusdname", err: fmt.Errorf("ent: validator failed for field \"statusdname\": %w", err)}
		}
	}

	var (
		err  error
		node *Statusd
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusdMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StatusdUpdateOne) SaveX(ctx context.Context) *Statusd {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *StatusdUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StatusdUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StatusdUpdateOne) sqlSave(ctx context.Context) (s *Statusd, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statusd.Table,
			Columns: statusd.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statusd.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Statusd.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.Statusdname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: statusd.FieldStatusdname,
		})
	}
	if nodes := suo.mutation.RemovedStatusdsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statusd.StatusdsTable,
			Columns: []string{statusd.StatusdsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deposit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.StatusdsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statusd.StatusdsTable,
			Columns: []string{statusd.StatusdsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deposit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Statusd{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statusd.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}