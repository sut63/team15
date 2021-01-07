// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/equipment"
	"github.com/team15/app/ent/facility"
	"github.com/team15/app/ent/nearbyplace"
	"github.com/team15/app/ent/quantity"
	"github.com/team15/app/ent/room"
	"github.com/team15/app/ent/staytype"
)

// RoomCreate is the builder for creating a Room entity.
type RoomCreate struct {
	config
	mutation *RoomMutation
	hooks    []Hook
}

// SetRoomprice sets the roomprice field.
func (rc *RoomCreate) SetRoomprice(i int) *RoomCreate {
	rc.mutation.SetRoomprice(i)
	return rc
}

// SetRoomtypename sets the roomtypename field.
func (rc *RoomCreate) SetRoomtypename(s string) *RoomCreate {
	rc.mutation.SetRoomtypename(s)
	return rc
}

// SetQuantityID sets the quantity edge to Quantity by id.
func (rc *RoomCreate) SetQuantityID(id int) *RoomCreate {
	rc.mutation.SetQuantityID(id)
	return rc
}

// SetNillableQuantityID sets the quantity edge to Quantity by id if the given value is not nil.
func (rc *RoomCreate) SetNillableQuantityID(id *int) *RoomCreate {
	if id != nil {
		rc = rc.SetQuantityID(*id)
	}
	return rc
}

// SetQuantity sets the quantity edge to Quantity.
func (rc *RoomCreate) SetQuantity(q *Quantity) *RoomCreate {
	return rc.SetQuantityID(q.ID)
}

// SetStaytypeID sets the staytype edge to StayType by id.
func (rc *RoomCreate) SetStaytypeID(id int) *RoomCreate {
	rc.mutation.SetStaytypeID(id)
	return rc
}

// SetNillableStaytypeID sets the staytype edge to StayType by id if the given value is not nil.
func (rc *RoomCreate) SetNillableStaytypeID(id *int) *RoomCreate {
	if id != nil {
		rc = rc.SetStaytypeID(*id)
	}
	return rc
}

// SetStaytype sets the staytype edge to StayType.
func (rc *RoomCreate) SetStaytype(s *StayType) *RoomCreate {
	return rc.SetStaytypeID(s.ID)
}

// AddFacilityIDs adds the facilities edge to Facility by ids.
func (rc *RoomCreate) AddFacilityIDs(ids ...int) *RoomCreate {
	rc.mutation.AddFacilityIDs(ids...)
	return rc
}

// AddFacilities adds the facilities edges to Facility.
func (rc *RoomCreate) AddFacilities(f ...*Facility) *RoomCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return rc.AddFacilityIDs(ids...)
}

// AddEquipmentIDs adds the equipments edge to Equipment by ids.
func (rc *RoomCreate) AddEquipmentIDs(ids ...int) *RoomCreate {
	rc.mutation.AddEquipmentIDs(ids...)
	return rc
}

// AddEquipments adds the equipments edges to Equipment.
func (rc *RoomCreate) AddEquipments(e ...*Equipment) *RoomCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rc.AddEquipmentIDs(ids...)
}

// AddNearbyplaceIDs adds the nearbyplace edge to NearbyPlace by ids.
func (rc *RoomCreate) AddNearbyplaceIDs(ids ...int) *RoomCreate {
	rc.mutation.AddNearbyplaceIDs(ids...)
	return rc
}

// AddNearbyplace adds the nearbyplace edges to NearbyPlace.
func (rc *RoomCreate) AddNearbyplace(n ...*NearbyPlace) *RoomCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return rc.AddNearbyplaceIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (rc *RoomCreate) Mutation() *RoomMutation {
	return rc.mutation
}

// Save creates the Room in the database.
func (rc *RoomCreate) Save(ctx context.Context) (*Room, error) {
	if _, ok := rc.mutation.Roomprice(); !ok {
		return nil, &ValidationError{Name: "roomprice", err: errors.New("ent: missing required field \"roomprice\"")}
	}
	if v, ok := rc.mutation.Roomprice(); ok {
		if err := room.RoompriceValidator(v); err != nil {
			return nil, &ValidationError{Name: "roomprice", err: fmt.Errorf("ent: validator failed for field \"roomprice\": %w", err)}
		}
	}
	if _, ok := rc.mutation.Roomtypename(); !ok {
		return nil, &ValidationError{Name: "roomtypename", err: errors.New("ent: missing required field \"roomtypename\"")}
	}
	var (
		err  error
		node *Room
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomMutation)
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
func (rc *RoomCreate) SaveX(ctx context.Context) *Room {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *RoomCreate) sqlSave(ctx context.Context) (*Room, error) {
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

func (rc *RoomCreate) createSpec() (*Room, *sqlgraph.CreateSpec) {
	var (
		r     = &Room{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: room.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: room.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Roomprice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldRoomprice,
		})
		r.Roomprice = value
	}
	if value, ok := rc.mutation.Roomtypename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: room.FieldRoomtypename,
		})
		r.Roomtypename = value
	}
	if nodes := rc.mutation.QuantityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   room.QuantityTable,
			Columns: []string{room.QuantityColumn},
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
			Inverse: false,
			Table:   room.StaytypeTable,
			Columns: []string{room.StaytypeColumn},
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
	if nodes := rc.mutation.FacilitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   room.FacilitiesTable,
			Columns: room.FacilitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facility.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.EquipmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   room.EquipmentsTable,
			Columns: room.EquipmentsPrimaryKey,
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
	if nodes := rc.mutation.NearbyplaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   room.NearbyplaceTable,
			Columns: room.NearbyplacePrimaryKey,
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
	return r, _spec
}
