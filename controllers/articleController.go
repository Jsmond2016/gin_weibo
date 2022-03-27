package controllers

import (
	"fmt"
	"gin_weibo/models"
	"gin_weibo/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/*
当访问/add路径的时候回触发AddArticleGet方法
响应的页面是通过HTML
*/
func AddArticleGet(c *gin.Context) {

	//获取session
	islogin := GetSession(c)

	c.HTML(http.StatusOK, "write_article.html", gin.H{"IsLogin": islogin})
}

func AddArticlePost(c *gin.Context) {

	//获取浏览器传输的数据，通过表单的name属性获取值
	//获取表单信息
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	//实例化model，将它出入到数据库中
	art := models.Article{Id: 0, Title: title, Tags: tags, Short: short, Content: content, Author: "孔壹学院", Createtime: time.Now().Unix()}
	_, err := models.AddArticle(art)

	//返回数据给浏览器
	response := gin.H{}
	if err == nil {
		//无误
		response = gin.H{"code": 1, "message": "ok"}
	} else {
		response = gin.H{"code": 0, "message": "error"}
	}

	c.JSON(http.StatusOK, response)

}

// 显示文章
func ShowArticleGet(c *gin.Context) {

	//获取session
	islogin := GetSession(c)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id:", id)

	//获取id所对应的文章信息
	art := models.QueryArticleWithId(id)
	//渲染HTML
	// c.HTML(http.StatusOK, "show_article.html", gin.H{"IsLogin": islogin, "Title": art.Title, "Content": art.Content})
	// 以 markdown 的格式渲染文档
	c.HTML(http.StatusOK, "show_article.html", gin.H{"IsLogin": islogin, "Title": art.Title, "Content": utils.SwitchMarkdownToHtml(art.Content)})
}
