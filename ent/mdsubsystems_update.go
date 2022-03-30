// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"fynegui/ent/mdsubsystems"
	"fynegui/ent/mdtabel"
	"fynegui/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MDSubSystemsUpdate is the builder for updating MDSubSystems entities.
type MDSubSystemsUpdate struct {
	config
	hooks    []Hook
	mutation *MDSubSystemsMutation
}

// Where appends a list predicates to the MDSubSystemsUpdate builder.
func (mssu *MDSubSystemsUpdate) Where(ps ...predicate.MDSubSystems) *MDSubSystemsUpdate {
	mssu.mutation.Where(ps...)
	return mssu
}

// SetNamerus sets the "namerus" field.
func (mssu *MDSubSystemsUpdate) SetNamerus(s string) *MDSubSystemsUpdate {
	mssu.mutation.SetNamerus(s)
	return mssu
}

// SetNameeng sets the "nameeng" field.
func (mssu *MDSubSystemsUpdate) SetNameeng(s string) *MDSubSystemsUpdate {
	mssu.mutation.SetNameeng(s)
	return mssu
}

// SetSynonym sets the "synonym" field.
func (mssu *MDSubSystemsUpdate) SetSynonym(s string) *MDSubSystemsUpdate {
	mssu.mutation.SetSynonym(s)
	return mssu
}

// SetParent sets the "parent" field.
func (mssu *MDSubSystemsUpdate) SetParent(s string) *MDSubSystemsUpdate {
	mssu.mutation.SetParent(s)
	return mssu
}

// SetNillableParent sets the "parent" field if the given value is not nil.
func (mssu *MDSubSystemsUpdate) SetNillableParent(s *string) *MDSubSystemsUpdate {
	if s != nil {
		mssu.SetParent(*s)
	}
	return mssu
}

// ClearParent clears the value of the "parent" field.
func (mssu *MDSubSystemsUpdate) ClearParent() *MDSubSystemsUpdate {
	mssu.mutation.ClearParent()
	return mssu
}

// AddChildMdsubsystemIDs adds the "child_mdsubsystems" edge to the MDSubSystems entity by IDs.
func (mssu *MDSubSystemsUpdate) AddChildMdsubsystemIDs(ids ...string) *MDSubSystemsUpdate {
	mssu.mutation.AddChildMdsubsystemIDs(ids...)
	return mssu
}

// AddChildMdsubsystems adds the "child_mdsubsystems" edges to the MDSubSystems entity.
func (mssu *MDSubSystemsUpdate) AddChildMdsubsystems(m ...*MDSubSystems) *MDSubSystemsUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mssu.AddChildMdsubsystemIDs(ids...)
}

// SetParentMdsubsystemsID sets the "parent_mdsubsystems" edge to the MDSubSystems entity by ID.
func (mssu *MDSubSystemsUpdate) SetParentMdsubsystemsID(id string) *MDSubSystemsUpdate {
	mssu.mutation.SetParentMdsubsystemsID(id)
	return mssu
}

// SetNillableParentMdsubsystemsID sets the "parent_mdsubsystems" edge to the MDSubSystems entity by ID if the given value is not nil.
func (mssu *MDSubSystemsUpdate) SetNillableParentMdsubsystemsID(id *string) *MDSubSystemsUpdate {
	if id != nil {
		mssu = mssu.SetParentMdsubsystemsID(*id)
	}
	return mssu
}

// SetParentMdsubsystems sets the "parent_mdsubsystems" edge to the MDSubSystems entity.
func (mssu *MDSubSystemsUpdate) SetParentMdsubsystems(m *MDSubSystems) *MDSubSystemsUpdate {
	return mssu.SetParentMdsubsystemsID(m.ID)
}

// AddMdtableIDs adds the "mdtables" edge to the MDTabel entity by IDs.
func (mssu *MDSubSystemsUpdate) AddMdtableIDs(ids ...string) *MDSubSystemsUpdate {
	mssu.mutation.AddMdtableIDs(ids...)
	return mssu
}

// AddMdtables adds the "mdtables" edges to the MDTabel entity.
func (mssu *MDSubSystemsUpdate) AddMdtables(m ...*MDTabel) *MDSubSystemsUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mssu.AddMdtableIDs(ids...)
}

