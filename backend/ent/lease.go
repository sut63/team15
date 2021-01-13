// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team15/app/ent/lease"
	"github.com/team15/app/ent/roomdetail"
	"github.com/team15/app/ent/wifi"
)

// Lease is the model entity for the Lease schema.
type Lease struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Addedtime holds the value of the "addedtime" field.
	Addedtime time.Time `json:"addedtime,omitempty"`
	// Tenant holds the value of the "tenant" field.
	Tenant string `json:"tenant,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LeaseQuery when eager-loading is set.
	Edges    LeaseEdges `json:"edges"`
	room_num *int
	wifi_id  *int
}

// LeaseEdges holds the relations/edges for other nodes in the graph.
type LeaseEdges struct {
	// Wifi holds the value of the Wifi edge.
	Wifi *Wifi
	// Roomdetail holds the value of the Roomdetail edge.
	Roomdetail *Roomdetail
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// WifiOrErr returns the Wifi value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LeaseEdges) WifiOrErr() (*Wifi, error) {
	if e.loadedTypes[0] {
		if e.Wifi == nil {
			// The edge Wifi was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: wifi.Label}
		}
		return e.Wifi, nil
	}
	return nil, &NotLoadedError{edge: "Wifi"}
}

// RoomdetailOrErr returns the Roomdetail value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LeaseEdges) RoomdetailOrErr() (*Roomdetail, error) {
	if e.loadedTypes[1] {
		if e.Roomdetail == nil {
			// The edge Roomdetail was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: roomdetail.Label}
		}
		return e.Roomdetail, nil
	}
	return nil, &NotLoadedError{edge: "Roomdetail"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Lease) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // addedtime
		&sql.NullString{}, // tenant
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Lease) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // room_num
		&sql.NullInt64{}, // wifi_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Lease fields.
func (l *Lease) assignValues(values ...interface{}) error {
	if m, n := len(values), len(lease.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	l.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field addedtime", values[0])
	} else if value.Valid {
		l.Addedtime = value.Time
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field tenant", values[1])
	} else if value.Valid {
		l.Tenant = value.String
	}
	values = values[2:]
	if len(values) == len(lease.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field room_num", value)
		} else if value.Valid {
			l.room_num = new(int)
			*l.room_num = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field wifi_id", value)
		} else if value.Valid {
			l.wifi_id = new(int)
			*l.wifi_id = int(value.Int64)
		}
	}
	return nil
}

// QueryWifi queries the Wifi edge of the Lease.
func (l *Lease) QueryWifi() *WifiQuery {
	return (&LeaseClient{config: l.config}).QueryWifi(l)
}

// QueryRoomdetail queries the Roomdetail edge of the Lease.
func (l *Lease) QueryRoomdetail() *RoomdetailQuery {
	return (&LeaseClient{config: l.config}).QueryRoomdetail(l)
}

// Update returns a builder for updating this Lease.
// Note that, you need to call Lease.Unwrap() before calling this method, if this Lease
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Lease) Update() *LeaseUpdateOne {
	return (&LeaseClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (l *Lease) Unwrap() *Lease {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Lease is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Lease) String() string {
	var builder strings.Builder
	builder.WriteString("Lease(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", addedtime=")
	builder.WriteString(l.Addedtime.Format(time.ANSIC))
	builder.WriteString(", tenant=")
	builder.WriteString(l.Tenant)
	builder.WriteByte(')')
	return builder.String()
}

// Leases is a parsable slice of Lease.
type Leases []*Lease

func (l Leases) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
