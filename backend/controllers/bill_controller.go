package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/lease"
	"github.com/team15/app/ent/payment"
	"github.com/team15/app/ent/situation"
)

// BillController defines the struct for the bill controller
type BillController struct {
	client *ent.Client
	router gin.IRouter
}

type Bill struct {
	Added     string
	Room      int
	Situation int
	Payment   int
	Total     string
	Tell      string
	Taxpayer  string
	Lease     int
}

// CreateBill handles POST requests for adding bill entities
// @Summary Create bill
// @Description Create bill
// @ID create-bill
// @Accept   json
// @Produce  json
// @Param bill body Bill true "Bill entity"
// @Success 200 {object} Bill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills [post]
func (ctl *BillController) CreateBill(c *gin.Context) {
	obj := Bill{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bill binding failed",
		})
		return
	}

	si, err := ctl.client.Situation.
		Query().
		Where(situation.IDEQ(int(obj.Situation))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "situation not found",
		})
		return
	}

	p, err := ctl.client.Payment.
		Query().
		Where(payment.IDEQ(int(obj.Payment))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "payment not found",
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
	ret, err := ctl.client.Bill.
		Create().
		SetAddedtime(time).
		SetSituation(si).
		SetTotal(obj.Total).
		SetTaxpayer(obj.Taxpayer).
		SetTell(obj.Tell).
		SetPayment(p).
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

// ListBill handles request to get a list of bill entities
// @Summary List bill entities
// @Description list bill entities
// @ID list-bill
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills [get]
func (ctl *BillController) ListBill(c *gin.Context) {
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

	bills, err := ctl.client.Bill.
		Query().
		WithSituation().
		WithPayment().
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

	c.JSON(200, bills)
}

// NewBillController creates and registers handles for the bill controller
func NewBillController(router gin.IRouter, client *ent.Client) *BillController {
	retc := &BillController{
		client: client,
		router: router,
	}

	retc.register()

	return retc

}

func (ctl *BillController) register() {
	bills := ctl.router.Group("/bills")

	bills.POST("", ctl.CreateBill)
	bills.GET("", ctl.ListBill)

}
