package server

import (
	"AynaAPI/config"
	"AynaAPI/server/api/v1/anime"
	"AynaAPI/server/api/v1/general"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(config.ServerConfig.GinMode)
}

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	var engine *gin.Engine = gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	apiV1 := engine.Group("/api/v1")
	{
		generalApi := apiV1.Group("/general")
		{
			generalApi.GET("/bypasscors", general.BypassCors)
			generalApi.POST("/upload/bilipic", general.UploadBiliPic)
		}
		animeApi := apiV1.Group("/anime")
		{
			animeApi.GET("/:provider/search", anime.Search)
			animeApi.GET("/:provider/playurl", anime.GetPlayUrl)
			animeApi.GET("/:provider/info", anime.GetInfo)
		}
	}
	return engine

}
