// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team15/app/ent/jobposition"
)

// Jobposition is the model entity for the Jobposition schema.
type Jobposition struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Positionname holds the value of the "positionname" field.
	Positionname string `json:"positionname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JobpositionQuery when eager-loading is set.
	Edges JobpositionEdges `json:"edges"`
}

// JobpositionEdges holds the relations/edges for other nodes in the graph.
type JobpositionEdges struct {
	// Employees holds the value of the employees edge.
	Employees []*Employee
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EmployeesOrErr returns the Employees value or an error if the edge
// was not loaded in eager-loading.
func (e JobpositionEdges) EmployeesOrErr() ([]*Employee, error) {
	if e.loadedTypes[0] {
		return e.Employees, nil
	}
	return nil, &NotLoadedError{edge: "employees"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Jobposition) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // positionname
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Jobposition fields.
func (j *Jobposition) assignValues(values ...interface{}) error {
	if m, n := len(values), len(jobposition.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	j.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field positionname", values[0])
	} else if value.Valid {
		j.Positionname = value.String
	}
	return nil
}

// QueryEmployees queries the employees edge of the Jobposition.
func (j *Jobposition) QueryEmployees() *EmployeeQuery {
	return (&JobpositionClient{config: j.config}).QueryEmployees(j)
}

// Update returns a builder for updating this Jobposition.
// Note that, you need to call Jobposition.Unwrap() before calling this method, if this Jobposition
// was returned from a transaction, and the transaction was committed or rolled back.
func (j *Jobposition) Update() *JobpositionUpdateOne {
	return (&JobpositionClient{config: j.config}).UpdateOne(j)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (j *Jobposition) Unwrap() *Jobposition {
	tx, ok := j.config.driver.(*txDriver)
	if !ok {
		panic("ent: Jobposition is not a transactional entity")
	}
	j.config.driver = tx.drv
	return j
}

// String implements the fmt.Stringer.
func (j *Jobposition) String() string {
	var builder strings.Builder
	builder.WriteString("Jobposition(")
	builder.WriteString(fmt.Sprintf("id=%v", j.ID))
	builder.WriteString(", positionname=")
	builder.WriteString(j.Positionname)
	builder.WriteByte(')')
	return builder.String()
}

// Jobpositions is a parsable slice of Jobposition.
type Jobpositions []*Jobposition

func (j Jobpositions) config(cfg config) {
	for _i := range j {
		j[_i].config = cfg
	}
}
