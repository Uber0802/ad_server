package models

import(
	"github.com/jinzhu/gorm"
	"github.com/Uber0802/ad_server/package/config"
	"time"
)

var db *gorm.DB

type Ad struct {
	gorm.Model
	Title      string    `json:"title"`
	StartAt    time.Time `json:"startAt"`
	EndAt      time.Time `json:"endAt"`
	Conditions struct {
		Age      []int    `json:"age,omitempty"`
		Gender   []string `json:"gender,omitempty"`
		Country  []string `json:"country,omitempty"`
		Platform []string `json:"platform,omitempty"`
	} `json:"conditions"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Ad{})
}

func (a *Ad) CreateAd() *Ad{
	db.NewRecord(a)
	db.Create(&a)
	return a
}

func GetAdByAge(Age int64) (*Ad, *gorm.DB){
	var getAd Ad
	db := db.Where("Age=?", Age).Find(&getAd)
	return &getAd, db
}

func GetAdByGender(Gender string) (*Ad, *gorm.DB){
	var getAd Ad
	db := db.Where("Gender=?", Gender).Find(&getAd)
	return &getAd, db
}

func GetAdByCountry(Country string) (*Ad, *gorm.DB){
	var getAd Ad
	db := db.Where("Country=?", Country).Find(&getAd)
	return &getAd, db
}

func GetAdByPlatform(Platform string) (*Ad, *gorm.DB){
	var getAd Ad
	db := db.Where("Platform=?", Platform).Find(&getAd)
	return &getAd, db
}
