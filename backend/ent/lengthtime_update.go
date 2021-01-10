// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/cleaningroom"
	"github.com/team15/app/ent/lengthtime"
	"github.com/team15/app/ent/predicate"
)

// LengthTimeUpdate is the builder for updating LengthTime entities.
type LengthTimeUpdate struct {
	config
	hooks      []Hook
	mutation   *LengthTimeMutation
	predicates []predicate.LengthTime
}

// Where adds a new predicate for the builder.
func (ltu *LengthTimeUpdate) Where(ps ...predicate.LengthTime) *LengthTimeUpdate {
	ltu.predicates = append(ltu.predicates, ps...)
	return ltu
}

// SetLengthtime sets the lengthtime field.
func (ltu *LengthTimeUpdate) SetLengthtime(s string) *LengthTimeUpdate {
	ltu.mutation.SetLengthtime(s)
	return ltu
}

// AddCleaningroomIDs adds the cleaningrooms edge to CleaningRoom by ids.
func (ltu *LengthTimeUpdate) AddCleaningroomIDs(ids ...int) *LengthTimeUpdate {
	ltu.mutation.AddCleaningroomIDs(ids...)
	return ltu
}

// AddCleaningrooms adds the cleaningrooms edges to CleaningRoom.
func (ltu *LengthTimeUpdate) AddCleaningrooms(c ...*CleaningRoom) *LengthTimeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ltu.AddCleaningroomIDs(ids...)
}

// Mutation returns the LengthTimeMutation object of the builder.
func (ltu *LengthTimeUpdate) Mutation() *LengthTimeMutation {
	return ltu.mutation
}

// RemoveCleaningroomIDs removes the cleaningrooms edge to CleaningRoom by ids.
func (ltu *LengthTimeUpdate) RemoveCleaningroomIDs(ids ...int) *LengthTimeUpdate {
	ltu.mutation.RemoveCleaningroomIDs(ids...)
	return ltu
}

// RemoveCleaningrooms removes cleaningrooms edges to CleaningRoom.
func (ltu *LengthTimeUpdate) RemoveCleaningrooms(c ...*CleaningRoom) *LengthTimeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ltu.RemoveCleaningroomIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ltu *LengthTimeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ltu.mutation.Lengthtime(); ok {
		if err := lengthtime.LengthtimeValidator(v); err != nil {
			return 0, &ValidationError{Name: "lengthtime", err: fmt.Errorf("ent: validator failed for field \"lengthtime\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(ltu.hooks) == 0 {
		affected, err = ltu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LengthTimeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ltu.mutation = mutation
			affected, err = ltu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ltu.hooks) - 1; i >= 0; i-- {
			mut = ltu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ltu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ltu *LengthTimeUpdate) SaveX(ctx context.Context) int {
	affected, err := ltu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ltu *LengthTimeUpdate) Exec(ctx context.Context) error {
	_, err := ltu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltu *LengthTimeUpdate) ExecX(ctx context.Context) {
	if err := ltu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ltu *LengthTimeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lengthtime.Table,
			Columns: lengthtime.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lengthtime.FieldID,
			},
		},
	}
	if ps := ltu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ltu.mutation.Lengthtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lengthtime.FieldLengthtime,
		})
	}
	if nodes := ltu.mutation.RemovedCleaningroomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lengthtime.CleaningroomsTable,
			Columns: []string{lengthtime.CleaningroomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cleaningroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ltu.mutation.CleaningroomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lengthtime.CleaningroomsTable,
			Columns: []string{lengthtime.CleaningroomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cleaningroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ltu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lengthtime.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// LengthTimeUpdateOne is the builder for updating a single LengthTime entity.
type LengthTimeUpdateOne struct {
	config
	hooks    []Hook
	mutation *LengthTimeMutation
}

// SetLengthtime sets the lengthtime field.
func (ltuo *LengthTimeUpdateOne) SetLengthtime(s string) *LengthTimeUpdateOne {
	ltuo.mutation.SetLengthtime(s)
	return ltuo
}

// AddCleaningroomIDs adds the cleaningrooms edge to CleaningRoom by ids.
func (ltuo *LengthTimeUpdateOne) AddCleaningroomIDs(ids ...int) *LengthTimeUpdateOne {
	ltuo.mutation.AddCleaningroomIDs(ids...)
	return ltuo
}

// AddCleaningrooms adds the cleaningrooms edges to CleaningRoom.
func (ltuo *LengthTimeUpdateOne) AddCleaningrooms(c ...*CleaningRoom) *LengthTimeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ltuo.AddCleaningroomIDs(ids...)
}

// Mutation returns the LengthTimeMutation object of the builder.
func (ltuo *LengthTimeUpdateOne) Mutation() *LengthTimeMutation {
	return ltuo.mutation
}

// RemoveCleaningroomIDs removes the cleaningrooms edge to CleaningRoom by ids.
func (ltuo *LengthTimeUpdateOne) RemoveCleaningroomIDs(ids ...int) *LengthTimeUpdateOne {
	ltuo.mutation.RemoveCleaningroomIDs(ids...)
	return ltuo
}

// RemoveCleaningrooms removes cleaningrooms edges to CleaningRoom.
func (ltuo *LengthTimeUpdateOne) RemoveCleaningrooms(c ...*CleaningRoom) *LengthTimeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ltuo.RemoveCleaningroomIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ltuo *LengthTimeUpdateOne) Save(ctx context.Context) (*LengthTime, error) {
	if v, ok := ltuo.mutation.Lengthtime(); ok {
		if err := lengthtime.LengthtimeValidator(v); err != nil {
			return nil, &ValidationError{Name: "lengthtime", err: fmt.Errorf("ent: validator failed for field \"lengthtime\": %w", err)}
		}
	}

	var (
		err  error
		node *LengthTime
	)
	if len(ltuo.hooks) == 0 {
		node, err = ltuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LengthTimeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ltuo.mutation = mutation
			node, err = ltuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ltuo.hooks) - 1; i >= 0; i-- {
			mut = ltuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ltuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ltuo *LengthTimeUpdateOne) SaveX(ctx context.Context) *LengthTime {
	lt, err := ltuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return lt
}

// Exec executes the query on the entity.
func (ltuo *LengthTimeUpdateOne) Exec(ctx context.Context) error {
	_, err := ltuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltuo *LengthTimeUpdateOne) ExecX(ctx context.Context) {
	if err := ltuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ltuo *LengthTimeUpdateOne) sqlSave(ctx context.Context) (lt *LengthTime, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lengthtime.Table,
			Columns: lengthtime.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lengthtime.FieldID,
			},
		},
	}
	id, ok := ltuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing LengthTime.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ltuo.mutation.Lengthtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lengthtime.FieldLengthtime,
		})
	}
	if nodes := ltuo.mutation.RemovedCleaningroomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lengthtime.CleaningroomsTable,
			Columns: []string{lengthtime.CleaningroomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cleaningroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ltuo.mutation.CleaningroomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lengthtime.CleaningroomsTable,
			Columns: []string{lengthtime.CleaningroomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cleaningroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	lt = &LengthTime{config: ltuo.config}
	_spec.Assign = lt.assignValues
	_spec.ScanValues = lt.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ltuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lengthtime.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return lt, nil
}
