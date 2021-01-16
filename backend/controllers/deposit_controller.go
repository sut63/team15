package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/statusd"
	"github.com/team15/app/ent/lease"
)

// DepositController defines the struct for the deposit controller
type DepositController struct {
	client *ent.Client
	router gin.IRouter
}

type Deposit struct {
	Added    string
	Info     string
	Employee int
	Statusd  int
	Lease  int
}

// CreateDeposit handles POST requests for adding deposit entities
// @Summary Create deposit
// @Description Create deposit
// @ID create-deposit
// @Accept   json
// @Produce  json
// @Param deposit body Deposit true "Deposit entity"
// @Success 200 {object} Deposit
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /deposits [post]
func (ctl *DepositController) CreateDeposit(c *gin.Context) {
	obj := Deposit{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "deposit binding failed",
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

	st, err := ctl.client.Statusd.
		Query().
		Where(statusd.IDEQ(int(obj.Statusd))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "statusd not found",
		})
		return
	}

	le, err := ctl.client.Lease.
		Query().
		Where(lease.IDEQ(int(obj.Lease))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "lease not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Added)
	ret, err := ctl.client.Deposit.
		Create().
		SetAddedtime(time).
		SetInfo(obj.Info).
		SetEmployee(em).
		SetStatusd(st).
		SetLease(le).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ret)
}

// ListDeposit handles request to get a list of deposit entities
// @Summary List deposit entities
// @Description list deposit entities
// @ID list-deposit
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Deposit
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /deposits [get]
func (ctl *DepositController) ListDeposit(c *gin.Context) {
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

	deposits, err := ctl.client.Deposit.
		Query().
		WithEmployee().
		WithStatusd().
		WithLease().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, deposits)
}

// NewDepositController creates and registers handles for the deposit controller
func NewDepositController(router gin.IRouter, client *ent.Client) *DepositController {
	retc := &DepositController{
		client: client,
		router: router,
	}

	retc.register()

	return retc

}

func (ctl *DepositController) register() {
	deposits := ctl.router.Group("/deposits")

	deposits.POST("", ctl.CreateDeposit)
	deposits.GET("", ctl.ListDeposit)

}
