// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/lease"
	"github.com/team15/app/ent/rentalstatus"
	"github.com/team15/app/ent/repairinvoice"
)

// RepairinvoiceCreate is the builder for creating a Repairinvoice entity.
type RepairinvoiceCreate struct {
	config
	mutation *RepairinvoiceMutation
	hooks    []Hook
}

// SetBequipment sets the bequipment field.
func (rc *RepairinvoiceCreate) SetBequipment(s string) *RepairinvoiceCreate {
	rc.mutation.SetBequipment(s)
	return rc
}

// SetEmtell sets the emtell field.
func (rc *RepairinvoiceCreate) SetEmtell(s string) *RepairinvoiceCreate {
	rc.mutation.SetEmtell(s)
	return rc
}

// SetNum sets the num field.
func (rc *RepairinvoiceCreate) SetNum(i int) *RepairinvoiceCreate {
	rc.mutation.SetNum(i)
	return rc
}

// SetEmployeeID sets the employee edge to Employee by id.
func (rc *RepairinvoiceCreate) SetEmployeeID(id int) *RepairinvoiceCreate {
	rc.mutation.SetEmployeeID(id)
	return rc
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (rc *RepairinvoiceCreate) SetNillableEmployeeID(id *int) *RepairinvoiceCreate {
	if id != nil {
		rc = rc.SetEmployeeID(*id)
	}
	return rc
}

// SetEmployee sets the employee edge to Employee.
func (rc *RepairinvoiceCreate) SetEmployee(e *Employee) *RepairinvoiceCreate {
	return rc.SetEmployeeID(e.ID)
}

// SetRentalstatusID sets the Rentalstatus edge to Rentalstatus by id.
func (rc *RepairinvoiceCreate) SetRentalstatusID(id int) *RepairinvoiceCreate {
	rc.mutation.SetRentalstatusID(id)
	return rc
}

// SetNillableRentalstatusID sets the Rentalstatus edge to Rentalstatus by id if the given value is not nil.
func (rc *RepairinvoiceCreate) SetNillableRentalstatusID(id *int) *RepairinvoiceCreate {
	if id != nil {
		rc = rc.SetRentalstatusID(*id)
	}
	return rc
}

// SetRentalstatus sets the Rentalstatus edge to Rentalstatus.
func (rc *RepairinvoiceCreate) SetRentalstatus(r *Rentalstatus) *RepairinvoiceCreate {
	return rc.SetRentalstatusID(r.ID)
}

// SetLeaseID sets the Lease edge to Lease by id.
func (rc *RepairinvoiceCreate) SetLeaseID(id int) *RepairinvoiceCreate {
	rc.mutation.SetLeaseID(id)
	return rc
}

// SetNillableLeaseID sets the Lease edge to Lease by id if the given value is not nil.
func (rc *RepairinvoiceCreate) SetNillableLeaseID(id *int) *RepairinvoiceCreate {
	if id != nil {
		rc = rc.SetLeaseID(*id)
	}
	return rc
}

// SetLease sets the Lease edge to Lease.
func (rc *RepairinvoiceCreate) SetLease(l *Lease) *RepairinvoiceCreate {
	return rc.SetLeaseID(l.ID)
}

// Mutation returns the RepairinvoiceMutation object of the builder.
func (rc *RepairinvoiceCreate) Mutation() *RepairinvoiceMutation {
	return rc.mutation
}

// Save creates the Repairinvoice in the database.
func (rc *RepairinvoiceCreate) Save(ctx context.Context) (*Repairinvoice, error) {
	if _, ok := rc.mutation.Bequipment(); !ok {
		return nil, &ValidationError{Name: "bequipment", err: errors.New("ent: missing required field \"bequipment\"")}
	}
	if v, ok := rc.mutation.Bequipment(); ok {
		if err := repairinvoice.BequipmentValidator(v); err != nil {
			return nil, &ValidationError{Name: "bequipment", err: fmt.Errorf("ent: validator failed for field \"bequipment\": %w", err)}
		}
	}
	if _, ok := rc.mutation.Emtell(); !ok {
		return nil, &ValidationError{Name: "emtell", err: errors.New("ent: missing required field \"emtell\"")}
	}
	if v, ok := rc.mutation.Emtell(); ok {
		if err := repairinvoice.EmtellValidator(v); err != nil {
			return nil, &ValidationError{Name: "emtell", err: fmt.Errorf("ent: validator failed for field \"emtell\": %w", err)}
		}
	}
	if _, ok := rc.mutation.Num(); !ok {
		return nil, &ValidationError{Name: "num", err: errors.New("ent: missing required field \"num\"")}
	}
	if v, ok := rc.mutation.Num(); ok {
		if err := repairinvoice.NumValidator(v); err != nil {
			return nil, &ValidationError{Name: "num", err: fmt.Errorf("ent: validator failed for field \"num\": %w", err)}
		}
	}
	var (
		err  error
		node *Repairinvoice
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepairinvoiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RepairinvoiceCreate) SaveX(ctx context.Context) *Repairinvoice {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *RepairinvoiceCreate) sqlSave(ctx context.Context) (*Repairinvoice, error) {
	r, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	r.ID = int(id)
	return r, nil
}

func (rc *RepairinvoiceCreate) createSpec() (*Repairinvoice, *sqlgraph.CreateSpec) {
	var (
		r     = &Repairinvoice{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: repairinvoice.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: repairinvoice.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Bequipment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repairinvoice.FieldBequipment,
		})
		r.Bequipment = value
	}
	if value, ok := rc.mutation.Emtell(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repairinvoice.FieldEmtell,
		})
		r.Emtell = value
	}
	if value, ok := rc.mutation.Num(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repairinvoice.FieldNum,
		})
		r.Num = value
	}
	if nodes := rc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repairinvoice.EmployeeTable,
			Columns: []string{repairinvoice.EmployeeColumn},
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
	if nodes := rc.mutation.RentalstatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repairinvoice.RentalstatusTable,
			Columns: []string{repairinvoice.RentalstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rentalstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.LeaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repairinvoice.LeaseTable,
			Columns: []string{repairinvoice.LeaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return r, _spec
}
