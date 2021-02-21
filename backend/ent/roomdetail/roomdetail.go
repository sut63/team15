// Code generated by entc, DO NOT EDIT.

package roomdetail

const (
	// Label holds the string label denoting the roomdetail type in the database.
	Label = "roomdetail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRoomnumber holds the string denoting the roomnumber field in the database.
	FieldRoomnumber = "roomnumber"
	// FieldRoomtypename holds the string denoting the roomtypename field in the database.
	FieldRoomtypename = "roomtypename"
	// FieldRoomprice holds the string denoting the roomprice field in the database.
	FieldRoomprice = "roomprice"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldSleep holds the string denoting the sleep field in the database.
	FieldSleep = "sleep"
	// FieldBed holds the string denoting the bed field in the database.
	FieldBed = "bed"

	// EdgePledge holds the string denoting the pledge edge name in mutations.
	EdgePledge = "pledge"
	// EdgePetrule holds the string denoting the petrule edge name in mutations.
	EdgePetrule = "petrule"
	// EdgeBedtype holds the string denoting the bedtype edge name in mutations.
	EdgeBedtype = "bedtype"
	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"
	// EdgeJobposition holds the string denoting the jobposition edge name in mutations.
	EdgeJobposition = "jobposition"
	// EdgeStaytype holds the string denoting the staytype edge name in mutations.
	EdgeStaytype = "staytype"
	// EdgeRoomdetails holds the string denoting the roomdetails edge name in mutations.
	EdgeRoomdetails = "roomdetails"
	// EdgeCleaningrooms holds the string denoting the cleaningrooms edge name in mutations.
	EdgeCleaningrooms = "cleaningrooms"

	// Table holds the table name of the roomdetail in the database.
	Table = "roomdetails"
	// PledgeTable is the table the holds the pledge relation/edge.
	PledgeTable = "roomdetails"
	// PledgeInverseTable is the table name for the Pledge entity.
	// It exists in this package in order to avoid circular dependency with the "pledge" package.
	PledgeInverseTable = "pledges"
	// PledgeColumn is the table column denoting the pledge relation/edge.
	PledgeColumn = "pledge_roomdetails"
	// PetruleTable is the table the holds the petrule relation/edge.
	PetruleTable = "roomdetails"
	// PetruleInverseTable is the table name for the Petrule entity.
	// It exists in this package in order to avoid circular dependency with the "petrule" package.
	PetruleInverseTable = "petrules"
	// PetruleColumn is the table column denoting the petrule relation/edge.
	PetruleColumn = "petrule_roomdetails"
	// BedtypeTable is the table the holds the bedtype relation/edge.
	BedtypeTable = "roomdetails"
	// BedtypeInverseTable is the table name for the Bedtype entity.
	// It exists in this package in order to avoid circular dependency with the "bedtype" package.
	BedtypeInverseTable = "bedtypes"
	// BedtypeColumn is the table column denoting the bedtype relation/edge.
	BedtypeColumn = "bedtype_roomdetails"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "roomdetails"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "employee_id"
	// JobpositionTable is the table the holds the jobposition relation/edge.
	JobpositionTable = "roomdetails"
	// JobpositionInverseTable is the table name for the Jobposition entity.
	// It exists in this package in order to avoid circular dependency with the "jobposition" package.
	JobpositionInverseTable = "jobpositions"
	// JobpositionColumn is the table column denoting the jobposition relation/edge.
	JobpositionColumn = "roomdetail_id"
	// StaytypeTable is the table the holds the staytype relation/edge.
	StaytypeTable = "roomdetails"
	// StaytypeInverseTable is the table name for the Staytype entity.
	// It exists in this package in order to avoid circular dependency with the "staytype" package.
	StaytypeInverseTable = "staytypes"
	// StaytypeColumn is the table column denoting the staytype relation/edge.
	StaytypeColumn = "staytype_roomdetails"
	// RoomdetailsTable is the table the holds the roomdetails relation/edge.
	RoomdetailsTable = "leases"
	// RoomdetailsInverseTable is the table name for the Lease entity.
	// It exists in this package in order to avoid circular dependency with the "lease" package.
	RoomdetailsInverseTable = "leases"
	// RoomdetailsColumn is the table column denoting the roomdetails relation/edge.
	RoomdetailsColumn = "room_num"
	// CleaningroomsTable is the table the holds the cleaningrooms relation/edge.
	CleaningroomsTable = "cleaningrooms"
	// CleaningroomsInverseTable is the table name for the Cleaningroom entity.
	// It exists in this package in order to avoid circular dependency with the "cleaningroom" package.
	CleaningroomsInverseTable = "cleaningrooms"
	// CleaningroomsColumn is the table column denoting the cleaningrooms relation/edge.
	CleaningroomsColumn = "roomdetail_id"
)

// Columns holds all SQL columns for roomdetail fields.
var Columns = []string{
	FieldID,
	FieldRoomnumber,
	FieldRoomtypename,
	FieldRoomprice,
	FieldPhone,
	FieldSleep,
	FieldBed,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Roomdetail type.
var ForeignKeys = []string{
	"bedtype_roomdetails",
	"employee_id",
	"roomdetail_id",
	"petrule_roomdetails",
	"pledge_roomdetails",
	"staytype_roomdetails",
}

var (
	// RoomnumberValidator is a validator for the "roomnumber" field. It is called by the builders before save.
	RoomnumberValidator func(string) error
	// RoomtypenameValidator is a validator for the "roomtypename" field. It is called by the builders before save.
	RoomtypenameValidator func(string) error
	// RoompriceValidator is a validator for the "roomprice" field. It is called by the builders before save.
	RoompriceValidator func(int) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// SleepValidator is a validator for the "sleep" field. It is called by the builders before save.
	SleepValidator func(int) error
	// BedValidator is a validator for the "bed" field. It is called by the builders before save.
	BedValidator func(int) error
)
