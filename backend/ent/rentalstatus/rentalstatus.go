// Code generated by entc, DO NOT EDIT.

package rentalstatus

const (
	// Label holds the string label denoting the rentalstatus type in the database.
	Label = "rentalstatus"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRentalstatus holds the string denoting the rentalstatus field in the database.
	FieldRentalstatus = "rentalstatus"

	// EdgeRepairinvoices holds the string denoting the repairinvoices edge name in mutations.
	EdgeRepairinvoices = "repairinvoices"

	// Table holds the table name of the rentalstatus in the database.
	Table = "rentalstatuses"
	// RepairinvoicesTable is the table the holds the repairinvoices relation/edge.
	RepairinvoicesTable = "repairinvoices"
	// RepairinvoicesInverseTable is the table name for the Repairinvoice entity.
	// It exists in this package in order to avoid circular dependency with the "repairinvoice" package.
	RepairinvoicesInverseTable = "repairinvoices"
	// RepairinvoicesColumn is the table column denoting the repairinvoices relation/edge.
	RepairinvoicesColumn = "rentalstatus_repairinvoices"
)

// Columns holds all SQL columns for rentalstatus fields.
var Columns = []string{
	FieldID,
	FieldRentalstatus,
}
