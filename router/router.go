package router

import (
	"github.com/gin-gonic/gin"
	"smallweb/handler"
)

func InitRouter(e *gin.Engine) {
	api := e.Group("/api")
	{
		api.GET("getChatReplay", handler.GetChatReplay)
	}
}
