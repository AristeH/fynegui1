// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"fynegui/ent/mdforms"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// MDForms is the model entity for the MDForms schema.
type MDForms struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"ссылка,omitempty"`
	// Idform holds the value of the "idform" field.
	Idform string `json:"Родитель,omitempty"`
	// Conteiner holds the value of the "conteiner" field.
	Conteiner string `json:"conteiner,omitempty"`
	// Parent holds the value of the "parent" field.
	Parent string `json:"Родитель,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MDFormsQuery when eager-loading is set.
	Edges MDFormsEdges `json:"edges"`
}

// MDFormsEdges holds the relations/edges for other nodes in the graph.
type MDFormsEdges struct {
	// ChildMdforms holds the value of the child_mdforms edge.
	ChildMdforms []*MDForms `json:"child_mdforms,omitempty"`
	// ParentMdforms holds the value of the parent_mdforms edge.
	ParentMdforms *MDForms `json:"родитель,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ChildMdformsOrErr returns the ChildMdforms value or an error if the edge
// was not loaded in eager-loading.
func (e MDFormsEdges) ChildMdformsOrErr() ([]*MDForms, error) {
	if e.loadedTypes[0] {
		return e.ChildMdforms, nil
	}
	return nil, &NotLoadedError{edge: "child_mdforms"}
}

// ParentMdformsOrErr returns the ParentMdforms value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MDFormsEdges) ParentMdformsOrErr() (*MDForms, error) {
	if e.loadedTypes[1] {
		if e.ParentMdforms == nil {
			// The edge parent_mdforms was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: mdforms.Label}
		}
		return e.ParentMdforms, nil
	}
	return nil, &NotLoadedError{edge: "parent_mdforms"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MDForms) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mdforms.FieldID, mdforms.FieldIdform, mdforms.FieldConteiner, mdforms.FieldParent:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MDForms", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MDForms fields.
func (mf *MDForms) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mdforms.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				mf.ID = value.String
			}
		case mdforms.FieldIdform:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field idform", values[i])
			} else if value.Valid {
				mf.Idform = value.String
			}
		case mdforms.FieldConteiner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field conteiner", values[i])
			} else if value.Valid {
				mf.Conteiner = value.String
			}
		case mdforms.FieldParent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent", values[i])
			} else if value.Valid {
				mf.Parent = value.String
			}
		}
	}
	return nil
}

// QueryChildMdforms queries the "child_mdforms" edge of the MDForms entity.
func (mf *MDForms) QueryChildMdforms() *MDFormsQuery {
	return (&MDFormsClient{config: mf.config}).QueryChildMdforms(mf)
}

// QueryParentMdforms queries the "parent_mdforms" edge of the MDForms entity.
func (mf *MDForms) QueryParentMdforms() *MDFormsQuery {
	return (&MDFormsClient{config: mf.config}).QueryParentMdforms(mf)
}

// Update returns a builder for updating this MDForms.
// Note that you need to call MDForms.Unwrap() before calling this method if this MDForms
// was returned from a transaction, and the transaction was committed or rolled back.
func (mf *MDForms) Update() *MDFormsUpdateOne {
	return (&MDFormsClient{config: mf.config}).UpdateOne(mf)
}

// Unwrap unwraps the MDForms entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mf *MDForms) Unwrap() *MDForms {
	tx, ok := mf.config.driver.(*txDriver)
	if !ok {
		panic("ent: MDForms is not a transactional entity")
	}
	mf.config.driver = tx.drv
	return mf
}

// String implements the fmt.Stringer.
func (mf *MDForms) String() string {
	var builder strings.Builder
	builder.WriteString("MDForms(")
	builder.WriteString(fmt.Sprintf("id=%v", mf.ID))
	builder.WriteString(", idform=")
	builder.WriteString(mf.Idform)
	builder.WriteString(", conteiner=")
	builder.WriteString(mf.Conteiner)
	builder.WriteString(", parent=")
	builder.WriteString(mf.Parent)
	builder.WriteByte(')')
	return builder.String()
}

// MDFormsSlice is a parsable slice of MDForms.
type MDFormsSlice []*MDForms

func (mf MDFormsSlice) config(cfg config) {
	for _i := range mf {
		mf[_i].config = cfg
	}
}
