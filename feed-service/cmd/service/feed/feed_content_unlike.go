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

func UnLikeFeedContentService(c *gin.Context) (interface{}, error) {
	feedId := c.Param("feedId")

	username := c.GetString("username")
	liked := false

	feedSeenLike := models.FeedSeenLike{
		Username: username,
		FeedId:   feedId,
		Liked:    &liked,
	}

	if err := validator.New().Struct(feedSeenLike); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}

	if _, err := config.UpdateOne(&feedSeenLike); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return config.CreateOne(&feedSeenLike)
		}

		return nil, err
	}

	res, err := config.GetOne(&models.Feed{
		Id: feedId,
	})
	if err != nil {
		return nil, err
	}

	if res.LikesCount > 0 {
		res.LikesCount -= 1
	}
	res.UpdatedAt = time.Now()

	return config.UpdateOne(&res)
}
