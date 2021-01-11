// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/cleanername"
	"github.com/team15/app/ent/cleaningroom"
	"github.com/team15/app/ent/lengthtime"
	"github.com/team15/app/ent/predicate"
	"github.com/team15/app/ent/room"
)

// CleaningRoomUpdate is the builder for updating CleaningRoom entities.
type CleaningRoomUpdate struct {
	config
	hooks      []Hook
	mutation   *CleaningRoomMutation
	predicates []predicate.CleaningRoom
}

// Where adds a new predicate for the builder.
func (cru *CleaningRoomUpdate) Where(ps ...predicate.CleaningRoom) *CleaningRoomUpdate {
	cru.predicates = append(cru.predicates, ps...)
	return cru
}

// SetDateandstarttime sets the dateandstarttime field.
func (cru *CleaningRoomUpdate) SetDateandstarttime(t time.Time) *CleaningRoomUpdate {
	cru.mutation.SetDateandstarttime(t)
	return cru
}

// SetNote sets the note field.
func (cru *CleaningRoomUpdate) SetNote(s string) *CleaningRoomUpdate {
	cru.mutation.SetNote(s)
	return cru
}

// SetRoomID sets the Room edge to Room by id.
func (cru *CleaningRoomUpdate) SetRoomID(id int) *CleaningRoomUpdate {
	cru.mutation.SetRoomID(id)
	return cru
}

// SetNillableRoomID sets the Room edge to Room by id if the given value is not nil.
func (cru *CleaningRoomUpdate) SetNillableRoomID(id *int) *CleaningRoomUpdate {
	if id != nil {
		cru = cru.SetRoomID(*id)
	}
	return cru
}

// SetRoom sets the Room edge to Room.
func (cru *CleaningRoomUpdate) SetRoom(r *Room) *CleaningRoomUpdate {
	return cru.SetRoomID(r.ID)
}

// SetCleanerNameID sets the CleanerName edge to CleanerName by id.
func (cru *CleaningRoomUpdate) SetCleanerNameID(id int) *CleaningRoomUpdate {
	cru.mutation.SetCleanerNameID(id)
	return cru
}

// SetNillableCleanerNameID sets the CleanerName edge to CleanerName by id if the given value is not nil.
func (cru *CleaningRoomUpdate) SetNillableCleanerNameID(id *int) *CleaningRoomUpdate {
	if id != nil {
		cru = cru.SetCleanerNameID(*id)
	}
	return cru
}

// SetCleanerName sets the CleanerName edge to CleanerName.
func (cru *CleaningRoomUpdate) SetCleanerName(c *CleanerName) *CleaningRoomUpdate {
	return cru.SetCleanerNameID(c.ID)
}

// SetLengthTimeID sets the LengthTime edge to LengthTime by id.
func (cru *CleaningRoomUpdate) SetLengthTimeID(id int) *CleaningRoomUpdate {
	cru.mutation.SetLengthTimeID(id)
	return cru
}

// SetNillableLengthTimeID sets the LengthTime edge to LengthTime by id if the given value is not nil.
func (cru *CleaningRoomUpdate) SetNillableLengthTimeID(id *int) *CleaningRoomUpdate {
	if id != nil {
		cru = cru.SetLengthTimeID(*id)
	}
	return cru
}

// SetLengthTime sets the LengthTime edge to LengthTime.
func (cru *CleaningRoomUpdate) SetLengthTime(l *LengthTime) *CleaningRoomUpdate {
	return cru.SetLengthTimeID(l.ID)
}

// Mutation returns the CleaningRoomMutation object of the builder.
func (cru *CleaningRoomUpdate) Mutation() *CleaningRoomMutation {
	return cru.mutation
}

// ClearRoom clears the Room edge to Room.
func (cru *CleaningRoomUpdate) ClearRoom() *CleaningRoomUpdate {
	cru.mutation.ClearRoom()
	return cru
}

