package service

import (
	"fmt"
	"super-feed-service/cmd/config"
	models "super-feed-service/cmd/models/pin_board"
	"time"

	"github.com/gin-gonic/gin"
)

func DeletePinBoardService(c *gin.Context) (interface{}, error) {
	res, err := GetPinBoardService(c)
	if err != nil {
		return nil, err
	}

	pinBoard, ok := res.(*models.PinBoard)
	if !ok {
		return nil, fmt.Errorf("internal error")
	}

	if pinBoard.IsDeleted {
		return nil, fmt.Errorf("board not found")
	}

	t := time.Now()
	pinBoard.IsDeleted = true
	pinBoard.DeletedAt = &t

	return config.UpdateOne(&pinBoard)
}
