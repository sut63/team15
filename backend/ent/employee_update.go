// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/deposit"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/predicate"
)

// EmployeeUpdate is the builder for updating Employee entities.
type EmployeeUpdate struct {
	config
	hooks      []Hook
	mutation   *EmployeeMutation
	predicates []predicate.Employee
}

// Where adds a new predicate for the builder.
func (eu *EmployeeUpdate) Where(ps ...predicate.Employee) *EmployeeUpdate {
	eu.predicates = append(eu.predicates, ps...)
	return eu
}

// SetEmployeename sets the employeename field.
func (eu *EmployeeUpdate) SetEmployeename(s string) *EmployeeUpdate {
	eu.mutation.SetEmployeename(s)
	return eu
}

// SetEmployeeemail sets the employeeemail field.
func (eu *EmployeeUpdate) SetEmployeeemail(s string) *EmployeeUpdate {
	eu.mutation.SetEmployeeemail(s)
	return eu
}

// SetPassword sets the password field.
func (eu *EmployeeUpdate) SetPassword(s string) *EmployeeUpdate {
	eu.mutation.SetPassword(s)
	return eu
}

// AddEmployeeIDs adds the employees edge to Deposit by ids.
func (eu *EmployeeUpdate) AddEmployeeIDs(ids ...int) *EmployeeUpdate {
	eu.mutation.AddEmployeeIDs(ids...)
	return eu
}

// AddEmployees adds the employees edges to Deposit.
func (eu *EmployeeUpdate) AddEmployees(d ...*Deposit) *EmployeeUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return eu.AddEmployeeIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (eu *EmployeeUpdate) Mutation() *EmployeeMutation {
	return eu.mutation
}

// RemoveEmployeeIDs removes the employees edge to Deposit by ids.
func (eu *EmployeeUpdate) RemoveEmployeeIDs(ids ...int) *EmployeeUpdate {
	eu.mutation.RemoveEmployeeIDs(ids...)
	return eu
}

// RemoveEmployees removes employees edges to Deposit.
func (eu *EmployeeUpdate) RemoveEmployees(d ...*Deposit) *EmployeeUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return eu.RemoveEmployeeIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (eu *EmployeeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := eu.mutation.Employeename(); ok {
		if err := employee.EmployeenameValidator(v); err != nil {
			return 0, &ValidationError{Name: "employeename", err: fmt.Errorf("ent: validator failed for field \"employeename\": %w", err)}
		}
	}
	if v, ok := eu.mutation.Employeeemail(); ok {
		if err := employee.EmployeeemailValidator(v); err != nil {
			return 0, &ValidationError{Name: "employeeemail", err: fmt.Errorf("ent: validator failed for field \"employeeemail\": %w", err)}
		}
	}
	if v, ok := eu.mutation.Password(); ok {
		if err := employee.PasswordValidator(v); err != nil {
			return 0, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmployeeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmployeeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmployeeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EmployeeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   employee.Table,
			Columns: employee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employee.FieldID,
			},
		},
	}
	if ps := eu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Employeename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEmployeename,
		})
	}
	if value, ok := eu.mutation.Employeeemail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEmployeeemail,
		})
	}
	if value, ok := eu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldPassword,
		})
	}
	if nodes := eu.mutation.RemovedEmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.EmployeesTable,
			Columns: []string{employee.EmployeesColumn},
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
	if nodes := eu.mutation.EmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.EmployeesTable,
			Columns: []string{employee.EmployeesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// EmployeeUpdateOne is the builder for updating a single Employee entity.
type EmployeeUpdateOne struct {
	config
	hooks    []Hook
	mutation *EmployeeMutation
}

// SetEmployeename sets the employeename field.
func (euo *EmployeeUpdateOne) SetEmployeename(s string) *EmployeeUpdateOne {
	euo.mutation.SetEmployeename(s)
	return euo
}

// SetEmployeeemail sets the employeeemail field.
func (euo *EmployeeUpdateOne) SetEmployeeemail(s string) *EmployeeUpdateOne {
	euo.mutation.SetEmployeeemail(s)
	return euo
}

// SetPassword sets the password field.
func (euo *EmployeeUpdateOne) SetPassword(s string) *EmployeeUpdateOne {
	euo.mutation.SetPassword(s)
	return euo
}

// AddEmployeeIDs adds the employees edge to Deposit by ids.
func (euo *EmployeeUpdateOne) AddEmployeeIDs(ids ...int) *EmployeeUpdateOne {
	euo.mutation.AddEmployeeIDs(ids...)
	return euo
}

// AddEmployees adds the employees edges to Deposit.
func (euo *EmployeeUpdateOne) AddEmployees(d ...*Deposit) *EmployeeUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return euo.AddEmployeeIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (euo *EmployeeUpdateOne) Mutation() *EmployeeMutation {
	return euo.mutation
}

// RemoveEmployeeIDs removes the employees edge to Deposit by ids.
func (euo *EmployeeUpdateOne) RemoveEmployeeIDs(ids ...int) *EmployeeUpdateOne {
	euo.mutation.RemoveEmployeeIDs(ids...)
	return euo
}

// RemoveEmployees removes employees edges to Deposit.
func (euo *EmployeeUpdateOne) RemoveEmployees(d ...*Deposit) *EmployeeUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return euo.RemoveEmployeeIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (euo *EmployeeUpdateOne) Save(ctx context.Context) (*Employee, error) {
	if v, ok := euo.mutation.Employeename(); ok {
		if err := employee.EmployeenameValidator(v); err != nil {
			return nil, &ValidationError{Name: "employeename", err: fmt.Errorf("ent: validator failed for field \"employeename\": %w", err)}
		}
	}
	if v, ok := euo.mutation.Employeeemail(); ok {
		if err := employee.EmployeeemailValidator(v); err != nil {
			return nil, &ValidationError{Name: "employeeemail", err: fmt.Errorf("ent: validator failed for field \"employeeemail\": %w", err)}
		}
	}
	if v, ok := euo.mutation.Password(); ok {
		if err := employee.PasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}

	var (
		err  error
		node *Employee
	)
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmployeeUpdateOne) SaveX(ctx context.Context) *Employee {
	e, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return e
}

// Exec executes the query on the entity.
func (euo *EmployeeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmployeeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EmployeeUpdateOne) sqlSave(ctx context.Context) (e *Employee, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   employee.Table,
			Columns: employee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employee.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Employee.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := euo.mutation.Employeename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEmployeename,
		})
	}
	if value, ok := euo.mutation.Employeeemail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEmployeeemail,
		})
	}
	if value, ok := euo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldPassword,
		})
	}
	if nodes := euo.mutation.RemovedEmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.EmployeesTable,
			Columns: []string{employee.EmployeesColumn},
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
	if nodes := euo.mutation.EmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.EmployeesTable,
			Columns: []string{employee.EmployeesColumn},
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
	e = &Employee{config: euo.config}
	_spec.Assign = e.assignValues
	_spec.ScanValues = e.scanValues()
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{employee.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return e, nil
}
