package controllers

import (
	"gin_weibo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AlbumGet(c *gin.Context) {
	//获取session
	islogin := GetSession(c)
	albums, _ := models.FindAllAlbums()

	c.HTML(http.StatusOK, "album.html", gin.H{"IsLogin": islogin, "Album": albums})
}
