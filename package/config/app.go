package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "Uber0802:uber0802!@/adtable?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
	err = db.DB().Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
}

func GetDB() *gorm.DB{
	return db
}