// Mutation returns the MDSubSystemsMutation object of the builder.
func (mssu *MDSubSystemsUpdate) Mutation() *MDSubSystemsMutation {
	return mssu.mutation
}

// ClearChildMdsubsystems clears all "child_mdsubsystems" edges to the MDSubSystems entity.
func (mssu *MDSubSystemsUpdate) ClearChildMdsubsystems() *MDSubSystemsUpdate {
	mssu.mutation.ClearChildMdsubsystems()
	return mssu
}

// RemoveChildMdsubsystemIDs removes the "child_mdsubsystems" edge to MDSubSystems entities by IDs.
func (mssu *MDSubSystemsUpdate) RemoveChildMdsubsystemIDs(ids ...string) *MDSubSystemsUpdate {
	mssu.mutation.RemoveChildMdsubsystemIDs(ids...)
	return mssu
}

// RemoveChildMdsubsystems removes "child_mdsubsystems" edges to MDSubSystems entities.
func (mssu *MDSubSystemsUpdate) RemoveChildMdsubsystems(m ...*MDSubSystems) *MDSubSystemsUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mssu.RemoveChildMdsubsystemIDs(ids...)
}

// ClearParentMdsubsystems clears the "parent_mdsubsystems" edge to the MDSubSystems entity.
func (mssu *MDSubSystemsUpdate) ClearParentMdsubsystems() *MDSubSystemsUpdate {
	mssu.mutation.ClearParentMdsubsystems()
	return mssu
}

// ClearMdtables clears all "mdtables" edges to the MDTabel entity.
func (mssu *MDSubSystemsUpdate) ClearMdtables() *MDSubSystemsUpdate {
	mssu.mutation.ClearMdtables()
	return mssu
}

// RemoveMdtableIDs removes the "mdtables" edge to MDTabel entities by IDs.
func (mssu *MDSubSystemsUpdate) RemoveMdtableIDs(ids ...string) *MDSubSystemsUpdate {
	mssu.mutation.RemoveMdtableIDs(ids...)
	return mssu
}

// RemoveMdtables removes "mdtables" edges to MDTabel entities.
func (mssu *MDSubSystemsUpdate) RemoveMdtables(m ...*MDTabel) *MDSubSystemsUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mssu.RemoveMdtableIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mssu *MDSubSystemsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mssu.hooks) == 0 {
		if err = mssu.check(); err != nil {
			return 0, err
		}
		affected, err = mssu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MDSubSystemsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mssu.check(); err != nil {
				return 0, err
			}
			mssu.mutation = mutation
			affected, err = mssu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mssu.hooks) - 1; i >= 0; i-- {
			if mssu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mssu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mssu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mssu *MDSubSystemsUpdate) SaveX(ctx context.Context) int {
	affected, err := mssu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mssu *MDSubSystemsUpdate) Exec(ctx context.Context) error {
	_, err := mssu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mssu *MDSubSystemsUpdate) ExecX(ctx context.Context) {
	if err := mssu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mssu *MDSubSystemsUpdate) check() error {
	if v, ok := mssu.mutation.Namerus(); ok {
		if err := mdsubsystems.NamerusValidator(v); err != nil {
			return &ValidationError{Name: "namerus", err: fmt.Errorf(`ent: validator failed for field "MDSubSystems.namerus": %w`, err)}
		}
	}
	if v, ok := mssu.mutation.Nameeng(); ok {
		if err := mdsubsystems.NameengValidator(v); err != nil {
			return &ValidationError{Name: "nameeng", err: fmt.Errorf(`ent: validator failed for field "MDSubSystems.nameeng": %w`, err)}
		}
	}
	if v, ok := mssu.mutation.Synonym(); ok {
		if err := mdsubsystems.SynonymValidator(v); err != nil {
			return &ValidationError{Name: "synonym", err: fmt.Errorf(`ent: validator failed for field "MDSubSystems.synonym": %w`, err)}
		}
	}
	return nil
}

