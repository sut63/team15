// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team15/app/ent/rentalstatus"
)

// Rentalstatus is the model entity for the Rentalstatus schema.
type Rentalstatus struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Rentalstatus holds the value of the "rentalstatus" field.
	Rentalstatus string `json:"rentalstatus,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RentalstatusQuery when eager-loading is set.
	Edges RentalstatusEdges `json:"edges"`
}

// RentalstatusEdges holds the relations/edges for other nodes in the graph.
type RentalstatusEdges struct {
	// Repairinvoices holds the value of the repairinvoices edge.
	Repairinvoices []*Repairinvoice
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RepairinvoicesOrErr returns the Repairinvoices value or an error if the edge
// was not loaded in eager-loading.
func (e RentalstatusEdges) RepairinvoicesOrErr() ([]*Repairinvoice, error) {
	if e.loadedTypes[0] {
		return e.Repairinvoices, nil
	}
	return nil, &NotLoadedError{edge: "repairinvoices"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Rentalstatus) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // rentalstatus
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Rentalstatus fields.
func (r *Rentalstatus) assignValues(values ...interface{}) error {
	if m, n := len(values), len(rentalstatus.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field rentalstatus", values[0])
	} else if value.Valid {
		r.Rentalstatus = value.String
	}
	return nil
}

// QueryRepairinvoices queries the repairinvoices edge of the Rentalstatus.
func (r *Rentalstatus) QueryRepairinvoices() *RepairinvoiceQuery {
	return (&RentalstatusClient{config: r.config}).QueryRepairinvoices(r)
}

// Update returns a builder for updating this Rentalstatus.
// Note that, you need to call Rentalstatus.Unwrap() before calling this method, if this Rentalstatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Rentalstatus) Update() *RentalstatusUpdateOne {
	return (&RentalstatusClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Rentalstatus) Unwrap() *Rentalstatus {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Rentalstatus is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Rentalstatus) String() string {
	var builder strings.Builder
	builder.WriteString("Rentalstatus(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", rentalstatus=")
	builder.WriteString(r.Rentalstatus)
	builder.WriteByte(')')
	return builder.String()
}

// Rentalstatuses is a parsable slice of Rentalstatus.
type Rentalstatuses []*Rentalstatus

func (r Rentalstatuses) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
