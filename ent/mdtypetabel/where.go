// Code generated by entc, DO NOT EDIT.

package mdtypetabel

import (
	"fynegui/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
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
func IDGT(id string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Nameeng applies equality check predicate on the "nameeng" field. It's identical to NameengEQ.
func Nameeng(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameeng), v))
	})
}

// Synonym applies equality check predicate on the "synonym" field. It's identical to SynonymEQ.
func Synonym(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSynonym), v))
	})
}

// Por applies equality check predicate on the "por" field. It's identical to PorEQ.
func Por(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPor), v))
	})
}

// Parent applies equality check predicate on the "parent" field. It's identical to ParentEQ.
func Parent(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParent), v))
	})
}

// NameengEQ applies the EQ predicate on the "nameeng" field.
func NameengEQ(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameeng), v))
	})
}

// NameengNEQ applies the NEQ predicate on the "nameeng" field.
func NameengNEQ(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNameeng), v))
	})
}

// NameengIn applies the In predicate on the "nameeng" field.
func NameengIn(vs ...string) predicate.MDTypeTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTypeTabel(func(s *sql.Selector) {
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
func NameengNotIn(vs ...string) predicate.MDTypeTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTypeTabel(func(s *sql.Selector) {
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
func NameengGT(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNameeng), v))
	})
}

// NameengGTE applies the GTE predicate on the "nameeng" field.
func NameengGTE(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNameeng), v))
	})
}

// NameengLT applies the LT predicate on the "nameeng" field.
func NameengLT(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNameeng), v))
	})
}

// NameengLTE applies the LTE predicate on the "nameeng" field.
func NameengLTE(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNameeng), v))
	})
}

// NameengContains applies the Contains predicate on the "nameeng" field.
func NameengContains(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNameeng), v))
	})
}

// NameengHasPrefix applies the HasPrefix predicate on the "nameeng" field.
func NameengHasPrefix(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNameeng), v))
	})
}

// NameengHasSuffix applies the HasSuffix predicate on the "nameeng" field.
func NameengHasSuffix(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNameeng), v))
	})
}

// NameengEqualFold applies the EqualFold predicate on the "nameeng" field.
func NameengEqualFold(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNameeng), v))
	})
}

// NameengContainsFold applies the ContainsFold predicate on the "nameeng" field.
func NameengContainsFold(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNameeng), v))
	})
}

// SynonymEQ applies the EQ predicate on the "synonym" field.
func SynonymEQ(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSynonym), v))
	})
}

// SynonymNEQ applies the NEQ predicate on the "synonym" field.
func SynonymNEQ(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSynonym), v))
	})
}

// SynonymIn applies the In predicate on the "synonym" field.
func SynonymIn(vs ...string) predicate.MDTypeTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTypeTabel(func(s *sql.Selector) {
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
func SynonymNotIn(vs ...string) predicate.MDTypeTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTypeTabel(func(s *sql.Selector) {
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
func SynonymGT(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSynonym), v))
	})
}

// SynonymGTE applies the GTE predicate on the "synonym" field.
func SynonymGTE(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSynonym), v))
	})
}

// SynonymLT applies the LT predicate on the "synonym" field.
func SynonymLT(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSynonym), v))
	})
}

// SynonymLTE applies the LTE predicate on the "synonym" field.
func SynonymLTE(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSynonym), v))
	})
}

// SynonymContains applies the Contains predicate on the "synonym" field.
func SynonymContains(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSynonym), v))
	})
}

// SynonymHasPrefix applies the HasPrefix predicate on the "synonym" field.
func SynonymHasPrefix(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSynonym), v))
	})
}

// SynonymHasSuffix applies the HasSuffix predicate on the "synonym" field.
func SynonymHasSuffix(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSynonym), v))
	})
}

// SynonymEqualFold applies the EqualFold predicate on the "synonym" field.
func SynonymEqualFold(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSynonym), v))
	})
}

// SynonymContainsFold applies the ContainsFold predicate on the "synonym" field.
func SynonymContainsFold(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSynonym), v))
	})
}

// PorEQ applies the EQ predicate on the "por" field.
func PorEQ(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPor), v))
	})
}

// PorNEQ applies the NEQ predicate on the "por" field.
func PorNEQ(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPor), v))
	})
}

// PorIn applies the In predicate on the "por" field.
func PorIn(vs ...string) predicate.MDTypeTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPor), v...))
	})
}

// PorNotIn applies the NotIn predicate on the "por" field.
func PorNotIn(vs ...string) predicate.MDTypeTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPor), v...))
	})
}

