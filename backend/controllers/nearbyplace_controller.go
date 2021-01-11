package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/nearbyplace"
)

type NearbyplaceController struct {
	client *ent.Client
	router gin.IRouter
}

type Nearbyplace struct {
	Nearbyplace string
}

// CreateNearbyplace handles POST requests for adding nearbyplace entities
// @Summary Create nearbyplace
// @Description Create nearbyplace
// @ID create-nearbyplace
// @Accept   json
// @Produce  json
// @Param nearbyplace body ent.Nearbyplace true "Nearbyplace entity"
// @Success 200 {object} ent.Nearbyplace
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nearbyplaces [post]
func (ctl *NearbyplaceController) CreateNearbyplace(c *gin.Context) {
	obj := Nearbyplace{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "nearbyplace binding failed",
		})
		return
	}

	st, err := ctl.client.Nearbyplace.
		Create().
		SetNearbyplace(obj.Nearbyplace).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, st)
}

// GetNearbyplace handles GET requests to retrieve a nearbyplace entity
// @Summary Get a nearbyplace entity by ID
// @Description get nearbyplace by ID
// @ID get-nearbyplace
// @Produce  json
// @Param id path int true "Nearbyplace ID"
// @Success 200 {object} ent.Nearbyplace
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nearbyplaces/{id} [get]
func (ctl *NearbyplaceController) GetNearbyplace(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	st, err := ctl.client.Nearbyplace.
		Query().
		Where(nearbyplace.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, st)
}

// ListNearbyplace handles request to get a list of nearbyplace entities
// @Summary List nearbyplace entities
// @Description list nearbyplace entities
// @ID list-nearbyplace
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Nearbyplace
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nearbyplaces [get]
func (ctl *NearbyplaceController) ListNearbyplace(c *gin.Context) {
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

	nearbyplaces, err := ctl.client.Nearbyplace.
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

// DeleteNearbyplace handles DELETE requests to delete a nearbyplace entity
// @Summary Delete a nearbyplace entity by ID
// @Description get nearbyplace by ID
// @ID delete-nearbyplace
// @Produce  json
// @Param id path int true "Nearbyplace ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nearbyplace/{id} [delete]
func (ctl *NearbyplaceController) DeleteNearbyplace(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Nearbyplace.
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

// NewNearbyplaceController creates and registers handles for the nearbyplace controller
func NewNearbyplaceController(router gin.IRouter, client *ent.Client) *NearbyplaceController {
	stc := &NearbyplaceController{
		client: client,
		router: router,
	}

	stc.register()

	return stc

}

func (ctl *NearbyplaceController) register() {
	nearbyplaces := ctl.router.Group("/nearbyplaces")

	nearbyplaces.POST("", ctl.CreateNearbyplace)
	nearbyplaces.GET(":id", ctl.GetNearbyplace)
	nearbyplaces.GET("", ctl.ListNearbyplace)
	nearbyplaces.DELETE("", ctl.DeleteNearbyplace)

}
