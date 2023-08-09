package controller

import (
	"fmt"

	"github.com/fcVcnte/crud-go/src/configuration/validation"
	"github.com/fcVcnte/crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)
		c.JSON(restError.HttpCode, restError)
		return
	}

	fmt.Println(userRequest)
}
