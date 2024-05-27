package main

import (
	"fmt"
)

var CurrenciesMap = make(map[string]ExchangeRate)
var CompaniesMap = make(map[int]Company)

func fetchCompanies(BasePath string, query string) ([]Company, error) {
	url := fmt.Sprintf(CommonEndpointFormat, BasePath, "/companies")
	if query != "" {
		url = fmt.Sprintf(CommonQueryFormat, url, query)
	}
	var companiesValues []Company
	err := FetchDataFromUrl(url, &companiesValues)
	if err != nil {
		return nil, err
	}
	fillMapCompanies(companiesValues)

	return companiesValues, nil
}

func FetchTransactionsFromUrl(rawQuery string, transactionEndpoint string) ([]Transaction, error) {
	transactionUrl := fmt.Sprintf(CommonEndpointFormat, BasePath, transactionEndpoint)

	if rawQuery != "" {
		transactionUrl = fmt.Sprintf(CommonQueryFormat, transactionUrl, rawQuery)
	}
	var transactions []Transaction
	err := FetchDataFromUrl(transactionUrl, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, err
}

func fetchExchangeRatesFromUrl() ([]ExchangeRate, error) {
	url := fmt.Sprintf(CommonEndpointFormat, BasePath, "/exchange-rates")
	var exchangeRatesValues []ExchangeRate
	err := FetchDataFromUrl(url, &exchangeRatesValues)
	if err != nil {
		return nil, err
	}
	fillExchangeRatesMap(exchangeRatesValues)
	return exchangeRatesValues, err
}
