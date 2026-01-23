package main

import (
	"fmt"
	"myApp/internal"
	"myApp/internal/handelers"
	"net/http"
)

func main() {

	http.HandleFunc("/", handelers.Home)
	http.HandleFunc("/About", handelers.About)

	fmt.Println("Server is Listenining on port", internal.PORT)
	http.ListenAndServe(internal.PORT, nil)
}
