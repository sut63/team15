// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/bedtype"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/lease"
	"github.com/team15/app/ent/petrule"
	"github.com/team15/app/ent/pledge"
	"github.com/team15/app/ent/roomdetail"
	"github.com/team15/app/ent/staytype"
)

// RoomdetailCreate is the builder for creating a Roomdetail entity.
type RoomdetailCreate struct {
	config
	mutation *RoomdetailMutation
	hooks    []Hook
}

// SetRoomnumber sets the roomnumber field.
func (rc *RoomdetailCreate) SetRoomnumber(s string) *RoomdetailCreate {
	rc.mutation.SetRoomnumber(s)
	return rc
}

// SetRoomtypename sets the roomtypename field.
func (rc *RoomdetailCreate) SetRoomtypename(s string) *RoomdetailCreate {
	rc.mutation.SetRoomtypename(s)
	return rc
}

// SetRoomprice sets the roomprice field.
func (rc *RoomdetailCreate) SetRoomprice(s string) *RoomdetailCreate {
	rc.mutation.SetRoomprice(s)
	return rc
}

// SetSleep sets the sleep field.
func (rc *RoomdetailCreate) SetSleep(s string) *RoomdetailCreate {
	rc.mutation.SetSleep(s)
	return rc
}

// SetBed sets the bed field.
func (rc *RoomdetailCreate) SetBed(s string) *RoomdetailCreate {
	rc.mutation.SetBed(s)
	return rc
}

// SetPledgeID sets the pledge edge to Pledge by id.
func (rc *RoomdetailCreate) SetPledgeID(id int) *RoomdetailCreate {
	rc.mutation.SetPledgeID(id)
	return rc
}

// SetNillablePledgeID sets the pledge edge to Pledge by id if the given value is not nil.
func (rc *RoomdetailCreate) SetNillablePledgeID(id *int) *RoomdetailCreate {
	if id != nil {
		rc = rc.SetPledgeID(*id)
	}
	return rc
}

// SetPledge sets the pledge edge to Pledge.
func (rc *RoomdetailCreate) SetPledge(p *Pledge) *RoomdetailCreate {
	return rc.SetPledgeID(p.ID)
}

// SetPetruleID sets the petrule edge to Petrule by id.
func (rc *RoomdetailCreate) SetPetruleID(id int) *RoomdetailCreate {
	rc.mutation.SetPetruleID(id)
	return rc
}

// SetNillablePetruleID sets the petrule edge to Petrule by id if the given value is not nil.
func (rc *RoomdetailCreate) SetNillablePetruleID(id *int) *RoomdetailCreate {
	if id != nil {
		rc = rc.SetPetruleID(*id)
	}
	return rc
}

// SetPetrule sets the petrule edge to Petrule.
func (rc *RoomdetailCreate) SetPetrule(p *Petrule) *RoomdetailCreate {
	return rc.SetPetruleID(p.ID)
}

// SetBedtypeID sets the bedtype edge to Bedtype by id.
func (rc *RoomdetailCreate) SetBedtypeID(id int) *RoomdetailCreate {
	rc.mutation.SetBedtypeID(id)
	return rc
}

// SetNillableBedtypeID sets the bedtype edge to Bedtype by id if the given value is not nil.
func (rc *RoomdetailCreate) SetNillableBedtypeID(id *int) *RoomdetailCreate {
	if id != nil {
		rc = rc.SetBedtypeID(*id)
	}
	return rc
}

// SetBedtype sets the bedtype edge to Bedtype.
func (rc *RoomdetailCreate) SetBedtype(b *Bedtype) *RoomdetailCreate {
	return rc.SetBedtypeID(b.ID)
}

