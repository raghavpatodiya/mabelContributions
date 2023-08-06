package service

import (
	"strconv"
	"super-feed-service/cmd/config"
	models "super-feed-service/cmd/models/feed"

	"github.com/gin-gonic/gin"
)

func GetUserClusterId(username string) string {
	return "dafe1243-6ac6-4e04-b0c1-dd90cc80e231"
}

func GetUserFeedsService(c *gin.Context) (interface{}, error) {

	// id := c.GetString("username")
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		return nil, err
	}
	// clusterid := GetUserClusterId(id)
	// wherequery := fmt.Sprintf("feed_cluster_map.cluster_id='%s'", clusterid)
	// var results *models.Feed
	// joins:=[]string{"JOIN feed_cluster_map on feeds.id = feed_cluster_map.feed_id"}
	// res, err := config.JoinGetAll("feeds", "*", joins, page, FEED_PER_PAGE, wherequery, results)
	res, err := config.GetAll(&models.Feed{}, page, FEED_PER_PAGE, "asset_ranking desc")
	return res, err
}
