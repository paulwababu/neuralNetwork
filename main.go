package main

import (

	"net/http"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("./public").HTTPBox()))
	http.ListenAndServe(":8080", router)
}