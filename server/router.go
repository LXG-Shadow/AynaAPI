package server

import (
	"AynaAPI/config"
	"AynaAPI/server/api/v1/auth"
	"AynaAPI/server/api/v1/general"
	"AynaAPI/server/api/v1/upload"
	"AynaAPI/server/api/v2/anime"
	"AynaAPI/server/api/v2/novel"
	"AynaAPI/server/fs"
	"AynaAPI/server/middleware/jwt"
	"AynaAPI/server/middleware/perm"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "AynaAPI/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	gin.SetMode(config.ServerConfig.GinMode)
}

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	var engine *gin.Engine = gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	corsconfig := cors.DefaultConfig()
	corsconfig.AllowOrigins = []string{"*"}

	engine.Use(cors.New(corsconfig))

	staticDirFs := engine.Group("/_static")
	{
		staticDirFs.Use(jwt.AuthUser(), perm.CheckPermission(128))
		staticDirFs.StaticFS("/", gin.Dir(config.ServerConfig.FileRoot, true))
	}

	staticFs := engine.Group("/static")
	{
		staticFs.Static("/file", config.ServerConfig.GetFilePath("file"))
		staticFs.Static(fs.GetUploadUrl(), fs.GetUploadPath())
	}

	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := engine.Group("/api/v1")
	{
		authApi := apiV1.Group("/auth")
		{
			authApi.GET("/login", auth.Login)
			authApi.GET("/info", jwt.GetUserOnly(), auth.GetInfo)
		}
		uploadApi := apiV1.Group("/upload")
		{
			uploadApi.Use(jwt.AuthUser())
			uploadApi.POST("/bilipic", upload.UploadBiliPic)
			uploadApi.POST("/image", upload.UploadImage)
		}
		generalApi := apiV1.Group("/general")
		{
			generalApi.GET("/bypasscors", general.BypassCors)
			generalApi.GET("/teamsplit", general.GetRandomSeparation)
		}
	}

	apiV2 := engine.Group("/api/v2")
	{
		animeApi := apiV2.Group("/anime")
		{
			animeApi.GET("/plist", anime.GetProviderList)
			animeApi.GET("/providerlist", anime.GetProviderList)

			animeApi.GET("/search", anime.SearchAll)
			animeApi.GET("/search/:provider", anime.Search)

			animeApi.GET("/playurl", anime.GetPlayUrl)
			animeApi.GET("/info", anime.GetInfo)
		}
		novelApi := apiV2.Group("/novel")
		{
			novelApi.GET("/plist", novel.GetProviderList)
			novelApi.GET("/providerlist", novel.GetProviderList)

			novelApi.GET("/info", novel.GetInfo)
			novelApi.GET("/content", novel.GetContent)

			novelApi.GET("/search/:provider", novel.Search)
			novelApi.GET("/search", novel.SearchAll)
		}
	}
	return engine
}
