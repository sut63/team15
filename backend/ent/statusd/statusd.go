// Code generated by entc, DO NOT EDIT.

package statusd

const (
	// Label holds the string label denoting the statusd type in the database.
	Label = "statusd"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatusdname holds the string denoting the statusdname field in the database.
	FieldStatusdname = "statusdname"

	// EdgeStatusds holds the string denoting the statusds edge name in mutations.
	EdgeStatusds = "statusds"

	// Table holds the table name of the statusd in the database.
	Table = "statusds"
	// StatusdsTable is the table the holds the statusds relation/edge.
	StatusdsTable = "deposits"
	// StatusdsInverseTable is the table name for the Deposit entity.
	// It exists in this package in order to avoid circular dependency with the "deposit" package.
	StatusdsInverseTable = "deposits"
	// StatusdsColumn is the table column denoting the statusds relation/edge.
	StatusdsColumn = "statusd_id"
)

// Columns holds all SQL columns for statusd fields.
var Columns = []string{
	FieldID,
	FieldStatusdname,
}

var (
	// StatusdnameValidator is a validator for the "statusdname" field. It is called by the builders before save.
	StatusdnameValidator func(string) error
)
