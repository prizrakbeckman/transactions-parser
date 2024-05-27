package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var companies []Company

func getCompanyBalance(w http.ResponseWriter, r *http.Request) {
	companyID := r.URL.Query().Get("company-id")
	companyIdIntValue, err := strconv.Atoi(companyID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	currentTransactions := FindTransactionsByQuery(r)
	balance := calculateCompanyBalance(currentTransactions, companyIdIntValue)
	balanceValue := fmt.Sprintf("%f", balance)

	companyStatus := CompanyStatus{
		ID:      companyID,
		Balance: balanceValue,
	}
	jsonData, err := json.Marshal(companyStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(ContentType, JsonContentType)
	w.Write(jsonData)
}

func calculateCompanyBalance(currentTransactions []Transaction, companyID int) float64 {
	balance := 0.0
	if currentTransactions == nil {
		return balance
	}
	companyIbans, ok := CompaniesMap[companyID].Ibans.(string)
	companyIbanCurrent := extractIban(companyIbans)
	if !ok {
		log.Fatal("Could not convert interface{} to []string")
	}
	for _, transaction := range currentTransactions {
		exchangeRate, exists := CurrenciesMap[transaction.Currency]
		if exists {
			rate := exchangeRate.EurRate
			if isValidDebitor(companyIbanCurrent, transaction) {
				balance -= transaction.Amount * rate
			} else if isValidCreditor(companyIbanCurrent, transaction) {
				balance += transaction.Amount * rate
			}
		}
	}
	return balance
}

func getCompanyTransactions(w http.ResponseWriter, r *http.Request) {
	currentTransactions := FindTransactionsByQuery(r)
	fmt.Println("length of transaction", len(currentTransactions))

	// Assume transactions is a global variable
	jsonData, err := json.Marshal(currentTransactions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(ContentType, JsonContentType)
	w.Write(jsonData)
}

func getCompanies(w http.ResponseWriter, r *http.Request) {
	companies, _ = fetchCompanies(BasePath, "")
	jsonData, err := json.Marshal(companies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(ContentType, JsonContentType)
	w.Write(jsonData)
}
