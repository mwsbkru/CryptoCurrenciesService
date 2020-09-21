package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func CurrencyByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currencyCode := vars["code"]
	fmt.Fprint(w, "CurrencyByCode: " + currencyCode)
}

func CachedCurrencies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "CachedCurrencies")
}