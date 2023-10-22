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

type World struct {
	ID    int32   `gorm:"primaryKey"`  // 大文字の "ID" に修正
	Date  string  `gorm:"column:date"` // DBのカラム名を明示的に指定
	Price float64 `gorm:"column:price"`
}

func GetAllStockData() (stock_datas []StockData) {
	result := Db.First(&stock_datas)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOneStockData() (world World) {
	// stock_dataテーブルにあるPKであるid=2のレコードを取得する。
	result := Db.First(&world, 2) // 主キーを用いてレコードを検索

	if result.Error != nil {
		panic(result.Error)
	}

	return
}

func GetAllWorld() (worlds []World) {
	result := Db.Find(&worlds)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

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
