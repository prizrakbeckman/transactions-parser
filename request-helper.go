package main

import (
	"log"
	"net/url"
	"strings"
)

func ParseParamsMap(paramsMap url.Values) string {
	var rawQuery = ""
	if len(paramsMap) != 0 {
		rawQuery = paramsMap.Encode()
	}
	return rawQuery
}

func validateParam(parameter string, parameterValue string, receiverParams url.Values) {
	if strings.TrimSpace(parameterValue) != "" {
		decodedValue, err := url.QueryUnescape(parameterValue)
		if err != nil {
			log.Fatal(err)
		}
		receiverParams.Add(parameter, decodedValue)
	}
}

func CreateParams(param, companyIban, afterTimestamp string) url.Values {
	params := url.Values{}
	validateParam(param, companyIban, params)
	validateParam(AfterTimestamParam, afterTimestamp, params)
	return params
}
