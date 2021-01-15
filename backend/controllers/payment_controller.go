package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
    "github.com/team15/app/ent/payment"
)

type PaymentController struct {
	client *ent.Client
	router gin.IRouter
}

type Payment struct {
	Paymentname string
}

// CreatePayment handles POST requests for adding payment entities
// @Summary Create payment
// @Description Create payment
// @ID create-payment
// @Accept   json
// @Produce  json
// @Param payment body ent.Payment true "Payment entity"
// @Success 200 {object} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payments [post]
func (ctl *PaymentController) CreatePayment(c *gin.Context) {
	obj := Payment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "payment binding failed",
		})
		return
	}

	p, err := ctl.client.Payment.
		Create().
		SetPaymentname(obj.Paymentname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200,p)
}

// GetPayment handles GET requests to retrieve a payment entity
// @Summary Get a payment entity by ID
// @Description get payment by ID
// @ID get-payment
// @Produce  json
// @Param id path int true "Payment ID"
// @Success 200 {object} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payments/{id} [get]
func (ctl *PaymentController) GetPayment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	p, err := ctl.client.Payment.
		Query().
		Where(payment.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// ListPayment handles request to get a list of payment entities
// @Summary List payment entities
// @Description list payment entities
// @ID list-payment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payments [get]
func (ctl *PaymentController) ListPayment(c *gin.Context) {
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

	payments, err := ctl.client.Payment.
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

	c.JSON(200, payments)
}

// DeletePayment handles DELETE requests to delete a payment entity
// @Summary Delete a payment entity by ID
// @Description get payment by ID
// @ID delete-payment
// @Produce  json
// @Param id path int true "Payment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payment/{id} [delete]
func (ctl *PaymentController) DeletePayment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Payment.
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

// NewPaymentController creates and registers handles for the payment controller
func NewPaymentController(router gin.IRouter, client *ent.Client) *PaymentController {
	pc := &PaymentController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *PaymentController) register() {
	payments := ctl.router.Group("/payments")

	payments.POST("", ctl.CreatePayment)
	payments.GET(":id", ctl.GetPayment)
	payments.GET("", ctl.ListPayment)
	payments.DELETE("", ctl.DeletePayment)

}