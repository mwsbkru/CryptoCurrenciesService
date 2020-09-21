package main

import (
	"./controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/rates/cached", controllers.CachedCurrencies)
	router.HandleFunc("/api/rates/{code:[A-Za-z]+}", controllers.CurrencyByCode)
	http.Handle("/", router)

	// TODO move port to env
	fmt.Println("CryptoService server is listening on 8181")
    http.ListenAndServe(":8181", nil)
}