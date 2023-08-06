package service

import (
	"fmt"
	"strconv"
	"super-feed-service/cmd/config"
	models "super-feed-service/cmd/models/pin_board"

	"github.com/gin-gonic/gin"
)

func GetPinsByBoardService(c *gin.Context) (interface{}, error) {
	if _, err := GetPinBoardService(c); err != nil {
		return nil, err
	}

	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		return nil, err
	}

	boardID := c.Param("boardId")
	if boardID == "" {
		return nil, fmt.Errorf("invalid board id")
	}

	return config.GetAll(&models.PinFeedContent{
		BoardID: boardID,
	}, page, pinBoardsPerPage, "")
}
