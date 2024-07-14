package main

import (
	controller "ListMahasiswa/Controller"
	connection "ListMahasiswa/Model/Connection"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := ":1211"
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	connection.ConnectDatabase()

	//###BEGIN WEB API
	// Get data
	r.GET("/api/ListMahasiswa/GetData", controller.GetData)

	// Insert data
	r.POST("/api/ListMahasiswa/InsertData", controller.InsertData)

	// Update data
	r.PUT("/api/ListMahasiswa/UpdateData", controller.UpdateData)

	// Delete data
	r.DELETE("/api/ListMahasiswa/DeleteData", controller.DeleteData)
	//###END WEB API

	r.Run(port)
}
