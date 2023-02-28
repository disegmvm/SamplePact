package main

import (
	"github.com/gin-gonic/gin"
	"transPact/pkg/server"
)

func main() {
	router := gin.Default()
	router.GET("/sample/endpoint/to/get/transformed/sku", server.GetUserByID)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
