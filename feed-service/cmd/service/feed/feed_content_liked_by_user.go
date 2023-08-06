package service

import (
	"fmt"
	"strconv"
	"super-feed-service/cmd/config"
	models "super-feed-service/cmd/models/feed"

	"github.com/gin-gonic/gin"
)

func GetFeedContentsLikedByUserService(c *gin.Context) (interface{}, error) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		return nil, err
	}

	username := c.GetString("username")
	if username == "" {
		return nil, fmt.Errorf("invalid username")
	}

	liked := true

	return config.GetAll(&models.FeedSeenLike{
		Username: username,
		Liked:    &liked,
	}, page, FEED_PER_PAGE, "")
}
