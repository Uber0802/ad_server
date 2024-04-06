package main

import (
	"net/http"

	"github.com/Uber0802/ad_server/package/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterAdRoutes(r)
	http.Handle("/", r)
	http.ListenAndServe("localhost:9010", r)
}

