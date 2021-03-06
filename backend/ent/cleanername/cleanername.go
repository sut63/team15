// Code generated by entc, DO NOT EDIT.

package cleanername

const (
	// Label holds the string label denoting the cleanername type in the database.
	Label = "cleanername"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCleanername holds the string denoting the cleanername field in the database.
	FieldCleanername = "cleanername"

	// EdgeCleaningrooms holds the string denoting the cleaningrooms edge name in mutations.
	EdgeCleaningrooms = "cleaningrooms"

	// Table holds the table name of the cleanername in the database.
	Table = "cleanernames"
	// CleaningroomsTable is the table the holds the cleaningrooms relation/edge.
	CleaningroomsTable = "cleaningrooms"
	// CleaningroomsInverseTable is the table name for the Cleaningroom entity.
	// It exists in this package in order to avoid circular dependency with the "cleaningroom" package.
	CleaningroomsInverseTable = "cleaningrooms"
	// CleaningroomsColumn is the table column denoting the cleaningrooms relation/edge.
	CleaningroomsColumn = "cleanerroom_id"
)

// Columns holds all SQL columns for cleanername fields.
var Columns = []string{
	FieldID,
	FieldCleanername,
}

var (
	// CleanernameValidator is a validator for the "cleanername" field. It is called by the builders before save.
	CleanernameValidator func(string) error
)
