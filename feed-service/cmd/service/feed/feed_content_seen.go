package service

import (
	"fmt"
	"super-feed-service/cmd/config"
	models "super-feed-service/cmd/models/feed"

	"github.com/gin-gonic/gin"
)

func SeenFeedContentService(c *gin.Context) (interface{}, error) {
	feedId := c.Param("feedId")
	if feedId == "" {
		return nil, fmt.Errorf("invalid feed id")
	}

	username := c.GetString("username")
	if username == "" {
		return nil, fmt.Errorf("invalid username")
	}

	return config.CreateOne(&models.FeedSeenLike{
		Username: username,
		FeedId:   feedId,
	})
}
