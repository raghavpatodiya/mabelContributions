package service

import (
	"fmt"
	"strconv"
	"super-feed-service/cmd/config"
	models "super-feed-service/cmd/models/pin_board"

	"github.com/gin-gonic/gin"
)

func GetPinBoardsByUserService(c *gin.Context) (interface{}, error) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		return nil, err
	}

	username := c.GetString("username")
	if username == "" {
		return nil, fmt.Errorf("invalid username")
	}

	return config.GetAll(&models.PinBoard{
		Username: username,
	}, page, pinBoardsPerPage, "")
}
