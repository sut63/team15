// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/equipment"
	"github.com/team15/app/ent/facilitie"
	"github.com/team15/app/ent/lease"
	"github.com/team15/app/ent/nearbyplace"
	"github.com/team15/app/ent/quantity"
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

// SetEquipmentsID sets the equipments edge to Equipment by id.
func (rc *RoomdetailCreate) SetEquipmentsID(id int) *RoomdetailCreate {
	rc.mutation.SetEquipmentsID(id)
	return rc
}

// SetNillableEquipmentsID sets the equipments edge to Equipment by id if the given value is not nil.
func (rc *RoomdetailCreate) SetNillableEquipmentsID(id *int) *RoomdetailCreate {
	if id != nil {
		rc = rc.SetEquipmentsID(*id)
	}
	return rc
}

// SetEquipments sets the equipments edge to Equipment.
func (rc *RoomdetailCreate) SetEquipments(e *Equipment) *RoomdetailCreate {
	return rc.SetEquipmentsID(e.ID)
}

// SetFacilitiesID sets the facilities edge to Facilitie by id.
func (rc *RoomdetailCreate) SetFacilitiesID(id int) *RoomdetailCreate {
	rc.mutation.SetFacilitiesID(id)
	return rc
}

// SetNillableFacilitiesID sets the facilities edge to Facilitie by id if the given value is not nil.
func (rc *RoomdetailCreate) SetNillableFacilitiesID(id *int) *RoomdetailCreate {
	if id != nil {
		rc = rc.SetFacilitiesID(*id)
	}
	return rc
}

// SetFacilities sets the facilities edge to Facilitie.
func (rc *RoomdetailCreate) SetFacilities(f *Facilitie) *RoomdetailCreate {
	return rc.SetFacilitiesID(f.ID)
}

// SetNearbyplacesID sets the nearbyplaces edge to Nearbyplace by id.
func (rc *RoomdetailCreate) SetNearbyplacesID(id int) *RoomdetailCreate {
	rc.mutation.SetNearbyplacesID(id)
	return rc
}

// SetNillableNearbyplacesID sets the nearbyplaces edge to Nearbyplace by id if the given value is not nil.
func (rc *RoomdetailCreate) SetNillableNearbyplacesID(id *int) *RoomdetailCreate {
	if id != nil {
		rc = rc.SetNearbyplacesID(*id)
	}
	return rc
}

// SetNearbyplaces sets the nearbyplaces edge to Nearbyplace.
func (rc *RoomdetailCreate) SetNearbyplaces(n *Nearbyplace) *RoomdetailCreate {
	return rc.SetNearbyplacesID(n.ID)
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

// SetQuantityID sets the quantity edge to Quantity by id.
func (rc *RoomdetailCreate) SetQuantityID(id int) *RoomdetailCreate {
	rc.mutation.SetQuantityID(id)
	return rc
}

// SetNillableQuantityID sets the quantity edge to Quantity by id if the given value is not nil.
func (rc *RoomdetailCreate) SetNillableQuantityID(id *int) *RoomdetailCreate {
	if id != nil {
		rc = rc.SetQuantityID(*id)
	}
	return rc
}

// SetQuantity sets the quantity edge to Quantity.
func (rc *RoomdetailCreate) SetQuantity(q *Quantity) *RoomdetailCreate {
	return rc.SetQuantityID(q.ID)
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
	if v, ok := rc.mutation.Roomnumber(); ok {
		if err := roomdetail.RoomnumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "roomnumber", err: fmt.Errorf("ent: validator failed for field \"roomnumber\": %w", err)}
		}
	}
	if _, ok := rc.mutation.Roomtypename(); !ok {
		return nil, &ValidationError{Name: "roomtypename", err: errors.New("ent: missing required field \"roomtypename\"")}
	}
	if v, ok := rc.mutation.Roomtypename(); ok {
		if err := roomdetail.RoomtypenameValidator(v); err != nil {
			return nil, &ValidationError{Name: "roomtypename", err: fmt.Errorf("ent: validator failed for field \"roomtypename\": %w", err)}
		}
	}
	if _, ok := rc.mutation.Roomprice(); !ok {
		return nil, &ValidationError{Name: "roomprice", err: errors.New("ent: missing required field \"roomprice\"")}
	}
	if v, ok := rc.mutation.Roomprice(); ok {
		if err := roomdetail.RoompriceValidator(v); err != nil {
			return nil, &ValidationError{Name: "roomprice", err: fmt.Errorf("ent: validator failed for field \"roomprice\": %w", err)}
		}
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
	if nodes := rc.mutation.EquipmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roomdetail.EquipmentsTable,
			Columns: []string{roomdetail.EquipmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.FacilitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roomdetail.FacilitiesTable,
			Columns: []string{roomdetail.FacilitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitie.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.NearbyplacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roomdetail.NearbyplacesTable,
			Columns: []string{roomdetail.NearbyplacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nearbyplace.FieldID,
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
	if nodes := rc.mutation.QuantityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roomdetail.QuantityTable,
			Columns: []string{roomdetail.QuantityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: quantity.FieldID,
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
