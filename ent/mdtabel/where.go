// Code generated by entc, DO NOT EDIT.

package mdtabel

import (
	"fynegui/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Namerus applies equality check predicate on the "namerus" field. It's identical to NamerusEQ.
func Namerus(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNamerus), v))
	})
}

// Nameeng applies equality check predicate on the "nameeng" field. It's identical to NameengEQ.
func Nameeng(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameeng), v))
	})
}

// Synonym applies equality check predicate on the "synonym" field. It's identical to SynonymEQ.
func Synonym(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSynonym), v))
	})
}

// File applies equality check predicate on the "file" field. It's identical to FileEQ.
func File(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFile), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// NamerusEQ applies the EQ predicate on the "namerus" field.
func NamerusEQ(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNamerus), v))
	})
}

// NamerusNEQ applies the NEQ predicate on the "namerus" field.
func NamerusNEQ(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNamerus), v))
	})
}

// NamerusIn applies the In predicate on the "namerus" field.
func NamerusIn(vs ...string) predicate.MDTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNamerus), v...))
	})
}

// NamerusNotIn applies the NotIn predicate on the "namerus" field.
func NamerusNotIn(vs ...string) predicate.MDTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNamerus), v...))
	})
}

// NamerusGT applies the GT predicate on the "namerus" field.
func NamerusGT(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNamerus), v))
	})
}

// NamerusGTE applies the GTE predicate on the "namerus" field.
func NamerusGTE(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNamerus), v))
	})
}

// NamerusLT applies the LT predicate on the "namerus" field.
func NamerusLT(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNamerus), v))
	})
}

// NamerusLTE applies the LTE predicate on the "namerus" field.
func NamerusLTE(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNamerus), v))
	})
}

// NamerusContains applies the Contains predicate on the "namerus" field.
func NamerusContains(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNamerus), v))
	})
}

// NamerusHasPrefix applies the HasPrefix predicate on the "namerus" field.
func NamerusHasPrefix(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNamerus), v))
	})
}

// NamerusHasSuffix applies the HasSuffix predicate on the "namerus" field.
func NamerusHasSuffix(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNamerus), v))
	})
}

// NamerusEqualFold applies the EqualFold predicate on the "namerus" field.
func NamerusEqualFold(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNamerus), v))
	})
}

// NamerusContainsFold applies the ContainsFold predicate on the "namerus" field.
func NamerusContainsFold(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNamerus), v))
	})
}

// NameengEQ applies the EQ predicate on the "nameeng" field.
func NameengEQ(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameeng), v))
	})
}

// NameengNEQ applies the NEQ predicate on the "nameeng" field.
func NameengNEQ(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNameeng), v))
	})
}

// NameengIn applies the In predicate on the "nameeng" field.
func NameengIn(vs ...string) predicate.MDTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNameeng), v...))
	})
}

// NameengNotIn applies the NotIn predicate on the "nameeng" field.
func NameengNotIn(vs ...string) predicate.MDTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNameeng), v...))
	})
}

// NameengGT applies the GT predicate on the "nameeng" field.
func NameengGT(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNameeng), v))
	})
}

// NameengGTE applies the GTE predicate on the "nameeng" field.
func NameengGTE(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNameeng), v))
	})
}

// NameengLT applies the LT predicate on the "nameeng" field.
func NameengLT(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNameeng), v))
	})
}

// NameengLTE applies the LTE predicate on the "nameeng" field.
func NameengLTE(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNameeng), v))
	})
}

// NameengContains applies the Contains predicate on the "nameeng" field.
func NameengContains(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNameeng), v))
	})
}

// NameengHasPrefix applies the HasPrefix predicate on the "nameeng" field.
func NameengHasPrefix(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNameeng), v))
	})
}

// NameengHasSuffix applies the HasSuffix predicate on the "nameeng" field.
func NameengHasSuffix(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNameeng), v))
	})
}

// NameengEqualFold applies the EqualFold predicate on the "nameeng" field.
func NameengEqualFold(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNameeng), v))
	})
}

// NameengContainsFold applies the ContainsFold predicate on the "nameeng" field.
func NameengContainsFold(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNameeng), v))
	})
}

// SynonymEQ applies the EQ predicate on the "synonym" field.
func SynonymEQ(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSynonym), v))
	})
}

