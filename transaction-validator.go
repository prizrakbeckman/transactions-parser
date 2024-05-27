package main

import (
	"regexp"
)

const ibanPattern = `^[A-Z]{2}[0-9]{2}[A-Z0-9]{1,30}$`

func isValidCreditor(companyIbans string, transaction Transaction) bool {
	return isValidTransaction(companyIbans, transaction.Receiver) || isValidTransaction(companyIbans, transaction.Beneficiary)
}

func isValidDebitor(companyIbans string, transaction Transaction) bool {
	return isValidTransaction(companyIbans, transaction.Sender) || isValidTransaction(companyIbans, transaction.Payer)
}

func isValidTransaction(companyIbans string, iban string) bool {
	return isIBAN(iban) && companyIbans == iban
}

func isIBAN(input string) bool {
	match, _ := regexp.MatchString(ibanPattern, input)
	return match
}
