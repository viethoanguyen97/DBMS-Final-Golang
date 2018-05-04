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
			"duration_time": duration,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "Get customer's info successfully",
		"data":     customerInfo,
		"duration": duration,
	})
}

func (e *CustomersController) GetAllCustomersInfo(c *gin.Context) {
	_, duration, err := customersDAO.GetAllCustomersInfo()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			//"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get all Customers info successfully",
		//"data":     customers,
		"duration": duration,
	})
}

/*
func (r *CustomersController) AddNewCustomer(c *gin.Context) {
	addCustomerData := &dataService.AddCustomer{}
	err := c.BindJSON(addCustomerData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	CustomerInfo, err := CustomersdaoMySQLAddNewCustomer(addCustomerData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Add new Customer successfully!",
		"data":    CustomerInfo,
	})
}

func (r *CustomersController) EditCustomer(c *gin.Context) { //TODO: edit Customer with optional field
	CustomerID, _ := strconv.ParseInt(c.Param("Customer_id"), 10, 64)

	editCustomerData := &dataService.EditCustomer{}
	err := c.BindJSON(editCustomerData)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	CustomerInfo, err := CustomersdaoMySQLEditCustomer(CustomerID, editCustomerData)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Edit Customer's info successfully",
		"data":    CustomerInfo,
	})
	return
}

func (r *CustomersController) BorrowCustomer(c *gin.Context) {
	CustomerID, _ := strconv.ParseInt(c.Param("Customer_id"), 10, 64)

	CustomerInfo, err := CustomersdaoMySQLBorrowCustomer(CustomerID)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  http.StatusForbidden,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Borrow Customer successfully",
		"data":    CustomerInfo,
	})

	return
}

func (r *CustomersController) DeleteCustomer(c *gin.Context) {
	CustomerID, _ := strconv.ParseInt(c.Param("Customer_id"), 10, 64)

	err := CustomersdaoMySQLDeleteCustomer(CustomerID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Delete Customer successfully",
	})
}
*/
