package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/Uber0802/ad_server/package/utils"
	"github.com/Uber0802/ad_server/package/models"
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

func GetAdByAge(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	Age := vars["Age"]
	age, err := strconv.ParseInt(Age, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	AdDetails, _ := models.GetAdByAge(age)
	res, _ := json.Marshal(AdDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAdByGender(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	gender := vars["Gender"]
	AdDetails, _ := models.GetAdByGender(gender)
	res, _ := json.Marshal(AdDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAdByCountry(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	country := vars["Country"]
	AdDetails, _ := models.GetAdByCountry(country)
	res, _ := json.Marshal(AdDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAdByPlatform(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	platform := vars["Platform"]
	AdDetails, _ := models.GetAdByPlatform(platform)
	res, _ := json.Marshal(AdDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
