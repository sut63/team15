// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/equipment"
	"github.com/team15/app/ent/facilitie"
	"github.com/team15/app/ent/lease"
	"github.com/team15/app/ent/nearbyplace"
	"github.com/team15/app/ent/quantity"
	"github.com/team15/app/ent/roomdetail"
	"github.com/team15/app/ent/staytype"
)

// Roomdetail is the model entity for the Roomdetail schema.
type Roomdetail struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Roomtypename holds the value of the "roomtypename" field.
	Roomtypename string `json:"roomtypename,omitempty"`
	// Roomprice holds the value of the "roomprice" field.
	Roomprice string `json:"roomprice,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomdetailQuery when eager-loading is set.
	Edges                  RoomdetailEdges `json:"edges"`
	employee_id            *int
	equipment_roomdetail   *int
	facilitie_roomdetail   *int
	nearbyplace_roomdetail *int
	quantity_roomdetails   *int
	staytype_roomdetails   *int
}

// RoomdetailEdges holds the relations/edges for other nodes in the graph.
type RoomdetailEdges struct {
	// Equipments holds the value of the equipments edge.
	Equipments *Equipment
	// Facilities holds the value of the facilities edge.
	Facilities *Facilitie
	// Nearbyplaces holds the value of the nearbyplaces edge.
	Nearbyplaces *Nearbyplace
	// Employee holds the value of the employee edge.
	Employee *Employee
	// Quantity holds the value of the quantity edge.
	Quantity *Quantity
	// Staytype holds the value of the staytype edge.
	Staytype *Staytype
	// Roomdetails holds the value of the roomdetails edge.
	Roomdetails *Lease
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// EquipmentsOrErr returns the Equipments value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomdetailEdges) EquipmentsOrErr() (*Equipment, error) {
	if e.loadedTypes[0] {
		if e.Equipments == nil {
			// The edge equipments was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Equipments, nil
	}
	return nil, &NotLoadedError{edge: "equipments"}
}

// FacilitiesOrErr returns the Facilities value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomdetailEdges) FacilitiesOrErr() (*Facilitie, error) {
	if e.loadedTypes[1] {
		if e.Facilities == nil {
			// The edge facilities was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: facilitie.Label}
		}
		return e.Facilities, nil
	}
	return nil, &NotLoadedError{edge: "facilities"}
}

// NearbyplacesOrErr returns the Nearbyplaces value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomdetailEdges) NearbyplacesOrErr() (*Nearbyplace, error) {
	if e.loadedTypes[2] {
		if e.Nearbyplaces == nil {
			// The edge nearbyplaces was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: nearbyplace.Label}
		}
		return e.Nearbyplaces, nil
	}
	return nil, &NotLoadedError{edge: "nearbyplaces"}
}

// EmployeeOrErr returns the Employee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomdetailEdges) EmployeeOrErr() (*Employee, error) {
	if e.loadedTypes[3] {
		if e.Employee == nil {
			// The edge employee was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: employee.Label}
		}
		return e.Employee, nil
	}
	return nil, &NotLoadedError{edge: "employee"}
}

// QuantityOrErr returns the Quantity value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomdetailEdges) QuantityOrErr() (*Quantity, error) {
	if e.loadedTypes[4] {
		if e.Quantity == nil {
			// The edge quantity was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: quantity.Label}
		}
		return e.Quantity, nil
	}
	return nil, &NotLoadedError{edge: "quantity"}
}

// StaytypeOrErr returns the Staytype value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomdetailEdges) StaytypeOrErr() (*Staytype, error) {
	if e.loadedTypes[5] {
		if e.Staytype == nil {
			// The edge staytype was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: staytype.Label}
		}
		return e.Staytype, nil
	}
	return nil, &NotLoadedError{edge: "staytype"}
}

// RoomdetailsOrErr returns the Roomdetails value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomdetailEdges) RoomdetailsOrErr() (*Lease, error) {
	if e.loadedTypes[6] {
		if e.Roomdetails == nil {
			// The edge roomdetails was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: lease.Label}
		}
		return e.Roomdetails, nil
	}
	return nil, &NotLoadedError{edge: "roomdetails"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Roomdetail) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // roomtypename
		&sql.NullString{}, // roomprice
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Roomdetail) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // employee_id
		&sql.NullInt64{}, // equipment_roomdetail
		&sql.NullInt64{}, // facilitie_roomdetail
		&sql.NullInt64{}, // nearbyplace_roomdetail
		&sql.NullInt64{}, // quantity_roomdetails
		&sql.NullInt64{}, // staytype_roomdetails
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Roomdetail fields.
func (r *Roomdetail) assignValues(values ...interface{}) error {
	if m, n := len(values), len(roomdetail.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field roomtypename", values[0])
	} else if value.Valid {
		r.Roomtypename = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field roomprice", values[1])
	} else if value.Valid {
		r.Roomprice = value.String
	}
	values = values[2:]
	if len(values) == len(roomdetail.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field employee_id", value)
		} else if value.Valid {
			r.employee_id = new(int)
			*r.employee_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field equipment_roomdetail", value)
		} else if value.Valid {
			r.equipment_roomdetail = new(int)
			*r.equipment_roomdetail = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field facilitie_roomdetail", value)
		} else if value.Valid {
			r.facilitie_roomdetail = new(int)
			*r.facilitie_roomdetail = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field nearbyplace_roomdetail", value)
		} else if value.Valid {
			r.nearbyplace_roomdetail = new(int)
			*r.nearbyplace_roomdetail = int(value.Int64)
		}
		if value, ok := values[4].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field quantity_roomdetails", value)
		} else if value.Valid {
			r.quantity_roomdetails = new(int)
			*r.quantity_roomdetails = int(value.Int64)
		}
		if value, ok := values[5].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field staytype_roomdetails", value)
		} else if value.Valid {
			r.staytype_roomdetails = new(int)
			*r.staytype_roomdetails = int(value.Int64)
		}
	}
	return nil
}

// QueryEquipments queries the equipments edge of the Roomdetail.
func (r *Roomdetail) QueryEquipments() *EquipmentQuery {
	return (&RoomdetailClient{config: r.config}).QueryEquipments(r)
}

// QueryFacilities queries the facilities edge of the Roomdetail.
func (r *Roomdetail) QueryFacilities() *FacilitieQuery {
	return (&RoomdetailClient{config: r.config}).QueryFacilities(r)
}

// QueryNearbyplaces queries the nearbyplaces edge of the Roomdetail.
func (r *Roomdetail) QueryNearbyplaces() *NearbyplaceQuery {
	return (&RoomdetailClient{config: r.config}).QueryNearbyplaces(r)
}

// QueryEmployee queries the employee edge of the Roomdetail.
func (r *Roomdetail) QueryEmployee() *EmployeeQuery {
	return (&RoomdetailClient{config: r.config}).QueryEmployee(r)
}

// QueryQuantity queries the quantity edge of the Roomdetail.
func (r *Roomdetail) QueryQuantity() *QuantityQuery {
	return (&RoomdetailClient{config: r.config}).QueryQuantity(r)
}

// QueryStaytype queries the staytype edge of the Roomdetail.
func (r *Roomdetail) QueryStaytype() *StaytypeQuery {
	return (&RoomdetailClient{config: r.config}).QueryStaytype(r)
}

// QueryRoomdetails queries the roomdetails edge of the Roomdetail.
func (r *Roomdetail) QueryRoomdetails() *LeaseQuery {
	return (&RoomdetailClient{config: r.config}).QueryRoomdetails(r)
}

// Update returns a builder for updating this Roomdetail.
// Note that, you need to call Roomdetail.Unwrap() before calling this method, if this Roomdetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Roomdetail) Update() *RoomdetailUpdateOne {
	return (&RoomdetailClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Roomdetail) Unwrap() *Roomdetail {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Roomdetail is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Roomdetail) String() string {
	var builder strings.Builder
	builder.WriteString("Roomdetail(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", roomtypename=")
	builder.WriteString(r.Roomtypename)
	builder.WriteString(", roomprice=")
	builder.WriteString(r.Roomprice)
	builder.WriteByte(')')
	return builder.String()
}

// Roomdetails is a parsable slice of Roomdetail.
type Roomdetails []*Roomdetail

func (r Roomdetails) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
