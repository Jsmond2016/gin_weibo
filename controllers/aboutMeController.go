package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AboutMeGet(c *gin.Context) {

	//获取session
	islogin := GetSession(c)

	c.HTML(http.StatusOK, "aboultme.html", gin.H{"IsLogin": islogin, "wechat": "微信：rubyhan1314", "qq": "QQ：79539705", "tel": "Tel：13910439137"})
}
