package main

import (
	"log"
	"net/http"
)

func FirstMiddlware(function func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("In middleware - Before handler")
		function(w, r) // this is orginal function call
		log.Println("In middleware - After handler")
	}
}

