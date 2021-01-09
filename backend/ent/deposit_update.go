// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/deposit"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/predicate"
	"github.com/team15/app/ent/statusd"
)

// DepositUpdate is the builder for updating Deposit entities.
type DepositUpdate struct {
	config
	hooks      []Hook
	mutation   *DepositMutation
	predicates []predicate.Deposit
}

// Where adds a new predicate for the builder.
func (du *DepositUpdate) Where(ps ...predicate.Deposit) *DepositUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetAddedtime sets the addedtime field.
func (du *DepositUpdate) SetAddedtime(t time.Time) *DepositUpdate {
	du.mutation.SetAddedtime(t)
	return du
}

// SetInfo sets the info field.
func (du *DepositUpdate) SetInfo(s string) *DepositUpdate {
	du.mutation.SetInfo(s)
	return du
}

// SetEmployeeID sets the Employee edge to Employee by id.
func (du *DepositUpdate) SetEmployeeID(id int) *DepositUpdate {
	du.mutation.SetEmployeeID(id)
	return du
}

// SetNillableEmployeeID sets the Employee edge to Employee by id if the given value is not nil.
func (du *DepositUpdate) SetNillableEmployeeID(id *int) *DepositUpdate {
	if id != nil {
		du = du.SetEmployeeID(*id)
	}
	return du
}

// SetEmployee sets the Employee edge to Employee.
func (du *DepositUpdate) SetEmployee(e *Employee) *DepositUpdate {
	return du.SetEmployeeID(e.ID)
}

// SetStatusdID sets the Statusd edge to Statusd by id.
func (du *DepositUpdate) SetStatusdID(id int) *DepositUpdate {
	du.mutation.SetStatusdID(id)
	return du
}

// SetNillableStatusdID sets the Statusd edge to Statusd by id if the given value is not nil.
func (du *DepositUpdate) SetNillableStatusdID(id *int) *DepositUpdate {
	if id != nil {
		du = du.SetStatusdID(*id)
	}
	return du
}

// SetStatusd sets the Statusd edge to Statusd.
func (du *DepositUpdate) SetStatusd(s *Statusd) *DepositUpdate {
	return du.SetStatusdID(s.ID)
}

// Mutation returns the DepositMutation object of the builder.
func (du *DepositUpdate) Mutation() *DepositMutation {
	return du.mutation
}

// ClearEmployee clears the Employee edge to Employee.
func (du *DepositUpdate) ClearEmployee() *DepositUpdate {
	du.mutation.ClearEmployee()
	return du
}

// ClearStatusd clears the Statusd edge to Statusd.
func (du *DepositUpdate) ClearStatusd() *DepositUpdate {
	du.mutation.ClearStatusd()
	return du
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DepositUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepositMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DepositUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DepositUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DepositUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DepositUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deposit.Table,
			Columns: deposit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deposit.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Addedtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deposit.FieldAddedtime,
		})
	}
	if value, ok := du.mutation.Info(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deposit.FieldInfo,
		})
	}
	if du.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deposit.EmployeeTable,
			Columns: []string{deposit.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deposit.EmployeeTable,
			Columns: []string{deposit.EmployeeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.StatusdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deposit.StatusdTable,
			Columns: []string{deposit.StatusdColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statusd.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.StatusdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deposit.StatusdTable,
			Columns: []string{deposit.StatusdColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statusd.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deposit.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DepositUpdateOne is the builder for updating a single Deposit entity.
type DepositUpdateOne struct {
	config
	hooks    []Hook
	mutation *DepositMutation
}

// SetAddedtime sets the addedtime field.
func (duo *DepositUpdateOne) SetAddedtime(t time.Time) *DepositUpdateOne {
	duo.mutation.SetAddedtime(t)
	return duo
}

// SetInfo sets the info field.
func (duo *DepositUpdateOne) SetInfo(s string) *DepositUpdateOne {
	duo.mutation.SetInfo(s)
	return duo
}

// SetEmployeeID sets the Employee edge to Employee by id.
func (duo *DepositUpdateOne) SetEmployeeID(id int) *DepositUpdateOne {
	duo.mutation.SetEmployeeID(id)
	return duo
}

// SetNillableEmployeeID sets the Employee edge to Employee by id if the given value is not nil.
func (duo *DepositUpdateOne) SetNillableEmployeeID(id *int) *DepositUpdateOne {
	if id != nil {
		duo = duo.SetEmployeeID(*id)
	}
	return duo
}

// SetEmployee sets the Employee edge to Employee.
func (duo *DepositUpdateOne) SetEmployee(e *Employee) *DepositUpdateOne {
	return duo.SetEmployeeID(e.ID)
}

// SetStatusdID sets the Statusd edge to Statusd by id.
func (duo *DepositUpdateOne) SetStatusdID(id int) *DepositUpdateOne {
	duo.mutation.SetStatusdID(id)
	return duo
}

// SetNillableStatusdID sets the Statusd edge to Statusd by id if the given value is not nil.
func (duo *DepositUpdateOne) SetNillableStatusdID(id *int) *DepositUpdateOne {
	if id != nil {
		duo = duo.SetStatusdID(*id)
	}
	return duo
}

// SetStatusd sets the Statusd edge to Statusd.
func (duo *DepositUpdateOne) SetStatusd(s *Statusd) *DepositUpdateOne {
	return duo.SetStatusdID(s.ID)
}

// Mutation returns the DepositMutation object of the builder.
func (duo *DepositUpdateOne) Mutation() *DepositMutation {
	return duo.mutation
}

// ClearEmployee clears the Employee edge to Employee.
func (duo *DepositUpdateOne) ClearEmployee() *DepositUpdateOne {
	duo.mutation.ClearEmployee()
	return duo
}

// ClearStatusd clears the Statusd edge to Statusd.
func (duo *DepositUpdateOne) ClearStatusd() *DepositUpdateOne {
	duo.mutation.ClearStatusd()
	return duo
}

// Save executes the query and returns the updated entity.
func (duo *DepositUpdateOne) Save(ctx context.Context) (*Deposit, error) {

	var (
		err  error
		node *Deposit
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepositMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DepositUpdateOne) SaveX(ctx context.Context) *Deposit {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DepositUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DepositUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DepositUpdateOne) sqlSave(ctx context.Context) (d *Deposit, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deposit.Table,
			Columns: deposit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deposit.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Deposit.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Addedtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deposit.FieldAddedtime,
		})
	}
	if value, ok := duo.mutation.Info(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deposit.FieldInfo,
		})
	}
	if duo.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deposit.EmployeeTable,
			Columns: []string{deposit.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deposit.EmployeeTable,
			Columns: []string{deposit.EmployeeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.StatusdCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deposit.StatusdTable,
			Columns: []string{deposit.StatusdColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statusd.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.StatusdIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deposit.StatusdTable,
			Columns: []string{deposit.StatusdColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statusd.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Deposit{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deposit.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
