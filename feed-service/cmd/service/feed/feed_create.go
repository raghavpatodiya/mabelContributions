package service

import (
	"encoding/json"
	"super-feed-service/cmd/config"
	models "super-feed-service/cmd/models/feed"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/google/uuid"
)

func CreateFeedService(c *gin.Context) (interface{}, error) {
	var feed models.Feed
	err := c.BindJSON(&feed)
	if err != nil {
		return nil, err
	}
	ProductListJsonData, err := json.Marshal(feed.Product_List)
	if err != nil {
		return nil, err
	}
	feed.Id = uuid.New().String()
	feed.Product_List = ProductListJsonData
	feed.Is_Deleted = false
	feed.CreatedAt = time.Now()
	if err := validator.New().Struct(feed); err != nil {
		return nil, err
	}
	res, err := config.CreateOne(&feed)
	return res, err
}
