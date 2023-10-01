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
	StockPrice int32
	Date       string
}

func GetAllStockData() []StockData {
	var stock_datas []StockData
	result := Db.Find(&stock_datas)
	if result.Error != nil {
		panic(result.Error)
	}
	return stock_datas
}

func GetAll() (books []Book) {
	result := Db.Find(&books)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOne(id int) (book Book) {
	result := Db.First(&book, id)
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
