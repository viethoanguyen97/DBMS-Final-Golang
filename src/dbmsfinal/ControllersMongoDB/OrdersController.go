package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	daoMongoDB "dbmsfinal/DAOMongoDB"
	dataMongoDB "dbmsfinal/dataMongoDB"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type OrdersController struct{}

var OrdersDAO *daoMongoDB.OrdersDAO

func (e *OrdersController) GetOrderInfo(c *gin.Context) {
	Order_id, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
	OrderInfo, duration, err := OrdersDAO.GetOrderInfo(Order_id)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":        http.StatusNotFound,
			"message":       err.Error(),
			"data":          nil,
			"rows":          0,
			"duration_time": duration,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "Get Order's info successfully",
		"data":     OrderInfo,
		"rows":     1,
		"duration": duration,
	})
}

func (e *OrdersController) GetAllOrdersInfo(c *gin.Context) {
	//_, duration, err := OrdersDAO.GetAllOrdersInfo()
	rows, duration, err := OrdersDAO.GetAllOrdersInfo()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			//"data":    nil,
			"rows":     0,
			"duration": duration,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get all Orders info successfully",
		//"data":     Orders,
		"rows":     rows,
		"duration": duration,
	})
}

func (r *OrdersController) EditOrder(c *gin.Context) { //TODO: edit Order with optional field
	orderID, _ := strconv.ParseInt(c.Param("order_id"), 10, 64)

	editOrderData := &dataMongoDB.Order{}
	err := c.BindJSON(editOrderData)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"rows":    0,
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
			"rows":     0,
			"duration": duration,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "Edit Order's info successfully",
		"data":     OrderInfo,
		"rows":     1,
		"duration": duration,
	})
	return
}
