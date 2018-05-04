package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	daoMySQL "dbmsfinal/DAOMySQL"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type CarsController struct{}

var CarsDAO *daoMySQL.CarsDAO

func (e *CarsController) GetCarInfo(c *gin.Context) {
	Car_id, err := strconv.ParseInt(c.Param("car_id"), 10, 64)
	CarInfo, duration, err := CarsDAO.GetCarInfo(Car_id)

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
		"message":  "Get Car's info successfully",
		"data":     CarInfo,
		"duration": duration,
	})
}

func (e *CarsController) GetAllCarsInfo(c *gin.Context) {
	_, duration, err := CarsDAO.GetAllCarsInfo()

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
		"status":  http.StatusOK,
		"message": "Get all Cars info successfully",
		//"data":     Cars,
		"duration": duration,
	})
}
