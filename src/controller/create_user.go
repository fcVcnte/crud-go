package controller

import (
	"net/http"

	"github.com/fcVcnte/crud-go/src/configuration/logger"
	"github.com/fcVcnte/crud-go/src/configuration/validation"
	"github.com/fcVcnte/crud-go/src/controller/model/request"
	"github.com/fcVcnte/crud-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		restError := validation.ValidateUserError(err)

		c.JSON(restError.HttpCode, restError)
		return
	}

	response := response.UserResponse{
		Id:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))
	c.JSON(http.StatusOK, response)
}
