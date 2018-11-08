package controllers

import (
	"strconv"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"../models"
	"../repository/BankAccount"
	"github.com/gorilla/mux"
	
)

type Controller struct{}

var bankAccounts []models.BankAccount

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetBanks(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type","application/json")
	
		var bankAccount models.BankAccount
		bankAccounts = []models.BankAccount{}
		bankAccountRepo := bankAccountRepository.BankAccountRepository{}

		bankAccounts = bankAccountRepo.GetBanks(db,bankAccount,bankAccounts)
		
		json.NewEncoder(w).Encode(bankAccounts)
	}
}

func (c Controller) GetBank(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type","application/json")

		var bankAccount models.BankAccount
		param := mux.Vars(r)
		id, err := strconv.Atoi(param["id"])
		logFatal(err)
		bankAccountRepo := bankAccountRepository.BankAccountRepository{}

		bankAccount = bankAccountRepo.GetBank(db, bankAccount, id)
	
		json.NewEncoder(w).Encode(bankAccount)
	}
}

func (c Controller) AddBank(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {	
		w.Header().Set("Content-type","application/json")

		var bankAccount models.BankAccount
		var bookID int 
		json.NewDecoder(r.Body).Decode(&bankAccount)
		bankAccountRepo := bankAccountRepository.BankAccountRepository{}

		bookID = bankAccountRepo.AddBank(db, bankAccount)

		json.NewEncoder(w).Encode(bookID)
	}
}

func (c Controller) UpdateBank(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type","application/json")
		var bankAccount models.BankAccount
		// var response apiResponse
		json.NewDecoder(r.Body).Decode(&bankAccount)
		bankAccountRepo := bankAccountRepository.BankAccountRepository{}

		rowsUpdated := bankAccountRepo.UpdateBank(db, bankAccount)
	
		// response.ID = rowsUpdated
		// response.Status = "Success"
		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) DeleteBank(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type","application/json")
		// var response apiResponse
		param := mux.Vars(r)
		id, err := strconv.Atoi(param["id"])
		logFatal(err)
		bankAccountRepo := bankAccountRepository.BankAccountRepository{}

		rowsDeleted := bankAccountRepo.DeleteBank(db, id)
	
		// response.ID = rowsDeleted
		// response.Status = "Success"
		json.NewEncoder(w).Encode(rowsDeleted)
	}
}
	







