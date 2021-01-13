package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/equipment"
	"github.com/team15/app/ent/facilitie"
	"github.com/team15/app/ent/nearbyplace"
	"github.com/team15/app/ent/quantity"
	"github.com/team15/app/ent/roomdetail"
	"github.com/team15/app/ent/staytype"
)

type RoomdetailController struct {
	client *ent.Client
	router gin.IRouter
}

type Roomdetail struct {
	Roomtypename string
	Roomprice    string
	Quantity     int
	Staytype     int
	Equipment    int
	Facilitie    int
	Nearbyplace  int
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

	qu, err := ctl.client.Quantity.
		Query().
		Where(quantity.IDEQ(int(obj.Quantity))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "quantity not found",
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

	eq, err := ctl.client.Equipment.
		Query().
		Where(equipment.IDEQ(int(obj.Equipment))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "equipment not found",
		})
		return
	}

	fa, err := ctl.client.Facilitie.
		Query().
		Where(facilitie.IDEQ(int(obj.Facilitie))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "facilitie not found",
		})
		return
	}

	nb, err := ctl.client.Nearbyplace.
		Query().
		Where(nearbyplace.IDEQ(int(obj.Nearbyplace))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "nearbyplace not found",
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
		SetRoomprice(obj.Roomprice).
		SetRoomtypename(obj.Roomtypename).
		SetQuantity(qu).
		SetStaytype(st).
		SetEmployee(em).
		SetEquipments(eq).
		SetFacilities(fa).
		SetNearbyplaces(nb).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, rd)
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
		WithEquipments().
		WithFacilities().
		WithNearbyplaces().
		WithQuantity().
		WithStaytype().
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
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	roomdetails, err := ctl.client.Roomdetail.
		Query().
		WithQuantity().
		WithStaytype().
		WithEquipments().
		WithFacilities().
		WithNearbyplaces().
		WithEmployee().
		Limit(limit).
		Offset(offset).
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
