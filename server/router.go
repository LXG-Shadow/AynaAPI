package server

import (
	"AynaAPI/config"
	"AynaAPI/server/api/v1/auth"
	"AynaAPI/server/api/v1/general"
	"AynaAPI/server/api/v2/anime"
	authV2 "AynaAPI/server/api/v2/auth"
	"AynaAPI/server/api/v2/music"
	"AynaAPI/server/api/v2/upload"
	"AynaAPI/server/controllers/chat"
	"AynaAPI/server/controllers/index"
	"AynaAPI/server/middleware/jwt"
	"AynaAPI/server/middleware/perm"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	// load all provider
	_ "AynaAPI/server/service/api_service"

	_ "AynaAPI/docs"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	var engine *gin.Engine = gin.New()

	engine.LoadHTMLGlob("./server/view/**/*")

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	corsconfig := cors.DefaultConfig()
	// https://www.cnblogs.com/cnxkey/articles/14259716.html
	//corsconfig.AllowOrigins = []string{"http://127.0.0.1:3000", "http://localhost:3000"}
	corsconfig.AllowOrigins = []string{"*"}
	corsconfig.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	corsconfig.AllowCredentials = true

	engine.Use(cors.New(corsconfig))

	staticDirFs := engine.Group("/_static")
	{
		staticDirFs.Use(jwt.AuthUser(), perm.CheckPermission(128))
		staticDirFs.StaticFS("/", gin.Dir(config.ServerConfig.FileRoot, true))
	}

	staticFs := engine.Group("/static")
	{
		staticFs.Static("/file", config.ServerConfig.GetFilePath("file"))
		staticFs.Static("/upload", config.ServerConfig.UploadSavePath)
	}

	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := engine.Group("/api/v1")
	{
		authApi := apiV1.Group("/auth")
		{
			authApi.GET("/login", auth.Login)
			authApi.GET("/info", jwt.GetUserOnly(), auth.GetInfo)
			authApi.GET("/logout", jwt.GetUserOnly(), auth.GetInfo)
		}
		generalApi := apiV1.Group("/general")
		{
			generalApi.GET("/bypasscors", general.BypassCorsFixFile)
			generalApi.GET("/bypasscors/:any", general.BypassCors)
			generalApi.GET("/teamsplit", general.GetRandomSeparation)
		}
	}

	apiV2 := engine.Group("/api/v2")
	{
		authV2Api := apiV2.Group("/auth")
		{
			authV2Api.GET("/login", authV2.Login)
			authV2Api.GET("/info", jwt.GetUserOnly(), authV2.GetInfo)
			authV2Api.GET("/logout", jwt.GetUserOnly(), authV2.Logout)
		}

		uploadApi := apiV1.Group("/upload")
		{
			//uploadApi.Use(jwt.AuthUser())
			uploadApi.POST("/image", upload.UploadFile)
		}

		animeApi := apiV2.Group("/anime")
		{
			animeApi.GET("/plist", anime.GetProviderList)
			animeApi.GET("/providerlist", anime.GetProviderList)

			animeApi.GET("/search", anime.SearchAll)
			animeApi.GET("/search/:provider", anime.Search)

			animeApi.GET("/playurl", anime.GetPlayUrl)
			animeApi.GET("/info", anime.GetInfo)
		}
		//novelApi := apiV2.Group("/novel")
		//{
		//	novelApi.GET("/plist", novel.GetProviderList)
		//	novelApi.GET("/providerlist", novel.GetProviderList)
		//
		//	novelApi.GET("/info", novel.GetInfo)
		//	novelApi.GET("/content", novel.GetContent)
		//
		//	novelApi.GET("/search/:provider", novel.Search)
		//	novelApi.GET("/search", novel.SearchAll)
		//}
		musicApi := apiV2.Group("/music")
		{
			musicApi.GET("/plist", music.GetProviderList)
			musicApi.GET("/providerlist", music.GetProviderList)

			musicApi.GET("/search", music.SearchAll)
			musicApi.GET("/search/:provider", music.Search)

			musicApi.GET("/url", music.GetUrl)
			musicApi.GET("/info", music.GetInfo)
		}
	}

	frontend := engine.Group("/frontend")
	{
		frontend.GET("/", index.Index)
		frontend.GET("/chat", jwt.GetUserOnly(), chat.Index)
	}

	return engine
}
