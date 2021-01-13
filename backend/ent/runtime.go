// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/team15/app/ent/cleanername"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/lengthtime"
	"github.com/team15/app/ent/roomdetail"
	"github.com/team15/app/ent/schema"
	"github.com/team15/app/ent/statusd"
	"github.com/team15/app/ent/wifi"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	cleanernameFields := schema.CleanerName{}.Fields()
	_ = cleanernameFields
	// cleanernameDescCleanername is the schema descriptor for cleanername field.
	cleanernameDescCleanername := cleanernameFields[0].Descriptor()
	// cleanername.CleanernameValidator is a validator for the "cleanername" field. It is called by the builders before save.
	cleanername.CleanernameValidator = cleanernameDescCleanername.Validators[0].(func(string) error)
	employeeFields := schema.Employee{}.Fields()
	_ = employeeFields
	// employeeDescName is the schema descriptor for name field.
	employeeDescName := employeeFields[0].Descriptor()
	// employee.NameValidator is a validator for the "name" field. It is called by the builders before save.
	employee.NameValidator = employeeDescName.Validators[0].(func(string) error)
	// employeeDescEmail is the schema descriptor for email field.
	employeeDescEmail := employeeFields[1].Descriptor()
	// employee.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	employee.EmailValidator = employeeDescEmail.Validators[0].(func(string) error)
	// employeeDescPassword is the schema descriptor for password field.
	employeeDescPassword := employeeFields[2].Descriptor()
	// employee.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	employee.PasswordValidator = employeeDescPassword.Validators[0].(func(string) error)
	lengthtimeFields := schema.LengthTime{}.Fields()
	_ = lengthtimeFields
	// lengthtimeDescLengthtime is the schema descriptor for lengthtime field.
	lengthtimeDescLengthtime := lengthtimeFields[0].Descriptor()
	// lengthtime.LengthtimeValidator is a validator for the "lengthtime" field. It is called by the builders before save.
	lengthtime.LengthtimeValidator = lengthtimeDescLengthtime.Validators[0].(func(string) error)
	roomdetailFields := schema.Roomdetail{}.Fields()
	_ = roomdetailFields
	// roomdetailDescRoomnumber is the schema descriptor for roomnumber field.
	roomdetailDescRoomnumber := roomdetailFields[0].Descriptor()
	// roomdetail.RoomnumberValidator is a validator for the "roomnumber" field. It is called by the builders before save.
	roomdetail.RoomnumberValidator = roomdetailDescRoomnumber.Validators[0].(func(string) error)
	// roomdetailDescRoomtypename is the schema descriptor for roomtypename field.
	roomdetailDescRoomtypename := roomdetailFields[1].Descriptor()
	// roomdetail.RoomtypenameValidator is a validator for the "roomtypename" field. It is called by the builders before save.
	roomdetail.RoomtypenameValidator = roomdetailDescRoomtypename.Validators[0].(func(string) error)
	// roomdetailDescRoomprice is the schema descriptor for roomprice field.
	roomdetailDescRoomprice := roomdetailFields[2].Descriptor()
	// roomdetail.RoompriceValidator is a validator for the "roomprice" field. It is called by the builders before save.
	roomdetail.RoompriceValidator = roomdetailDescRoomprice.Validators[0].(func(string) error)
	statusdFields := schema.Statusd{}.Fields()
	_ = statusdFields
	// statusdDescStatusdname is the schema descriptor for statusdname field.
	statusdDescStatusdname := statusdFields[0].Descriptor()
	// statusd.StatusdnameValidator is a validator for the "statusdname" field. It is called by the builders before save.
	statusd.StatusdnameValidator = statusdDescStatusdname.Validators[0].(func(string) error)
	wifiFields := schema.Wifi{}.Fields()
	_ = wifiFields
	// wifiDescWifiname is the schema descriptor for wifiname field.
	wifiDescWifiname := wifiFields[0].Descriptor()
	// wifi.WifinameValidator is a validator for the "wifiname" field. It is called by the builders before save.
	wifi.WifinameValidator = wifiDescWifiname.Validators[0].(func(string) error)
}
