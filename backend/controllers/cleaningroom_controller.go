package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/cleanername"
	"github.com/team15/app/ent/cleaningroom"
	"github.com/team15/app/ent/lengthtime"
	"github.com/team15/app/ent/room"
)

// CleaningRoomController defines the struct for the cleaningroom controller
type CleaningRoomController struct {
	client *ent.Client
	router gin.IRouter
}

type CleaningRoom struct {
	Dateandstarttime string
	Note             string
	Room             int
	CleanerName      int
	LengthTime       int
}

// CreateCleaningRoom handles POST requests for adding cleaningroom entities
// @Summary Create cleaningroom
// @Description Create cleaningroom
// @ID create-cleaningroom
// @Accept   json
// @Produce  json
// @Param cleaningroom body ent.CleaningRoom true "CleaningRoom entity"
// @Success 200 {object} ent.CleaningRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cleaningrooms [post]
func (ctl *CleaningRoomController) CreateCleaningRoom(c *gin.Context) {
	obj := CleaningRoom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "cleaningroom binding failed",
		})
		return
	}

	r, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(obj.Room))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "room not found",
		})
		return
	}

	cn, err := ctl.client.CleanerName.
		Query().
		Where(cleanername.IDEQ(int(obj.CleanerName))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cleanername not found",
		})
		return
	}

	lt, err := ctl.client.LengthTime.
		Query().
		Where(lengthtime.IDEQ(int(obj.LengthTime))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "lengthtime not found",
		})
		return
	}
	time, err := time.Parse(time.RFC3339, obj.Dateandstarttime)
	cr, err := ctl.client.CleaningRoom.
		Create().
		SetDateandstarttime(time).
		SetNote(obj.Note).
		SetRoom(r).
		SetCleanerName(cn).
		SetLengthTime(lt).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   cr,
	})
}

// ListCleaningRoom handles request to get a list of cleaningroom entities
// @Summary List cleaningroom entities
// @Description list cleaningroom entities
// @ID list-cleaningroom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.CleaningRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cleaningrooms [get]
func (ctl *CleaningRoomController) ListCleaningRoom(c *gin.Context) {
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

	cleaningrooms, err := ctl.client.CleaningRoom.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cleaningrooms)
}

// GetCleaningRoom handles GET requests to retrieve a cleaningroom entity
// @Summary Get a cleaningroom entity by ID
// @Description get cleaningroom by ID
// @ID get-cleaningroom
// @Produce  json
// @Param id path int true "CleaningRoom ID"
// @Success 200 {object} ent.CleaningRoom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cleaningrooms/{id} [get]
func (ctl *CleaningRoomController) GetCleaningRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.CleaningRoom.
		Query().
		Where(cleaningroom.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// NewCleaningRoomController creates and registers handles for the user controller
func NewCleaningRoomController(router gin.IRouter, client *ent.Client) *CleaningRoomController {
	uc := &CleaningRoomController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitCleaningRoomController registers routes to the main engine
 func (ctl *CleaningRoomController) register() {
	cleaningrooms := ctl.router.Group("/cleaningrooms")
  
	cleaningrooms.GET("", ctl.ListCleaningRoom)
  
	// CRUD
	cleaningrooms.POST("", ctl.CreateCleaningRoom)
	cleaningrooms.GET(":id", ctl.GetCleaningRoom)
 }
 
