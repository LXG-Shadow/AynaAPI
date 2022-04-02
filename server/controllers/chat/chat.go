package chat

import (
	"AynaAPI/server/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	appG := app.AppGin{C: context}
	//user, pl = appG.GetUser()
	appG.C.HTML(http.StatusOK, "chat/chat.html", gin.H{})
}
