package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"payment_gateway/model"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	login, err := model.CheckLogin(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.Header().Set("Token", login)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(login)
	}
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var transaction model.Transaction
	err := decoder.Decode(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := model.CreateTransaction(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func ChangeTransactionStatus(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var changeStatusRequest model.ChangeStatusRequest
	err := decoder.Decode(&changeStatusRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := model.ChangeTransactionStatus(&changeStatusRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := model.GetAllTransactions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		if transactions != nil {
			json.NewEncoder(w).Encode(transactions)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Not found")
			return
		}
	}
}
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	transaction, err := model.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if transaction.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Not found")
		return
	}

	json.NewEncoder(w).Encode(transaction)
}
func GetTransactionByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(param, 10, 64)
	transactions, err := model.GetTransactionByUserId(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		if transactions != nil {
			json.NewEncoder(w).Encode(transactions)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Not found")
			return
		}
	}
}
func GetTransactionByUserEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)["email"]
	transactions, err := model.GetTransactionByUserEmail(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		if transactions != nil {
			json.NewEncoder(w).Encode(transactions)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Not found")
			return
		}
	}
}
func CloseTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	response, err := model.CloseTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	status, err := model.GetTransactionStatus(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if status.Title == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Not found")
		return
	}

	json.NewEncoder(w).Encode(status)
}
