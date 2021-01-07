package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/staytype"
)

type StayTypeController struct {
	client *ent.Client
	router gin.IRouter
}

type StayType struct {
	StayType string
}

// CreateStayType handles POST requests for adding staytype entities
// @Summary Create staytype
// @Description Create staytype
// @ID create-staytype
// @Accept   json
// @Produce  json
// @Param staytype body ent.StayType true "StayType entity"
// @Success 200 {object} ent.StayType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /staytypes [post]
func (ctl *StayTypeController) CreateStayType(c *gin.Context) {
	obj := StayType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "stay type binding failed",
		})
		return
	}

	st, err := ctl.client.StayType.
		Create().
		SetStaytype(obj.StayType).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, st)
}

// GetStayType handles GET requests to retrieve a staytype entity
// @Summary Get a staytype entity by ID
// @Description get staytype by ID
// @ID get-staytype
// @Produce  json
// @Param id path int true "StayType ID"
// @Success 200 {object} ent.StayType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /staytypes/{id} [get]
func (ctl *StayTypeController) GetStayType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	st, err := ctl.client.StayType.
		Query().
		Where(staytype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, st)
}

// ListStayType handles request to get a list of staytype entities
// @Summary List staytype entities
// @Description list staytype entities
// @ID list-staytype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.StayType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /staytypes [get]
func (ctl *StayTypeController) ListStayType(c *gin.Context) {
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

	staytypes, err := ctl.client.StayType.
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

	c.JSON(200, staytypes)
}

// DeleteStayType handles DELETE requests to delete a staytype entity
// @Summary Delete a staytype entity by ID
// @Description get staytype by ID
// @ID delete-staytype
// @Produce  json
// @Param id path int true "StayType ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /staytype/{id} [delete]
func (ctl *StayTypeController) DeleteStayType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.StayType.
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

// NewStayTypeController creates and registers handles for the staytype controller
func NewStayTypeController(router gin.IRouter, client *ent.Client) *StayTypeController {
	stc := &StayTypeController{
		client: client,
		router: router,
	}

	stc.register()

	return stc

}

func (ctl *StayTypeController) register() {
	staytypes := ctl.router.Group("/staytypes")

	staytypes.POST("", ctl.CreateStayType)
	staytypes.GET(":id", ctl.GetStayType)
	staytypes.GET("", ctl.ListStayType)
	staytypes.DELETE("", ctl.DeleteStayType)

}
