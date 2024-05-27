package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func fetchCompaniesV2(afterID string) ([]Company, error) {
	url := fmt.Sprintf(CommonEndpointFormat, BasePath, "/companies")

	resp, err := http.Get(fmt.Sprintf(url+"?after-id=%s", afterID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var companies []Company
	err = json.Unmarshal(body, &companies)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func worker(afterID string) {
	defer wg.Done()

	comps, err := fetchCompanies(BasePath, afterID)
	if err != nil {
		fmt.Println("Error fetching companies:", err)
		return
	}

	mutex.Lock()
	companies = append(companies, comps...)
	mutex.Unlock()

	if len(comps) == 500 {
		wg.Add(1)
		go worker(fmt.Sprintf("%d", 500+(500*(len(companies)/500-1))))
	}
}
