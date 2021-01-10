// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team15/app/ent/lengthtime"
	"github.com/team15/app/ent/predicate"
)

// LengthTimeDelete is the builder for deleting a LengthTime entity.
type LengthTimeDelete struct {
	config
	hooks      []Hook
	mutation   *LengthTimeMutation
	predicates []predicate.LengthTime
}

// Where adds a new predicate to the delete builder.
func (ltd *LengthTimeDelete) Where(ps ...predicate.LengthTime) *LengthTimeDelete {
	ltd.predicates = append(ltd.predicates, ps...)
	return ltd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ltd *LengthTimeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ltd.hooks) == 0 {
		affected, err = ltd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LengthTimeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ltd.mutation = mutation
			affected, err = ltd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ltd.hooks) - 1; i >= 0; i-- {
			mut = ltd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ltd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltd *LengthTimeDelete) ExecX(ctx context.Context) int {
	n, err := ltd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ltd *LengthTimeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: lengthtime.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lengthtime.FieldID,
			},
		},
	}
	if ps := ltd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ltd.driver, _spec)
}

// LengthTimeDeleteOne is the builder for deleting a single LengthTime entity.
type LengthTimeDeleteOne struct {
	ltd *LengthTimeDelete
}

// Exec executes the deletion query.
func (ltdo *LengthTimeDeleteOne) Exec(ctx context.Context) error {
	n, err := ltdo.ltd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{lengthtime.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ltdo *LengthTimeDeleteOne) ExecX(ctx context.Context) {
	ltdo.ltd.ExecX(ctx)
}
