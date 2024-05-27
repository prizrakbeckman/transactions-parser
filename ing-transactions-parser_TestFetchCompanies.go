package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchCompanies(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(response()))
	}))
	defer server.Close()

	companies, err := fetchCompanies(server.URL, "")
	if err != nil {
		t.Errorf("Error was not expected while fetching companies: %s", err)
	}

	if len(companies) != 2 {
		t.Errorf("Expected 2 companies, got %d", len(companies))
	}
}

func response() string {
	return `[{
  "id": 1,
  "ibans": "['RO97UOEW8394440003813323']",
  "name": "Bouvet SARL",
  "address": "7, boulevard de Alexandre\n80687 Sainte Jeannine-les-Bains"
 },
 {
  "id": 2,
  "ibans": "['NO5997936829634']",
  "name": "Mahe Chrétien S.A.S.",
  "address": "Stefandreef 7\n5685BS\nVreeland"
 }]`
}
