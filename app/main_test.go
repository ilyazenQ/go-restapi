package main

import (
	"github.com/gorilla/mux"
	"net/http"
	controller "payment_gateway/controller/transaction"
	"testing"
)

func Test_GetAllTransactions(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/transactions", controller.GetAllTransactions).Methods("GET")

	res, err := http.Get("http://localhost:3200/" + "api/transactions")

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}

func Test_GetTransaction(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/transactions/{id}", controller.GetTransaction).Methods("GET")

	t.Run("not found", func(t *testing.T) {
		res, err := http.Get("http://localhost:3200/" + "api/transactions/1000000")
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusNotFound {
			t.Errorf("Expected %d, received %d", http.StatusNotFound, res.StatusCode)
		}
	})

	t.Run("found", func(t *testing.T) {
		res, err := http.Get("http://localhost:3200/" + "api/transactions/1")
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
		}
	})
}

func Test_GetStatus(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/transactions/status/{id}", controller.GetStatus).Methods("GET")

	t.Run("not found", func(t *testing.T) {
		res, err := http.Get("http://localhost:3200/" + "api/transactions/status/1000000")
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusNotFound {
			t.Errorf("Expected %d, received %d", http.StatusNotFound, res.StatusCode)
		}
	})

	t.Run("found", func(t *testing.T) {
		res, err := http.Get("http://localhost:3200/" + "api/transactions/status/1")
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
		}
	})
}

func Test_GetTransactionByUserId(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/transactions/user-id/{id}", controller.GetStatus).Methods("GET")

	t.Run("not found", func(t *testing.T) {
		res, err := http.Get("http://localhost:3200/" + "/api/transactions/user-id/1000000")
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusNotFound {
			t.Errorf("Expected %d, received %d", http.StatusNotFound, res.StatusCode)
		}
	})

	t.Run("found", func(t *testing.T) {
		res, err := http.Get("http://localhost:3200/" + "api/transactions/user-id/1")
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
		}
	})
}

func Test_GetTransactionByEmail(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/transactions/user-email/{email}", controller.GetStatus).Methods("GET")

	t.Run("not found", func(t *testing.T) {
		res, err := http.Get("http://localhost:3200/" + "/api/transactions/user-email/asdasdasd")
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusNotFound {
			t.Errorf("Expected %d, received %d", http.StatusNotFound, res.StatusCode)
		}
	})

	t.Run("found", func(t *testing.T) {
		res, err := http.Get("http://localhost:3200/" + "api/transactions/user-email/email@email.com")
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
		}
	})
}
