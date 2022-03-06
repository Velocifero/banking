package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer", addNewCustomer).Methods(http.MethodPost)
	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
