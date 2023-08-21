package controller

import (
	"app/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	stock_datas := model.GetAll()
	c.HTML(200, "index.html", gin.H{"stock_datas": stock_datas})
}

func Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	stock_data := model.GetOne(id)
	c.HTML(200, "show.html", gin.H{"stock_data": stock_data})
}

func GetCreate(c *gin.Context) {
	c.HTML(200, "create.html", gin.H{})
}

//formから送信された内容をDBに登録せずにそのまま返す
func PostCreate(c *gin.Context) {
	title := c.PostForm("title")
	date := c.PostForm("date")

	c.HTML(200, "create.html", gin.H{"title": title, "date": date})
}
