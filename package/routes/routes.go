package routes

import (
	"github.com/gorilla/mux"
	"github.com/Uber0802/ad_server/package/controllers"

)

var RegisterAdRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/v1/ad", controllers.CreateAd).Methods("POST")
	router.HandleFunc("/api/v1/ad", controllers.GetAdByAge).Methods("GET")
	router.HandleFunc("/api/v1/ad", controllers.GetAdByGender).Methods("GET")
	router.HandleFunc("/api/v1/ad", controllers.GetAdByCountry).Methods("GET")
	router.HandleFunc("/api/v1/ad", controllers.GetAdByPlatform).Methods("GET")
}

