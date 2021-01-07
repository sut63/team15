package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/facility"
)

type FacilityController struct {
	client *ent.Client
	router gin.IRouter
}

type Facility struct {
	Facility string
}

// CreateFacility handles POST requests for adding facility entities
// @Summary Create facility
// @Description Create facility
// @ID create-facility
// @Accept   json
// @Produce  json
// @Param facility body ent.Facility true "Facility entity"
// @Success 200 {object} ent.Facility
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /facilitys [post]
func (ctl *FacilityController) CreateFacility(c *gin.Context) {
	obj := Facility{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "facility binding failed",
		})
		return
	}

	fa, err := ctl.client.Facility.
		Create().
		SetFacility(obj.Facility).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, fa)
}

// GetFacility handles GET requests to retrieve a facility entity
// @Summary Get a facility entity by ID
// @Description get facility by ID
// @ID get-facility
// @Produce  json
// @Param id path int true "Facility ID"
// @Success 200 {object} ent.Facility
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /facilitys/{id} [get]
func (ctl *FacilityController) GetFacility(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	fa, err := ctl.client.Facility.
		Query().
		Where(facility.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, fa)
}

// ListFacility handles request to get a list of facility entities
// @Summary List facility entities
// @Description list facility entities
// @ID list-facility
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Facility
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /facilitys [get]
func (ctl *FacilityController) ListFacility(c *gin.Context) {
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

	facilitys, err := ctl.client.Facility.
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

	c.JSON(200, facilitys)
}

// DeleteFacility handles DELETE requests to delete a facility entity
// @Summary Delete a facility entity by ID
// @Description get facility by ID
// @ID delete-facility
// @Produce  json
// @Param id path int true "Facility ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /facility/{id} [delete]
func (ctl *FacilityController) DeleteFacility(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Facility.
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

// NewFacilityController creates and registers handles for the facility controller
func NewFacilityController(router gin.IRouter, client *ent.Client) *FacilityController {
	fac := &FacilityController{
		client: client,
		router: router,
	}

	fac.register()

	return fac

}

func (ctl *FacilityController) register() {
	facilitys := ctl.router.Group("/facilitys")

	facilitys.POST("", ctl.CreateFacility)
	facilitys.GET(":id", ctl.GetFacility)
	facilitys.GET("", ctl.ListFacility)
	facilitys.DELETE("", ctl.DeleteFacility)

}
