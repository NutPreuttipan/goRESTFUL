package models

type BankAccount struct {
	ID          int    `json:"id"`
	BankName    string `json:"bankName"`
	BankBranch  string `json:"bankBranch"`
	AccountNo   string `json:"accountNo"`
	AccountName string `json:"accountName"`
	AccountType string `json:"accountType"`
}
