package controller

import (
	"app/model"
	"fmt"
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

func PostCreate(c *gin.Context) {
	title := c.PostForm("title")
	date := c.PostForm("date")
	stock_data := model.StockData{Title: title, Date: date}
	stock_data.Create()
	c.Redirect(301, "/")
}

func GetEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	stock_data := model.GetOne(id)
	c.HTML(200, "edit.html", gin.H{"stock_data": stock_data})
}

func PostEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	stock_data := model.GetOne(id)
	title := c.PostForm("title")
	stock_data.Title = title
	date := c.PostForm("date")
	stock_data.Date = date
	stock_data.Update()
	c.Redirect(301, "/")
}

func GetDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	stock_data := model.GetOne(id)
	c.HTML(200, "delete.html", gin.H{"stock_data": stock_data})
}

func PostDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	fmt.Printf("ここに削除用のidを入れる%d", id)
	stock_data := model.GetOne(id)
	stock_data.Delete()
	c.Redirect(301, "/")
}
