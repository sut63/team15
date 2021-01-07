package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/nearbyplace"
)

type NearbyPlaceController struct {
	client *ent.Client
	router gin.IRouter
}

type NearbyPlace struct {
	Placename string
}

// CreateNearbyPlace handles POST requests for adding nearbyplace entities
// @Summary Create nearbyplace
// @Description Create nearbyplace
// @ID create-nearbyplace
// @Accept   json
// @Produce  json
// @Param nearbyplace body ent.NearbyPlace true "NearbyPlace entity"
// @Success 200 {object} ent.NearbyPlace
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nearbyplaces [post]
func (ctl *NearbyPlaceController) CreateNearbyPlace(c *gin.Context) {
	obj := NearbyPlace{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "place binding failed",
		})
		return
	}

	np, err := ctl.client.NearbyPlace.
		Create().
		SetPlacename(obj.Placename).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, np)
}

// GetNearbyPlace handles GET requests to retrieve a nearbyplace entity
// @Summary Get a nearbyplace entity by ID
// @Description get nearbyplace by ID
// @ID get-nearbyplace
// @Produce  json
// @Param id path int true "NearbyPlace ID"
// @Success 200 {object} ent.NearbyPlace
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nearbyplaces/{id} [get]
func (ctl *NearbyPlaceController) GetNearbyPlace(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	np, err := ctl.client.NearbyPlace.
		Query().
		Where(nearbyplace.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, np)
}

// ListNearbyPlace handles request to get a list of nearbyplace entities
// @Summary List nearbyplace entities
// @Description list nearbyplace entities
// @ID list-nearbyplace
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.NearbyPlace
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nearbyplaces [get]
func (ctl *NearbyPlaceController) ListNearbyPlace(c *gin.Context) {
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

	nearbyplaces, err := ctl.client.NearbyPlace.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, nearbyplaces)
}

// DeleteNearbyPlace handles DELETE requests to delete a nearbyplace entity
// @Summary Delete a nearbyplace entity by ID
// @Description get nearbyplace by ID
// @ID delete-nearbyplace
// @Produce  json
// @Param id path int true "NearbyPlace ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nearbyplace/{id} [delete]
func (ctl *NearbyPlaceController) DeleteNearbyPlace(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.NearbyPlace.
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

// NewNearbyPlaceController creates and registers handles for the nearbyplace controller
func NewNearbyPlaceController(router gin.IRouter, client *ent.Client) *NearbyPlaceController {
	npc := &NearbyPlaceController{
		client: client,
		router: router,
	}

	npc.register()

	return npc

}

func (ctl *NearbyPlaceController) register() {
	nearbyplaces := ctl.router.Group("/nearbyplaces")

	nearbyplaces.POST("", ctl.CreateNearbyPlace)
	nearbyplaces.GET(":id", ctl.GetNearbyPlace)
	nearbyplaces.GET("", ctl.ListNearbyPlace)
	nearbyplaces.DELETE("", ctl.DeleteNearbyPlace)

}
