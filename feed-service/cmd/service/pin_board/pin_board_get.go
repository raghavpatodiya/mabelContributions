package service

import (
	"fmt"
	"super-feed-service/cmd/config"
	models "super-feed-service/cmd/models/pin_board"

	"github.com/gin-gonic/gin"
)

func GetPinBoardService(c *gin.Context) (interface{}, error) {
	boardId := c.Param("boardId")
	if boardId == "" {
		return nil, fmt.Errorf("invalid board id")
	}

	pinBoard, err := config.GetOne(&models.PinBoard{
		BoardID: boardId,
	})
	if err != nil {
		return nil, err
	}

	username := c.GetString("username")
	if username == "" {
		return nil, fmt.Errorf("invalid username")
	}

	if pinBoard.IsDeleted {
		return nil, fmt.Errorf("board not found")
	}

	if !pinBoard.IsPublic {
		if pinBoard.Username != username {
			return nil, fmt.Errorf("board not found")
		}
	}

	return pinBoard, nil
}
