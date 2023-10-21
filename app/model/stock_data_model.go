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
	gorm.Model
	Date      string
	MsciWorld float64
}

func GetAllStockData() (stock_datas []StockData) {
	result := Db.First(&stock_datas)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOneStockData() (stock_data StockData) {
	result := Db.Where("price  = ?", 591347.0871464558).Find(&stock_data)

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
