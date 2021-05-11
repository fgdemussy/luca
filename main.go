package main

import (
	"log"
	"net/http"

	"github.com/fgdemussy/luca/data"
	"github.com/fgdemussy/luca/ledger"
	"github.com/gorilla/mux"
)

var accounts []ledger.Account

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/accounts", getAccounts).Methods(http.MethodGet)
	// r.HandleFunc("/accounts", createAccount).Methods(http.MethodPost)
	r.HandleFunc("/accounts/{id}", getAccount).Methods(http.MethodGet)
	// r.HandleFunc("/accounts/{id}", updateAccount).Methods(http.MethodPut)
	// r.HandleFunc("/accounts/{id}", deleteAccount).Methods(http.MethodDelete)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	data.ToJSON(accounts, w)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range accounts {
		if item.ID == params["id"] {
			data.ToJSON(item, w)
			break
		}
	}
	data.ToJSON(&ledger.Account{}, w)
}
