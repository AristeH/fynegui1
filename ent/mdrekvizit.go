// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"fynegui/ent/mdrekvizit"
	"fynegui/ent/mdtabel"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// MDRekvizit is the model entity for the MDRekvizit schema.
type MDRekvizit struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"ссылка,omitempty"`
	// Namerus holds the value of the "namerus" field.
	Namerus string `json:"ИмяРус,omitempty"`
	// Nameeng holds the value of the "nameeng" field.
	Nameeng string `json:"ИмяАнгл,omitempty"`
	// Synonym holds the value of the "synonym" field.
	Synonym string `json:"Синоним,omitempty"`
	// Por holds the value of the "por" field.
	Por string `json:"ПорядокВывода,omitempty"`
	// WidthElem holds the value of the "widthElem" field.
	WidthElem float64 `json:"ШиринаЭлемента,omitempty"`
	// WidthSpisok holds the value of the "widthSpisok" field.
	WidthSpisok float64 `json:"ШиринаКолонки,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"Тип,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID string `json:"owner_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MDRekvizitQuery when eager-loading is set.
	Edges MDRekvizitEdges `json:"edges"`
}

// MDRekvizitEdges holds the relations/edges for other nodes in the graph.
type MDRekvizitEdges struct {
	// Owner holds the value of the owner edge.
	Owner *MDTabel `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MDRekvizitEdges) OwnerOrErr() (*MDTabel, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: mdtabel.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MDRekvizit) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mdrekvizit.FieldWidthElem, mdrekvizit.FieldWidthSpisok:
			values[i] = new(sql.NullFloat64)
		case mdrekvizit.FieldID, mdrekvizit.FieldNamerus, mdrekvizit.FieldNameeng, mdrekvizit.FieldSynonym, mdrekvizit.FieldPor, mdrekvizit.FieldType, mdrekvizit.FieldOwnerID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MDRekvizit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MDRekvizit fields.
func (mr *MDRekvizit) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mdrekvizit.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				mr.ID = value.String
			}
		case mdrekvizit.FieldNamerus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namerus", values[i])
			} else if value.Valid {
				mr.Namerus = value.String
			}
		case mdrekvizit.FieldNameeng:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nameeng", values[i])
			} else if value.Valid {
				mr.Nameeng = value.String
			}
		case mdrekvizit.FieldSynonym:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field synonym", values[i])
			} else if value.Valid {
				mr.Synonym = value.String
			}
		case mdrekvizit.FieldPor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field por", values[i])
			} else if value.Valid {
				mr.Por = value.String
			}
		case mdrekvizit.FieldWidthElem:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field widthElem", values[i])
			} else if value.Valid {
				mr.WidthElem = value.Float64
			}
		case mdrekvizit.FieldWidthSpisok:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field widthSpisok", values[i])
			} else if value.Valid {
				mr.WidthSpisok = value.Float64
			}
		case mdrekvizit.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				mr.Type = value.String
			}
		case mdrekvizit.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				mr.OwnerID = value.String
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the MDRekvizit entity.
func (mr *MDRekvizit) QueryOwner() *MDTabelQuery {
	return (&MDRekvizitClient{config: mr.config}).QueryOwner(mr)
}

// Update returns a builder for updating this MDRekvizit.
// Note that you need to call MDRekvizit.Unwrap() before calling this method if this MDRekvizit
// was returned from a transaction, and the transaction was committed or rolled back.
func (mr *MDRekvizit) Update() *MDRekvizitUpdateOne {
	return (&MDRekvizitClient{config: mr.config}).UpdateOne(mr)
}

// Unwrap unwraps the MDRekvizit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mr *MDRekvizit) Unwrap() *MDRekvizit {
	tx, ok := mr.config.driver.(*txDriver)
	if !ok {
		panic("ent: MDRekvizit is not a transactional entity")
	}
	mr.config.driver = tx.drv
	return mr
}

// String implements the fmt.Stringer.
func (mr *MDRekvizit) String() string {
	var builder strings.Builder
	builder.WriteString("MDRekvizit(")
	builder.WriteString(fmt.Sprintf("id=%v", mr.ID))
	builder.WriteString(", namerus=")
	builder.WriteString(mr.Namerus)
	builder.WriteString(", nameeng=")
	builder.WriteString(mr.Nameeng)
	builder.WriteString(", synonym=")
	builder.WriteString(mr.Synonym)
	builder.WriteString(", por=")
	builder.WriteString(mr.Por)
	builder.WriteString(", widthElem=")
	builder.WriteString(fmt.Sprintf("%v", mr.WidthElem))
	builder.WriteString(", widthSpisok=")
	builder.WriteString(fmt.Sprintf("%v", mr.WidthSpisok))
	builder.WriteString(", type=")
	builder.WriteString(mr.Type)
	builder.WriteString(", owner_id=")
	builder.WriteString(mr.OwnerID)
	builder.WriteByte(')')
	return builder.String()
}

// MDRekvizits is a parsable slice of MDRekvizit.
type MDRekvizits []*MDRekvizit

func (mr MDRekvizits) config(cfg config) {
	for _i := range mr {
		mr[_i].config = cfg
	}
}