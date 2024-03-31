package routes

import (
	"github.com/gorilla/mux"
	"github.com/Uber0802/ad_server/package/controllers"

)

var RegisterAdRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/v1/ad", controllers.CreateAd).Methods("POST")
	router.HandleFunc("/api/v1/ad", controllers.ListAds).Methods("GET")
}

