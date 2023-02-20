package server

import (
	"github.com/gin-gonic/gin"
	"transPact/pkg/server"
)

func main() {
	router := gin.Default()
	router.GET("/users/:userId", server.GetUserByID)
	router.Run(":8080")
}
