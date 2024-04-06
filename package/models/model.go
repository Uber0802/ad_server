package models

import(
	jsoniter "github.com/json-iterator/go"
	"github.com/jinzhu/gorm"
	"github.com/Uber0802/ad_server/package/config"
	"time"
	"sync"
	"fmt"
)
var json = jsoniter.ConfigCompatibleWithStandardLibrary

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

var conditionsCache sync.Map

func (a *Ad) MatchesConditions(age int, gender, country, platform string) bool {
    // Unmarshal Conditions JSON
	cacheKey := fmt.Sprintf("%d-%s-%s-%s", age, gender, country, platform)
    if cachedResult, ok := conditionsCache.Load(cacheKey); ok {
        return cachedResult.(bool)
    }
    var conditions Conditions
    json.Unmarshal([]byte(a.Conditions), &conditions)

	result := true
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
    conditionsCache.Store(cacheKey, result)
    return result
}

func contains(slice []string, str string) bool {
    for _, v := range slice {
        if v == str {
            return true
        }
    }
    return false
}

