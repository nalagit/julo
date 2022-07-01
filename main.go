package main

import (
	"fmt"
	"julo/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server Start")
	InitializeRouter()
}

func InitializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/regist", db.Registration).Methods("POST")
	r.HandleFunc("/api/v1/initialize", db.InitializeUser).Methods("POST")
	r.HandleFunc("/api/v1/validate", db.Validation).Methods("GET")
	r.HandleFunc("/api/v1/wallet", db.EnableWallet).Methods("POST")
	r.HandleFunc("/api/v1/wallet", db.DisableWallet).Methods("PATCH")
	r.HandleFunc("/api/v1/wallet", db.GetWallet).Methods("GET")
	r.HandleFunc("/api/v1/wallet/deposit", db.DepositWallet).Methods("POST")
	r.HandleFunc("/api/v1/wallet/withdraw", db.WithdrawWallet).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
