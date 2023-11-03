package controller

import (
	"app/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCreate(c *gin.Context) {
	c.HTML(200, "create.html", gin.H{})
}

func PostCreate(c *gin.Context) {
	// string_investment_duration_year := c.PostForm("investment_duration_year")
	string_initial_investment := c.PostForm("initial_investment")
	string_monthly_amount := c.PostForm("monthly_amount")
	string_accumulation_period := c.DefaultPostForm("year", "10")
	string_rate := c.PostForm("rate")
	// TODO: 為替による調整を追加したときに実装
	// string_exchange_rate := c.PostForm("exchange_order")

	// 投資開始年度（ex.1970年）から30年度分のレコードを取得できるプログラムを書く

	var float_initial_investment float64
	var float_monthly_amount float64
	var int_accumulation_period int

	float_initial_investment, _ = strconv.ParseFloat(string_initial_investment, 32)
	float_monthly_amount, _ = strconv.ParseFloat(string_monthly_amount, 32)
	int_accumulation_period, _ = strconv.Atoi(string_accumulation_period)

	// 選択した投資対象に応じてリターンを分岐
	var float_rate float64
	switch string_rate {
	case "sp_500":
		float_rate = 1.07
	case "world_stock":
		float_rate = 1.06
	case "nikkei225":
		float_rate = 1.04
	case "us_bond":
		float_rate = 1.03
	case "jpn_bond":
		float_rate = 1.01
	}

	for i := 0; i < int_accumulation_period; i++ {
		float_initial_investment = float_initial_investment + float_monthly_amount*12
		float_initial_investment = float_initial_investment * float_rate
	}
	fmt.Println(float_initial_investment)

	result := int32(float_initial_investment)

	c.HTML(200, "create.html", gin.H{
		"result": result,
		// TODO: 為替による調整を追加したときに実装
		// "string_exchange_rate":       string_exchange_rate,
		"string_initial_investment":  string_initial_investment,
		"string_monthly_amount":      string_monthly_amount,
		"string_accumulation_period": string_accumulation_period,
		"string_rate":                string_rate,
	})
}

func Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOne(id)
	c.HTML(200, "show.html", gin.H{"data": data})
}

func GetStockData(c *gin.Context) {
	stock_data := model.GetAllStockData()
	c.HTML(200, "stock_data.html", gin.H{
		"stock_data": stock_data,
	})
}

func GetWorld(c *gin.Context) {
	world := model.GetAllWorld()

	var sum float64
	var totalMoney []int32

	for _, v := range world {
		r := v.Price
		// 何株あるかのカウント（oneは保有株数）
		one := 33333 / r

		sum = sum + one

		totalMoneyFloat := sum * v.Price

		totalMoneyValue := int32(totalMoneyFloat)

		totalMoney = append(totalMoney, totalMoneyValue)

	}

	c.HTML(200, "world.html", gin.H{
		"totalMoney": totalMoney,
	})
}

func GetOneWorld(c *gin.Context) {
	price := model.GetOneStockData()
	c.HTML(200, "oneworld.html", gin.H{
		"price": price,
	})
}
