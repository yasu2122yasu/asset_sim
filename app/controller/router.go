package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.Static("../js", "./js")
	r.Static("../css", "./css")
	r.GET("/show/:id", Show)
	r.GET("/create", GetCreate)
	// r.POST("/create", PostCreate)

	// TODO: データベースからレコードを取得できるか検証
	r.GET("/get", GetStockData)
	r.POST("/get/world", PostWorld)
	r.GET("/get/world", GetWorld)
	// r.GET("/one/world", GetOneWorld)
	// r.GET("/edit/:id", GetEdit)
	// r.POST("/edit", PostEdit)
	// r.GET("/delete/:id", GetDelete)
	// r.POST("/delete", PostDelete)

	return r
}
