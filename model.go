package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func (*Database) Init() {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", username, password, addr, DBName)
	db, err := gorm.Open("mysql", config)
	if err != nil {
		panic(err)
	}
	DB = &Database{
		Self: db,
	}
}
