package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var (
	b *gorm.DB
)

func Connect (){
	d.err := gorm.Open("mysql", "partha:Password/tablename?charset=utf&parseTime=True&loc=Local")
	if err != nil {
		panic((err))
	}
	db=d;

}

func GetDB() *gorm.DB{
	return db
}