package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team15/app/controllers"
	_ "github.com/team15/app/docs"
	"github.com/team15/app/ent"
)

type Quantitys struct {
	Quantity []Quantity
}

type Quantity struct {
	Quantity int
}

type StayTypes struct {
	StayType []StayType
}

type StayType struct {
	StayType string
}

type Equipments struct {
	Equipment []Equipment
}

type Equipment struct {
	Equipment string
}

type Facilitys struct {
	Facility []Facility
}

type Facility struct {
	Facility string
}

type NearbyPlaces struct {
	NearbyPlace []NearbyPlace
}

type NearbyPlace struct {
	Placename string
}

type Employees struct {
	Employee []Employee
}

type Employee struct {
	Employeename  string
	Employeeemail string
	Password      string
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
	controllers.NewRoomController(v1, client)
	controllers.NewQuantityController(v1, client)
	controllers.NewStayTypeController(v1, client)
	controllers.NewFacilityController(v1, client)
	controllers.NewEquipmentController(v1, client)
	controllers.NewNearbyPlaceController(v1, client)
	controllers.NewStatusdController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewDepositController(v1, client)
	controllers.NewCleaningRoomController(v1, client)
	controllers.NewCleanerNameController(v1, client)
	controllers.NewLengthTimeController(v1, client)

	// Set Quantity Data
	quantitys := Quantitys{
		Quantity: []Quantity{
			Quantity{2},
			Quantity{4},
			Quantity{6},
		},
	}
	for _, q := range quantitys.Quantity {
		client.Quantity.
			Create().
			SetQuantity(q.Quantity).
			Save(context.Background())
	}

	// Set StayType Data
	staytypes := StayTypes{
		StayType: []StayType{
			StayType{"Days"},
			StayType{"Month"},
			StayType{"Days & Month"},
		},
	}
	for _, st := range staytypes.StayType {
		client.StayType.
			Create().
			SetStaytype(st.StayType).
			Save(context.Background())
	}

	// Set Equipment Data
	equipments := Equipments{
		Equipment: []Equipment{
			Equipment{"Bed"},
			Equipment{"Wardrobe"},
			Equipment{"Hot kettle"},
		},
	}
	for _, eq := range equipments.Equipment {
		client.Equipment.
			Create().
			SetEquipment(eq.Equipment).
			Save(context.Background())
	}
	// Set Facility Data
	facilitys := Facilitys{
		Facility: []Facility{
			Facility{"Swimming pool"},
			Facility{"Breakfast"},
			Facility{"Fitness"},
		},
	}
	for _, fa := range facilitys.Facility {
		client.Facility.
			Create().
			SetFacility(fa.Facility).
			Save(context.Background())
	}

	// Set NearbyPlace Data
	nearbyplaces := NearbyPlaces{
		NearbyPlace: []NearbyPlace{
			NearbyPlace{"Shopping mall"},
			NearbyPlace{"Supermarket"},
			NearbyPlace{"BTS"},
		},
	}
	for _, np := range nearbyplaces.NearbyPlace {
		client.NearbyPlace.
			Create().
			SetPlacename(np.Placename).
			Save(context.Background())
	}

	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"John Marston", "1111@gmail.com", "1111"},
			Employee{"Arthur Morgan", "2222@gmail.com", "2222"},
			Employee{"Dutch Vanderlinde", "3333@gmail.com", "3333"},
		},
	}
	for _, em := range employees.Employee {
		client.Employee.
			Create().
			SetEmployeename(em.Employeename).
			SetEmployeeemail(em.Employeeemail).
			SetPassword(em.Password).
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
