package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Uber0802/ad_server/package/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	fmt.Println("start")
	r := mux.NewRouter()
	routes.RegisterAdRoutes(r)
	http.Handle("/", r)
	fmt.Println("Listening")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	fmt.Println("here")
}

