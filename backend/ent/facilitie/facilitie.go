// Code generated by entc, DO NOT EDIT.

package facilitie

const (
	// Label holds the string label denoting the facilitie type in the database.
	Label = "facilitie"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFacilitie holds the string denoting the facilitie field in the database.
	FieldFacilitie = "facilitie"

	// EdgeRoomdetail holds the string denoting the roomdetail edge name in mutations.
	EdgeRoomdetail = "roomdetail"

	// Table holds the table name of the facilitie in the database.
	Table = "facilities"
	// RoomdetailTable is the table the holds the roomdetail relation/edge.
	RoomdetailTable = "roomdetails"
	// RoomdetailInverseTable is the table name for the Roomdetail entity.
	// It exists in this package in order to avoid circular dependency with the "roomdetail" package.
	RoomdetailInverseTable = "roomdetails"
	// RoomdetailColumn is the table column denoting the roomdetail relation/edge.
	RoomdetailColumn = "facilitie_roomdetail"
)

// Columns holds all SQL columns for facilitie fields.
var Columns = []string{
	FieldID,
	FieldFacilitie,
}
