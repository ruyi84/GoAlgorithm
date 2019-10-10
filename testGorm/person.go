package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Person struct {
	gorm.Model
	Name string
}

func init() {
	var err error

	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8&loc=Local&parseTime=True")
	if err != nil {
		fmt.Println("gorm.Open Faild", err)
		return
	}

	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(100)
}

func GetDB() *gorm.DB {
	return db

}

func main() {
	p := Person{}
	p.Name = "2"
	p.Add()
}

func (item *Person) Add() {
	db := GetDB()
	db.Create(&item)
}

func (item *Person) Create() {
	db := GetDB()
	db.CreateTable(item)
}

func (item *Person) Update() {
	db := GetDB()
}
