package schema

import (
   "errors"
   "regexp"

   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)

// Deposit holds the schema definition for the Deposit entity.
type Deposit struct {
	ent.Schema
}

// Fields of the Deposit.
func (Deposit) Fields() []ent.Field {
	return []ent.Field{
	   field.Time("addedtime"),
	   field.String("info").NotEmpty(),
	   field.String("depositorname").NotEmpty(),
	   field.String("depositortell").MaxLen(12).MinLen(12),
	   field.String("recipienttell").MaxLen(12).MinLen(12),
	   field.String("parcelcode").Validate(func(s string) error {
			match, _ := regexp.MatchString("^([A-Z]{3})-([0-9]{4})$", s)
			if !match {
				return errors.New("The package code format is not valid.")
			}
			return nil
		}),
   }
}

// Edges of the Deposit.
func (Deposit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Employee", Employee.Type).
            Ref("employees").
            Unique(),

		edge.From("Statusd", Statusd.Type).
            Ref("statusds").
			Unique(),

	    edge.From("Lease", Lease.Type).
            Ref("leases").
            Unique(),
	}
}
