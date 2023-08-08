package controller

import (
	"fmt"

	rest_error "github.com/fcVcnte/crud-go/src/configuration/err"
	"github.com/fcVcnte/crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := rest_error.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error()))

		c.JSON(restError.HttpCode, restError)
		return
	}

	fmt.Println(userRequest)
}
