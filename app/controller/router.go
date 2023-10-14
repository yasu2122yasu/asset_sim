package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	// r.GET("/show/:id", Show)
	r.GET("/create", GetCreate)
	r.POST("/create", PostCreate)

	// TODO: データベースからレコードを取得できるか検証
	r.GET("/get", GetStockData)
	// r.GET("/edit/:id", GetEdit)
	// r.POST("/edit", PostEdit)
	// r.GET("/delete/:id", GetDelete)
	// r.POST("/delete", PostDelete)

	return r
}
