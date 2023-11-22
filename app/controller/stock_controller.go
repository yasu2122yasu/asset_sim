package controller

import (
	"app/model"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCreate(c *gin.Context) {
	c.HTML(200, "create.html", gin.H{})
}

func Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOne(id)
	c.HTML(200, "show.html", gin.H{"data": data})
}

func GetStockData(c *gin.Context) {
	c.HTML(200, "world.html", gin.H{})
}

func PostWorld(c *gin.Context) {
	stock := c.PostForm("stock")
	currency := c.PostForm("currency")
	string_monthly_amount := c.PostForm(("monthly_amount"))

	float_monthly_amount, _ := strconv.ParseFloat(string_monthly_amount, 64)

	// どの株価指数を利用するか判定するフラグを作成する。
	var stockFlag string

	if stock == "sp_500" && currency == "dollar" {
		stockFlag = "sp500_dollar"
	} else if stock == "sp_500" && currency == "yen" {
		stockFlag = "sp500_yen"
	} else if stock == "world_stock" && currency == "dollar" {
		stockFlag = "world_dollar"
	} else if stock == "world_stock" && currency == "yen" {
		stockFlag = "world_yen"
	} else {
		errors.New("何らかのエラーが発生しました。入力値を変更してください。")
	}

	total := model.GetAllStockData(stockFlag)

	var sum float64
	var totalMoney []int32

	fmt.Println(float_monthly_amount)

	for _, v := range total {
		// 何株購入するかのカウント（oneは保有株数）
		one := float_monthly_amount / v

		sum = sum + one

		totalMoneyFloat := sum * v

		totalMoneyValue := int32(totalMoneyFloat)

		totalMoney = append(totalMoney, totalMoneyValue)

	}

	c.HTML(200, "world.html", gin.H{
		"string_monthly_amount": string_monthly_amount,
		"totalMoney":            totalMoney,
	})

}

func GetWorld(c *gin.Context) {
	c.HTML(200, "world.html", gin.H{})
}

// func GetOneWorld(c *gin.Context) {
// 	price := model.GetOneStockData()
// 	c.HTML(200, "oneworld.html", gin.H{
// 		"price": price,
// 	})
// }
