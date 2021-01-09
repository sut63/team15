// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/deposit"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/statusd"
)

// DepositCreate is the builder for creating a Deposit entity.
type DepositCreate struct {
	config
	mutation *DepositMutation
	hooks    []Hook
}

// SetAddedtime sets the addedtime field.
func (dc *DepositCreate) SetAddedtime(t time.Time) *DepositCreate {
	dc.mutation.SetAddedtime(t)
	return dc
}

// SetInfo sets the info field.
func (dc *DepositCreate) SetInfo(s string) *DepositCreate {
	dc.mutation.SetInfo(s)
	return dc
}

// SetEmployeeID sets the Employee edge to Employee by id.
func (dc *DepositCreate) SetEmployeeID(id int) *DepositCreate {
	dc.mutation.SetEmployeeID(id)
	return dc
}

// SetNillableEmployeeID sets the Employee edge to Employee by id if the given value is not nil.
func (dc *DepositCreate) SetNillableEmployeeID(id *int) *DepositCreate {
	if id != nil {
		dc = dc.SetEmployeeID(*id)
	}
	return dc
}

// SetEmployee sets the Employee edge to Employee.
func (dc *DepositCreate) SetEmployee(e *Employee) *DepositCreate {
	return dc.SetEmployeeID(e.ID)
}

// SetStatusdID sets the Statusd edge to Statusd by id.
func (dc *DepositCreate) SetStatusdID(id int) *DepositCreate {
	dc.mutation.SetStatusdID(id)
	return dc
}

// SetNillableStatusdID sets the Statusd edge to Statusd by id if the given value is not nil.
func (dc *DepositCreate) SetNillableStatusdID(id *int) *DepositCreate {
	if id != nil {
		dc = dc.SetStatusdID(*id)
	}
	return dc
}

// SetStatusd sets the Statusd edge to Statusd.
func (dc *DepositCreate) SetStatusd(s *Statusd) *DepositCreate {
	return dc.SetStatusdID(s.ID)
}

// Mutation returns the DepositMutation object of the builder.
func (dc *DepositCreate) Mutation() *DepositMutation {
	return dc.mutation
}

// Save creates the Deposit in the database.
func (dc *DepositCreate) Save(ctx context.Context) (*Deposit, error) {
	if _, ok := dc.mutation.Addedtime(); !ok {
		return nil, &ValidationError{Name: "addedtime", err: errors.New("ent: missing required field \"addedtime\"")}
	}
	if _, ok := dc.mutation.Info(); !ok {
		return nil, &ValidationError{Name: "info", err: errors.New("ent: missing required field \"info\"")}
	}
	var (
		err  error
		node *Deposit
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepositMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DepositCreate) SaveX(ctx context.Context) *Deposit {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DepositCreate) sqlSave(ctx context.Context) (*Deposit, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DepositCreate) createSpec() (*Deposit, *sqlgraph.CreateSpec) {
	var (
		d     = &Deposit{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deposit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deposit.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Addedtime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deposit.FieldAddedtime,
		})
		d.Addedtime = value
	}
	if value, ok := dc.mutation.Info(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deposit.FieldInfo,
		})
		d.Info = value
	}
	if nodes := dc.mutation.EmployeeIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.StatusdIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return d, _spec
}
