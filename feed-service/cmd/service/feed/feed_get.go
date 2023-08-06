package service

import (
	"super-feed-service/cmd/config"
	"super-feed-service/cmd/models/feed"
	"github.com/gin-gonic/gin"
)

func GetFeedService(c *gin.Context) (interface{}, error) {

	id := c.Param("id")
	feed := models.Feed{Id: id}
	res, err := config.GetOne(&feed)
	return res, err
}
