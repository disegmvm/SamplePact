package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type User2 struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Title     string `json:"title"`
}

func GetUserByID(ctx *gin.Context) {
	id := ctx.Param("userId")

	ctx.JSON(http.StatusOK, User2{
		ID:        id,
		FirstName: "Default first name",
		LastName:  "Default last name",
		Title:     "River Island",
	})
}
