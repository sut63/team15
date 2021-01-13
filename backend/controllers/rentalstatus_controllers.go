package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/rentalstatus"
)

type RentalstatusController struct {
	client *ent.Client
	router gin.IRouter
}

type Rentalstatus struct {
	Rentalstatus string
}

// CreateRentalstatus handles POST requests for adding rentalstatus entities
// @Summary Create rentalstatus
// @Description Create rentalstatus
// @ID create-rentalstatus
// @Accept   json
// @Produce  json
// @Param rentalstatus body ent.Rentalstatus true "Rentalstatus entity"
// @Success 200 {object} ent.Rentalstatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rentalstatuss [post]
func (ctl *RentalstatusController) CreateRentalstatus(c *gin.Context) {
	obj := Rentalstatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "rentalstatus binding failed",
		})
		return
	}

	rs, err := ctl.client.Rentalstatus.
		Create().
		SetRentalstatus(obj.Rentalstatus).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, rs)
}

// GetRentalstatus handles GET requests to retrieve a rentalstatus entity
// @Summary Get a rentalstatus entity by ID
// @Description get rentalstatus by ID
// @ID get-rentalstatus
// @Produce  json
// @Param id path int true "Rentalstatus ID"
// @Success 200 {object} ent.Rentalstatus
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Rentalstatuss/{id} [get]
func (ctl *RentalstatusController) GetRentalstatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	rs, err := ctl.client.Rentalstatus.
		Query().
		Where(rentalstatus.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, rs)
}

// ListRentalstatus handles request to get a list of rentalstatus entities
// @Summary List rentalstatus entities
// @Description list rentalstatus entities
// @ID list-rentalstatus
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Rentalstatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rentalstatuss [get]
func (ctl *RentalstatusController) ListRentalstatus(c *gin.Context) {
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

	rentalstatuss, err := ctl.client.Rentalstatus.
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

	c.JSON(200, rentalstatuss)
}

// NewRentalstatusController creates and registers handles for the rentalstatus controller
func NewRentalstatusController(router gin.IRouter, client *ent.Client) *RentalstatusController {
	rst := &RentalstatusController{
		client: client,
		router: router,
	}

	rst.register()

	return rst

}

func (ctl *RentalstatusController) register() {
	rentalstatuss := ctl.router.Group("/rentalstatuss")

	rentalstatuss.POST("", ctl.CreateRentalstatus)
	rentalstatuss.GET(":id", ctl.GetRentalstatus)
	rentalstatuss.GET("", ctl.ListRentalstatus)
}
