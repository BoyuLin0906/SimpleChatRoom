package router

import (
	"net/http"

	"simple_chatroom/controllers"

	"github.com/gin-gonic/gin"
)

func home_page(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "template/index.html")
}

func chat(c *gin.Context) {
	controllers.ServeChatting(c.Writer, c.Request)
}

func InitServer() *gin.Engine {
	router := gin.Default()
	router.Static("/js", "template/js")
	URLGroup := router.Group("/")
	{
		URLGroup.GET("/", home_page)
		URLGroup.GET("/ws", chat)
	}

	return router
}
