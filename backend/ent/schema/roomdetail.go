package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Roomdetail holds the schema definition for the Roomdetail entity.
type Roomdetail struct {
	ent.Schema
}

// Fields of the Roomdetail.
func (Roomdetail) Fields() []ent.Field {
	return []ent.Field{
		field.String("roomnumber").Unique().Validate(func(s string) error {
			match, _ := regexp.MatchString("[ABD]\\d{3}", s)
			if !match {
				return errors.New("รูปแบบเลขห้องไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("roomtypename").NotEmpty(),
		field.Int("roomprice").Positive().Min(0),
		field.String("phone").MaxLen(12).MinLen(12),
		field.Int("sleep").Min(0).Max(10),
		field.Int("bed").Min(0).Max(4),
	}
}

// Edges of the Roomdetail.
func (Roomdetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pledge", Pledge.Type).Ref("roomdetails").Unique(),
		edge.From("petrule", Petrule.Type).Ref("roomdetails").Unique(),
		edge.From("bedtype", Bedtype.Type).Ref("roomdetails").Unique(),
		edge.From("employee", Employee.Type).Ref("roomdetails").Unique(),
		edge.From("jobposition", Jobposition.Type).Ref("roomdetails").Unique(),
		edge.From("staytype", Staytype.Type).Ref("roomdetails").Unique(),
		edge.To("roomdetails", Lease.Type).Unique().StorageKey(edge.Column("room_num")),
		edge.To("cleaningrooms", Cleaningroom.Type).StorageKey(edge.Column("roomdetail_id")),
	}
}
