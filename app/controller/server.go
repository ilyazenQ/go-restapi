package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	controller "payment_gateway/controller/transaction"
	"payment_gateway/middleware"
)

var router *mux.Router

func initHandlers() {
	router.HandleFunc("/api/transactions", controller.GetAllTransactions).Methods("GET")
	router.HandleFunc("/api/transactions/status/{id}", controller.GetStatus).Methods("GET")
	router.HandleFunc("/api/transactions/{id}", controller.GetTransaction).Methods("GET")
	router.HandleFunc("/api/transactions/user-id/{id}", controller.GetTransactionByUserId).Methods("GET")
	router.HandleFunc("/api/transactions/user-email/{email}", controller.GetTransactionByUserEmail).Methods("GET")

	router.HandleFunc("/api/transactions/close/{id}", controller.CloseTransaction).Methods("PUT")
	router.Handle("/api/transactions/change-status", middleware.IsAuthorized(controller.ChangeTransactionStatus)).Methods("PUT")

	router.HandleFunc("/api/transactions/new", controller.CreateTransaction).Methods("POST")

	router.HandleFunc("/api/login", controller.Login).Methods("POST")
}

func Start() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
