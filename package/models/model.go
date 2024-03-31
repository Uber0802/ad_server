package models

import(
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/Uber0802/ad_server/package/config"
	"time"
)

var db *gorm.DB
type Conditions struct {
	AgeStart  int      `json:"ageStart,omitempty"`
	AgeEnd    int      `json:"ageEnd,omitempty"`
	Gender    []string `json:"gender,omitempty"`
	Country   []string `json:"country,omitempty"`
	Platform  []string `json:"platform,omitempty"`
}

type Ad struct {
	gorm.Model
	Title      string    `json:"title"`
	StartAt    time.Time `json:"startAt"`
	EndAt      time.Time `json:"endAt"`
	Conditions string    `json:"conditions"`
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

func (a *Ad) MatchesConditions(age int, gender, country, platform string) bool {
    // Unmarshal Conditions JSON
    var conditions Conditions
    json.Unmarshal([]byte(a.Conditions), &conditions)

    if !(age >= conditions.AgeStart && age <= conditions.AgeEnd) && (conditions.AgeStart != 0 || conditions.AgeEnd != 0) {
		return false
    }

    if gender != "" && !contains(conditions.Gender, gender) && len(conditions.Gender) > 0 {
		return false
    }

    if country != "" && !contains(conditions.Country, country) && len(conditions.Country) > 0 {
		return false
    }

    if platform != "" && !contains(conditions.Platform, platform) && len(conditions.Platform) > 0 {
		return false
    }
    return true
}

func contains(slice []string, str string) bool {
    for _, v := range slice {
        if v == str {
            return true
        }
    }
    return false
}

