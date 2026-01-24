package main

import (
	"log"
	"myApp/pkg"
	"myApp/pkg/handelers"
	"net/http"
)

func main() {

	http.HandleFunc("/", handelers.Home)
	http.HandleFunc("/About", handelers.About)

	log.Println("Server is Listenining on port", pkg.PORT)
	http.ListenAndServe(pkg.PORT, nil)
}
