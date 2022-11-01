package controller

import (
	"assignment_2/database"
	"assignment_2/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db database.Database
}

func New(db database.Database) Controller {
	return Controller{
		db: db,
	}
}

func (c Controller) GetOrders(ctx *gin.Context) {
	orders, err := c.db.GetOrders()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "error get orders data",
		})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (c Controller) CreateOrder(ctx *gin.Context) {
	var newOrder model.Order

	err := ctx.BindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "error bind json request",
		})
		return
	}

	orderResult, err := c.db.CreateOrder(newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "error create order",
		})
		return
	}

	ctx.JSON(http.StatusOK, orderResult)
}

func (c Controller) UpdateOrderById(ctx *gin.Context) {
	var newOrder model.Order

	orderId := ctx.Param("orderId")
	err := ctx.BindJSON(&newOrder)

	log.Println("print new order (controller):", newOrder)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "error bind json request",
		})
		return
	}

	id, err := strconv.Atoi(orderId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "error order id not recognized",
		})
		return
	}

	updatedOrder, err := c.db.UpdateOrderById(id, &newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "order id not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedOrder)
}

func (c Controller) DeleteOrderById(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	id, err := strconv.Atoi(orderId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "error order id not recognized",
		})
		return
	}

	deletedId, err := c.db.DeleteOrderById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "order id not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"order_id": deletedId})
}