func (mssu *MDSubSystemsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mdsubsystems.Table,
			Columns: mdsubsystems.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: mdsubsystems.FieldID,
			},
		},
	}
	if ps := mssu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mssu.mutation.Namerus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mdsubsystems.FieldNamerus,
		})
	}
	if value, ok := mssu.mutation.Nameeng(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mdsubsystems.FieldNameeng,
		})
	}
	if value, ok := mssu.mutation.Synonym(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mdsubsystems.FieldSynonym,
		})
	}
	if mssu.mutation.ChildMdsubsystemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mdsubsystems.ChildMdsubsystemsTable,
			Columns: []string{mdsubsystems.ChildMdsubsystemsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdsubsystems.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mssu.mutation.RemovedChildMdsubsystemsIDs(); len(nodes) > 0 && !mssu.mutation.ChildMdsubsystemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mdsubsystems.ChildMdsubsystemsTable,
			Columns: []string{mdsubsystems.ChildMdsubsystemsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdsubsystems.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mssu.mutation.ChildMdsubsystemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mdsubsystems.ChildMdsubsystemsTable,
			Columns: []string{mdsubsystems.ChildMdsubsystemsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdsubsystems.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mssu.mutation.ParentMdsubsystemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mdsubsystems.ParentMdsubsystemsTable,
			Columns: []string{mdsubsystems.ParentMdsubsystemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdsubsystems.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mssu.mutation.ParentMdsubsystemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mdsubsystems.ParentMdsubsystemsTable,
			Columns: []string{mdsubsystems.ParentMdsubsystemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdsubsystems.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mssu.mutation.MdtablesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mdsubsystems.MdtablesTable,
			Columns: mdsubsystems.MdtablesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdtabel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mssu.mutation.RemovedMdtablesIDs(); len(nodes) > 0 && !mssu.mutation.MdtablesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mdsubsystems.MdtablesTable,
			Columns: mdsubsystems.MdtablesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdtabel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mssu.mutation.MdtablesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mdsubsystems.MdtablesTable,
			Columns: mdsubsystems.MdtablesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdtabel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mssu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mdsubsystems.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MDSubSystemsUpdateOne is the builder for updating a single MDSubSystems entity.
type MDSubSystemsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MDSubSystemsMutation
}

// SetNamerus sets the "namerus" field.
func (mssuo *MDSubSystemsUpdateOne) SetNamerus(s string) *MDSubSystemsUpdateOne {
	mssuo.mutation.SetNamerus(s)
	return mssuo
}

// SetNameeng sets the "nameeng" field.
func (mssuo *MDSubSystemsUpdateOne) SetNameeng(s string) *MDSubSystemsUpdateOne {
	mssuo.mutation.SetNameeng(s)
	return mssuo
}

// SetSynonym sets the "synonym" field.
func (mssuo *MDSubSystemsUpdateOne) SetSynonym(s string) *MDSubSystemsUpdateOne {
	mssuo.mutation.SetSynonym(s)
	return mssuo
}

// SetParent sets the "parent" field.
func (mssuo *MDSubSystemsUpdateOne) SetParent(s string) *MDSubSystemsUpdateOne {
	mssuo.mutation.SetParent(s)
	return mssuo
}

// SetNillableParent sets the "parent" field if the given value is not nil.
func (mssuo *MDSubSystemsUpdateOne) SetNillableParent(s *string) *MDSubSystemsUpdateOne {
	if s != nil {
		mssuo.SetParent(*s)
	}
	return mssuo
}

// ClearParent clears the value of the "parent" field.
func (mssuo *MDSubSystemsUpdateOne) ClearParent() *MDSubSystemsUpdateOne {
	mssuo.mutation.ClearParent()
	return mssuo
}

// AddChildMdsubsystemIDs adds the "child_mdsubsystems" edge to the MDSubSystems entity by IDs.
func (mssuo *MDSubSystemsUpdateOne) AddChildMdsubsystemIDs(ids ...string) *MDSubSystemsUpdateOne {
	mssuo.mutation.AddChildMdsubsystemIDs(ids...)
	return mssuo
}

// AddChildMdsubsystems adds the "child_mdsubsystems" edges to the MDSubSystems entity.
func (mssuo *MDSubSystemsUpdateOne) AddChildMdsubsystems(m ...*MDSubSystems) *MDSubSystemsUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mssuo.AddChildMdsubsystemIDs(ids...)
}

// SetParentMdsubsystemsID sets the "parent_mdsubsystems" edge to the MDSubSystems entity by ID.
func (mssuo *MDSubSystemsUpdateOne) SetParentMdsubsystemsID(id string) *MDSubSystemsUpdateOne {
	mssuo.mutation.SetParentMdsubsystemsID(id)
	return mssuo
}

// SetNillableParentMdsubsystemsID sets the "parent_mdsubsystems" edge to the MDSubSystems entity by ID if the given value is not nil.
func (mssuo *MDSubSystemsUpdateOne) SetNillableParentMdsubsystemsID(id *string) *MDSubSystemsUpdateOne {
	if id != nil {
		mssuo = mssuo.SetParentMdsubsystemsID(*id)
	}
	return mssuo
}

// SetParentMdsubsystems sets the "parent_mdsubsystems" edge to the MDSubSystems entity.
func (mssuo *MDSubSystemsUpdateOne) SetParentMdsubsystems(m *MDSubSystems) *MDSubSystemsUpdateOne {
	return mssuo.SetParentMdsubsystemsID(m.ID)
}

// AddMdtableIDs adds the "mdtables" edge to the MDTabel entity by IDs.
func (mssuo *MDSubSystemsUpdateOne) AddMdtableIDs(ids ...string) *MDSubSystemsUpdateOne {
	mssuo.mutation.AddMdtableIDs(ids...)
	return mssuo
}

// AddMdtables adds the "mdtables" edges to the MDTabel entity.
func (mssuo *MDSubSystemsUpdateOne) AddMdtables(m ...*MDTabel) *MDSubSystemsUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mssuo.AddMdtableIDs(ids...)
}

// Mutation returns the MDSubSystemsMutation object of the builder.
func (mssuo *MDSubSystemsUpdateOne) Mutation() *MDSubSystemsMutation {
	return mssuo.mutation
}

// ClearChildMdsubsystems clears all "child_mdsubsystems" edges to the MDSubSystems entity.
func (mssuo *MDSubSystemsUpdateOne) ClearChildMdsubsystems() *MDSubSystemsUpdateOne {
	mssuo.mutation.ClearChildMdsubsystems()
	return mssuo
}

// RemoveChildMdsubsystemIDs removes the "child_mdsubsystems" edge to MDSubSystems entities by IDs.
func (mssuo *MDSubSystemsUpdateOne) RemoveChildMdsubsystemIDs(ids ...string) *MDSubSystemsUpdateOne {
	mssuo.mutation.RemoveChildMdsubsystemIDs(ids...)
	return mssuo
}

// RemoveChildMdsubsystems removes "child_mdsubsystems" edges to MDSubSystems entities.
func (mssuo *MDSubSystemsUpdateOne) RemoveChildMdsubsystems(m ...*MDSubSystems) *MDSubSystemsUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mssuo.RemoveChildMdsubsystemIDs(ids...)
}

// ClearParentMdsubsystems clears the "parent_mdsubsystems" edge to the MDSubSystems entity.
func (mssuo *MDSubSystemsUpdateOne) ClearParentMdsubsystems() *MDSubSystemsUpdateOne {
	mssuo.mutation.ClearParentMdsubsystems()
	return mssuo
}

// ClearMdtables clears all "mdtables" edges to the MDTabel entity.
func (mssuo *MDSubSystemsUpdateOne) ClearMdtables() *MDSubSystemsUpdateOne {
	mssuo.mutation.ClearMdtables()
	return mssuo
}

// RemoveMdtableIDs removes the "mdtables" edge to MDTabel entities by IDs.
func (mssuo *MDSubSystemsUpdateOne) RemoveMdtableIDs(ids ...string) *MDSubSystemsUpdateOne {
	mssuo.mutation.RemoveMdtableIDs(ids...)
	return mssuo
}

// RemoveMdtables removes "mdtables" edges to MDTabel entities.
func (mssuo *MDSubSystemsUpdateOne) RemoveMdtables(m ...*MDTabel) *MDSubSystemsUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mssuo.RemoveMdtableIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mssuo *MDSubSystemsUpdateOne) Select(field string, fields ...string) *MDSubSystemsUpdateOne {
	mssuo.fields = append([]string{field}, fields...)
	return mssuo
}

