package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("mysql", "root:LICHEN00@tcp(127.0.0.1:3306)/lastsummer?parseTime=true&charset=utf8&loc=Local")
	if err != nil {
		log.Panicf("Panic while connecting the gorm. Error: %s", err)
	}
	DB = db
	//之前一直写的DropTableIfExists
	if !DB.HasTable(&Auth{}) {
		DB.CreateTable(&Auth{})
	}
	if !DB.HasTable(&User{}) {
		DB.CreateTable(&User{})
	}
	if !DB.HasTable(&File{}) {
		DB.CreateTable(&File{})
	}

}
