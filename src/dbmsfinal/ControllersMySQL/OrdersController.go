package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	daoMySQL "dbmsfinal/DAOMySQL"
	dataMySQL "dbmsfinal/dataMySQL"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type OrdersController struct{}

var OrdersDAO *daoMySQL.OrdersDAO

func (e *OrdersController) GetOrderInfo(c *gin.Context) {
	Order_id, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
	OrderInfo, duration, err := OrdersDAO.GetOrderInfo(Order_id)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":        http.StatusNotFound,
			"message":       err.Error(),
			"data":          nil,
			"duration_time": duration,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "Get Order's info successfully",
		"data":     OrderInfo,
		"duration": duration,
	})
}

func (e *OrdersController) GetAllOrdersInfo(c *gin.Context) {
	Orders, duration, err := OrdersDAO.GetAllOrdersInfo()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "Get all Orders info successfully",
		"data":     Orders,
		"duration": duration,
	})
}

func (r *OrdersController) EditOrder(c *gin.Context) { //TODO: edit Order with optional field
	orderID, _ := strconv.ParseInt(c.Param("Order_id"), 10, 64)

	editOrderData := &dataMySQL.Order{}
	err := c.BindJSON(editOrderData)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	OrderInfo, duration, err := OrdersDAO.EditOrder(orderID, editOrderData)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"message":  err.Error(),
			"data":     nil,
			"duration": duration,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "Edit Order's info successfully",
		"data":     OrderInfo,
		"duration": duration,
	})
	return
}
