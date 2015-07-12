package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db gorm.DB

func init() {
	var err error
	db, err = gorm.Open("postgres", "dbname=passr_debug sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)

	db.AutoMigrate(&Credential{})
}
