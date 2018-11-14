package controllers

import (
	"text/template"
	"crypto/md5"
	"time"
	"fmt"
	"strings"
	"os"
	"io"
	"strconv"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"../models"
	"../repository/BankAccount"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
	
)

type Controller struct{}

var bankAccounts []models.BankAccount

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func after(value string, a string) string {
    // Get substring after a string.
    pos := strings.LastIndex(value, a)
    if pos == -1 {
        return ""
    }
    adjustedPos := pos + len(a)
    if adjustedPos >= len(value) {
        return ""
    }
    return value[adjustedPos:len(value)]
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

func (c Controller) Upload(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type","application/json")
		var imageName models.Upload

		fmt.Println("method:", r.Method)
		if r.Method == "GET" {
			crutime := time.Now().Unix()
			h := md5.New()
			io.WriteString(h, strconv.FormatInt(crutime, 10))
			token := fmt.Sprintf("%x", h.Sum(nil))
	
			t, _ := template.ParseFiles("upload.gtpl")
			t.Execute(w, token)
		} else {
			r.ParseMultipartForm(32 << 20)
			file, handler, err := r.FormFile("file")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			fmt.Fprintf(w, "%v", handler.Header)

			guid := ksuid.New()
			fileType := handler.Filename

			f, err := os.OpenFile("uploadfile/" + guid.String() + "." + after(fileType, "."), os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			io.Copy(f, file)
			
			name := f.Name()
			imageName.ImagePath = after(name,"/")

			json.NewEncoder(w).Encode(imageName)
		}
	}
}
	







