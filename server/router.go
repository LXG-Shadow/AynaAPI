package server

import (
	"AynaAPI/config"
	"AynaAPI/server/api/v1/anime"
	"AynaAPI/server/api/v1/auth"
	"AynaAPI/server/api/v1/general"
	"AynaAPI/server/fs"
	"AynaAPI/server/middleware/jwt"
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

	staticDirFs := engine.Group("/_static")
	{
		staticDirFs.Use(jwt.JWT())
		staticDirFs.StaticFS("/", gin.Dir(config.ServerConfig.FileRoot, true))
	}

	staticFs := engine.Group("/static")
	{
		staticFs.Static("/file", config.ServerConfig.GetFilePath("file"))
		staticFs.Static(fs.GetUploadUrl(), fs.GetUploadPath())
	}

	apiV1 := engine.Group("/api/v1")
	{
		authApi := apiV1.Group("/auth")
		{
			authApi.GET("/login", auth.Login)
		}
		generalApi := apiV1.Group("/general")
		{
			generalApi.GET("/bypasscors", general.BypassCors)
			generalApi.GET("/teamsplit", general.GetRandomSeparation)
			generalApi.POST("/upload/bilipic", jwt.JWT(), general.UploadBiliPic)
			generalApi.POST("/upload/image", jwt.JWT(), general.UploadImage)
		}
		animeApi := apiV1.Group("/anime")
		{
			animeApi.GET("/search", anime.SearchAll)
			animeApi.GET("/playurl", anime.GetPlayUrlAll)
			animeApi.GET("/info", anime.InfoAll)
			animeApi.GET("/resolve", anime.ResolveAll)

			animeApi.GET("/:provider/search", anime.Search)
			animeApi.GET("/:provider/playurl", anime.GetPlayUrl)
			animeApi.GET("/:provider/info", anime.GetInfo)
			animeApi.GET("/:provider/resolve", anime.Resolve)
		}
	}
	return engine

}