// SynonymNEQ applies the NEQ predicate on the "synonym" field.
func SynonymNEQ(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSynonym), v))
	})
}

// SynonymIn applies the In predicate on the "synonym" field.
func SynonymIn(vs ...string) predicate.MDTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSynonym), v...))
	})
}

// SynonymNotIn applies the NotIn predicate on the "synonym" field.
func SynonymNotIn(vs ...string) predicate.MDTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSynonym), v...))
	})
}

// SynonymGT applies the GT predicate on the "synonym" field.
func SynonymGT(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSynonym), v))
	})
}

// SynonymGTE applies the GTE predicate on the "synonym" field.
func SynonymGTE(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSynonym), v))
	})
}

// SynonymLT applies the LT predicate on the "synonym" field.
func SynonymLT(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSynonym), v))
	})
}

// SynonymLTE applies the LTE predicate on the "synonym" field.
func SynonymLTE(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSynonym), v))
	})
}

// SynonymContains applies the Contains predicate on the "synonym" field.
func SynonymContains(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSynonym), v))
	})
}

// SynonymHasPrefix applies the HasPrefix predicate on the "synonym" field.
func SynonymHasPrefix(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSynonym), v))
	})
}

// SynonymHasSuffix applies the HasSuffix predicate on the "synonym" field.
func SynonymHasSuffix(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSynonym), v))
	})
}

// SynonymEqualFold applies the EqualFold predicate on the "synonym" field.
func SynonymEqualFold(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSynonym), v))
	})
}

// SynonymContainsFold applies the ContainsFold predicate on the "synonym" field.
func SynonymContainsFold(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSynonym), v))
	})
}

// FileEQ applies the EQ predicate on the "file" field.
func FileEQ(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFile), v))
	})
}

// FileNEQ applies the NEQ predicate on the "file" field.
func FileNEQ(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFile), v))
	})
}

// FileIn applies the In predicate on the "file" field.
func FileIn(vs ...string) predicate.MDTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFile), v...))
	})
}

// FileNotIn applies the NotIn predicate on the "file" field.
func FileNotIn(vs ...string) predicate.MDTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFile), v...))
	})
}

// FileGT applies the GT predicate on the "file" field.
func FileGT(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFile), v))
	})
}

// FileGTE applies the GTE predicate on the "file" field.
func FileGTE(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFile), v))
	})
}

// FileLT applies the LT predicate on the "file" field.
func FileLT(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFile), v))
	})
}

// FileLTE applies the LTE predicate on the "file" field.
func FileLTE(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFile), v))
	})
}

// FileContains applies the Contains predicate on the "file" field.
func FileContains(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFile), v))
	})
}

// FileHasPrefix applies the HasPrefix predicate on the "file" field.
func FileHasPrefix(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFile), v))
	})
}

// FileHasSuffix applies the HasSuffix predicate on the "file" field.
func FileHasSuffix(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFile), v))
	})
}

// FileEqualFold applies the EqualFold predicate on the "file" field.
func FileEqualFold(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFile), v))
	})
}

// FileContainsFold applies the ContainsFold predicate on the "file" field.
func FileContainsFold(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFile), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.MDTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.MDTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// HasMdsubsystems applies the HasEdge predicate on the "mdsubsystems" edge.
func HasMdsubsystems() predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MdsubsystemsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MdsubsystemsTable, MdsubsystemsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMdsubsystemsWith applies the HasEdge predicate on the "mdsubsystems" edge with a given conditions (other predicates).
func HasMdsubsystemsWith(preds ...predicate.MDSubSystems) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MdsubsystemsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MdsubsystemsTable, MdsubsystemsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMdrekvizits applies the HasEdge predicate on the "mdrekvizits" edge.
func HasMdrekvizits() predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MdrekvizitsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MdrekvizitsTable, MdrekvizitsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMdrekvizitsWith applies the HasEdge predicate on the "mdrekvizits" edge with a given conditions (other predicates).
func HasMdrekvizitsWith(preds ...predicate.MDRekvizit) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MdrekvizitsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MdrekvizitsTable, MdrekvizitsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MDTabel) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MDTabel) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MDTabel) predicate.MDTabel {
	return predicate.MDTabel(func(s *sql.Selector) {
		p(s.Not())
	})
}