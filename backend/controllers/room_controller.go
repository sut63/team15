package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/equipment"
	"github.com/team15/app/ent/facility"
	"github.com/team15/app/ent/nearbyplace"
	"github.com/team15/app/ent/quantity"
	"github.com/team15/app/ent/staytype"
)

type RoomController struct {
	client *ent.Client
	router gin.IRouter
}

type Room struct {
	Roomprice    int
	Roomtypename string
	Quantity     int
	StayType     int
	Equipment    int
	Facility     int
	NearbyPlace  int
}

// CreateRoom handles POST requests for adding room entities
// @Summary Create room
// @Description Create room
// @ID create-room
// @Accept   json
// @Produce  json
// @Param room body Room true "Room entity"
// @Success 200 {object} Room
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rooms [post]
func (ctl *RoomController) CreateRoom(c *gin.Context) {
	obj := Room{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "room binding failed",
		})
		return
	}

	q, err := ctl.client.Quantity.
		Query().
		Where(quantity.IDEQ(int(obj.Quantity))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "quantity not found",
		})
		return
	}

	st, err := ctl.client.StayType.
		Query().
		Where(staytype.IDEQ(int(obj.StayType))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "stay type not found",
		})
		return
	}

	eq, err := ctl.client.Equipment.
		Query().
		Where(equipment.IDEQ(int(obj.Equipment))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "equipment in room not found",
		})
		return
	}

	fa, err := ctl.client.Facility.
		Query().
		Where(facility.IDEQ(int(obj.Facility))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "facility not found",
		})
		return
	}

	np, err := ctl.client.NearbyPlace.
		Query().
		Where(nearbyplace.IDEQ(int(obj.NearbyPlace))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "nearby place not found",
		})
		return
	}

	r, err := ctl.client.Room.
		Create().
		SetRoomprice(obj.Roomprice).
		SetRoomtypename(obj.Roomtypename).
		SetQuantity(q).
		SetStaytype(st).
		AddEquipments(eq).
		AddFacilities(fa).
		AddNearbyplace(np).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, r)
}

// DeleteRoom handles DELETE requests to delete a room entity
// @Summary Delete a room entity by ID
// @Description get room by ID
// @ID delete-room
// @Produce  json
// @Param id path int true "Room ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rooms/{id} [delete]
func (ctl *RoomController) DeleteRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Room.
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

// ListRoom handles request to get a list of room entities
// @Summary List room entities
// @Description list room entities
// @ID list-room
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Room
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rooms [get]
func (ctl *RoomController) ListRoom(c *gin.Context) {

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	rooms, err := ctl.client.Room.
		Query().
		WithQuantity().
		WithStaytype().
		WithEquipments().
		WithFacilities().
		WithNearbyplace().
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, rooms)
}

// NewRoomController creates and registers handles for the room controller
func NewRoomController(router gin.IRouter, client *ent.Client) *RoomController {
	rc := &RoomController{
		client: client,
		router: router,
	}

	rc.register()

	return rc

}

func (ctl *RoomController) register() {
	rooms := ctl.router.Group("/rooms")

	rooms.POST("", ctl.CreateRoom)
	rooms.GET("", ctl.ListRoom)
	rooms.DELETE(":id", ctl.DeleteRoom)

}
