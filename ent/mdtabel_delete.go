// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"fynegui/ent/mdtabel"
	"fynegui/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MDTabelDelete is the builder for deleting a MDTabel entity.
type MDTabelDelete struct {
	config
	hooks    []Hook
	mutation *MDTabelMutation
}

// Where appends a list predicates to the MDTabelDelete builder.
func (mtd *MDTabelDelete) Where(ps ...predicate.MDTabel) *MDTabelDelete {
	mtd.mutation.Where(ps...)
	return mtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mtd *MDTabelDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mtd.hooks) == 0 {
		affected, err = mtd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MDTabelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mtd.mutation = mutation
			affected, err = mtd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mtd.hooks) - 1; i >= 0; i-- {
			if mtd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mtd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mtd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtd *MDTabelDelete) ExecX(ctx context.Context) int {
	n, err := mtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mtd *MDTabelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: mdtabel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: mdtabel.FieldID,
			},
		},
	}
	if ps := mtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, mtd.driver, _spec)
}

// MDTabelDeleteOne is the builder for deleting a single MDTabel entity.
type MDTabelDeleteOne struct {
	mtd *MDTabelDelete
}

// Exec executes the deletion query.
func (mtdo *MDTabelDeleteOne) Exec(ctx context.Context) error {
	n, err := mtdo.mtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{mdtabel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mtdo *MDTabelDeleteOne) ExecX(ctx context.Context) {
	mtdo.mtd.ExecX(ctx)
}