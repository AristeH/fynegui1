// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"fynegui/ent/mdrekvizit"
	"fynegui/ent/mdtabel"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MDRekvizitCreate is the builder for creating a MDRekvizit entity.
type MDRekvizitCreate struct {
	config
	mutation *MDRekvizitMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNamerus sets the "namerus" field.
func (mrc *MDRekvizitCreate) SetNamerus(s string) *MDRekvizitCreate {
	mrc.mutation.SetNamerus(s)
	return mrc
}

// SetNameeng sets the "nameeng" field.
func (mrc *MDRekvizitCreate) SetNameeng(s string) *MDRekvizitCreate {
	mrc.mutation.SetNameeng(s)
	return mrc
}

// SetSynonym sets the "synonym" field.
func (mrc *MDRekvizitCreate) SetSynonym(s string) *MDRekvizitCreate {
	mrc.mutation.SetSynonym(s)
	return mrc
}

// SetPor sets the "por" field.
func (mrc *MDRekvizitCreate) SetPor(s string) *MDRekvizitCreate {
	mrc.mutation.SetPor(s)
	return mrc
}

// SetWidthElem sets the "widthElem" field.
func (mrc *MDRekvizitCreate) SetWidthElem(f float64) *MDRekvizitCreate {
	mrc.mutation.SetWidthElem(f)
	return mrc
}

// SetWidthSpisok sets the "widthSpisok" field.
func (mrc *MDRekvizitCreate) SetWidthSpisok(f float64) *MDRekvizitCreate {
	mrc.mutation.SetWidthSpisok(f)
	return mrc
}

// SetType sets the "type" field.
func (mrc *MDRekvizitCreate) SetType(s string) *MDRekvizitCreate {
	mrc.mutation.SetType(s)
	return mrc
}

// SetOwnerID sets the "owner_id" field.
func (mrc *MDRekvizitCreate) SetOwnerID(s string) *MDRekvizitCreate {
	mrc.mutation.SetOwnerID(s)
	return mrc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (mrc *MDRekvizitCreate) SetNillableOwnerID(s *string) *MDRekvizitCreate {
	if s != nil {
		mrc.SetOwnerID(*s)
	}
	return mrc
}

// SetID sets the "id" field.
func (mrc *MDRekvizitCreate) SetID(s string) *MDRekvizitCreate {
	mrc.mutation.SetID(s)
	return mrc
}

// SetOwner sets the "owner" edge to the MDTabel entity.
func (mrc *MDRekvizitCreate) SetOwner(m *MDTabel) *MDRekvizitCreate {
	return mrc.SetOwnerID(m.ID)
}

// Mutation returns the MDRekvizitMutation object of the builder.
func (mrc *MDRekvizitCreate) Mutation() *MDRekvizitMutation {
	return mrc.mutation
}

// Save creates the MDRekvizit in the database.
func (mrc *MDRekvizitCreate) Save(ctx context.Context) (*MDRekvizit, error) {
	var (
		err  error
		node *MDRekvizit
	)
	if len(mrc.hooks) == 0 {
		if err = mrc.check(); err != nil {
			return nil, err
		}
		node, err = mrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MDRekvizitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mrc.check(); err != nil {
				return nil, err
			}
			mrc.mutation = mutation
			if node, err = mrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mrc.hooks) - 1; i >= 0; i-- {
			if mrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mrc *MDRekvizitCreate) SaveX(ctx context.Context) *MDRekvizit {
	v, err := mrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrc *MDRekvizitCreate) Exec(ctx context.Context) error {
	_, err := mrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrc *MDRekvizitCreate) ExecX(ctx context.Context) {
	if err := mrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mrc *MDRekvizitCreate) check() error {
	if _, ok := mrc.mutation.Namerus(); !ok {
		return &ValidationError{Name: "namerus", err: errors.New(`ent: missing required field "MDRekvizit.namerus"`)}
	}
	if v, ok := mrc.mutation.Namerus(); ok {
		if err := mdrekvizit.NamerusValidator(v); err != nil {
			return &ValidationError{Name: "namerus", err: fmt.Errorf(`ent: validator failed for field "MDRekvizit.namerus": %w`, err)}
		}
	}
	if _, ok := mrc.mutation.Nameeng(); !ok {
		return &ValidationError{Name: "nameeng", err: errors.New(`ent: missing required field "MDRekvizit.nameeng"`)}
	}
	if v, ok := mrc.mutation.Nameeng(); ok {
		if err := mdrekvizit.NameengValidator(v); err != nil {
			return &ValidationError{Name: "nameeng", err: fmt.Errorf(`ent: validator failed for field "MDRekvizit.nameeng": %w`, err)}
		}
	}
	if _, ok := mrc.mutation.Synonym(); !ok {
		return &ValidationError{Name: "synonym", err: errors.New(`ent: missing required field "MDRekvizit.synonym"`)}
	}
	if v, ok := mrc.mutation.Synonym(); ok {
		if err := mdrekvizit.SynonymValidator(v); err != nil {
			return &ValidationError{Name: "synonym", err: fmt.Errorf(`ent: validator failed for field "MDRekvizit.synonym": %w`, err)}
		}
	}
	if _, ok := mrc.mutation.Por(); !ok {
		return &ValidationError{Name: "por", err: errors.New(`ent: missing required field "MDRekvizit.por"`)}
	}
	if v, ok := mrc.mutation.Por(); ok {
		if err := mdrekvizit.PorValidator(v); err != nil {
			return &ValidationError{Name: "por", err: fmt.Errorf(`ent: validator failed for field "MDRekvizit.por": %w`, err)}
		}
	}
	if _, ok := mrc.mutation.WidthElem(); !ok {
		return &ValidationError{Name: "widthElem", err: errors.New(`ent: missing required field "MDRekvizit.widthElem"`)}
	}
	if _, ok := mrc.mutation.WidthSpisok(); !ok {
		return &ValidationError{Name: "widthSpisok", err: errors.New(`ent: missing required field "MDRekvizit.widthSpisok"`)}
	}
	if _, ok := mrc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "MDRekvizit.type"`)}
	}
	if v, ok := mrc.mutation.GetType(); ok {
		if err := mdrekvizit.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "MDRekvizit.type": %w`, err)}
		}
	}
	if v, ok := mrc.mutation.OwnerID(); ok {
		if err := mdrekvizit.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`ent: validator failed for field "MDRekvizit.owner_id": %w`, err)}
		}
	}
	if v, ok := mrc.mutation.ID(); ok {
		if err := mdrekvizit.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "MDRekvizit.id": %w`, err)}
		}
	}
	return nil
}

func (mrc *MDRekvizitCreate) sqlSave(ctx context.Context) (*MDRekvizit, error) {
	_node, _spec := mrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected MDRekvizit.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (mrc *MDRekvizitCreate) createSpec() (*MDRekvizit, *sqlgraph.CreateSpec) {
	var (
		_node = &MDRekvizit{config: mrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mdrekvizit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: mdrekvizit.FieldID,
			},
		}
	)
	_spec.OnConflict = mrc.conflict
	if id, ok := mrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mrc.mutation.Namerus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mdrekvizit.FieldNamerus,
		})
		_node.Namerus = value
	}
	if value, ok := mrc.mutation.Nameeng(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mdrekvizit.FieldNameeng,
		})
		_node.Nameeng = value
	}
	if value, ok := mrc.mutation.Synonym(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mdrekvizit.FieldSynonym,
		})
		_node.Synonym = value
	}
	if value, ok := mrc.mutation.Por(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mdrekvizit.FieldPor,
		})
		_node.Por = value
	}
	if value, ok := mrc.mutation.WidthElem(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mdrekvizit.FieldWidthElem,
		})
		_node.WidthElem = value
	}
	if value, ok := mrc.mutation.WidthSpisok(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: mdrekvizit.FieldWidthSpisok,
		})
		_node.WidthSpisok = value
	}
	if value, ok := mrc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mdrekvizit.FieldType,
		})
		_node.Type = value
	}
	if nodes := mrc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mdrekvizit.OwnerTable,
			Columns: []string{mdrekvizit.OwnerColumn},
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
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MDRekvizit.Create().
//		SetNamerus(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MDRekvizitUpsert) {
//			SetNamerus(v+v).
//		}).
//		Exec(ctx)
//
func (mrc *MDRekvizitCreate) OnConflict(opts ...sql.ConflictOption) *MDRekvizitUpsertOne {
	mrc.conflict = opts
	return &MDRekvizitUpsertOne{
		create: mrc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MDRekvizit.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mrc *MDRekvizitCreate) OnConflictColumns(columns ...string) *MDRekvizitUpsertOne {
	mrc.conflict = append(mrc.conflict, sql.ConflictColumns(columns...))
	return &MDRekvizitUpsertOne{
		create: mrc,
	}
}

type (
	// MDRekvizitUpsertOne is the builder for "upsert"-ing
	//  one MDRekvizit node.
	MDRekvizitUpsertOne struct {
		create *MDRekvizitCreate
	}

	// MDRekvizitUpsert is the "OnConflict" setter.
	MDRekvizitUpsert struct {
		*sql.UpdateSet
	}
)

// SetNamerus sets the "namerus" field.
func (u *MDRekvizitUpsert) SetNamerus(v string) *MDRekvizitUpsert {
	u.Set(mdrekvizit.FieldNamerus, v)
	return u
}

// UpdateNamerus sets the "namerus" field to the value that was provided on create.
func (u *MDRekvizitUpsert) UpdateNamerus() *MDRekvizitUpsert {
	u.SetExcluded(mdrekvizit.FieldNamerus)
	return u
}

// SetNameeng sets the "nameeng" field.
func (u *MDRekvizitUpsert) SetNameeng(v string) *MDRekvizitUpsert {
	u.Set(mdrekvizit.FieldNameeng, v)
	return u
}

// UpdateNameeng sets the "nameeng" field to the value that was provided on create.
func (u *MDRekvizitUpsert) UpdateNameeng() *MDRekvizitUpsert {
	u.SetExcluded(mdrekvizit.FieldNameeng)
	return u
}

// SetSynonym sets the "synonym" field.
func (u *MDRekvizitUpsert) SetSynonym(v string) *MDRekvizitUpsert {
	u.Set(mdrekvizit.FieldSynonym, v)
	return u
}

// UpdateSynonym sets the "synonym" field to the value that was provided on create.
func (u *MDRekvizitUpsert) UpdateSynonym() *MDRekvizitUpsert {
	u.SetExcluded(mdrekvizit.FieldSynonym)
	return u
}

// SetPor sets the "por" field.
func (u *MDRekvizitUpsert) SetPor(v string) *MDRekvizitUpsert {
	u.Set(mdrekvizit.FieldPor, v)
	return u
}

// UpdatePor sets the "por" field to the value that was provided on create.
func (u *MDRekvizitUpsert) UpdatePor() *MDRekvizitUpsert {
	u.SetExcluded(mdrekvizit.FieldPor)
	return u
}

// SetWidthElem sets the "widthElem" field.
func (u *MDRekvizitUpsert) SetWidthElem(v float64) *MDRekvizitUpsert {
	u.Set(mdrekvizit.FieldWidthElem, v)
	return u
}

// UpdateWidthElem sets the "widthElem" field to the value that was provided on create.
func (u *MDRekvizitUpsert) UpdateWidthElem() *MDRekvizitUpsert {
	u.SetExcluded(mdrekvizit.FieldWidthElem)
	return u
}

// AddWidthElem adds v to the "widthElem" field.
func (u *MDRekvizitUpsert) AddWidthElem(v float64) *MDRekvizitUpsert {
	u.Add(mdrekvizit.FieldWidthElem, v)
	return u
}

// SetWidthSpisok sets the "widthSpisok" field.
func (u *MDRekvizitUpsert) SetWidthSpisok(v float64) *MDRekvizitUpsert {
	u.Set(mdrekvizit.FieldWidthSpisok, v)
	return u
}

// UpdateWidthSpisok sets the "widthSpisok" field to the value that was provided on create.
func (u *MDRekvizitUpsert) UpdateWidthSpisok() *MDRekvizitUpsert {
	u.SetExcluded(mdrekvizit.FieldWidthSpisok)
	return u
}

// AddWidthSpisok adds v to the "widthSpisok" field.
func (u *MDRekvizitUpsert) AddWidthSpisok(v float64) *MDRekvizitUpsert {
	u.Add(mdrekvizit.FieldWidthSpisok, v)
	return u
}

// SetType sets the "type" field.
func (u *MDRekvizitUpsert) SetType(v string) *MDRekvizitUpsert {
	u.Set(mdrekvizit.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *MDRekvizitUpsert) UpdateType() *MDRekvizitUpsert {
	u.SetExcluded(mdrekvizit.FieldType)
	return u
}

// SetOwnerID sets the "owner_id" field.
func (u *MDRekvizitUpsert) SetOwnerID(v string) *MDRekvizitUpsert {
	u.Set(mdrekvizit.FieldOwnerID, v)
	return u
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *MDRekvizitUpsert) UpdateOwnerID() *MDRekvizitUpsert {
	u.SetExcluded(mdrekvizit.FieldOwnerID)
	return u
}

// ClearOwnerID clears the value of the "owner_id" field.
func (u *MDRekvizitUpsert) ClearOwnerID() *MDRekvizitUpsert {
	u.SetNull(mdrekvizit.FieldOwnerID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MDRekvizit.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mdrekvizit.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *MDRekvizitUpsertOne) UpdateNewValues() *MDRekvizitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(mdrekvizit.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.MDRekvizit.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *MDRekvizitUpsertOne) Ignore() *MDRekvizitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MDRekvizitUpsertOne) DoNothing() *MDRekvizitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MDRekvizitCreate.OnConflict
// documentation for more info.
func (u *MDRekvizitUpsertOne) Update(set func(*MDRekvizitUpsert)) *MDRekvizitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MDRekvizitUpsert{UpdateSet: update})
	}))
	return u
}

// SetNamerus sets the "namerus" field.
func (u *MDRekvizitUpsertOne) SetNamerus(v string) *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetNamerus(v)
	})
}

// UpdateNamerus sets the "namerus" field to the value that was provided on create.
func (u *MDRekvizitUpsertOne) UpdateNamerus() *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateNamerus()
	})
}

// SetNameeng sets the "nameeng" field.
func (u *MDRekvizitUpsertOne) SetNameeng(v string) *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetNameeng(v)
	})
}

// UpdateNameeng sets the "nameeng" field to the value that was provided on create.
func (u *MDRekvizitUpsertOne) UpdateNameeng() *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateNameeng()
	})
}

// SetSynonym sets the "synonym" field.
func (u *MDRekvizitUpsertOne) SetSynonym(v string) *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetSynonym(v)
	})
}

// UpdateSynonym sets the "synonym" field to the value that was provided on create.
func (u *MDRekvizitUpsertOne) UpdateSynonym() *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateSynonym()
	})
}

// SetPor sets the "por" field.
func (u *MDRekvizitUpsertOne) SetPor(v string) *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetPor(v)
	})
}

// UpdatePor sets the "por" field to the value that was provided on create.
func (u *MDRekvizitUpsertOne) UpdatePor() *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdatePor()
	})
}

// SetWidthElem sets the "widthElem" field.
func (u *MDRekvizitUpsertOne) SetWidthElem(v float64) *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetWidthElem(v)
	})
}

// AddWidthElem adds v to the "widthElem" field.
func (u *MDRekvizitUpsertOne) AddWidthElem(v float64) *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.AddWidthElem(v)
	})
}

// UpdateWidthElem sets the "widthElem" field to the value that was provided on create.
func (u *MDRekvizitUpsertOne) UpdateWidthElem() *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateWidthElem()
	})
}

// SetWidthSpisok sets the "widthSpisok" field.
func (u *MDRekvizitUpsertOne) SetWidthSpisok(v float64) *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetWidthSpisok(v)
	})
}

// AddWidthSpisok adds v to the "widthSpisok" field.
func (u *MDRekvizitUpsertOne) AddWidthSpisok(v float64) *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.AddWidthSpisok(v)
	})
}

// UpdateWidthSpisok sets the "widthSpisok" field to the value that was provided on create.
func (u *MDRekvizitUpsertOne) UpdateWidthSpisok() *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateWidthSpisok()
	})
}

// SetType sets the "type" field.
func (u *MDRekvizitUpsertOne) SetType(v string) *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *MDRekvizitUpsertOne) UpdateType() *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateType()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *MDRekvizitUpsertOne) SetOwnerID(v string) *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *MDRekvizitUpsertOne) UpdateOwnerID() *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateOwnerID()
	})
}

// ClearOwnerID clears the value of the "owner_id" field.
func (u *MDRekvizitUpsertOne) ClearOwnerID() *MDRekvizitUpsertOne {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.ClearOwnerID()
	})
}

// Exec executes the query.
func (u *MDRekvizitUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MDRekvizitCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MDRekvizitUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MDRekvizitUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MDRekvizitUpsertOne.ID is not supported by MySQL driver. Use MDRekvizitUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MDRekvizitUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MDRekvizitCreateBulk is the builder for creating many MDRekvizit entities in bulk.
type MDRekvizitCreateBulk struct {
	config
	builders []*MDRekvizitCreate
	conflict []sql.ConflictOption
}

// Save creates the MDRekvizit entities in the database.
func (mrcb *MDRekvizitCreateBulk) Save(ctx context.Context) ([]*MDRekvizit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mrcb.builders))
	nodes := make([]*MDRekvizit, len(mrcb.builders))
	mutators := make([]Mutator, len(mrcb.builders))
	for i := range mrcb.builders {
		func(i int, root context.Context) {
			builder := mrcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MDRekvizitMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mrcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mrcb *MDRekvizitCreateBulk) SaveX(ctx context.Context) []*MDRekvizit {
	v, err := mrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrcb *MDRekvizitCreateBulk) Exec(ctx context.Context) error {
	_, err := mrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrcb *MDRekvizitCreateBulk) ExecX(ctx context.Context) {
	if err := mrcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MDRekvizit.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MDRekvizitUpsert) {
//			SetNamerus(v+v).
//		}).
//		Exec(ctx)
//
func (mrcb *MDRekvizitCreateBulk) OnConflict(opts ...sql.ConflictOption) *MDRekvizitUpsertBulk {
	mrcb.conflict = opts
	return &MDRekvizitUpsertBulk{
		create: mrcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MDRekvizit.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mrcb *MDRekvizitCreateBulk) OnConflictColumns(columns ...string) *MDRekvizitUpsertBulk {
	mrcb.conflict = append(mrcb.conflict, sql.ConflictColumns(columns...))
	return &MDRekvizitUpsertBulk{
		create: mrcb,
	}
}

// MDRekvizitUpsertBulk is the builder for "upsert"-ing
// a bulk of MDRekvizit nodes.
type MDRekvizitUpsertBulk struct {
	create *MDRekvizitCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MDRekvizit.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mdrekvizit.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *MDRekvizitUpsertBulk) UpdateNewValues() *MDRekvizitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(mdrekvizit.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MDRekvizit.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *MDRekvizitUpsertBulk) Ignore() *MDRekvizitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MDRekvizitUpsertBulk) DoNothing() *MDRekvizitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MDRekvizitCreateBulk.OnConflict
// documentation for more info.
func (u *MDRekvizitUpsertBulk) Update(set func(*MDRekvizitUpsert)) *MDRekvizitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MDRekvizitUpsert{UpdateSet: update})
	}))
	return u
}

// SetNamerus sets the "namerus" field.
func (u *MDRekvizitUpsertBulk) SetNamerus(v string) *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetNamerus(v)
	})
}

// UpdateNamerus sets the "namerus" field to the value that was provided on create.
func (u *MDRekvizitUpsertBulk) UpdateNamerus() *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateNamerus()
	})
}

// SetNameeng sets the "nameeng" field.
func (u *MDRekvizitUpsertBulk) SetNameeng(v string) *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetNameeng(v)
	})
}

// UpdateNameeng sets the "nameeng" field to the value that was provided on create.
func (u *MDRekvizitUpsertBulk) UpdateNameeng() *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateNameeng()
	})
}

// SetSynonym sets the "synonym" field.
func (u *MDRekvizitUpsertBulk) SetSynonym(v string) *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetSynonym(v)
	})
}

// UpdateSynonym sets the "synonym" field to the value that was provided on create.
func (u *MDRekvizitUpsertBulk) UpdateSynonym() *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateSynonym()
	})
}

// SetPor sets the "por" field.
func (u *MDRekvizitUpsertBulk) SetPor(v string) *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetPor(v)
	})
}

// UpdatePor sets the "por" field to the value that was provided on create.
func (u *MDRekvizitUpsertBulk) UpdatePor() *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdatePor()
	})
}

// SetWidthElem sets the "widthElem" field.
func (u *MDRekvizitUpsertBulk) SetWidthElem(v float64) *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetWidthElem(v)
	})
}

// AddWidthElem adds v to the "widthElem" field.
func (u *MDRekvizitUpsertBulk) AddWidthElem(v float64) *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.AddWidthElem(v)
	})
}

// UpdateWidthElem sets the "widthElem" field to the value that was provided on create.
func (u *MDRekvizitUpsertBulk) UpdateWidthElem() *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateWidthElem()
	})
}

// SetWidthSpisok sets the "widthSpisok" field.
func (u *MDRekvizitUpsertBulk) SetWidthSpisok(v float64) *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetWidthSpisok(v)
	})
}

// AddWidthSpisok adds v to the "widthSpisok" field.
func (u *MDRekvizitUpsertBulk) AddWidthSpisok(v float64) *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.AddWidthSpisok(v)
	})
}

// UpdateWidthSpisok sets the "widthSpisok" field to the value that was provided on create.
func (u *MDRekvizitUpsertBulk) UpdateWidthSpisok() *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateWidthSpisok()
	})
}

// SetType sets the "type" field.
func (u *MDRekvizitUpsertBulk) SetType(v string) *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *MDRekvizitUpsertBulk) UpdateType() *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateType()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *MDRekvizitUpsertBulk) SetOwnerID(v string) *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.SetOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *MDRekvizitUpsertBulk) UpdateOwnerID() *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.UpdateOwnerID()
	})
}

// ClearOwnerID clears the value of the "owner_id" field.
func (u *MDRekvizitUpsertBulk) ClearOwnerID() *MDRekvizitUpsertBulk {
	return u.Update(func(s *MDRekvizitUpsert) {
		s.ClearOwnerID()
	})
}

// Exec executes the query.
func (u *MDRekvizitUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MDRekvizitCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MDRekvizitCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MDRekvizitUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}