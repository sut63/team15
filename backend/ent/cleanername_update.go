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

// CleanernameUpdate is the builder for updating Cleanername entities.
type CleanernameUpdate struct {
	config
	hooks      []Hook
	mutation   *CleanernameMutation
	predicates []predicate.Cleanername
}

// Where adds a new predicate for the builder.
func (cu *CleanernameUpdate) Where(ps ...predicate.Cleanername) *CleanernameUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetCleanername sets the cleanername field.
func (cu *CleanernameUpdate) SetCleanername(s string) *CleanernameUpdate {
	cu.mutation.SetCleanername(s)
	return cu
}

// AddCleaningroomIDs adds the cleaningrooms edge to Cleaningroom by ids.
func (cu *CleanernameUpdate) AddCleaningroomIDs(ids ...int) *CleanernameUpdate {
	cu.mutation.AddCleaningroomIDs(ids...)
	return cu
}

// AddCleaningrooms adds the cleaningrooms edges to Cleaningroom.
func (cu *CleanernameUpdate) AddCleaningrooms(c ...*Cleaningroom) *CleanernameUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCleaningroomIDs(ids...)
}

// Mutation returns the CleanernameMutation object of the builder.
func (cu *CleanernameUpdate) Mutation() *CleanernameMutation {
	return cu.mutation
}

// RemoveCleaningroomIDs removes the cleaningrooms edge to Cleaningroom by ids.
func (cu *CleanernameUpdate) RemoveCleaningroomIDs(ids ...int) *CleanernameUpdate {
	cu.mutation.RemoveCleaningroomIDs(ids...)
	return cu
}

// RemoveCleaningrooms removes cleaningrooms edges to Cleaningroom.
func (cu *CleanernameUpdate) RemoveCleaningrooms(c ...*Cleaningroom) *CleanernameUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCleaningroomIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CleanernameUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := cu.mutation.Cleanername(); ok {
		if err := cleanername.CleanernameValidator(v); err != nil {
			return 0, &ValidationError{Name: "cleanername", err: fmt.Errorf("ent: validator failed for field \"cleanername\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CleanernameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CleanernameUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CleanernameUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CleanernameUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CleanernameUpdate) sqlSave(ctx context.Context) (n int, err error) {
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
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Cleanername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cleanername.FieldCleanername,
		})
	}
	if nodes := cu.mutation.RemovedCleaningroomsIDs(); len(nodes) > 0 {
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
	if nodes := cu.mutation.CleaningroomsIDs(); len(nodes) > 0 {
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cleanername.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CleanernameUpdateOne is the builder for updating a single Cleanername entity.
type CleanernameUpdateOne struct {
	config
	hooks    []Hook
	mutation *CleanernameMutation
}

// SetCleanername sets the cleanername field.
func (cuo *CleanernameUpdateOne) SetCleanername(s string) *CleanernameUpdateOne {
	cuo.mutation.SetCleanername(s)
	return cuo
}

// AddCleaningroomIDs adds the cleaningrooms edge to Cleaningroom by ids.
func (cuo *CleanernameUpdateOne) AddCleaningroomIDs(ids ...int) *CleanernameUpdateOne {
	cuo.mutation.AddCleaningroomIDs(ids...)
	return cuo
}

// AddCleaningrooms adds the cleaningrooms edges to Cleaningroom.
func (cuo *CleanernameUpdateOne) AddCleaningrooms(c ...*Cleaningroom) *CleanernameUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCleaningroomIDs(ids...)
}

// Mutation returns the CleanernameMutation object of the builder.
func (cuo *CleanernameUpdateOne) Mutation() *CleanernameMutation {
	return cuo.mutation
}

// RemoveCleaningroomIDs removes the cleaningrooms edge to Cleaningroom by ids.
func (cuo *CleanernameUpdateOne) RemoveCleaningroomIDs(ids ...int) *CleanernameUpdateOne {
	cuo.mutation.RemoveCleaningroomIDs(ids...)
	return cuo
}

// RemoveCleaningrooms removes cleaningrooms edges to Cleaningroom.
func (cuo *CleanernameUpdateOne) RemoveCleaningrooms(c ...*Cleaningroom) *CleanernameUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCleaningroomIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *CleanernameUpdateOne) Save(ctx context.Context) (*Cleanername, error) {
	if v, ok := cuo.mutation.Cleanername(); ok {
		if err := cleanername.CleanernameValidator(v); err != nil {
			return nil, &ValidationError{Name: "cleanername", err: fmt.Errorf("ent: validator failed for field \"cleanername\": %w", err)}
		}
	}

	var (
		err  error
		node *Cleanername
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CleanernameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CleanernameUpdateOne) SaveX(ctx context.Context) *Cleanername {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CleanernameUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CleanernameUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CleanernameUpdateOne) sqlSave(ctx context.Context) (c *Cleanername, err error) {
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
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Cleanername.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.Cleanername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cleanername.FieldCleanername,
		})
	}
	if nodes := cuo.mutation.RemovedCleaningroomsIDs(); len(nodes) > 0 {
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
	if nodes := cuo.mutation.CleaningroomsIDs(); len(nodes) > 0 {
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
	c = &Cleanername{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cleanername.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
