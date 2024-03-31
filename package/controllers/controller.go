package controllers

import(
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	"strconv"
	"github.com/Uber0802/ad_server/package/utils"
	"github.com/Uber0802/ad_server/package/models"
    "github.com/Uber0802/ad_server/package/config"
)

var NewAd models.Ad

func CreateAd(w http.ResponseWriter, r *http.Request){
	CreateAd := &models.Ad{} // recieve json
	utils.ParseBody(r, CreateAd) // turn json to ad struct
	a := CreateAd.CreateAd() // send to database
	res, _ := json.Marshal(a)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func ListAds(w http.ResponseWriter, r *http.Request) {
    db := config.GetDB()
    var ads []models.Ad
    var resultAds []models.Ad

    // Parse query parameters
    query := r.URL.Query()
    offset, _ := strconv.Atoi(query.Get("offset"))
    limit, _ := strconv.Atoi(query.Get("limit"))
    age, _ := strconv.Atoi(query.Get("age"))
    gender := query.Get("gender")
    country := query.Get("country")
    platform := query.Get("platform")

    // Set defaults for offset and limit
    if offset < 0 {
        offset = 0
    }
    if limit < 1 || limit > 100 {
        limit = 5
    }

    currentTime := time.Now()
    db.Where("start_at <= ? AND end_at >= ?", currentTime, currentTime).
        Order("end_at asc").Offset(offset).Limit(limit).Find(&ads)
    
    for _, ad := range ads {
		fmt.Println(ad)
        if ad.MatchesConditions(age, gender, country, platform) {
            resultAds = append(resultAds, ad)
        }
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{"items": resultAds})
}


