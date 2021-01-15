package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/roomdetail"
	"github.com/team15/app/ent/wifi"
)

// LeaseController defines the struct for the lease controller
type LeaseController struct {
	client *ent.Client
	router gin.IRouter
}

type Lease struct {
	Added      string
	Tenant     string
	Employee   int
	Roomdetail int
	Wifi       int
}

// CreateLease handles POST requests for adding lease entities
// @Summary Create lease
// @Description Create lease
// @ID create-lease
// @Accept   json
// @Produce  json
// @Param lease body Lease true "Lease entity"
// @Success 200 {object} Lease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /leases [post]
func (ctl *LeaseController) CreateLease(c *gin.Context) {
	obj := Lease{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "lease binding failed",
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

	rdc, err := ctl.client.Roomdetail.
		Query().
		Where(roomdetail.IDEQ(int(obj.Roomdetail))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "roomnumber not found",
		})
		return
	}

	wf, err := ctl.client.Wifi.
		Query().
		Where(wifi.IDEQ(int(obj.Wifi))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "wifi not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Added)
	lea, err := ctl.client.Lease.
		Create().
		SetAddedtime(time).
		SetTenant(obj.Tenant).
		SetRoomdetail(rdc).
		SetWifi(wf).
		SetEmployee(em).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, lea)
}

// ListLease handles request to get a list of lease entities
// @Summary List lease entities
// @Description list lease entities
// @ID list-lease
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Lease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /leases [get]
func (ctl *LeaseController) ListLease(c *gin.Context) {
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

	leases, err := ctl.client.Lease.
		Query().
		WithRoomdetail().
		WithWifi().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, leases)
}

// NewLeaseController creates and registers handles for the lease controller
func NewLeaseController(router gin.IRouter, client *ent.Client) *LeaseController {
	leas := &LeaseController{
		client: client,
		router: router,
	}

	leas.register()

	return leas

}

func (ctl *LeaseController) register() {
	leases := ctl.router.Group("/leases")

	leases.POST("", ctl.CreateLease)
	leases.GET("", ctl.ListLease)

}
