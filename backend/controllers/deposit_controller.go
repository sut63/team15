package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/employee"
	"github.com/team15/app/ent/statusd"
	"github.com/team15/app/ent/lease"
	"github.com/team15/app/ent/deposit"
)

// DepositController defines the struct for the deposit controller
type DepositController struct {
	client *ent.Client
	router gin.IRouter
}

type Deposit struct {
	Added    string
	Info     string
	Depositorname     string
	Employee int
	Statusd  int
	Lease  int
	Depositortell     string
	Recipienttell     string
	Parcelcode     string
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
		SetDepositorname(obj.Depositorname).
		SetDepositortell(obj.Depositortell).
		SetRecipienttell(obj.Recipienttell).
		SetParcelcode(obj.Parcelcode).
		SetEmployee(em).
		SetStatusd(st).
		SetLease(le).
		Save(context.Background())
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   ret,
	})
}

// GetDepositBySearch handles GET requests to retrieve a deposit entity
// @Summary Get a deposit entity by Depositid
// @Description get deposit by Depositid
// @ID get-deposit-by-depositid
// @Produce  json
// @Param parcelcode query string false "Parcelcode Search"
// @Param lease query int false "Lease Search"
// @Param statusd query int false "Statusd Search"
// @Success 200 {object} ent.Deposit
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /searchdeposits [get]
func (ctl *DepositController) GetDepositBySearch(c *gin.Context) {
	pcsearch := c.Query("parcelcode")
	lesearch, err := strconv.ParseInt(c.Query("lease"), 10, 64)
	stsearch, err := strconv.ParseInt(c.Query("statusd"), 10, 64)

	leases := ""
	le, err := ctl.client.Lease.
		Query().
		Where(lease.IDEQ(int(lesearch))).
		Only(context.Background())

	if le != nil {
		leases = le.Tenant
	}

	statusds := ""
	st, err := ctl.client.Statusd.
		Query().
		Where(statusd.IDEQ(int(stsearch))).
		Only(context.Background())

	if st != nil {
		statusds = st.Statusdname
	}

	de, err := ctl.client.Deposit.
		Query().
		WithStatusd().
		WithLease().
		Where(deposit.ParcelcodeContains(pcsearch)).
		Where(deposit.HasLeaseWith(lease.TenantContains(leases))).
		Where(deposit.HasStatusdWith(statusd.StatusdnameContains(statusds))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	if lesearch == 0 && pcsearch == "" && stsearch == 0 {
		c.JSON(200, gin.H{
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"data": de,
	})
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
	depositsearch := ctl.router.Group("/searchdeposits")

	deposits.POST("", ctl.CreateDeposit)
	deposits.GET("", ctl.ListDeposit)

	depositsearch.GET("", ctl.GetDepositBySearch)

}