// Save executes the query and returns the updated MDSubSystems entity.
func (mssuo *MDSubSystemsUpdateOne) Save(ctx context.Context) (*MDSubSystems, error) {
	var (
		err  error
		node *MDSubSystems
	)
	if len(mssuo.hooks) == 0 {
		if err = mssuo.check(); err != nil {
			return nil, err
		}
		node, err = mssuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MDSubSystemsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mssuo.check(); err != nil {
				return nil, err
			}
			mssuo.mutation = mutation
			node, err = mssuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mssuo.hooks) - 1; i >= 0; i-- {
			if mssuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mssuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mssuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mssuo *MDSubSystemsUpdateOne) SaveX(ctx context.Context) *MDSubSystems {
	node, err := mssuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mssuo *MDSubSystemsUpdateOne) Exec(ctx context.Context) error {
	_, err := mssuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mssuo *MDSubSystemsUpdateOne) ExecX(ctx context.Context) {
	if err := mssuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mssuo *MDSubSystemsUpdateOne) check() error {
	if v, ok := mssuo.mutation.Namerus(); ok {
		if err := mdsubsystems.NamerusValidator(v); err != nil {
			return &ValidationError{Name: "namerus", err: fmt.Errorf(`ent: validator failed for field "MDSubSystems.namerus": %w`, err)}
		}
	}
	if v, ok := mssuo.mutation.Nameeng(); ok {
		if err := mdsubsystems.NameengValidator(v); err != nil {
			return &ValidationError{Name: "nameeng", err: fmt.Errorf(`ent: validator failed for field "MDSubSystems.nameeng": %w`, err)}
		}
	}
	if v, ok := mssuo.mutation.Synonym(); ok {
		if err := mdsubsystems.SynonymValidator(v); err != nil {
			return &ValidationError{Name: "synonym", err: fmt.Errorf(`ent: validator failed for field "MDSubSystems.synonym": %w`, err)}
		}
	}
	return nil
}

func (mssuo *MDSubSystemsUpdateOne) sqlSave(ctx context.Context) (_node *MDSubSystems, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mdsubsystems.Table,
			Columns: mdsubsystems.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: mdsubsystems.FieldID,
			},
		},
	}
	id, ok := mssuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MDSubSystems.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mssuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mdsubsystems.FieldID)
		for _, f := range fields {
			if !mdsubsystems.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mdsubsystems.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mssuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mssuo.mutation.Namerus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mdsubsystems.FieldNamerus,
		})
	}
	if value, ok := mssuo.mutation.Nameeng(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mdsubsystems.FieldNameeng,
		})
	}
	if value, ok := mssuo.mutation.Synonym(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mdsubsystems.FieldSynonym,
		})
	}
	if mssuo.mutation.ChildMdsubsystemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mdsubsystems.ChildMdsubsystemsTable,
			Columns: []string{mdsubsystems.ChildMdsubsystemsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdsubsystems.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mssuo.mutation.RemovedChildMdsubsystemsIDs(); len(nodes) > 0 && !mssuo.mutation.ChildMdsubsystemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mdsubsystems.ChildMdsubsystemsTable,
			Columns: []string{mdsubsystems.ChildMdsubsystemsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdsubsystems.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mssuo.mutation.ChildMdsubsystemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mdsubsystems.ChildMdsubsystemsTable,
			Columns: []string{mdsubsystems.ChildMdsubsystemsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdsubsystems.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mssuo.mutation.ParentMdsubsystemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mdsubsystems.ParentMdsubsystemsTable,
			Columns: []string{mdsubsystems.ParentMdsubsystemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdsubsystems.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mssuo.mutation.ParentMdsubsystemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mdsubsystems.ParentMdsubsystemsTable,
			Columns: []string{mdsubsystems.ParentMdsubsystemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdsubsystems.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mssuo.mutation.MdtablesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mdsubsystems.MdtablesTable,
			Columns: mdsubsystems.MdtablesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdtabel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mssuo.mutation.RemovedMdtablesIDs(); len(nodes) > 0 && !mssuo.mutation.MdtablesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mdsubsystems.MdtablesTable,
			Columns: mdsubsystems.MdtablesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdtabel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mssuo.mutation.MdtablesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   mdsubsystems.MdtablesTable,
			Columns: mdsubsystems.MdtablesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: mdtabel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MDSubSystems{config: mssuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mssuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mdsubsystems.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
