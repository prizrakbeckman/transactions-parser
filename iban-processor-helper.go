package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func parseIban(err error, companyIdValue int, companyIbanParamValue string) string {
	if companyIbanParamValue != "" {
		return companyIbanParamValue
	}
	if err != nil {
		fmt.Println(err)
	}
	company, exists := CompaniesMap[companyIdValue]
	if !exists {
		return ""
	}

	return extractIban(company.Ibans.(string))
}

func extractIban(str string) string {

	str = strings.ReplaceAll(str, "'", "\"")
	var ibans []string
	err := json.Unmarshal([]byte(str), &ibans)
	if err != nil {
		log.Fatal(err)
	}

	iban := ibans[0]
	return iban
}
