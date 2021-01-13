package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team15/app/controllers"
	_ "github.com/team15/app/docs"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/jobposition"
)

type Employees struct {
	Employee []Employee
}

type Employee struct {
	name          string
	email         string
	password      string
	jobpositionID int
}

type Jobpositions struct {
	Jobposition []Jobposition
}

type Jobposition struct {
	name string
}
type Statusds struct {
	Statusd []Statusd
}

type Statusd struct {
	Statusdname string
}

type CleanerNames struct {
	CleanerName []CleanerName
}

type CleanerName struct {
	cleanername string
}

type LengthTimes struct {
	LengthTime []LengthTime
}

type LengthTime struct {
	lengthtime string
}

type Quantitys struct {
	Quantity []Quantity
}

type Quantity struct {
	Quantity string
}

type Staytypes struct {
	Staytype []Staytype
}

type Staytype struct {
	Staytype string
}

type Equipments struct {
	Equipment []Equipment
}

type Equipment struct {
	Equipment string
}

type Facilities struct {
	Facilitie []Facilitie
}

type Facilitie struct {
	Facilitie string
}

type Nearbyplaces struct {
	Nearbyplace []Nearbyplace
}

type Nearbyplace struct {
	Nearbyplace string
}

type Wifis struct {
	Wifi []Wifi
}

type Wifi struct {
	Wifiname string
}
type Rentalstatuss struct {
	Rentalstatus []Rentalstatus
}

type Rentalstatus struct {
	Rentalstatus string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:team15.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	v1 := router.Group("/api/v1")
	controllers.NewStatusdController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewDepositController(v1, client)
	controllers.NewCleaningRoomController(v1, client)
	controllers.NewCleanerNameController(v1, client)
	controllers.NewLengthTimeController(v1, client)
	controllers.NewRoomdetailController(v1, client)
	controllers.NewQuantityController(v1, client)
	controllers.NewEquipmentController(v1, client)
	controllers.NewFacilitieController(v1, client)
	controllers.NewNearbyplaceController(v1, client)
	controllers.NewStaytypeController(v1, client)
	controllers.NewWifiController(v1, client)
	controllers.NewRepairinvoiceController(v1, client)
	controllers.NewRentalstatusController(v1, client)

	// Set Employees Data

	jobpositions := []string{"พนักงานหอพัก1", "พนักงานหอพัก2", "พนักงานหอพัก3", "พนักงานหอพัก4", "พนักงานหอพัก5", "พนักงานหอพัก6"}
	for _, jp := range jobpositions {
		client.Jobposition.
			Create().
			SetPositionName(jp).
			Save(context.Background())
	}

	employees := Employees{
		Employee: []Employee{
			Employee{"โรเจอร์", "rogerkung@gmail.com", "1234", 1},
			Employee{"นาตาชา", "natasha@gmail.com", "1234", 2},
			Employee{"โทนี่ สตาร์ค", "stark@gmail.com", "1234", 3},
			Employee{"สเตฟานี่ โรเจอร์", "step@gmail.com", "1234", 4},
			Employee{"นาตาชา โรไปนอฟ", "romanoff@gmail.com", "1234", 5},
			Employee{"นาตาชา โรมานอฟ", "romania@gmail.com", "1234", 6},
		},
	}

	for _, em := range employees.Employee {
		jp, err := client.Jobposition.
			Query().
			Where(jobposition.IDEQ(int(em.jobpositionID))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Employee.
			Create().
			SetName(em.name).
			SetEmail(em.email).
			SetPassword(em.password).
			SetJobposition(jp).
			Save(context.Background())
	}
	// Set Statusds Data
	statusds := Statusds{
		Statusd: []Statusd{
			Statusd{"uncompleted"},
			Statusd{"completed"},
		},
	}
	for _, s := range statusds.Statusd {
		client.Statusd.
			Create().
			SetStatusdname(s.Statusdname).
			Save(context.Background())
	}

	// Set CleanerNames Data
	cleanernames := CleanerNames{
		CleanerName: []CleanerName{
			CleanerName{"supaporn"},
			CleanerName{"sutida"},
			CleanerName{"sompong"},
		},
	}
	for _, cn := range cleanernames.CleanerName {
		client.CleanerName.
			Create().
			SetCleanername(cn.cleanername).
			Save(context.Background())
	}

	// Set LengthTimes Data
	lengthtimes := LengthTimes{
		LengthTime: []LengthTime{
			LengthTime{"30 min."},
			LengthTime{"1 hr."},
			LengthTime{"2 hr."},
			LengthTime{"3 hr."},
		},
	}
	for _, lt := range lengthtimes.LengthTime {
		client.LengthTime.
			Create().
			SetLengthtime(lt.lengthtime).
			Save(context.Background())
	}

	// Set Quantity Data
	quantitys := Quantitys{
		Quantity: []Quantity{
			Quantity{"2 people"},
			Quantity{"4 people"},
			Quantity{"6 people"},
		},
	}
	for _, qu := range quantitys.Quantity {
		client.Quantity.
			Create().
			SetQuantity(qu.Quantity).
			Save(context.Background())
	}

	// Set StayType Data
	staytypes := Staytypes{
		Staytype: []Staytype{
			Staytype{"Days"},
			Staytype{"Month"},
			Staytype{"Days & Month"},
		},
	}
	for _, st := range staytypes.Staytype {
		client.Staytype.
			Create().
			SetStaytype(st.Staytype).
			Save(context.Background())
	}

	// Set Equipment Data
	equipments := Equipments{
		Equipment: []Equipment{
			Equipment{"Ensuite bathroom"},
			Equipment{"Living zone in room"},
			Equipment{"Small Kitchen"},
		},
	}
	for _, eq := range equipments.Equipment {
		client.Equipment.
			Create().
			SetEquipment(eq.Equipment).
			Save(context.Background())
	}
	// Set Facility Data
	facilities := Facilities{
		Facilitie: []Facilitie{
			Facilitie{"Swimming pool"},
			Facilitie{"Fitness"},
		},
	}
	for _, fa := range facilities.Facilitie {
		client.Facilitie.
			Create().
			SetFacilitie(fa.Facilitie).
			Save(context.Background())
	}

	// Set NearbyPlace Data
	nearbyplaces := Nearbyplaces{
		Nearbyplace: []Nearbyplace{
			Nearbyplace{"Shopping mall"},
			Nearbyplace{"Supermarket"},
			Nearbyplace{"BTS"},
		},
	}
	for _, np := range nearbyplaces.Nearbyplace {
		client.Nearbyplace.
			Create().
			SetNearbyplace(np.Nearbyplace).
			Save(context.Background())
	}

	// Set Wifis Data
	wifis := Wifis{
		Wifi: []Wifi{
			Wifi{"no service"},
			Wifi{"service"},
		},
	}
	for _, s := range wifis.Wifi {
		client.Wifi.
			Create().
			SetWifiname(s.Wifiname).
			Save(context.Background())
	}
	// Set Rentalstatus Data
	rentalstatuss := Rentalstatuss{
		Rentalstatus: []Rentalstatus{
			Rentalstatus{"1"},
			Rentalstatus{"30"},
		},
	}
	for _, rs := range rentalstatuss.Rentalstatus {
		client.Rentalstatus.
			Create().
			SetRentalstatus(rs.Rentalstatus).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
