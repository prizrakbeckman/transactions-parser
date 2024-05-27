package main

import "time"

type CompanyStatus struct {
	ID      string `json:"id"`
	Balance string `json:"balance"`
}

type Company struct {
	ID      int         `json:"id"`
	Ibans   interface{} `json:"ibans"`
	Name    string      `json:"name"`
	Address string      `json:"address"`
}

type Transaction struct {
	ID          string    `json:"id"`
	Payer       string    `json:"payer"`
	Receiver    string    `json:"receiver"`
	Sender      string    `json:"sender"`
	Beneficiary string    `json:"beneficiary"`
	Amount      float64   `json:"amount"`
	Currency    string    `json:"currency"`
	Timestamp   time.Time `json:"timestamp"`
}

type ExchangeRate struct {
	Currency string  `json:"currency"`
	UsdRate  float64 `json:"usd_rate"`
	EurRate  float64 `json:"eur_rate"`
}
