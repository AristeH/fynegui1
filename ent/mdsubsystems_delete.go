// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"fynegui/ent/mdsubsystems"
	"fynegui/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MDSubSystemsDelete is the builder for deleting a MDSubSystems entity.
type MDSubSystemsDelete struct {
	config
	hooks    []Hook
	mutation *MDSubSystemsMutation
}

// Where appends a list predicates to the MDSubSystemsDelete builder.
func (mssd *MDSubSystemsDelete) Where(ps ...predicate.MDSubSystems) *MDSubSystemsDelete {
	mssd.mutation.Where(ps...)
	return mssd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mssd *MDSubSystemsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mssd.hooks) == 0 {
		affected, err = mssd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MDSubSystemsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mssd.mutation = mutation
			affected, err = mssd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mssd.hooks) - 1; i >= 0; i-- {
			if mssd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mssd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mssd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mssd *MDSubSystemsDelete) ExecX(ctx context.Context) int {
	n, err := mssd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mssd *MDSubSystemsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: mdsubsystems.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: mdsubsystems.FieldID,
			},
		},
	}
	if ps := mssd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, mssd.driver, _spec)
}

// MDSubSystemsDeleteOne is the builder for deleting a single MDSubSystems entity.
type MDSubSystemsDeleteOne struct {
	mssd *MDSubSystemsDelete
}

// Exec executes the deletion query.
func (mssdo *MDSubSystemsDeleteOne) Exec(ctx context.Context) error {
	n, err := mssdo.mssd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{mdsubsystems.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mssdo *MDSubSystemsDeleteOne) ExecX(ctx context.Context) {
	mssdo.mssd.ExecX(ctx)
}