package main

import (
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

func FindTransactionsByQuery(request *http.Request) []Transaction {
	var transactions []Transaction
	companyId := request.URL.Query().Get(CompanyIDParam)
	companyIbanParamValue := request.URL.Query().Get(CompanyIban)

	companyIdValue, err := strconv.Atoi(companyId)
	companyIban := parseIban(err, companyIdValue, companyIbanParamValue)

	afterTimestamp := url.QueryEscape(request.URL.Query().Get(AfterTimestamParam))

	transactionChan := make(chan []Transaction)

	var wg sync.WaitGroup

	fetch := func(params url.Values, endpoint string) {
		defer wg.Done()
		rawQuery := ParseParamsMap(params)
		currentTransactions, _ := FetchTransactionsFromUrl(rawQuery, endpoint)
		transactionChan <- currentTransactions
	}

	wg.Add(NumberOfNecessaryCallsForBalance)
	go fetch(CreateParams(PayerParam, companyIban, afterTimestamp), SepaEndpoint)
	go fetch(CreateParams(ReceiverParam, companyIban, afterTimestamp), SepaEndpoint)
	go fetch(CreateParams(SenderParam, companyIban, afterTimestamp), SwiftEndpoint)
	go fetch(CreateParams(BeneficiaryParam, companyIban, afterTimestamp), SwiftEndpoint)

	go func() {
		wg.Wait()
		close(transactionChan)
	}()

	for currentTransactions := range transactionChan {
		transactions = append(transactions, currentTransactions...)
	}

	return transactions
}
