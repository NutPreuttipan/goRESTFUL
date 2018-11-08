package bankAccountRepository

import (
	"../../models"
	"database/sql"
	"log"
)

type BankAccountRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (b BankAccountRepository) GetBanks(db *sql.DB, bankAccount models.BankAccount, bankAccounts []models.BankAccount) []models.BankAccount {

	rows, err := db.Query("select * from \"TBBankAccount\" order by id")
		logFatal(err)
		defer rows.Close()
	
		for rows.Next() {
			err := rows.Scan(&bankAccount.ID, 
				&bankAccount.BankName, 
				&bankAccount.BankBranch,
				&bankAccount.AccountNo, 
				&bankAccount.AccountName,
				&bankAccount.AccountType, 
			)
			logFatal(err)
			
			bankAccounts = append(bankAccounts, bankAccount)
		}

		return bankAccounts
}

func (b BankAccountRepository) GetBank(db *sql.DB, bankAccount models.BankAccount, id int) models.BankAccount {
	rows := db.QueryRow("select * from \"TBBankAccount\" where id = $1", id)
	
		err := rows.Scan(
			&bankAccount.ID, 
			&bankAccount.BankName, 
			&bankAccount.BankBranch, 
			&bankAccount.AccountNo, 
			&bankAccount.AccountName,
			&bankAccount.AccountType,
		)
		logFatal(err)
		
	return bankAccount
}

func (b BankAccountRepository) AddBank(db *sql.DB, bankAccount models.BankAccount) int {
		
	err := db.QueryRow("insert into \"TBBankAccount\" (\"bankName\",\"bankBranch\",\"accountNo\",\"accountName\",\"accountType\") values ($1, $2, $3 , $4, $5) RETURNING id;",
			bankAccount.BankName, bankAccount.BankBranch, bankAccount.AccountNo, bankAccount.AccountName, bankAccount.AccountType).Scan(&bankAccount.ID)
	logFatal(err)

	return bankAccount.ID
}

func (b BankAccountRepository) UpdateBank(db *sql.DB, bankAccount models.BankAccount) int64 {
	result, err := db.Exec("update \"TBBankAccount\" set \"bankName\"=$1, \"bankBranch\"=$2, \"accountNo\"=$3, \"accountName\"=$4, \"accountType\"=$5 where id=$6",
		bankAccount.BankName, bankAccount.BankBranch, bankAccount.AccountNo, bankAccount.AccountName, bankAccount.AccountType, bankAccount.ID)
	
	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	return rowsUpdated
}

func (b BankAccountRepository) DeleteBank(db *sql.DB, id int) int64 {
	result, err := db.Exec("delete from \"TBBankAccount\" where id=$1", id)
		logFatal(err)
	
		rowsDeleted, err := result.RowsAffected()
		logFatal(err)

	return rowsDeleted
}	
