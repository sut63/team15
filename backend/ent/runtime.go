// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/team15/app/ent/bill"
	"github.com/team15/app/ent/cleanername"
	"github.com/team15/app/ent/deposit"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/lengthtime"
	"github.com/team15/app/ent/payment"
	"github.com/team15/app/ent/roomdetail"
	"github.com/team15/app/ent/schema"
	"github.com/team15/app/ent/situation"
	"github.com/team15/app/ent/statusd"
	"github.com/team15/app/ent/wifi"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	billFields := schema.Bill{}.Fields()
	_ = billFields
	// billDescTell is the schema descriptor for tell field.
	billDescTell := billFields[1].Descriptor()
	// bill.TellValidator is a validator for the "tell" field. It is called by the builders before save.
	bill.TellValidator = func() func(string) error {
		validators := billDescTell.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(tell string) error {
			for _, fn := range fns {
				if err := fn(tell); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// billDescTaxpayer is the schema descriptor for taxpayer field.
	billDescTaxpayer := billFields[2].Descriptor()
	// bill.TaxpayerValidator is a validator for the "taxpayer" field. It is called by the builders before save.
	bill.TaxpayerValidator = func() func(string) error {
		validators := billDescTaxpayer.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(taxpayer string) error {
			for _, fn := range fns {
				if err := fn(taxpayer); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// billDescTotal is the schema descriptor for total field.
	billDescTotal := billFields[3].Descriptor()
	// bill.TotalValidator is a validator for the "total" field. It is called by the builders before save.
	bill.TotalValidator = billDescTotal.Validators[0].(func(string) error)
	cleanernameFields := schema.CleanerName{}.Fields()
	_ = cleanernameFields
	// cleanernameDescCleanername is the schema descriptor for cleanername field.
	cleanernameDescCleanername := cleanernameFields[0].Descriptor()
	// cleanername.CleanernameValidator is a validator for the "cleanername" field. It is called by the builders before save.
	cleanername.CleanernameValidator = cleanernameDescCleanername.Validators[0].(func(string) error)
	depositFields := schema.Deposit{}.Fields()
	_ = depositFields
	// depositDescInfo is the schema descriptor for info field.
	depositDescInfo := depositFields[1].Descriptor()
	// deposit.InfoValidator is a validator for the "info" field. It is called by the builders before save.
	deposit.InfoValidator = depositDescInfo.Validators[0].(func(string) error)
	// depositDescDepositor is the schema descriptor for depositor field.
	depositDescDepositor := depositFields[2].Descriptor()
	// deposit.DepositorValidator is a validator for the "depositor" field. It is called by the builders before save.
	deposit.DepositorValidator = depositDescDepositor.Validators[0].(func(string) error)
	// depositDescDepositortell is the schema descriptor for depositortell field.
	depositDescDepositortell := depositFields[3].Descriptor()
	// deposit.DepositortellValidator is a validator for the "depositortell" field. It is called by the builders before save.
	deposit.DepositortellValidator = func() func(string) error {
		validators := depositDescDepositortell.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(depositortell string) error {
			for _, fn := range fns {
				if err := fn(depositortell); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// depositDescRecipienttell is the schema descriptor for recipienttell field.
	depositDescRecipienttell := depositFields[4].Descriptor()
	// deposit.RecipienttellValidator is a validator for the "recipienttell" field. It is called by the builders before save.
	deposit.RecipienttellValidator = func() func(string) error {
		validators := depositDescRecipienttell.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(recipienttell string) error {
			for _, fn := range fns {
				if err := fn(recipienttell); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// depositDescParcelcode is the schema descriptor for parcelcode field.
	depositDescParcelcode := depositFields[5].Descriptor()
	// deposit.ParcelcodeValidator is a validator for the "parcelcode" field. It is called by the builders before save.
	deposit.ParcelcodeValidator = depositDescParcelcode.Validators[0].(func(string) error)
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
	paymentFields := schema.Payment{}.Fields()
	_ = paymentFields
	// paymentDescPaymentname is the schema descriptor for paymentname field.
	paymentDescPaymentname := paymentFields[0].Descriptor()
	// payment.PaymentnameValidator is a validator for the "paymentname" field. It is called by the builders before save.
	payment.PaymentnameValidator = paymentDescPaymentname.Validators[0].(func(string) error)
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
	roomdetail.RoompriceValidator = func() func(string) error {
		validators := roomdetailDescRoomprice.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(roomprice string) error {
			for _, fn := range fns {
				if err := fn(roomprice); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// roomdetailDescPhone is the schema descriptor for phone field.
	roomdetailDescPhone := roomdetailFields[3].Descriptor()
	// roomdetail.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	roomdetail.PhoneValidator = func() func(string) error {
		validators := roomdetailDescPhone.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(phone string) error {
			for _, fn := range fns {
				if err := fn(phone); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// roomdetailDescSleep is the schema descriptor for sleep field.
	roomdetailDescSleep := roomdetailFields[4].Descriptor()
	// roomdetail.SleepValidator is a validator for the "sleep" field. It is called by the builders before save.
	roomdetail.SleepValidator = func() func(int) error {
		validators := roomdetailDescSleep.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(sleep int) error {
			for _, fn := range fns {
				if err := fn(sleep); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// roomdetailDescBed is the schema descriptor for bed field.
	roomdetailDescBed := roomdetailFields[5].Descriptor()
	// roomdetail.BedValidator is a validator for the "bed" field. It is called by the builders before save.
	roomdetail.BedValidator = func() func(int) error {
		validators := roomdetailDescBed.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(bed int) error {
			for _, fn := range fns {
				if err := fn(bed); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	situationFields := schema.Situation{}.Fields()
	_ = situationFields
	// situationDescSituationname is the schema descriptor for situationname field.
	situationDescSituationname := situationFields[0].Descriptor()
	// situation.SituationnameValidator is a validator for the "situationname" field. It is called by the builders before save.
	situation.SituationnameValidator = situationDescSituationname.Validators[0].(func(string) error)
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
