package main

import (
	"fmt"

	controllersMongoDB "dbmsfinal/ControllersMongoDB"
	controllersMySQL "dbmsfinal/ControllersMySQL"
	daoMongoDB "dbmsfinal/DAOMongoDB"
	daoMySQL "dbmsfinal/DAOMySQL"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, sessionkey")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	router := gin.Default()

	MySQLCarsController := &controllersMySQL.CarsController{}
	MySQLCustomersController := &controllersMySQL.CustomersController{}
	MySQLOrdersController := &controllersMySQL.OrdersController{}
	MySQLOrderdetailsController := &controllersMySQL.OrderdetailsController{}

	MongoDBCarsController := &controllersMongoDB.CarsController{}
	MongoDBCustomersController := &controllersMongoDB.CustomersController{}
	MongoDBOrdersController := &controllersMongoDB.OrdersController{}
	MongoDBOrderdetailsController := &controllersMongoDB.OrderdetailsController{}

	v1 := router.Group("/api/mysql")
	{
		v1.GET("/customers/", MySQLCustomersController.GetAllCustomersInfo)
		v1.GET("/customers/:customer_id", MySQLCustomersController.GetCustomerInfo)
		v1.GET("/cars/", MySQLCarsController.GetCarInfo)
		v1.GET("/cars/:car_id", MySQLCarsController.GetAllCarsInfo)
		v1.GET("/orders/", MySQLOrdersController.GetAllOrdersInfo)
		v1.GET("/orders/:order_id", MySQLOrdersController.GetOrderInfo)
		v1.GET("/orders/:order_id", MySQLOrdersController.GetOrderInfo)
		v1.GET("/orders/:order_id/details", MySQLOrderdetailsController.GetOrderdetailsOfOrderID)
		v1.GET("/orders/:order_id/cars", MySQLOrderdetailsController.GetOrderCardetailsOfOrderID)
		v1.GET("/customers/:customer_id/orders", MySQLOrderdetailsController.GetOrderCardetailsOfCustomerID)
	}

	v2 := router.Group("/api/mongodb")
	{
		v2.GET("/customers/", MongoDBCustomersController.GetAllCustomersInfo)
		v2.GET("/customers/:customer_id", MongoDBCustomersController.GetCustomerInfo)
		v2.GET("/cars/", MongoDBCarsController.GetCarInfo)
		v2.GET("/cars/:car_id", MongoDBCarsController.GetAllCarsInfo)
		v2.GET("/orders/", MongoDBOrdersController.GetAllOrdersInfo)
		v2.GET("/orders/:order_id", MongoDBOrdersController.GetOrderInfo)
		v2.GET("/orders/:order_id", MongoDBOrdersController.GetOrderInfo)
		v2.GET("/orders/:order_id/details", MongoDBOrderdetailsController.GetOrderdetailsOfOrderID)
		v2.GET("/orders/:order_id/cars", MongoDBOrderdetailsController.GetOrderCardetailsOfOrderID)
		v2.GET("/customers/:customer_id/orders", MongoDBOrderdetailsController.GetOrderCardetailsOfCustomerID)
	}

	/*config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("sessionkey")
	config.AddAllowHeaders("Access-Control-Allow-Headers")
	config.AddAllowHeaders("Access-Control-Allow-Origin")
	config.AddAllowHeaders("application/type")
	router.Use(cors.New(config))*/
	router.Use(CORSMiddleware())
	router.Run(":3001")

	defer daoMongoDB.CloseDB()
	defer daoMySQL.CloseDB()
}
