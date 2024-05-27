package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		companies, _ = fetchCompanies(BasePath, "")
		fmt.Println("companies fetches is ", len(companies))
	}()

	go func() {
		defer wg.Done()
		exchangeRates, _ := fetchExchangeRatesFromUrl()
		fmt.Println("currencies fetches is ", len(exchangeRates))
	}()

	wg.Wait()

	http.HandleFunc("/company/transactions", getCompanyTransactions)
	http.HandleFunc("/company/balance", getCompanyBalance)
	http.HandleFunc("/companies", getCompanies)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		return
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
