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

type OrderdetailsController struct{}

var OrderdetailsDAO *daoMongoDB.OrderDetailsDAO

func (e *OrderdetailsController) GetOrderdetailsOfOrderID(c *gin.Context) {
	order_id, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
	OrderdetailsInfo, duration, err := OrderdetailsDAO.GetOrderdetailsOfOrderID(order_id)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":        http.StatusNotFound,
			"message":       err.Error(),
			"data":          nil,
			"rows":          0, //len(OrderdetailsInfo)
			"duration_time": duration,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "Get Orderdetails's info successfully",
		"data":     OrderdetailsInfo,
		"rows":     len(OrderdetailsInfo),
		"duration": duration,
	})
}

func (e *OrderdetailsController) GetOrderCardetailsOfOrderID(c *gin.Context) {
	order_id, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
	Orderdetails, duration, err := OrderdetailsDAO.GetOrderCardetailsOfOrderID(order_id)

	if err != nil {
		fmt.Println(err)
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
		"message":  "Get all Orderdetails info successfully",
		"data":     Orderdetails,
		"rows":     len(Orderdetails),
		"duration": duration,
	})
}

func (e *OrderdetailsController) GetOrderCardetailsOfCustomerID(c *gin.Context) {
	customer_id, err := strconv.ParseInt(c.Param("customer_id"), 10, 64)
	Orderdetails, duration, err := OrderdetailsDAO.GetOrderCardetailsOfCustomerID(customer_id)

	if err != nil {
		fmt.Println(err)
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
		"message":  "Get all Orderdetails info successfully",
		"data":     Orderdetails,
		"rows":     len(Orderdetails),
		"duration": duration,
	})
}

func (r *OrderdetailsController) InsertNewOrderdetail(c *gin.Context) {
	insertOrderdetailsData := &dataMongoDB.OrderdetailCSV{}
	err := c.BindJSON(insertOrderdetailsData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"message":  err.Error(),
			"data":     nil,
			"rows":     0,
			"duration": 0,
		})
		return
	}

	_, duration, err := OrderdetailsDAO.InsertNewOrderdetail(insertOrderdetailsData)

	if err != nil {
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
		"message":  "insert new Orderdetails successfully!",
		"data":     insertOrderdetailsData,
		"rows":     1,
		"duration": duration,
	})
}

func (r *OrderdetailsController) DeleteOrderdetails(c *gin.Context) {
	order_id, _ := strconv.ParseInt(c.Param("order_id"), 10, 64)
	car_id, _ := strconv.ParseInt(c.Query("car_id"), 10, 64)

	fmt.Println(order_id, car_id)
	duration, err := OrderdetailsDAO.DeleteOrderdetail(order_id, car_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"message":  err.Error(),
			"rows":     0,
			"duration": duration,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "Delete Orderdetails successfully",
		"rows":     1,
		"duration": duration,
	})
}
