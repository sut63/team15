// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/jobposition"
)

// JobpositionCreate is the builder for creating a Jobposition entity.
type JobpositionCreate struct {
	config
	mutation *JobpositionMutation
	hooks    []Hook
}

// SetPositionName sets the position_name field.
func (jc *JobpositionCreate) SetPositionName(s string) *JobpositionCreate {
	jc.mutation.SetPositionName(s)
	return jc
}

// AddEmployeeIDs adds the employees edge to Employee by ids.
func (jc *JobpositionCreate) AddEmployeeIDs(ids ...int) *JobpositionCreate {
	jc.mutation.AddEmployeeIDs(ids...)
	return jc
}

// AddEmployees adds the employees edges to Employee.
func (jc *JobpositionCreate) AddEmployees(e ...*Employee) *JobpositionCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return jc.AddEmployeeIDs(ids...)
}

// Mutation returns the JobpositionMutation object of the builder.
func (jc *JobpositionCreate) Mutation() *JobpositionMutation {
	return jc.mutation
}

// Save creates the Jobposition in the database.
func (jc *JobpositionCreate) Save(ctx context.Context) (*Jobposition, error) {
	if _, ok := jc.mutation.PositionName(); !ok {
		return nil, &ValidationError{Name: "position_name", err: errors.New("ent: missing required field \"position_name\"")}
	}
	if v, ok := jc.mutation.PositionName(); ok {
		if err := jobposition.PositionNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "position_name", err: fmt.Errorf("ent: validator failed for field \"position_name\": %w", err)}
		}
	}
	var (
		err  error
		node *Jobposition
	)
	if len(jc.hooks) == 0 {
		node, err = jc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobpositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			jc.mutation = mutation
			node, err = jc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(jc.hooks) - 1; i >= 0; i-- {
			mut = jc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (jc *JobpositionCreate) SaveX(ctx context.Context) *Jobposition {
	v, err := jc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jc *JobpositionCreate) sqlSave(ctx context.Context) (*Jobposition, error) {
	j, _spec := jc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	j.ID = int(id)
	return j, nil
}

func (jc *JobpositionCreate) createSpec() (*Jobposition, *sqlgraph.CreateSpec) {
	var (
		j     = &Jobposition{config: jc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: jobposition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: jobposition.FieldID,
			},
		}
	)
	if value, ok := jc.mutation.PositionName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: jobposition.FieldPositionName,
		})
		j.PositionName = value
	}
	if nodes := jc.mutation.EmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.EmployeesTable,
			Columns: []string{jobposition.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return j, _spec
}
