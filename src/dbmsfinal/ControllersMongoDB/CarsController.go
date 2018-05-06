package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	daoMongoDB "dbmsfinal/DAOMongoDB"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type CarsController struct{}

var CarsDAO *daoMongoDB.CarsDAO

func (e *CarsController) GetCarInfo(c *gin.Context) {
	Car_id, err := strconv.ParseInt(c.Param("car_id"), 10, 64)
	CarInfo, duration, err := CarsDAO.GetCarInfo(Car_id)

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
		"message":  "Get Car's info successfully",
		"data":     CarInfo,
		"rows":     1,
		"duration": duration,
	})
}

func (e *CarsController) GetAllCarsInfo(c *gin.Context) {
	//_, duration, err := CarsDAO.GetAllCarsInfo()

	rows, duration, err := CarsDAO.GetAllCarsInfo()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			//"data":    nil,
			"rows": 0,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get all Cars info successfully",
		//"data":     Cars,
		"rows":     rows,
		"duration": duration,
	})
}
