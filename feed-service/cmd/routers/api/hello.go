package api

import (
	"net/http"
	hello_service "super-feed-service/cmd/service/hello"
	"super-feed-service/cmd/utils"

	"github.com/gin-gonic/gin"
)

// Hello World Its Me  ... Get Greeting
// @Summary     Greets with Hello World
// @Description get if server is up
// @Tags        Hello World
// @Success     200
// @Failure     404
// @Router      /hello/hello_world [get]
// @Security    BearerAuth
func GetHelloWorld(c *gin.Context) {
	res, err := hello_service.GetHelloWorld()
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}
