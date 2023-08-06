package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// TODO: Set environmental vars
const (
	MYSQL_CONN_STR = "root:Lyconchanhh2001@tcp(localhost:3306)/go-book-ms?charset=utf8&parseTime=True&loc=Local"
)

func Connect() {
	d, err := gorm.Open("mysql", MYSQL_CONN_STR)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
