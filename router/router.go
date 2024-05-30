package router

import (
	"github.com/cfindlayisme/wmb/requesthandlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/message", requesthandlers.PostMessage)
	router.GET("/message", requesthandlers.QueryMessage)
	router.POST("/directedMessage", requesthandlers.PostDirectedMessage)
	router.POST("/subscribe/message", requesthandlers.PostSubscribePrivmsg)
	router.POST("/unsubscribe/message", requesthandlers.PostUnsubscribePrivmsg)

	return router
}
