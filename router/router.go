package router

import (
	"github.com/cfindlayisme/wmb/env"
	"github.com/cfindlayisme/wmb/requesthandlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if env.GetDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.POST("/message", requesthandlers.PostMessage)
	router.GET("/message", requesthandlers.QueryMessage)
	router.POST("/directedMessage", requesthandlers.PostDirectedMessage)
	router.POST("/subscribe/message", requesthandlers.PostSubscribePrivmsg)
	router.POST("/unsubscribe/message", requesthandlers.PostUnsubscribePrivmsg)

	return router
}
