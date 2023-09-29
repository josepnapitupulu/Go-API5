package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/josepnapitupulu/API_Kegiatan/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterKegiatansRoutes(r)
	http.Handle("/", r)

	fmt.Print("Starting Server localhost:7075")
	log.Fatal(http.ListenAndServe("localhost:7075", r))
}