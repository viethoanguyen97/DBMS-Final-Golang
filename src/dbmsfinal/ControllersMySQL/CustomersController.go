package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	daoMySQL "dbmsfinal/DAOMySQL"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type CustomersController struct{}

var customersDAO *daoMySQL.CustomersDAO

func (e *CustomersController) GetCustomerInfo(c *gin.Context) {
	customer_id, err := strconv.ParseInt(c.Param("customer_id"), 10, 64)
	customerInfo, duration, err := customersDAO.GetCustomerInfo(customer_id)

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
		"message":  "Get customer's info successfully",
		"data":     customerInfo,
		"rows":     1,
		"duration": duration,
	})
}

func (e *CustomersController) GetAllCustomersInfo(c *gin.Context) {
	//_, duration, err := customersDAO.GetAllCustomersInfo()
	rows, duration, err := customersDAO.GetAllCustomersInfo()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			//"data":    nil,
			"rows":     rows,
			"duration": duration,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get all Customers info successfully",
		//"data":     customers,
		"rows":     rows,
		"duration": duration,
	})
}
