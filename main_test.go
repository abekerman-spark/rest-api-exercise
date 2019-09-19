package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

// Test #1 - Test all exchange rates are returned
func TestAllExchangeRates(t *testing.T) {

	SetupExchangeRates()

	r := SetupRouter()

	w := performRequest(r, "GET", "/exchangeRates")

	assert.Equal(t, 200, w.Code)
	var response map[string]ExchangeRate
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	// Assert key count is 3 (USD, EUR, JPY)
	assert.Equal(t, 3, len(response))
}

// Test #2 - Test existing currency as querystring param
func TestExchangeRateEURQuerystring(t *testing.T) {

	SetupExchangeRates()

	r := SetupRouter()

	w := performRequest(r, "GET", "/exchangeRates?currency=EUR")

	assert.Equal(t, 200, w.Code)
	var response ExchangeRate
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
}

// Test #3 - Test non-existing currency as querystring param
func TestNonExistingExchangeRate(t *testing.T) {

	SetupExchangeRates()

	r := SetupRouter()

	w := performRequest(r, "GET", "/exchangeRates?currency=XYZ")

	assert.Equal(t, 404, w.Code)

	ct := []string{"text/plain; charset=utf-8"}
	assert.Equal(t, ct, w.HeaderMap["Content-Type"])
}

// Test #4 - Test existing currency as URI
func TestExchangeRateEURURI(t *testing.T) {

	SetupExchangeRates()

	r := SetupRouter()

	w := performRequest(r, "GET", "/exchangeRates/EUR")

	assert.Equal(t, 200, w.Code)
	var response ExchangeRate
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
}

// Test #5 - Test non-existing currency as URI
func TestNonExistingExchangeRateURI(t *testing.T) {

	SetupExchangeRates()

	r := SetupRouter()

	w := performRequest(r, "GET", "/exchangeRates/XYZ")

	assert.Equal(t, 404, w.Code)

	ct := []string{"text/plain; charset=utf-8"}
	assert.Equal(t, ct, w.HeaderMap["Content-Type"])
}
