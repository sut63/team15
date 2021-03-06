// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team15/app/ent/bedtype"
)

// Bedtype is the model entity for the Bedtype schema.
type Bedtype struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Bedtypename holds the value of the "bedtypename" field.
	Bedtypename string `json:"bedtypename,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BedtypeQuery when eager-loading is set.
	Edges BedtypeEdges `json:"edges"`
}

// BedtypeEdges holds the relations/edges for other nodes in the graph.
type BedtypeEdges struct {
	// Roomdetails holds the value of the roomdetails edge.
	Roomdetails []*Roomdetail
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoomdetailsOrErr returns the Roomdetails value or an error if the edge
// was not loaded in eager-loading.
func (e BedtypeEdges) RoomdetailsOrErr() ([]*Roomdetail, error) {
	if e.loadedTypes[0] {
		return e.Roomdetails, nil
	}
	return nil, &NotLoadedError{edge: "roomdetails"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bedtype) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // bedtypename
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bedtype fields.
func (b *Bedtype) assignValues(values ...interface{}) error {
	if m, n := len(values), len(bedtype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	b.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field bedtypename", values[0])
	} else if value.Valid {
		b.Bedtypename = value.String
	}
	return nil
}

// QueryRoomdetails queries the roomdetails edge of the Bedtype.
func (b *Bedtype) QueryRoomdetails() *RoomdetailQuery {
	return (&BedtypeClient{config: b.config}).QueryRoomdetails(b)
}

// Update returns a builder for updating this Bedtype.
// Note that, you need to call Bedtype.Unwrap() before calling this method, if this Bedtype
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bedtype) Update() *BedtypeUpdateOne {
	return (&BedtypeClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Bedtype) Unwrap() *Bedtype {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bedtype is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bedtype) String() string {
	var builder strings.Builder
	builder.WriteString("Bedtype(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", bedtypename=")
	builder.WriteString(b.Bedtypename)
	builder.WriteByte(')')
	return builder.String()
}

// Bedtypes is a parsable slice of Bedtype.
type Bedtypes []*Bedtype

func (b Bedtypes) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
