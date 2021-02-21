package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/bedtype"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/petrule"
	"github.com/team15/app/ent/pledge"
	"github.com/team15/app/ent/roomdetail"
	"github.com/team15/app/ent/staytype"
)

type RoomdetailController struct {
	client *ent.Client
	router gin.IRouter
}

type Roomdetail struct {
	Roomnumber   string
	Roomtypename string
	Roomprice    int
	Phone        string
	Sleep        int
	Bed          int
	Staytype     int
	Pledge       int
	Petrule      int
	Bedtype      int
	Employee     int
}

// CreateRoomdetail handles POST requests for adding roomdetail entities
// @Summary Create roomdetail
// @Description Create roomdetail
// @ID create-roomdetail
// @Accept   json
// @Produce  json
// @Param roomdetail body Roomdetail true "Roomdetail entity"
// @Success 200 {object} ent.Roomdetail
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomdetails [post]
func (ctl *RoomdetailController) CreateRoomdetail(c *gin.Context) {
	obj := Roomdetail{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "room detail binding failed",
		})
		return
	}

	st, err := ctl.client.Staytype.
		Query().
		Where(staytype.IDEQ(int(obj.Staytype))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "staytype not found",
		})
		return
	}

	p, err := ctl.client.Pledge.
		Query().
		Where(pledge.IDEQ(int(obj.Pledge))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "pledge not found",
		})
		return
	}

	pr, err := ctl.client.Petrule.
		Query().
		Where(petrule.IDEQ(int(obj.Petrule))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "petrule not found",
		})
		return
	}

	bt, err := ctl.client.Bedtype.
		Query().
		Where(bedtype.IDEQ(int(obj.Bedtype))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "bedtype not found",
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

	rd, err := ctl.client.Roomdetail.
		Create().
		SetRoomnumber(obj.Roomnumber).
		SetRoomprice(obj.Roomprice).
		SetRoomtypename(obj.Roomtypename).
		SetPhone(obj.Phone).
		SetSleep(obj.Sleep).
		SetBed(obj.Bed).
		SetBedtype(bt).
		SetPetrule(pr).
		SetPledge(p).
		SetStaytype(st).
		SetEmployee(em).
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
		"data":   rd,
	})
}

// DeleteRoomdetail handles DELETE requests to delete a roomdetail entity
// @Summary Delete a roomdetail entity by ID
// @Description get roomdetail by ID
// @ID delete-roomdetail
// @Produce  json
// @Param id path int true "Roomdetail ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomdetails/{id} [delete]
func (ctl *RoomdetailController) DeleteRoomdetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Roomdetail.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// GetRoomdetail handles GET requests to retrieve a roomdetail entity
// @Summary Get a roomdetail entity by ID
// @Description get roomdetail by ID
// @ID get-roomdetail
// @Produce  json
// @Param id path int true "Roomdetail ID"
// @Success 200 {object} ent.Roomdetail
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomdetails/{id} [get]
func (ctl *RoomdetailController) GetRoomdetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	rd, err := ctl.client.Roomdetail.
		Query().
		WithEmployee().
		WithJobposition().
		WithStaytype().
		WithBedtype().
		WithPetrule().
		WithPledge().
		Where(roomdetail.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, rd)
}

// ListRoomdetail handles request to get a list of roomdetail entities
// @Summary List roomdetail entities
// @Description list roomdetail entities
// @ID list-roomdetail
// @Produce json
// @Success 200 {array} ent.Roomdetail
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomdetails [get]
func (ctl *RoomdetailController) ListRoomdetail(c *gin.Context) {
	roomdetails, err := ctl.client.Roomdetail.
		Query().
		WithBedtype().
		WithPetrule().
		WithPledge().
		WithStaytype().
		WithEmployee().
		WithJobposition().
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, roomdetails)
}

// NewRoomdetailController creates and registers handles for the roomdetail controller
func NewRoomdetailController(router gin.IRouter, client *ent.Client) *RoomdetailController {
	rdc := &RoomdetailController{
		client: client,
		router: router,
	}

	rdc.register()

	return rdc

}

func (ctl *RoomdetailController) register() {
	roomdetails := ctl.router.Group("/roomdetails")

	roomdetails.POST("", ctl.CreateRoomdetail)
	roomdetails.GET("", ctl.ListRoomdetail)
	roomdetails.DELETE(":id", ctl.DeleteRoomdetail)

}
