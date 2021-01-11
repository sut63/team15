package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/facilitie"
)

type FacilitieController struct {
	client *ent.Client
	router gin.IRouter
}

type Facilitie struct {
	Facilitie string
}

// CreateFacilitie handles POST requests for adding facilitie entities
// @Summary Create facilitie
// @Description Create facilitie
// @ID create-facilitie
// @Accept   json
// @Produce  json
// @Param facilitie body ent.Facilitie true "Facilitie entity"
// @Success 200 {object} ent.Facilitie
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /facilities [post]
func (ctl *FacilitieController) CreateFacilitie(c *gin.Context) {
	obj := Facilitie{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "facilitie binding failed",
		})
		return
	}

	fa, err := ctl.client.Facilitie.
		Create().
		SetFacilitie(obj.Facilitie).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, fa)
}

// GetFacilitie handles GET requests to retrieve a facilitie entity
// @Summary Get a facilitie entity by ID
// @Description get facilitie by ID
// @ID get-facilitie
// @Produce  json
// @Param id path int true "Facilitie ID"
// @Success 200 {object} ent.Facilitie
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /facilities/{id} [get]
func (ctl *FacilitieController) GetFacilitie(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	fa, err := ctl.client.Facilitie.
		Query().
		Where(facilitie.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, fa)
}

// ListFacilitie handles request to get a list of facilitie entities
// @Summary List facilitie entities
// @Description list facilitie entities
// @ID list-facilitie
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Facilitie
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /facilities [get]
func (ctl *FacilitieController) ListFacilitie(c *gin.Context) {
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

	facilities, err := ctl.client.Facilitie.
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

	c.JSON(200, facilities)
}

// DeleteFacilitie handles DELETE requests to delete a facilitie entity
// @Summary Delete a facilitie entity by ID
// @Description get facilitie by ID
// @ID delete-facilitie
// @Produce  json
// @Param id path int true "Facilitie ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /facilitie/{id} [delete]
func (ctl *FacilitieController) DeleteFacilitie(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Facilitie.
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

// NewFacilitieController creates and registers handles for the facilitie controller
func NewFacilitieController(router gin.IRouter, client *ent.Client) *FacilitieController {
	fac := &FacilitieController{
		client: client,
		router: router,
	}

	fac.register()

	return fac

}

func (ctl *FacilitieController) register() {
	facilities := ctl.router.Group("/facilities")

	facilities.POST("", ctl.CreateFacilitie)
	facilities.GET(":id", ctl.GetFacilitie)
	facilities.GET("", ctl.ListFacilitie)
	facilities.DELETE("", ctl.DeleteFacilitie)

}
