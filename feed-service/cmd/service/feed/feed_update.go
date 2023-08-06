package service

import (
	"super-feed-service/cmd/config"
	"super-feed-service/cmd/models/feed"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateFeedService(c *gin.Context) (interface{}, error) {
	var feed models.Feed

	id := c.Param("id")
	feed.Id = id

	feed_existing, err := config.GetOne(&feed)
	if err != nil {
		return nil, err
	}
	err = c.BindJSON(&feed_existing)
	if err != nil {
		return nil, err
	}

	feed_existing.Is_Deleted = false
	feed_existing.UpdatedAt = time.Now()

	resUpdate, err := config.UpdateOne(&feed_existing)
	return resUpdate, err
}

// 