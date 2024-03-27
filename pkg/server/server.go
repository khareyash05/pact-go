package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func GetUserByID(ctx *gin.Context) {
	id := ctx.Param("userId")

	ctx.JSON(http.StatusOK, User{
		ID:        id,
		FirstName: fmt.Sprintf("first%s", id),
		LastName:  fmt.Sprintf("last%s", id),
	})
}
