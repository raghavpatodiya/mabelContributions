package routers

import (
	"super-feed-service/cmd/routers/api"

	auth "super-feed-service/cmd/middleware/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "super-feed-service/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default()) //allows all origin
	router.GET("/", api.GetHelloWorld)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1/")
	{
		hello_world_api := v1.Group("/hello/")
		{
			hello_world_api.GET("/hello_world", auth.TokenAuthMiddleware("admin:read"), api.GetHelloWorld)
		}
		feed_service_api := v1.Group("/feed/")
		{
			feed_service_api.GET("/:id", auth.TokenAuthMiddleware("customer:read"), api.GetFeed)
			feed_service_api.GET("/get_feed_by_asset_code/:asset_code", auth.TokenAuthMiddleware("customer:read"), api.GetFeedByAssetCode)
			feed_service_api.POST("/create", auth.TokenAuthMiddleware("admin:write"), api.CreateFeed)
			feed_service_api.DELETE("/delete/:id", auth.TokenAuthMiddleware("admin:write"), api.DeleteFeed)
			feed_service_api.PUT("/update/:id", auth.TokenAuthMiddleware("admin:write"), api.UpdateFeed)
			feed_service_api.GET("/user/:page", auth.TokenAuthMiddleware("customer:read"), api.GetUserFeeds)
			feed_service_api.POST("/seen/:feedId", auth.TokenAuthMiddleware("customer:write"), api.SeenFeedContent)
			feed_service_api.POST("/like/:feedId", auth.TokenAuthMiddleware("customer:write"), api.LikeFeedContent)
			feed_service_api.POST("/unlike/:feedId", auth.TokenAuthMiddleware("customer:write"), api.UnLikeFeedContent)
			feed_service_api.GET("/likes/:page", auth.TokenAuthMiddleware("customer:read"), api.GetFeedContentsLikedByUser)
			feed_service_api.POST("/shared/:feedId", auth.TokenAuthMiddleware("customer:write"), api.ShareFeedContent)
		}

		pin_board_service_api := v1.Group("/pinboards/")
		{
			pin_board_service_api.POST("/", auth.TokenAuthMiddleware("customer:write"), api.CreatePinBoard)
			pin_board_service_api.GET("/:boardId", auth.TokenAuthMiddleware("customer:read"), api.GetPinBoard)
			pin_board_service_api.PUT("/:boardId", auth.TokenAuthMiddleware("customer:write"), api.UpdatePinBoard)
			pin_board_service_api.DELETE("/:boardId", auth.TokenAuthMiddleware("customer:write"), api.DeletePinBoard)
			pin_board_service_api.GET("/user/:page", auth.TokenAuthMiddleware("customer:read"), api.GetPinBoardsByUser)

			pin_board_service_api.POST("/pins/:feedId/boards/:boardId", auth.TokenAuthMiddleware("customer:write"), api.PinFeedContent)
			pin_board_service_api.DELETE("/unpin/:feedId/boards/:boardId", auth.TokenAuthMiddleware("customer:write"), api.UnpinFeedContent)
			pin_board_service_api.GET("/pins/boards/:boardId/:page", auth.TokenAuthMiddleware("customer:read"), api.PinsByBoard)
			pin_board_service_api.POST("/shared/:boardId", auth.TokenAuthMiddleware("customer:write"), api.SharePinBoard)
			pin_board_service_api.POST("/like/:boardId", auth.TokenAuthMiddleware("customer:write"), api.LikePinBoard)
			pin_board_service_api.POST("/unlike/:boardId", auth.TokenAuthMiddleware("customer:write"), api.UnlikePinBoard)
		}
	}

	return router
}
