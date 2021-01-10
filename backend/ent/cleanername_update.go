// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/cleanername"
	"github.com/team15/app/ent/cleaningroom"
	"github.com/team15/app/ent/predicate"
)

// CleanerNameUpdate is the builder for updating CleanerName entities.
type CleanerNameUpdate struct {
	config
	hooks      []Hook
	mutation   *CleanerNameMutation
	predicates []predicate.CleanerName
}

// Where adds a new predicate for the builder.
func (cnu *CleanerNameUpdate) Where(ps ...predicate.CleanerName) *CleanerNameUpdate {
	cnu.predicates = append(cnu.predicates, ps...)
	return cnu
}

// SetCleanername sets the cleanername field.
func (cnu *CleanerNameUpdate) SetCleanername(s string) *CleanerNameUpdate {
	cnu.mutation.SetCleanername(s)
	return cnu
}

// AddCleaningroomIDs adds the cleaningrooms edge to CleaningRoom by ids.
func (cnu *CleanerNameUpdate) AddCleaningroomIDs(ids ...int) *CleanerNameUpdate {
	cnu.mutation.AddCleaningroomIDs(ids...)
	return cnu
}

// AddCleaningrooms adds the cleaningrooms edges to CleaningRoom.
func (cnu *CleanerNameUpdate) AddCleaningrooms(c ...*CleaningRoom) *CleanerNameUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cnu.AddCleaningroomIDs(ids...)
}

// Mutation returns the CleanerNameMutation object of the builder.
func (cnu *CleanerNameUpdate) Mutation() *CleanerNameMutation {
	return cnu.mutation
}

// RemoveCleaningroomIDs removes the cleaningrooms edge to CleaningRoom by ids.
func (cnu *CleanerNameUpdate) RemoveCleaningroomIDs(ids ...int) *CleanerNameUpdate {
	cnu.mutation.RemoveCleaningroomIDs(ids...)
	return cnu
}

// RemoveCleaningrooms removes cleaningrooms edges to CleaningRoom.
func (cnu *CleanerNameUpdate) RemoveCleaningrooms(c ...*CleaningRoom) *CleanerNameUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cnu.RemoveCleaningroomIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cnu *CleanerNameUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := cnu.mutation.Cleanername(); ok {
		if err := cleanername.CleanernameValidator(v); err != nil {
			return 0, &ValidationError{Name: "cleanername", err: fmt.Errorf("ent: validator failed for field \"cleanername\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(cnu.hooks) == 0 {
		affected, err = cnu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CleanerNameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cnu.mutation = mutation
			affected, err = cnu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cnu.hooks) - 1; i >= 0; i-- {
			mut = cnu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cnu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cnu *CleanerNameUpdate) SaveX(ctx context.Context) int {
	affected, err := cnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cnu *CleanerNameUpdate) Exec(ctx context.Context) error {
	_, err := cnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cnu *CleanerNameUpdate) ExecX(ctx context.Context) {
	if err := cnu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cnu *CleanerNameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cleanername.Table,
			Columns: cleanername.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cleanername.FieldID,
			},
		},
	}
	if ps := cnu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cnu.mutation.Cleanername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cleanername.FieldCleanername,
		})
	}
	if nodes := cnu.mutation.RemovedCleaningroomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cleanername.CleaningroomsTable,
			Columns: []string{cleanername.CleaningroomsColumn},
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
	if nodes := cnu.mutation.CleaningroomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cleanername.CleaningroomsTable,
			Columns: []string{cleanername.CleaningroomsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cleanername.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CleanerNameUpdateOne is the builder for updating a single CleanerName entity.
type CleanerNameUpdateOne struct {
	config
	hooks    []Hook
	mutation *CleanerNameMutation
}

// SetCleanername sets the cleanername field.
func (cnuo *CleanerNameUpdateOne) SetCleanername(s string) *CleanerNameUpdateOne {
	cnuo.mutation.SetCleanername(s)
	return cnuo
}

// AddCleaningroomIDs adds the cleaningrooms edge to CleaningRoom by ids.
func (cnuo *CleanerNameUpdateOne) AddCleaningroomIDs(ids ...int) *CleanerNameUpdateOne {
	cnuo.mutation.AddCleaningroomIDs(ids...)
	return cnuo
}

// AddCleaningrooms adds the cleaningrooms edges to CleaningRoom.
func (cnuo *CleanerNameUpdateOne) AddCleaningrooms(c ...*CleaningRoom) *CleanerNameUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cnuo.AddCleaningroomIDs(ids...)
}

// Mutation returns the CleanerNameMutation object of the builder.
func (cnuo *CleanerNameUpdateOne) Mutation() *CleanerNameMutation {
	return cnuo.mutation
}

// RemoveCleaningroomIDs removes the cleaningrooms edge to CleaningRoom by ids.
func (cnuo *CleanerNameUpdateOne) RemoveCleaningroomIDs(ids ...int) *CleanerNameUpdateOne {
	cnuo.mutation.RemoveCleaningroomIDs(ids...)
	return cnuo
}

// RemoveCleaningrooms removes cleaningrooms edges to CleaningRoom.
func (cnuo *CleanerNameUpdateOne) RemoveCleaningrooms(c ...*CleaningRoom) *CleanerNameUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cnuo.RemoveCleaningroomIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cnuo *CleanerNameUpdateOne) Save(ctx context.Context) (*CleanerName, error) {
	if v, ok := cnuo.mutation.Cleanername(); ok {
		if err := cleanername.CleanernameValidator(v); err != nil {
			return nil, &ValidationError{Name: "cleanername", err: fmt.Errorf("ent: validator failed for field \"cleanername\": %w", err)}
		}
	}

	var (
		err  error
		node *CleanerName
	)
	if len(cnuo.hooks) == 0 {
		node, err = cnuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CleanerNameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cnuo.mutation = mutation
			node, err = cnuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cnuo.hooks) - 1; i >= 0; i-- {
			mut = cnuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cnuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cnuo *CleanerNameUpdateOne) SaveX(ctx context.Context) *CleanerName {
	cn, err := cnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return cn
}

// Exec executes the query on the entity.
func (cnuo *CleanerNameUpdateOne) Exec(ctx context.Context) error {
	_, err := cnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cnuo *CleanerNameUpdateOne) ExecX(ctx context.Context) {
	if err := cnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cnuo *CleanerNameUpdateOne) sqlSave(ctx context.Context) (cn *CleanerName, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cleanername.Table,
			Columns: cleanername.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cleanername.FieldID,
			},
		},
	}
	id, ok := cnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CleanerName.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cnuo.mutation.Cleanername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cleanername.FieldCleanername,
		})
	}
	if nodes := cnuo.mutation.RemovedCleaningroomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cleanername.CleaningroomsTable,
			Columns: []string{cleanername.CleaningroomsColumn},
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
	if nodes := cnuo.mutation.CleaningroomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cleanername.CleaningroomsTable,
			Columns: []string{cleanername.CleaningroomsColumn},
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
	cn = &CleanerName{config: cnuo.config}
	_spec.Assign = cn.assignValues
	_spec.ScanValues = cn.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cleanername.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return cn, nil
}
