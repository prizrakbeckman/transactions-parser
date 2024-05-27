package main

func fillExchangeRatesMap(exchangeRatesValues []ExchangeRate) {
	for _, exchangeRate := range exchangeRatesValues {
		CurrenciesMap[exchangeRate.Currency] = exchangeRate
	}
}

func fillMapCompanies(companiesValues []Company) {
	for _, company := range companiesValues {
		CompaniesMap[company.ID] = company
	}
}
