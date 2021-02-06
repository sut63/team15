package schema

import (
	
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Cleaningroom holds the schema definition for the Cleaningroom entity.
type Cleaningroom struct {
	ent.Schema
}

// Fields of the Cleaningroom.
func (Cleaningroom) Fields() []ent.Field {
	return []ent.Field{
		field.String("note").Validate(func(s string) error {
            match, _ := regexp.MatchString("^[ก-๏]+$", s)
            if !match {
                return errors.New("กรุณากรอกเฉพาะภาษาไทยเท่านั้น")
            }
            return nil
		}),
		field.Time("dateandstarttime"),
		field.String("phonenumber").MaxLen(12).MinLen(12),
		field.Int("numofem").Min(0).Max(4),
	}
}

// Edges of the Cleaningroom.
func (Cleaningroom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roomdetail", Roomdetail.Type).
			Ref("cleaningrooms").
			Unique(),

		edge.From("Cleanername", Cleanername.Type).
			Ref("cleaningrooms").
			Unique(),

		edge.From("Lengthtime", Lengthtime.Type).
			Ref("cleaningrooms").
			Unique(),

		edge.From("Employee", Employee.Type).
			Ref("cleaningrooms").
			Unique(),
	
	}
}
