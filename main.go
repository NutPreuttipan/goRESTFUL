package main

import (
	"log"
	"github.com/subosito/gotenv"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"database/sql"
	"./models"
	"./driver"
	"./controllers"
)

var bankAccounts []models.BankAccount
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/bank/{id}", controller.GetBank(db)).Methods("GET")
	router.HandleFunc("/bank", controller.GetBanks(db)).Methods("GET")
	router.HandleFunc("/bank", controller.AddBank(db)).Methods("POST")
	router.HandleFunc("/bank", controller.UpdateBank(db)).Methods("PUT")
	router.HandleFunc("/bank/{id}", controller.DeleteBank(db)).Methods("DELETE")
	router.HandleFunc("/upload", controller.Upload(db)).Methods("POST")

	fmt.Println("Starting....")

	log.Fatal(http.ListenAndServe(":8080",router))
}

// func getBanks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type","application/json")

// 	var bankAccount models.BankAccount
// 	bankAccounts = []models.BankAccount{}

// 	rows, err := db.Query("select * from \"TBBankAccount\"")
// 	logFatal(err)
// 	defer rows.Close()

// 	for rows.Next() {
// 		err := rows.Scan(&bankAccount.ID, 
// 			&bankAccount.BankName, 
// 			&bankAccount.BankBranch,
// 			&bankAccount.AccountNo, 
// 			&bankAccount.AccountName,
// 			&bankAccount.AccountType, 
// 		)
// 		logFatal(err)
		
// 		bankAccounts = append(bankAccounts, bankAccount)
// 	}

// 	json.NewEncoder(w).Encode(bankAccounts)
// }

// func getBank(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type","application/json")
// 	var bankAccount models.BankAccount
// 	param := mux.Vars(r)
// 	id := param["id"]

// 	rows := db.QueryRow("select * from \"TBBankAccount\" where id = $1", id)

// 	err := rows.Scan(&bankAccount.ID, 
// 		&bankAccount.BankName, 
// 		&bankAccount.BankBranch, 
// 		&bankAccount.AccountNo, 
// 		&bankAccount.AccountName,
// 		&bankAccount.AccountType,
// 	)
// 	logFatal(err)

// 	json.NewEncoder(w).Encode(bankAccount)
// }

// func addBank(w http.ResponseWriter, r *http.Request) {	
// 	w.Header().Set("Content-type","application/json")
// 	var bankAccount models.BankAccount
// 	var bookID int 

// 	json.NewDecoder(r.Body).Decode(&bankAccount)

// 	err := db.QueryRow("insert into \"TBBankAccount\" (\"bankName\",\"bankBranch\",\"accountNo\",\"accountName\",\"accountType\") values ($1, $2, $3 , $4, $5) RETURNING id;",
// 		bankAccount.BankName, bankAccount.BankBranch, bankAccount.AccountNo, bankAccount.AccountName, bankAccount.AccountType).Scan(&bookID)
// 	logFatal(err)

// 	json.NewEncoder(w).Encode(bookID)
// }

// func updateBank(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type","application/json")
// 	var bankAccount models.BankAccount
// 	var response apiResponse

// 	json.NewDecoder(r.Body).Decode(&bankAccount)

// 	result, err := db.Exec("update \"TBBankAccount\" set \"bankName\"=$1, \"bankBranch\"=$2, \"accountNo\"=$3, \"accountName\"=$4, \"accountType\"=$5 where id=$6",
// 	bankAccount.BankName, bankAccount.BankBranch, bankAccount.AccountNo, bankAccount.AccountName, bankAccount.AccountType, bankAccount.ID)

// 	rowsUpdated, err := result.RowsAffected()
// 	logFatal(err)

// 	response.ID = rowsUpdated
// 	response.Status = "Success"
// 	json.NewEncoder(w).Encode(response)
// }

// func deleteBank(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type","application/json")
// 	var response apiResponse
// 	param := mux.Vars(r)
// 	id := param["id"]

// 	result, err := db.Exec("delete from \"TBBankAccount\" where id=$1", id)
// 	logFatal(err)

// 	rowsDeleted, err := result.RowsAffected()
// 	logFatal(err)

// 	response.ID = rowsDeleted
// 	response.Status = "Success"
// 	json.NewEncoder(w).Encode(response)

// }
