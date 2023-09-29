package model

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	dsn := "port=5432 user=root password=password dbname=asset_sim sslmode=disable"

	var err error
	_ , err = sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println(err.Error())
    }
  fmt.Println("DB接続成功")
}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(dialector, count)
			return
		}
		panic(err.Error())
	}
}
