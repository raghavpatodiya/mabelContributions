package api

import (
	"net/http"
	feed_service "super-feed-service/cmd/service/feed"
	"super-feed-service/cmd/utils"

	"github.com/gin-gonic/gin"
)

func CreateFeed(c *gin.Context) {
	res, err := feed_service.CreateFeedService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusCreated)
	c.JSON(statusCode, data)
}

func DeleteFeed(c *gin.Context) {
	res, err := feed_service.DeleteFeedService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusNoContent)
	c.JSON(statusCode, data)
}

func GetFeed(c *gin.Context) {
	res, err := feed_service.GetFeedService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetFeedByAssetCode(c *gin.Context) {
	res, err := feed_service.GetFeedByAssetCodeService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func UpdateFeed(c *gin.Context) {
	res, err := feed_service.UpdateFeedService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetUserFeeds(c *gin.Context) {
	res, err := feed_service.GetUserFeedsService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func SeenFeedContent(c *gin.Context) {
	res, err := feed_service.SeenFeedContentService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusCreated)
	c.JSON(statusCode, data)
}

func LikeFeedContent(c *gin.Context) {
	res, err := feed_service.LikeFeedContentService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusNoContent)
	c.JSON(statusCode, data)
}

func UnLikeFeedContent(c *gin.Context) {
	res, err := feed_service.UnLikeFeedContentService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusNoContent)
	c.JSON(statusCode, data)
}

func GetFeedContentsLikedByUser(c *gin.Context) {
	res, err := feed_service.GetFeedContentsLikedByUserService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func ShareFeedContent(c *gin.Context) {
	res, err := feed_service.ShareFeedContentService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusCreated)
	c.JSON(statusCode, data)
}
