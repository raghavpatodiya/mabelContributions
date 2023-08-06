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

func CreatePinBoardService(c *gin.Context) (interface{}, error) {
	var pinBoard models.PinBoard

	err := c.BindJSON(&pinBoard)
	if err != nil {
		return nil, err
	}

	if err := validator.New().Struct(pinBoard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}

	username := c.GetString("username")
	if username == "" {
		return nil, fmt.Errorf("invalid username")
	}

	pinBoard.Username = username
	pinBoard.CreatedAt = time.Now()
	pinBoard.LikesCount = 0
	pinBoard.SharesCount = 0

	return config.CreateOne(&pinBoard)
}
