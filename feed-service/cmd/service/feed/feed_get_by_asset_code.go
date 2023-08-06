package service

import (
	"super-feed-service/cmd/config"
	models "super-feed-service/cmd/models/feed"

	"github.com/gin-gonic/gin"
)

func GetFeedByAssetCodeService(c *gin.Context) (interface{}, error) {
	asset_code := c.Param("asset_code")
	res, err := config.GetOne(&models.Feed{
		AssetCode: asset_code,
	})
	return res, err
}