// ClearCleanerName clears the CleanerName edge to CleanerName.
func (cru *CleaningRoomUpdate) ClearCleanerName() *CleaningRoomUpdate {
	cru.mutation.ClearCleanerName()
	return cru
}

// ClearLengthTime clears the LengthTime edge to LengthTime.
func (cru *CleaningRoomUpdate) ClearLengthTime() *CleaningRoomUpdate {
	cru.mutation.ClearLengthTime()
	return cru
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cru *CleaningRoomUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(cru.hooks) == 0 {
		affected, err = cru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CleaningRoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cru.mutation = mutation
			affected, err = cru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cru.hooks) - 1; i >= 0; i-- {
			mut = cru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cru *CleaningRoomUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *CleaningRoomUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *CleaningRoomUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cru *CleaningRoomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cleaningroom.Table,
			Columns: cleaningroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cleaningroom.FieldID,
			},
		},
	}
	if ps := cru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cru.mutation.Dateandstarttime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cleaningroom.FieldDateandstarttime,
		})
	}
	if value, ok := cru.mutation.Note(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cleaningroom.FieldNote,
		})
	}
	if cru.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cleaningroom.RoomTable,
			Columns: []string{cleaningroom.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cru.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cleaningroom.RoomTable,
			Columns: []string{cleaningroom.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cru.mutation.CleanerNameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cleaningroom.CleanerNameTable,
			Columns: []string{cleaningroom.CleanerNameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cleanername.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cru.mutation.CleanerNameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cleaningroom.CleanerNameTable,
			Columns: []string{cleaningroom.CleanerNameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cleanername.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cru.mutation.LengthTimeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cleaningroom.LengthTimeTable,
			Columns: []string{cleaningroom.LengthTimeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lengthtime.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cru.mutation.LengthTimeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cleaningroom.LengthTimeTable,
			Columns: []string{cleaningroom.LengthTimeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lengthtime.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cleaningroom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CleaningRoomUpdateOne is the builder for updating a single CleaningRoom entity.
type CleaningRoomUpdateOne struct {
	config
	hooks    []Hook
	mutation *CleaningRoomMutation
}

// SetDateandstarttime sets the dateandstarttime field.
func (cruo *CleaningRoomUpdateOne) SetDateandstarttime(t time.Time) *CleaningRoomUpdateOne {
	cruo.mutation.SetDateandstarttime(t)
	return cruo
}

// SetNote sets the note field.
func (cruo *CleaningRoomUpdateOne) SetNote(s string) *CleaningRoomUpdateOne {
	cruo.mutation.SetNote(s)
	return cruo
}

// SetRoomID sets the Room edge to Room by id.
func (cruo *CleaningRoomUpdateOne) SetRoomID(id int) *CleaningRoomUpdateOne {
	cruo.mutation.SetRoomID(id)
	return cruo
}

// SetNillableRoomID sets the Room edge to Room by id if the given value is not nil.
func (cruo *CleaningRoomUpdateOne) SetNillableRoomID(id *int) *CleaningRoomUpdateOne {
	if id != nil {
		cruo = cruo.SetRoomID(*id)
	}
	return cruo
}

// SetRoom sets the Room edge to Room.
func (cruo *CleaningRoomUpdateOne) SetRoom(r *Room) *CleaningRoomUpdateOne {
	return cruo.SetRoomID(r.ID)
}

// SetCleanerNameID sets the CleanerName edge to CleanerName by id.
func (cruo *CleaningRoomUpdateOne) SetCleanerNameID(id int) *CleaningRoomUpdateOne {
	cruo.mutation.SetCleanerNameID(id)
	return cruo
}

// SetNillableCleanerNameID sets the CleanerName edge to CleanerName by id if the given value is not nil.
func (cruo *CleaningRoomUpdateOne) SetNillableCleanerNameID(id *int) *CleaningRoomUpdateOne {
	if id != nil {
		cruo = cruo.SetCleanerNameID(*id)
	}
	return cruo
}

// SetCleanerName sets the CleanerName edge to CleanerName.
func (cruo *CleaningRoomUpdateOne) SetCleanerName(c *CleanerName) *CleaningRoomUpdateOne {
	return cruo.SetCleanerNameID(c.ID)
}

// SetLengthTimeID sets the LengthTime edge to LengthTime by id.
func (cruo *CleaningRoomUpdateOne) SetLengthTimeID(id int) *CleaningRoomUpdateOne {
	cruo.mutation.SetLengthTimeID(id)
	return cruo
}

// SetNillableLengthTimeID sets the LengthTime edge to LengthTime by id if the given value is not nil.
func (cruo *CleaningRoomUpdateOne) SetNillableLengthTimeID(id *int) *CleaningRoomUpdateOne {
	if id != nil {
		cruo = cruo.SetLengthTimeID(*id)
	}
	return cruo
}

// SetLengthTime sets the LengthTime edge to LengthTime.
func (cruo *CleaningRoomUpdateOne) SetLengthTime(l *LengthTime) *CleaningRoomUpdateOne {
	return cruo.SetLengthTimeID(l.ID)
}

// Mutation returns the CleaningRoomMutation object of the builder.
func (cruo *CleaningRoomUpdateOne) Mutation() *CleaningRoomMutation {
	return cruo.mutation
}

// ClearRoom clears the Room edge to Room.
func (cruo *CleaningRoomUpdateOne) ClearRoom() *CleaningRoomUpdateOne {
	cruo.mutation.ClearRoom()
	return cruo
}

// ClearCleanerName clears the CleanerName edge to CleanerName.
func (cruo *CleaningRoomUpdateOne) ClearCleanerName() *CleaningRoomUpdateOne {
	cruo.mutation.ClearCleanerName()
	return cruo
}

// ClearLengthTime clears the LengthTime edge to LengthTime.
func (cruo *CleaningRoomUpdateOne) ClearLengthTime() *CleaningRoomUpdateOne {
	cruo.mutation.ClearLengthTime()
	return cruo
}

// Save executes the query and returns the updated entity.
func (cruo *CleaningRoomUpdateOne) Save(ctx context.Context) (*CleaningRoom, error) {

	var (
		err  error
		node *CleaningRoom
	)
	if len(cruo.hooks) == 0 {
		node, err = cruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CleaningRoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cruo.mutation = mutation
			node, err = cruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cruo.hooks) - 1; i >= 0; i-- {
			mut = cruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *CleaningRoomUpdateOne) SaveX(ctx context.Context) *CleaningRoom {
	cr, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return cr
}

// Exec executes the query on the entity.
func (cruo *CleaningRoomUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *CleaningRoomUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cruo *CleaningRoomUpdateOne) sqlSave(ctx context.Context) (cr *CleaningRoom, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cleaningroom.Table,
			Columns: cleaningroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cleaningroom.FieldID,
			},
		},
	}
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CleaningRoom.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cruo.mutation.Dateandstarttime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cleaningroom.FieldDateandstarttime,
		})
	}
	if value, ok := cruo.mutation.Note(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cleaningroom.FieldNote,
		})
	}
	if cruo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cleaningroom.RoomTable,
			Columns: []string{cleaningroom.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cruo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cleaningroom.RoomTable,
			Columns: []string{cleaningroom.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cruo.mutation.CleanerNameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cleaningroom.CleanerNameTable,
			Columns: []string{cleaningroom.CleanerNameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cleanername.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cruo.mutation.CleanerNameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cleaningroom.CleanerNameTable,
			Columns: []string{cleaningroom.CleanerNameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cleanername.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cruo.mutation.LengthTimeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cleaningroom.LengthTimeTable,
			Columns: []string{cleaningroom.LengthTimeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lengthtime.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cruo.mutation.LengthTimeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cleaningroom.LengthTimeTable,
			Columns: []string{cleaningroom.LengthTimeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lengthtime.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	cr = &CleaningRoom{config: cruo.config}
	_spec.Assign = cr.assignValues
	_spec.ScanValues = cr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cleaningroom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return cr, nil
}