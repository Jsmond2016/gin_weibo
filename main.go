package main

import (
	"gin_weibo/routers"
	_ "gin_weibo/routers"
)

func main() {
	router := routers.InitRouter()
	//静态资源
	router.Static("/static", "./static")
	router.Run(":8081")
}
