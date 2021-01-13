// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/rentalstatus"
	"github.com/team15/app/ent/repairinvoice"
)

// RentalstatusCreate is the builder for creating a Rentalstatus entity.
type RentalstatusCreate struct {
	config
	mutation *RentalstatusMutation
	hooks    []Hook
}

// SetRentalstatus sets the rentalstatus field.
func (rc *RentalstatusCreate) SetRentalstatus(s string) *RentalstatusCreate {
	rc.mutation.SetRentalstatus(s)
	return rc
}

// AddRepairinvoiceIDs adds the repairinvoices edge to Repairinvoice by ids.
func (rc *RentalstatusCreate) AddRepairinvoiceIDs(ids ...int) *RentalstatusCreate {
	rc.mutation.AddRepairinvoiceIDs(ids...)
	return rc
}

// AddRepairinvoices adds the repairinvoices edges to Repairinvoice.
func (rc *RentalstatusCreate) AddRepairinvoices(r ...*Repairinvoice) *RentalstatusCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddRepairinvoiceIDs(ids...)
}

// Mutation returns the RentalstatusMutation object of the builder.
func (rc *RentalstatusCreate) Mutation() *RentalstatusMutation {
	return rc.mutation
}

// Save creates the Rentalstatus in the database.
func (rc *RentalstatusCreate) Save(ctx context.Context) (*Rentalstatus, error) {
	if _, ok := rc.mutation.Rentalstatus(); !ok {
		return nil, &ValidationError{Name: "rentalstatus", err: errors.New("ent: missing required field \"rentalstatus\"")}
	}
	var (
		err  error
		node *Rentalstatus
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RentalstatusMutation)
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
func (rc *RentalstatusCreate) SaveX(ctx context.Context) *Rentalstatus {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *RentalstatusCreate) sqlSave(ctx context.Context) (*Rentalstatus, error) {
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

func (rc *RentalstatusCreate) createSpec() (*Rentalstatus, *sqlgraph.CreateSpec) {
	var (
		r     = &Rentalstatus{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rentalstatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rentalstatus.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Rentalstatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rentalstatus.FieldRentalstatus,
		})
		r.Rentalstatus = value
	}
	if nodes := rc.mutation.RepairinvoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rentalstatus.RepairinvoicesTable,
			Columns: []string{rentalstatus.RepairinvoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repairinvoice.FieldID,
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
