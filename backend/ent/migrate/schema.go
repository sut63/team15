// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// BedtypesColumns holds the columns for the "bedtypes" table.
	BedtypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "bedtypename", Type: field.TypeString, Unique: true},
	}
	// BedtypesTable holds the schema information for the "bedtypes" table.
	BedtypesTable = &schema.Table{
		Name:        "bedtypes",
		Columns:     BedtypesColumns,
		PrimaryKey:  []*schema.Column{BedtypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// BillsColumns holds the columns for the "bills" table.
	BillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "addedtime", Type: field.TypeTime},
		{Name: "tell", Type: field.TypeString, Size: 12},
		{Name: "taxpayer", Type: field.TypeString, Size: 17},
		{Name: "total", Type: field.TypeString},
		{Name: "lease_id", Type: field.TypeInt, Nullable: true},
		{Name: "payment_id", Type: field.TypeInt, Nullable: true},
		{Name: "situation_id", Type: field.TypeInt, Nullable: true},
	}
	// BillsTable holds the schema information for the "bills" table.
	BillsTable = &schema.Table{
		Name:       "bills",
		Columns:    BillsColumns,
		PrimaryKey: []*schema.Column{BillsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "bills_leases_bill",
				Columns: []*schema.Column{BillsColumns[5]},

				RefColumns: []*schema.Column{LeasesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bills_payments_payments",
				Columns: []*schema.Column{BillsColumns[6]},

				RefColumns: []*schema.Column{PaymentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bills_situations_situations",
				Columns: []*schema.Column{BillsColumns[7]},

				RefColumns: []*schema.Column{SituationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CleanernamesColumns holds the columns for the "cleanernames" table.
	CleanernamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "cleanername", Type: field.TypeString, Unique: true},
	}
	// CleanernamesTable holds the schema information for the "cleanernames" table.
	CleanernamesTable = &schema.Table{
		Name:        "cleanernames",
		Columns:     CleanernamesColumns,
		PrimaryKey:  []*schema.Column{CleanernamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CleaningroomsColumns holds the columns for the "cleaningrooms" table.
	CleaningroomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "note", Type: field.TypeString},
		{Name: "dateandstarttime", Type: field.TypeTime},
		{Name: "phonenumber", Type: field.TypeString, Size: 12},
		{Name: "numofem", Type: field.TypeInt},
		{Name: "cleanerroom_id", Type: field.TypeInt, Nullable: true},
		{Name: "employee_id", Type: field.TypeInt, Nullable: true},
		{Name: "lengthtime_id", Type: field.TypeInt, Nullable: true},
		{Name: "roomdetail_id", Type: field.TypeInt, Nullable: true},
	}
	// CleaningroomsTable holds the schema information for the "cleaningrooms" table.
	CleaningroomsTable = &schema.Table{
		Name:       "cleaningrooms",
		Columns:    CleaningroomsColumns,
		PrimaryKey: []*schema.Column{CleaningroomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "cleaningrooms_cleanernames_cleaningrooms",
				Columns: []*schema.Column{CleaningroomsColumns[5]},

				RefColumns: []*schema.Column{CleanernamesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "cleaningrooms_employees_cleaningrooms",
				Columns: []*schema.Column{CleaningroomsColumns[6]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "cleaningrooms_lengthtimes_cleaningrooms",
				Columns: []*schema.Column{CleaningroomsColumns[7]},

				RefColumns: []*schema.Column{LengthtimesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "cleaningrooms_roomdetails_cleaningrooms",
				Columns: []*schema.Column{CleaningroomsColumns[8]},

				RefColumns: []*schema.Column{RoomdetailsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DepositsColumns holds the columns for the "deposits" table.
	DepositsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "addedtime", Type: field.TypeTime},
		{Name: "info", Type: field.TypeString},
		{Name: "depositor", Type: field.TypeString},
		{Name: "depositortell", Type: field.TypeString, Size: 12},
		{Name: "recipienttell", Type: field.TypeString, Size: 12},
		{Name: "parcelcode", Type: field.TypeString},
		{Name: "employee_id", Type: field.TypeInt, Nullable: true},
		{Name: "lease_id", Type: field.TypeInt, Nullable: true},
		{Name: "statusd_id", Type: field.TypeInt, Nullable: true},
	}
	// DepositsTable holds the schema information for the "deposits" table.
	DepositsTable = &schema.Table{
		Name:       "deposits",
		Columns:    DepositsColumns,
		PrimaryKey: []*schema.Column{DepositsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "deposits_employees_employees",
				Columns: []*schema.Column{DepositsColumns[7]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "deposits_leases_leases",
				Columns: []*schema.Column{DepositsColumns[8]},

				RefColumns: []*schema.Column{LeasesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "deposits_statusds_statusds",
				Columns: []*schema.Column{DepositsColumns[9]},

				RefColumns: []*schema.Column{StatusdsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EmployeesColumns holds the columns for the "employees" table.
	EmployeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "jobposition_employees", Type: field.TypeInt, Nullable: true},
	}
	// EmployeesTable holds the schema information for the "employees" table.
	EmployeesTable = &schema.Table{
		Name:       "employees",
		Columns:    EmployeesColumns,
		PrimaryKey: []*schema.Column{EmployeesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "employees_jobpositions_employees",
				Columns: []*schema.Column{EmployeesColumns[4]},

				RefColumns: []*schema.Column{JobpositionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// JobpositionsColumns holds the columns for the "jobpositions" table.
	JobpositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "positionname", Type: field.TypeString, Unique: true},
	}
	// JobpositionsTable holds the schema information for the "jobpositions" table.
	JobpositionsTable = &schema.Table{
		Name:        "jobpositions",
		Columns:     JobpositionsColumns,
		PrimaryKey:  []*schema.Column{JobpositionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// LeasesColumns holds the columns for the "leases" table.
	LeasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "addedtime", Type: field.TypeTime},
		{Name: "tenant", Type: field.TypeString},
		{Name: "numbtenant", Type: field.TypeString, Size: 12},
		{Name: "idtenant", Type: field.TypeString, Size: 17},
		{Name: "agetenant", Type: field.TypeInt},
		{Name: "employee_id", Type: field.TypeInt, Nullable: true},
		{Name: "room_num", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "wifi_id", Type: field.TypeInt, Nullable: true},
	}
	// LeasesTable holds the schema information for the "leases" table.
	LeasesTable = &schema.Table{
		Name:       "leases",
		Columns:    LeasesColumns,
		PrimaryKey: []*schema.Column{LeasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "leases_employees_leasess",
				Columns: []*schema.Column{LeasesColumns[6]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "leases_roomdetails_roomdetails",
				Columns: []*schema.Column{LeasesColumns[7]},

				RefColumns: []*schema.Column{RoomdetailsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "leases_wifis_wifis",
				Columns: []*schema.Column{LeasesColumns[8]},

				RefColumns: []*schema.Column{WifisColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LengthtimesColumns holds the columns for the "lengthtimes" table.
	LengthtimesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "lengthtime", Type: field.TypeString, Unique: true},
	}
	// LengthtimesTable holds the schema information for the "lengthtimes" table.
	LengthtimesTable = &schema.Table{
		Name:        "lengthtimes",
		Columns:     LengthtimesColumns,
		PrimaryKey:  []*schema.Column{LengthtimesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PaymentsColumns holds the columns for the "payments" table.
	PaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "paymentname", Type: field.TypeString, Unique: true},
	}
	// PaymentsTable holds the schema information for the "payments" table.
	PaymentsTable = &schema.Table{
		Name:        "payments",
		Columns:     PaymentsColumns,
		PrimaryKey:  []*schema.Column{PaymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PetrulesColumns holds the columns for the "petrules" table.
	PetrulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "petrule", Type: field.TypeString, Unique: true},
	}
	// PetrulesTable holds the schema information for the "petrules" table.
	PetrulesTable = &schema.Table{
		Name:        "petrules",
		Columns:     PetrulesColumns,
		PrimaryKey:  []*schema.Column{PetrulesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PledgesColumns holds the columns for the "pledges" table.
	PledgesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "provision", Type: field.TypeString, Unique: true},
	}
	// PledgesTable holds the schema information for the "pledges" table.
	PledgesTable = &schema.Table{
		Name:        "pledges",
		Columns:     PledgesColumns,
		PrimaryKey:  []*schema.Column{PledgesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RentalstatusesColumns holds the columns for the "rentalstatuses" table.
	RentalstatusesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "rentalstatus", Type: field.TypeString, Unique: true},
	}
	// RentalstatusesTable holds the schema information for the "rentalstatuses" table.
	RentalstatusesTable = &schema.Table{
		Name:        "rentalstatuses",
		Columns:     RentalstatusesColumns,
		PrimaryKey:  []*schema.Column{RentalstatusesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RepairinvoicesColumns holds the columns for the "repairinvoices" table.
	RepairinvoicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "bequipment", Type: field.TypeString},
		{Name: "employee_id", Type: field.TypeInt, Nullable: true},
		{Name: "rentalstatus_repairinvoices", Type: field.TypeInt, Nullable: true},
	}
	// RepairinvoicesTable holds the schema information for the "repairinvoices" table.
	RepairinvoicesTable = &schema.Table{
		Name:       "repairinvoices",
		Columns:    RepairinvoicesColumns,
		PrimaryKey: []*schema.Column{RepairinvoicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "repairinvoices_employees_repairinvoices",
				Columns: []*schema.Column{RepairinvoicesColumns[2]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "repairinvoices_rentalstatuses_repairinvoices",
				Columns: []*schema.Column{RepairinvoicesColumns[3]},

				RefColumns: []*schema.Column{RentalstatusesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RoomdetailsColumns holds the columns for the "roomdetails" table.
	RoomdetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "roomnumber", Type: field.TypeString, Unique: true},
		{Name: "roomtypename", Type: field.TypeString},
		{Name: "roomprice", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Size: 12},
		{Name: "sleep", Type: field.TypeInt},
		{Name: "bed", Type: field.TypeInt},
		{Name: "bedtype_roomdetails", Type: field.TypeInt, Nullable: true},
		{Name: "employee_id", Type: field.TypeInt, Nullable: true},
		{Name: "roomdetail_id", Type: field.TypeInt, Nullable: true},
		{Name: "petrule_roomdetails", Type: field.TypeInt, Nullable: true},
		{Name: "pledge_roomdetails", Type: field.TypeInt, Nullable: true},
		{Name: "staytype_roomdetails", Type: field.TypeInt, Nullable: true},
	}
	// RoomdetailsTable holds the schema information for the "roomdetails" table.
	RoomdetailsTable = &schema.Table{
		Name:       "roomdetails",
		Columns:    RoomdetailsColumns,
		PrimaryKey: []*schema.Column{RoomdetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "roomdetails_bedtypes_roomdetails",
				Columns: []*schema.Column{RoomdetailsColumns[7]},

				RefColumns: []*schema.Column{BedtypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "roomdetails_employees_roomdetails",
				Columns: []*schema.Column{RoomdetailsColumns[8]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "roomdetails_jobpositions_roomdetails",
				Columns: []*schema.Column{RoomdetailsColumns[9]},

				RefColumns: []*schema.Column{JobpositionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "roomdetails_petrules_roomdetails",
				Columns: []*schema.Column{RoomdetailsColumns[10]},

				RefColumns: []*schema.Column{PetrulesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "roomdetails_pledges_roomdetails",
				Columns: []*schema.Column{RoomdetailsColumns[11]},

				RefColumns: []*schema.Column{PledgesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "roomdetails_staytypes_roomdetails",
				Columns: []*schema.Column{RoomdetailsColumns[12]},

				RefColumns: []*schema.Column{StaytypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SituationsColumns holds the columns for the "situations" table.
	SituationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "situationname", Type: field.TypeString, Unique: true},
	}
	// SituationsTable holds the schema information for the "situations" table.
	SituationsTable = &schema.Table{
		Name:        "situations",
		Columns:     SituationsColumns,
		PrimaryKey:  []*schema.Column{SituationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// StatusdsColumns holds the columns for the "statusds" table.
	StatusdsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "statusdname", Type: field.TypeString, Unique: true},
	}
	// StatusdsTable holds the schema information for the "statusds" table.
	StatusdsTable = &schema.Table{
		Name:        "statusds",
		Columns:     StatusdsColumns,
		PrimaryKey:  []*schema.Column{StatusdsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// StaytypesColumns holds the columns for the "staytypes" table.
	StaytypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "staytype", Type: field.TypeString, Unique: true},
	}
	// StaytypesTable holds the schema information for the "staytypes" table.
	StaytypesTable = &schema.Table{
		Name:        "staytypes",
		Columns:     StaytypesColumns,
		PrimaryKey:  []*schema.Column{StaytypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// WifisColumns holds the columns for the "wifis" table.
	WifisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "wifiname", Type: field.TypeString, Unique: true},
	}
	// WifisTable holds the schema information for the "wifis" table.
	WifisTable = &schema.Table{
		Name:        "wifis",
		Columns:     WifisColumns,
		PrimaryKey:  []*schema.Column{WifisColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BedtypesTable,
		BillsTable,
		CleanernamesTable,
		CleaningroomsTable,
		DepositsTable,
		EmployeesTable,
		JobpositionsTable,
		LeasesTable,
		LengthtimesTable,
		PaymentsTable,
		PetrulesTable,
		PledgesTable,
		RentalstatusesTable,
		RepairinvoicesTable,
		RoomdetailsTable,
		SituationsTable,
		StatusdsTable,
		StaytypesTable,
		WifisTable,
	}
)

func init() {
	BillsTable.ForeignKeys[0].RefTable = LeasesTable
	BillsTable.ForeignKeys[1].RefTable = PaymentsTable
	BillsTable.ForeignKeys[2].RefTable = SituationsTable
	CleaningroomsTable.ForeignKeys[0].RefTable = CleanernamesTable
	CleaningroomsTable.ForeignKeys[1].RefTable = EmployeesTable
	CleaningroomsTable.ForeignKeys[2].RefTable = LengthtimesTable
	CleaningroomsTable.ForeignKeys[3].RefTable = RoomdetailsTable
	DepositsTable.ForeignKeys[0].RefTable = EmployeesTable
	DepositsTable.ForeignKeys[1].RefTable = LeasesTable
	DepositsTable.ForeignKeys[2].RefTable = StatusdsTable
	EmployeesTable.ForeignKeys[0].RefTable = JobpositionsTable
	LeasesTable.ForeignKeys[0].RefTable = EmployeesTable
	LeasesTable.ForeignKeys[1].RefTable = RoomdetailsTable
	LeasesTable.ForeignKeys[2].RefTable = WifisTable
	RepairinvoicesTable.ForeignKeys[0].RefTable = EmployeesTable
	RepairinvoicesTable.ForeignKeys[1].RefTable = RentalstatusesTable
	RoomdetailsTable.ForeignKeys[0].RefTable = BedtypesTable
	RoomdetailsTable.ForeignKeys[1].RefTable = EmployeesTable
	RoomdetailsTable.ForeignKeys[2].RefTable = JobpositionsTable
	RoomdetailsTable.ForeignKeys[3].RefTable = PetrulesTable
	RoomdetailsTable.ForeignKeys[4].RefTable = PledgesTable
	RoomdetailsTable.ForeignKeys[5].RefTable = StaytypesTable
}