// SetEmployeeID sets the employee edge to Employee by id.
func (rc *RoomdetailCreate) SetEmployeeID(id int) *RoomdetailCreate {
	rc.mutation.SetEmployeeID(id)
	return rc
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (rc *RoomdetailCreate) SetNillableEmployeeID(id *int) *RoomdetailCreate {
	if id != nil {
		rc = rc.SetEmployeeID(*id)
	}
	return rc
}

// SetEmployee sets the employee edge to Employee.
func (rc *RoomdetailCreate) SetEmployee(e *Employee) *RoomdetailCreate {
	return rc.SetEmployeeID(e.ID)
}

// SetStaytypeID sets the staytype edge to Staytype by id.
func (rc *RoomdetailCreate) SetStaytypeID(id int) *RoomdetailCreate {
	rc.mutation.SetStaytypeID(id)
	return rc
}

// SetNillableStaytypeID sets the staytype edge to Staytype by id if the given value is not nil.
func (rc *RoomdetailCreate) SetNillableStaytypeID(id *int) *RoomdetailCreate {
	if id != nil {
		rc = rc.SetStaytypeID(*id)
	}
	return rc
}

// SetStaytype sets the staytype edge to Staytype.
func (rc *RoomdetailCreate) SetStaytype(s *Staytype) *RoomdetailCreate {
	return rc.SetStaytypeID(s.ID)
}

// SetRoomdetailsID sets the roomdetails edge to Lease by id.
func (rc *RoomdetailCreate) SetRoomdetailsID(id int) *RoomdetailCreate {
	rc.mutation.SetRoomdetailsID(id)
	return rc
}

// SetNillableRoomdetailsID sets the roomdetails edge to Lease by id if the given value is not nil.
func (rc *RoomdetailCreate) SetNillableRoomdetailsID(id *int) *RoomdetailCreate {
	if id != nil {
		rc = rc.SetRoomdetailsID(*id)
	}
	return rc
}

// SetRoomdetails sets the roomdetails edge to Lease.
func (rc *RoomdetailCreate) SetRoomdetails(l *Lease) *RoomdetailCreate {
	return rc.SetRoomdetailsID(l.ID)
}

// Mutation returns the RoomdetailMutation object of the builder.
func (rc *RoomdetailCreate) Mutation() *RoomdetailMutation {
	return rc.mutation
}

// Save creates the Roomdetail in the database.
func (rc *RoomdetailCreate) Save(ctx context.Context) (*Roomdetail, error) {
	if _, ok := rc.mutation.Roomnumber(); !ok {
		return nil, &ValidationError{Name: "roomnumber", err: errors.New("ent: missing required field \"roomnumber\"")}
	}
	if _, ok := rc.mutation.Roomtypename(); !ok {
		return nil, &ValidationError{Name: "roomtypename", err: errors.New("ent: missing required field \"roomtypename\"")}
	}
	if _, ok := rc.mutation.Roomprice(); !ok {
		return nil, &ValidationError{Name: "roomprice", err: errors.New("ent: missing required field \"roomprice\"")}
	}
	if _, ok := rc.mutation.Sleep(); !ok {
		return nil, &ValidationError{Name: "sleep", err: errors.New("ent: missing required field \"sleep\"")}
	}
	if _, ok := rc.mutation.Bed(); !ok {
		return nil, &ValidationError{Name: "bed", err: errors.New("ent: missing required field \"bed\"")}
	}
	var (
		err  error
		node *Roomdetail
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomdetailMutation)
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
func (rc *RoomdetailCreate) SaveX(ctx context.Context) *Roomdetail {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *RoomdetailCreate) sqlSave(ctx context.Context) (*Roomdetail, error) {
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

func (rc *RoomdetailCreate) createSpec() (*Roomdetail, *sqlgraph.CreateSpec) {
	var (
		r     = &Roomdetail{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: roomdetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomdetail.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Roomnumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roomdetail.FieldRoomnumber,
		})
		r.Roomnumber = value
	}
	if value, ok := rc.mutation.Roomtypename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roomdetail.FieldRoomtypename,
		})
		r.Roomtypename = value
	}
	if value, ok := rc.mutation.Roomprice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roomdetail.FieldRoomprice,
		})
		r.Roomprice = value
	}
	if value, ok := rc.mutation.Sleep(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roomdetail.FieldSleep,
		})
		r.Sleep = value
	}
	if value, ok := rc.mutation.Bed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roomdetail.FieldBed,
		})
		r.Bed = value
	}
	if nodes := rc.mutation.PledgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roomdetail.PledgeTable,
			Columns: []string{roomdetail.PledgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pledge.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.PetruleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roomdetail.PetruleTable,
			Columns: []string{roomdetail.PetruleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: petrule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.BedtypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roomdetail.BedtypeTable,
			Columns: []string{roomdetail.BedtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bedtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roomdetail.EmployeeTable,
			Columns: []string{roomdetail.EmployeeColumn},
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
	if nodes := rc.mutation.StaytypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roomdetail.StaytypeTable,
			Columns: []string{roomdetail.StaytypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: staytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.RoomdetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   roomdetail.RoomdetailsTable,
			Columns: []string{roomdetail.RoomdetailsColumn},
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
