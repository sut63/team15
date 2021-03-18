package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/cleanername"
	"github.com/team15/app/ent/cleaningroom"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/lengthtime"
	"github.com/team15/app/ent/roomdetail"
)

type CleaningroomController struct {
	client *ent.Client
	router gin.IRouter
}

type Cleaningroom struct {
	Employee         int
	Cleanername      int
	Numofem          int
	Roomdetail       int
	Lengthtime       int
	Dateandstarttime string
	Phonenumber      string
	Note             string
}

// CreateCleaningroom handles POST requests for adding cleaningroom entities
// @Summary Create cleaningroom
// @Description Create cleaningroom
// @ID create-cleaningroom
// @Accept   json
// @Produce  json
// @Param cleaningroom body Cleaningroom true "Cleaningroom entity"
// @Success 200 {object} ent.Cleaningroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cleaningrooms [post]
func (ctl *CleaningroomController) CreateCleaningroom(c *gin.Context) {
	obj := Cleaningroom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "cleaningroom binding failed",
		})
		return
	}

	em, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(obj.Employee))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "employee not found",
		})
		return
	}

	cn, err := ctl.client.Cleanername.
		Query().
		Where(cleanername.IDEQ(int(obj.Cleanername))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cleanername not found",
		})
		return
	}

	rt, err := ctl.client.Roomdetail.
		Query().
		Where(roomdetail.IDEQ(int(obj.Roomdetail))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "roomdetail not found",
		})
		return
	}

	lt, err := ctl.client.Lengthtime.
		Query().
		Where(lengthtime.IDEQ(int(obj.Lengthtime))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "lengthtime not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Dateandstarttime)
	cr, err := ctl.client.Cleaningroom.
		Create().
		SetEmployee(em).
		SetCleanername(cn).
		SetNumofem(obj.Numofem).
		SetRoomdetail(rt).
		SetLengthtime(lt).
		SetDateandstarttime(time).
		SetPhonenumber(obj.Phonenumber).
		SetNote(obj.Note).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   cr,
	})
}


// GetCleaningroom handles GET requests to retrieve a cleaningroom entity
// @Summary Get a cleaningroom entity by ID
// @Description get cleaningroom by ID
// @ID get-cleaningroom
// @Produce  json
// @Param id path int true "Cleaningroom ID"
// @Success 200 {object} ent.Cleaningroom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cleaningrooms/{id} [get]
func (ctl *CleaningroomController) GetCleaningroom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cr, err := ctl.client.Cleaningroom.
		Query().
		WithEmployee().
		WithCleanername().
		WithLengthtime().
		WithRoomdetail().
		Where(cleaningroom.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cr)
}

// ListCleaningroom handles request to get a list of cleaningroom entities
// @Summary List cleaningroom entities
// @Description list cleaningroom entities
// @ID list-cleaningroom
// @Produce json
// @Success 200 {array} ent.Cleaningroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cleaningrooms [get]
func (ctl *CleaningroomController) ListCleaningroom(c *gin.Context) {
	cleaningrooms, err := ctl.client.Cleaningroom.
		Query().
		WithEmployee().
		WithCleanername().
		WithLengthtime().
		WithRoomdetail().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cleaningrooms)
}

// GetCleaningroomBySearch handles GET requests to retrieve a cleaningroom entity
// @Summary Get a cleaningroom entity by Roomnumber
// @Description get cleaningroom by Roomnumber
// @ID get-cleaningroom-by-roomnumber
// @Produce  json
// @Param roomdetail query string false "Roomnumber Search"
// @Param phonenumber query int false "Phonenumber Search"
// @Success 200 {object} ent.Cleaningroom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /searchcleaningrooms [get]
func (ctl *CleaningroomController) GetCleaningroomBySearch(c *gin.Context) {
	nsearch := c.Query("phonenumber")
	rdsearch, err := strconv.ParseInt(c.Query("roomdetail"), 10, 64)

	roomdetails := ""
	rd, err := ctl.client.Roomdetail.
		Query().
		Where(roomdetail.IDEQ(int(rdsearch))).
		Only(context.Background())

	if rd != nil {
		roomdetails = rd.Roomnumber
	}


	cr, err := ctl.client.Cleaningroom.
		Query().
		WithEmployee().
		WithCleanername().
		WithLengthtime().
		WithRoomdetail().
		Where(cleaningroom.PhonenumberContains(nsearch)).
		Where(cleaningroom.HasRoomdetailWith(roomdetail.RoomnumberContains(roomdetails))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	if nsearch == "" && rdsearch == 0 {
		c.JSON(200, gin.H{
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"data": cr,
	})
}


// NewCleaningroomController creates and registers handles for the cleaningroom controller
func NewCleaningroomController(router gin.IRouter, client *ent.Client) *CleaningroomController {
	crc := &CleaningroomController{
		client: client,
		router: router,
	}
	crc.register()
	return crc
}

func (ctl *CleaningroomController) register() {
	cleaningrooms := ctl.router.Group("/cleaningrooms")
	cleaningroomsearch := ctl.router.Group("/searchcleaningrooms")

	
	cleaningrooms.POST("", ctl.CreateCleaningroom)
	cleaningrooms.GET("", ctl.ListCleaningroom)
	cleaningrooms.GET(":id", ctl.GetCleaningroom)
	cleaningroomsearch.GET("", ctl.GetCleaningroomBySearch)
}
