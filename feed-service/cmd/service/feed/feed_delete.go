package service

import (
	"super-feed-service/cmd/models/feed"
	"super-feed-service/cmd/config"
	"time"
	"github.com/gin-gonic/gin"
)

func DeleteFeedService(c *gin.Context) (interface{}, error) {
	var feed models.Feed

	id:=c.Param("id")
	feed.Id = id
	res, err := config.GetOne(&feed)
	feed = *res
	if err != nil {
		return nil, err
	}
	feed.UpdatedAt = time.Now()
	feed.Is_Deleted = true

	resDelete, Err := config.UpdateOne(&feed)
	return resDelete, Err
}