// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team15/app/ent/facilitie"
	"github.com/team15/app/ent/roomdetail"
)

// Facilitie is the model entity for the Facilitie schema.
type Facilitie struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Facilitie holds the value of the "facilitie" field.
	Facilitie string `json:"facilitie,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FacilitieQuery when eager-loading is set.
	Edges                 FacilitieEdges `json:"edges"`
	roomdetail_facilities *int
}

// FacilitieEdges holds the relations/edges for other nodes in the graph.
type FacilitieEdges struct {
	// Roomdetail holds the value of the roomdetail edge.
	Roomdetail *Roomdetail
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoomdetailOrErr returns the Roomdetail value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FacilitieEdges) RoomdetailOrErr() (*Roomdetail, error) {
	if e.loadedTypes[0] {
		if e.Roomdetail == nil {
			// The edge roomdetail was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: roomdetail.Label}
		}
		return e.Roomdetail, nil
	}
	return nil, &NotLoadedError{edge: "roomdetail"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Facilitie) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // facilitie
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Facilitie) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // roomdetail_facilities
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Facilitie fields.
func (f *Facilitie) assignValues(values ...interface{}) error {
	if m, n := len(values), len(facilitie.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	f.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field facilitie", values[0])
	} else if value.Valid {
		f.Facilitie = value.String
	}
	values = values[1:]
	if len(values) == len(facilitie.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field roomdetail_facilities", value)
		} else if value.Valid {
			f.roomdetail_facilities = new(int)
			*f.roomdetail_facilities = int(value.Int64)
		}
	}
	return nil
}

// QueryRoomdetail queries the roomdetail edge of the Facilitie.
func (f *Facilitie) QueryRoomdetail() *RoomdetailQuery {
	return (&FacilitieClient{config: f.config}).QueryRoomdetail(f)
}

// Update returns a builder for updating this Facilitie.
// Note that, you need to call Facilitie.Unwrap() before calling this method, if this Facilitie
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Facilitie) Update() *FacilitieUpdateOne {
	return (&FacilitieClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (f *Facilitie) Unwrap() *Facilitie {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Facilitie is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Facilitie) String() string {
	var builder strings.Builder
	builder.WriteString("Facilitie(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", facilitie=")
	builder.WriteString(f.Facilitie)
	builder.WriteByte(')')
	return builder.String()
}

// Facilities is a parsable slice of Facilitie.
type Facilities []*Facilitie

func (f Facilities) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}