package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/rentalstatus"
)

// RepairinvoiceController defines the struct for the repairinvoice controller
type RepairinvoiceController struct {
	client *ent.Client
	router gin.IRouter
}

type Repairinvoice struct {
	Bequipment   string
	Employee      int
	Rentalstatus  int
}

// CreateRepairinvoice handles POST requests for adding Repairinvoice entities
// @Summary Create repairinvoice
// @Description Create repairinvoice
// @ID create-repairinvoice
// @Accept   json
// @Produce  json
// @Param repairinvoice body Repairinvoice true "Repairinvoice entity"
// @Success 200 {object} Repairinvoice
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /repairinvoices [post]
func (ctl *RepairinvoiceController) CreateRepairinvoice(c *gin.Context) {
	obj := Repairinvoice{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "repairinvoice binding failed",
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

	rs, err := ctl.client.Rentalstatus.
		Query().
		Where(rentalstatus.IDEQ(int(obj.Rentalstatus))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "rentalstatus not found",
		})
		return
	}
	ret, err := ctl.client.Repairinvoice.
		Create().
		SetBequipment(obj.Bequipment).
		SetEmployee(em).
		SetRentalstatus(rs).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ret)
}

// ListRepairinvoice handles request to get a list of repairinvoicet entities
// @Summary List repairinvoice entities
// @Description list repairinvoice entities
// @ID list-repairinvoice
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Repairinvoice
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /repairinvoices [get]
func (ctl *RepairinvoiceController) ListRepairinvoice(c *gin.Context) {
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

	repairinvoices, err := ctl.client.Repairinvoice.
		Query().
		WithEmployee().
		WithRentalstatus().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, repairinvoices)
}

// NewRepairinvoiceController creates and registers handles for the repairinvoice controller
func NewRepairinvoiceController(router gin.IRouter, client *ent.Client) *RepairinvoiceController {
	rpv := &RepairinvoiceController{
		client: client,
		router: router,
	}

	rpv.register()

	return rpv

}

func (ctl *RepairinvoiceController) register() {
	repairinvoices := ctl.router.Group("/repairinvoices")

	repairinvoices.POST("", ctl.CreateRepairinvoice)
	repairinvoices.GET("", ctl.ListRepairinvoice)

}
