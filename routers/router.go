package routers

import (
	"gin_weibo/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	//注册：
	router.GET("/register", controllers.RegisterGet)
	return router

}
