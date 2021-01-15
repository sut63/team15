// Code generated by entc, DO NOT EDIT.

package situation

const (
	// Label holds the string label denoting the situation type in the database.
	Label = "situation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSituationname holds the string denoting the situationname field in the database.
	FieldSituationname = "situationname"

	// EdgeSituations holds the string denoting the situations edge name in mutations.
	EdgeSituations = "situations"

	// Table holds the table name of the situation in the database.
	Table = "situations"
	// SituationsTable is the table the holds the situations relation/edge.
	SituationsTable = "bills"
	// SituationsInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	SituationsInverseTable = "bills"
	// SituationsColumn is the table column denoting the situations relation/edge.
	SituationsColumn = "situation_id"
)

// Columns holds all SQL columns for situation fields.
var Columns = []string{
	FieldID,
	FieldSituationname,
}

var (
	// SituationnameValidator is a validator for the "situationname" field. It is called by the builders before save.
	SituationnameValidator func(string) error
)
