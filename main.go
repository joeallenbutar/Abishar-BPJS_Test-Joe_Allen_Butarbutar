// Author : Joe Allen Butarbutar

package main

import (
	"Abishar-BPJS_Test-Joe_Allen_Butarbutar/controllers"
	"Abishar-BPJS_Test-Joe_Allen_Butarbutar/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("initialize function.")
	db.InitializeDB()
}

func main() {
	route := gin.Default()
	orderRoute := route.Group("/api")
	{
		orderRoute.POST("/create", controllers.CreateOrder)
		orderRoute.GET("/get", controllers.GetOrder)
	}
	fmt.Println("Server run at port 8080 (http://localhost:8080).")
	route.Run(":8080")
}
