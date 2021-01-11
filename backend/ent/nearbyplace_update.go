// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/nearbyplace"
	"github.com/team15/app/ent/predicate"
	"github.com/team15/app/ent/roomdetail"
)

// NearbyplaceUpdate is the builder for updating Nearbyplace entities.
type NearbyplaceUpdate struct {
	config
	hooks      []Hook
	mutation   *NearbyplaceMutation
	predicates []predicate.Nearbyplace
}

// Where adds a new predicate for the builder.
func (nu *NearbyplaceUpdate) Where(ps ...predicate.Nearbyplace) *NearbyplaceUpdate {
	nu.predicates = append(nu.predicates, ps...)
	return nu
}

// SetNearbyplace sets the nearbyplace field.
func (nu *NearbyplaceUpdate) SetNearbyplace(s string) *NearbyplaceUpdate {
	nu.mutation.SetNearbyplace(s)
	return nu
}

// SetRoomdetailID sets the roomdetail edge to Roomdetail by id.
func (nu *NearbyplaceUpdate) SetRoomdetailID(id int) *NearbyplaceUpdate {
	nu.mutation.SetRoomdetailID(id)
	return nu
}

// SetNillableRoomdetailID sets the roomdetail edge to Roomdetail by id if the given value is not nil.
func (nu *NearbyplaceUpdate) SetNillableRoomdetailID(id *int) *NearbyplaceUpdate {
	if id != nil {
		nu = nu.SetRoomdetailID(*id)
	}
	return nu
}

// SetRoomdetail sets the roomdetail edge to Roomdetail.
func (nu *NearbyplaceUpdate) SetRoomdetail(r *Roomdetail) *NearbyplaceUpdate {
	return nu.SetRoomdetailID(r.ID)
}

// Mutation returns the NearbyplaceMutation object of the builder.
func (nu *NearbyplaceUpdate) Mutation() *NearbyplaceMutation {
	return nu.mutation
}

// ClearRoomdetail clears the roomdetail edge to Roomdetail.
func (nu *NearbyplaceUpdate) ClearRoomdetail() *NearbyplaceUpdate {
	nu.mutation.ClearRoomdetail()
	return nu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (nu *NearbyplaceUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(nu.hooks) == 0 {
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NearbyplaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NearbyplaceUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NearbyplaceUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NearbyplaceUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NearbyplaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nearbyplace.Table,
			Columns: nearbyplace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nearbyplace.FieldID,
			},
		},
	}
	if ps := nu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Nearbyplace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nearbyplace.FieldNearbyplace,
		})
	}
	if nu.mutation.RoomdetailCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nearbyplace.RoomdetailTable,
			Columns: []string{nearbyplace.RoomdetailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomdetail.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RoomdetailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nearbyplace.RoomdetailTable,
			Columns: []string{nearbyplace.RoomdetailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomdetail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nearbyplace.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// NearbyplaceUpdateOne is the builder for updating a single Nearbyplace entity.
type NearbyplaceUpdateOne struct {
	config
	hooks    []Hook
	mutation *NearbyplaceMutation
}

// SetNearbyplace sets the nearbyplace field.
func (nuo *NearbyplaceUpdateOne) SetNearbyplace(s string) *NearbyplaceUpdateOne {
	nuo.mutation.SetNearbyplace(s)
	return nuo
}

// SetRoomdetailID sets the roomdetail edge to Roomdetail by id.
func (nuo *NearbyplaceUpdateOne) SetRoomdetailID(id int) *NearbyplaceUpdateOne {
	nuo.mutation.SetRoomdetailID(id)
	return nuo
}

// SetNillableRoomdetailID sets the roomdetail edge to Roomdetail by id if the given value is not nil.
func (nuo *NearbyplaceUpdateOne) SetNillableRoomdetailID(id *int) *NearbyplaceUpdateOne {
	if id != nil {
		nuo = nuo.SetRoomdetailID(*id)
	}
	return nuo
}

// SetRoomdetail sets the roomdetail edge to Roomdetail.
func (nuo *NearbyplaceUpdateOne) SetRoomdetail(r *Roomdetail) *NearbyplaceUpdateOne {
	return nuo.SetRoomdetailID(r.ID)
}

// Mutation returns the NearbyplaceMutation object of the builder.
func (nuo *NearbyplaceUpdateOne) Mutation() *NearbyplaceMutation {
	return nuo.mutation
}

// ClearRoomdetail clears the roomdetail edge to Roomdetail.
func (nuo *NearbyplaceUpdateOne) ClearRoomdetail() *NearbyplaceUpdateOne {
	nuo.mutation.ClearRoomdetail()
	return nuo
}

// Save executes the query and returns the updated entity.
func (nuo *NearbyplaceUpdateOne) Save(ctx context.Context) (*Nearbyplace, error) {

	var (
		err  error
		node *Nearbyplace
	)
	if len(nuo.hooks) == 0 {
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NearbyplaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			mut = nuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NearbyplaceUpdateOne) SaveX(ctx context.Context) *Nearbyplace {
	n, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

// Exec executes the query on the entity.
func (nuo *NearbyplaceUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NearbyplaceUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NearbyplaceUpdateOne) sqlSave(ctx context.Context) (n *Nearbyplace, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nearbyplace.Table,
			Columns: nearbyplace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nearbyplace.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Nearbyplace.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := nuo.mutation.Nearbyplace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nearbyplace.FieldNearbyplace,
		})
	}
	if nuo.mutation.RoomdetailCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nearbyplace.RoomdetailTable,
			Columns: []string{nearbyplace.RoomdetailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomdetail.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RoomdetailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nearbyplace.RoomdetailTable,
			Columns: []string{nearbyplace.RoomdetailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomdetail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	n = &Nearbyplace{config: nuo.config}
	_spec.Assign = n.assignValues
	_spec.ScanValues = n.scanValues()
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nearbyplace.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return n, nil
}
