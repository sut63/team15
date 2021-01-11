package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/staytype"
)

type StaytypeController struct {
	client *ent.Client
	router gin.IRouter
}

type Staytype struct {
	Staytype string
}

// CreateStaytype handles POST requests for adding staytype entities
// @Summary Create staytype
// @Description Create staytype
// @ID create-staytype
// @Accept   json
// @Produce  json
// @Param staytype body ent.Staytype true "Staytype entity"
// @Success 200 {object} ent.Staytype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /staytypes [post]
func (ctl *StaytypeController) CreateStaytype(c *gin.Context) {
	obj := Staytype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "staytype binding failed",
		})
		return
	}

	st, err := ctl.client.Staytype.
		Create().
		SetStaytype(obj.Staytype).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, st)
}

// GetStaytype handles GET requests to retrieve a staytype entity
// @Summary Get a staytype entity by ID
// @Description get staytype by ID
// @ID get-staytype
// @Produce  json
// @Param id path int true "Staytype ID"
// @Success 200 {object} ent.Staytype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /staytypes/{id} [get]
func (ctl *StaytypeController) GetStaytype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	st, err := ctl.client.Staytype.
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

// ListStaytype handles request to get a list of staytype entities
// @Summary List staytype entities
// @Description list staytype entities
// @ID list-staytype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Staytype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /staytypes [get]
func (ctl *StaytypeController) ListStaytype(c *gin.Context) {
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

	staytypes, err := ctl.client.Staytype.
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

// DeleteStaytype handles DELETE requests to delete a staytype entity
// @Summary Delete a staytype entity by ID
// @Description get staytype by ID
// @ID delete-staytype
// @Produce  json
// @Param id path int true "Staytype ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /staytype/{id} [delete]
func (ctl *StaytypeController) DeleteStaytype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Staytype.
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

// NewStaytypeController creates and registers handles for the staytype controller
func NewStaytypeController(router gin.IRouter, client *ent.Client) *StaytypeController {
	stc := &StaytypeController{
		client: client,
		router: router,
	}

	stc.register()

	return stc

}

func (ctl *StaytypeController) register() {
	staytypes := ctl.router.Group("/staytypes")

	staytypes.POST("", ctl.CreateStaytype)
	staytypes.GET(":id", ctl.GetStaytype)
	staytypes.GET("", ctl.ListStaytype)
	staytypes.DELETE("", ctl.DeleteStaytype)

}
