package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gourlshort/handler"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("POST").Path("/urls/create").HandlerFunc(handler.CreateUrl)
	router.Methods("GET").Path("/urls/{name}").HandlerFunc(handler.GetRedirect)

	fmt.Println("Starting the server.")
	http.ListenAndServe(":3000", router)
}
