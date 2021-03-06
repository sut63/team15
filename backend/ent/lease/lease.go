// Code generated by entc, DO NOT EDIT.

package lease

const (
	// Label holds the string label denoting the lease type in the database.
	Label = "lease"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAddedtime holds the string denoting the addedtime field in the database.
	FieldAddedtime = "addedtime"
	// FieldTenant holds the string denoting the tenant field in the database.
	FieldTenant = "tenant"
	// FieldNumbtenant holds the string denoting the numbtenant field in the database.
	FieldNumbtenant = "numbtenant"
	// FieldIdtenant holds the string denoting the idtenant field in the database.
	FieldIdtenant = "idtenant"
	// FieldAgetenant holds the string denoting the agetenant field in the database.
	FieldAgetenant = "agetenant"

	// EdgeWifi holds the string denoting the wifi edge name in mutations.
	EdgeWifi = "Wifi"
	// EdgeRoomdetail holds the string denoting the roomdetail edge name in mutations.
	EdgeRoomdetail = "Roomdetail"
	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"
	// EdgeLeases holds the string denoting the leases edge name in mutations.
	EdgeLeases = "leases"
	// EdgeBill holds the string denoting the bill edge name in mutations.
	EdgeBill = "bill"
	// EdgeRepairinvoices holds the string denoting the repairinvoices edge name in mutations.
	EdgeRepairinvoices = "repairinvoices"

	// Table holds the table name of the lease in the database.
	Table = "leases"
	// WifiTable is the table the holds the Wifi relation/edge.
	WifiTable = "leases"
	// WifiInverseTable is the table name for the Wifi entity.
	// It exists in this package in order to avoid circular dependency with the "wifi" package.
	WifiInverseTable = "wifis"
	// WifiColumn is the table column denoting the Wifi relation/edge.
	WifiColumn = "wifi_id"
	// RoomdetailTable is the table the holds the Roomdetail relation/edge.
	RoomdetailTable = "leases"
	// RoomdetailInverseTable is the table name for the Roomdetail entity.
	// It exists in this package in order to avoid circular dependency with the "roomdetail" package.
	RoomdetailInverseTable = "roomdetails"
	// RoomdetailColumn is the table column denoting the Roomdetail relation/edge.
	RoomdetailColumn = "room_num"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "leases"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "employee_id"
	// LeasesTable is the table the holds the leases relation/edge.
	LeasesTable = "deposits"
	// LeasesInverseTable is the table name for the Deposit entity.
	// It exists in this package in order to avoid circular dependency with the "deposit" package.
	LeasesInverseTable = "deposits"
	// LeasesColumn is the table column denoting the leases relation/edge.
	LeasesColumn = "lease_id"
	// BillTable is the table the holds the bill relation/edge.
	BillTable = "bills"
	// BillInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	BillInverseTable = "bills"
	// BillColumn is the table column denoting the bill relation/edge.
	BillColumn = "lease_id"
	// RepairinvoicesTable is the table the holds the repairinvoices relation/edge.
	RepairinvoicesTable = "repairinvoices"
	// RepairinvoicesInverseTable is the table name for the Repairinvoice entity.
	// It exists in this package in order to avoid circular dependency with the "repairinvoice" package.
	RepairinvoicesInverseTable = "repairinvoices"
	// RepairinvoicesColumn is the table column denoting the repairinvoices relation/edge.
	RepairinvoicesColumn = "lease_id"
)

// Columns holds all SQL columns for lease fields.
var Columns = []string{
	FieldID,
	FieldAddedtime,
	FieldTenant,
	FieldNumbtenant,
	FieldIdtenant,
	FieldAgetenant,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Lease type.
var ForeignKeys = []string{
	"employee_id",
	"room_num",
	"wifi_id",
}

var (
	// TenantValidator is a validator for the "tenant" field. It is called by the builders before save.
	TenantValidator func(string) error
	// NumbtenantValidator is a validator for the "numbtenant" field. It is called by the builders before save.
	NumbtenantValidator func(string) error
	// IdtenantValidator is a validator for the "idtenant" field. It is called by the builders before save.
	IdtenantValidator func(string) error
	// AgetenantValidator is a validator for the "agetenant" field. It is called by the builders before save.
	AgetenantValidator func(int) error
)
