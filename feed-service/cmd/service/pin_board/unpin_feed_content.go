package service

import (
	"fmt"
	"net/http"
	"super-feed-service/cmd/config"
	models "super-feed-service/cmd/models/pin_board"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UnpinFeedContentService(c *gin.Context) (interface{}, error) {
	username := c.GetString("username")
	if username == "" {
		return nil, fmt.Errorf("invalid username")
	}

	feedID := c.Param("feedId")
	boardID := c.Param("boardId")

	feedPin := models.PinFeedContent{
		FeedID:  feedID,
		BoardID: boardID,
	}

	if err := validator.New().Struct(feedPin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}

	if _, err := config.GetOne(&models.PinBoard{
		BoardID:  boardID,
		Username: username,
	}); err != nil {
		return nil, err
	}

	return config.DeleteOne(&models.PinFeedContent{
		FeedID:  feedID,
		BoardID: boardID,
	})
}
