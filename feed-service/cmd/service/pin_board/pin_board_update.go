package service

import (
	"fmt"
	"net/http"
	"super-feed-service/cmd/config"
	models "super-feed-service/cmd/models/pin_board"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UpdatePinBoardService(c *gin.Context) (interface{}, error) {
	var updatePinBoardRequest models.PinBoard
	if err := c.BindJSON(&updatePinBoardRequest); err != nil {
		return nil, err
	}

	if err := validator.New().Struct(updatePinBoardRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}

	username := c.GetString("username")
	if username == "" {
		return nil, fmt.Errorf("invalid username")
	}

	res, err := config.GetOne(&models.PinBoard{
		BoardID:  updatePinBoardRequest.BoardID,
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	res.BoardName = updatePinBoardRequest.BoardName
	res.IsPublic = updatePinBoardRequest.IsPublic

	res.IsUpdated = true
	res.UpdatedAt = time.Now()

	return config.UpdateOne(&res)
}
