package service

import (
	"errors"
	"net/http"
	"super-feed-service/cmd/config"
	models "super-feed-service/cmd/models/feed"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func ShareFeedContentService(c *gin.Context) (interface{}, error) {
	feedId := c.Param("feedId")
	username := c.GetString("username")

	feedShare := models.FeedShare{
		Username: username,
		FeedId:   feedId,
	}

	if err := validator.New().Struct(feedShare); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}

	if _, err := config.GetOne(&feedShare); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return config.CreateOne(&feedShare)
		}
		return nil, err
	}

	res, err := config.GetOne(&models.Feed{
		Id: feedId,
	})
	if err != nil {
		return nil, err
	}

	res.SharesCount += 1
	res.UpdatedAt = time.Now()

	return config.UpdateOne(&res)
}
