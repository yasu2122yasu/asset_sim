package model

import (
	"fmt"

	"gorm.io/gorm"
)

type StockData struct {
	gorm.Model
	Title string
	Count int
	Date string
}

func GetAll() (stock_datas []StockData) {
	result := Db.Find(&stock_datas)
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

func (b *StockData) Create() {
	result := Db.Create(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *StockData) Update() {
	result := Db.Save(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *StockData) Delete() {
	result := Db.Delete(b)
	fmt.Println(result)
	if result.Error != nil {
		panic(result.Error)
	}
}
