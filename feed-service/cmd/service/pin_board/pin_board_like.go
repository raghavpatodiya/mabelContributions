package service

import (
	"errors"
	"net/http"
	"super-feed-service/cmd/config"
	models "super-feed-service/cmd/models/pin_board"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func LikePinBoardService(c *gin.Context) (interface{}, error) {
	boardId := c.Param("boardId")
	username := c.GetString("username")
	liked := true

	boardSeenLike := models.PinBoardSeenLike{
		Username: username,
		BoardId:  boardId,
		Liked:    &liked,
	}

	if err := validator.New().Struct(boardSeenLike); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}

	if _, err := config.UpdateOne(&boardSeenLike); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return config.CreateOne(&boardSeenLike)
		}

		return nil, err
	}

	res, err := config.GetOne(&models.PinBoard{
		BoardID: boardId,
	})
	if err != nil {
		return nil, err
	}

	res.LikesCount += 1

	res.IsUpdated = true
	res.UpdatedAt = time.Now()

	return config.UpdateOne(&res)
}
