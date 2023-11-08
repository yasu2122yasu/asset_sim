package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string
	Price int32
}

type StockData struct {
	gorm.Model
	Title string
	Price int32
}

type Stock struct {
	Id          int32   `gorm:"primaryKey"`
	Date        string  `gorm:"column:date"`
	WorldDollar float64 `gorm:"column:world_dollar"`
	WorldYen    float64 `gorm:"column:world_yen"`
	Sp500Dollar float64 `gorm:"column:sp500_dollar"`
	Sp500Yen    float64 `gorm:"column:sp500_yen"`
}

func GetAllStockData(stock_flag string) []float64 {

	var stock_flags []float64
	Db.Model(&Stock{}).Select(stock_flag).Find(&stock_flags)

	return stock_flags
}

// func GetOneStockData(initial_year string) (world []World) {
// 	// stock_dataテーブルにあるPKであるid=2のレコードを取得する。
// 	// result := Db.First(&world, 2) // 主キーを用いてレコード検索

// 	result := Db.Where("date >= ?", initial_year).Order("date ASC").Limit(30).Find(&world)

// 	if result.Error != nil {
// 		panic(result.Error)
// 	}

// 	return
// }

// func GetAllWorld() (worlds []World) {
// 	result := Db.Find(&worlds)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return
// }

func GetAll() (books []Book) {
	result := Db.Find(&books)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOne(id int) (stock_data StockData) {
	result := Db.First(&stock_data, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func (b *Book) Create() {
	result := Db.Create(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *Book) Update() {
	result := Db.Save(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *Book) Delete() {
	result := Db.Delete(b)
	fmt.Println(result)
	if result.Error != nil {
		panic(result.Error)
	}
}