// PorGT applies the GT predicate on the "por" field.
func PorGT(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPor), v))
	})
}

// PorGTE applies the GTE predicate on the "por" field.
func PorGTE(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPor), v))
	})
}

// PorLT applies the LT predicate on the "por" field.
func PorLT(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPor), v))
	})
}

// PorLTE applies the LTE predicate on the "por" field.
func PorLTE(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPor), v))
	})
}

// PorContains applies the Contains predicate on the "por" field.
func PorContains(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPor), v))
	})
}

// PorHasPrefix applies the HasPrefix predicate on the "por" field.
func PorHasPrefix(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPor), v))
	})
}

// PorHasSuffix applies the HasSuffix predicate on the "por" field.
func PorHasSuffix(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPor), v))
	})
}

// PorEqualFold applies the EqualFold predicate on the "por" field.
func PorEqualFold(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPor), v))
	})
}

// PorContainsFold applies the ContainsFold predicate on the "por" field.
func PorContainsFold(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPor), v))
	})
}

// ParentEQ applies the EQ predicate on the "parent" field.
func ParentEQ(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParent), v))
	})
}

// ParentNEQ applies the NEQ predicate on the "parent" field.
func ParentNEQ(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldParent), v))
	})
}

// ParentIn applies the In predicate on the "parent" field.
func ParentIn(vs ...string) predicate.MDTypeTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldParent), v...))
	})
}

// ParentNotIn applies the NotIn predicate on the "parent" field.
func ParentNotIn(vs ...string) predicate.MDTypeTabel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldParent), v...))
	})
}

// ParentGT applies the GT predicate on the "parent" field.
func ParentGT(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldParent), v))
	})
}

// ParentGTE applies the GTE predicate on the "parent" field.
func ParentGTE(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldParent), v))
	})
}

// ParentLT applies the LT predicate on the "parent" field.
func ParentLT(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldParent), v))
	})
}

// ParentLTE applies the LTE predicate on the "parent" field.
func ParentLTE(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldParent), v))
	})
}

// ParentContains applies the Contains predicate on the "parent" field.
func ParentContains(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldParent), v))
	})
}

// ParentHasPrefix applies the HasPrefix predicate on the "parent" field.
func ParentHasPrefix(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldParent), v))
	})
}

// ParentHasSuffix applies the HasSuffix predicate on the "parent" field.
func ParentHasSuffix(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldParent), v))
	})
}

// ParentIsNil applies the IsNil predicate on the "parent" field.
func ParentIsNil() predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldParent)))
	})
}

// ParentNotNil applies the NotNil predicate on the "parent" field.
func ParentNotNil() predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldParent)))
	})
}

// ParentEqualFold applies the EqualFold predicate on the "parent" field.
func ParentEqualFold(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldParent), v))
	})
}

// ParentContainsFold applies the ContainsFold predicate on the "parent" field.
func ParentContainsFold(v string) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldParent), v))
	})
}

// HasChildMdtypetabels applies the HasEdge predicate on the "child_mdtypetabels" edge.
func HasChildMdtypetabels() predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChildMdtypetabelsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildMdtypetabelsTable, ChildMdtypetabelsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildMdtypetabelsWith applies the HasEdge predicate on the "child_mdtypetabels" edge with a given conditions (other predicates).
func HasChildMdtypetabelsWith(preds ...predicate.MDTypeTabel) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildMdtypetabelsTable, ChildMdtypetabelsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParentMdtypetabels applies the HasEdge predicate on the "parent_mdtypetabels" edge.
func HasParentMdtypetabels() predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentMdtypetabelsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentMdtypetabelsTable, ParentMdtypetabelsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentMdtypetabelsWith applies the HasEdge predicate on the "parent_mdtypetabels" edge with a given conditions (other predicates).
func HasParentMdtypetabelsWith(preds ...predicate.MDTypeTabel) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentMdtypetabelsTable, ParentMdtypetabelsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMdtypetabels applies the HasEdge predicate on the "mdtypetabels" edge.
func HasMdtypetabels() predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MdtypetabelsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MdtypetabelsTable, MdtypetabelsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMdtypetabelsWith applies the HasEdge predicate on the "mdtypetabels" edge with a given conditions (other predicates).
func HasMdtypetabelsWith(preds ...predicate.MDTabel) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MdtypetabelsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MdtypetabelsTable, MdtypetabelsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MDTypeTabel) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MDTypeTabel) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
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
func Not(p predicate.MDTypeTabel) predicate.MDTypeTabel {
	return predicate.MDTypeTabel(func(s *sql.Selector) {
		p(s.Not())
	})
}
