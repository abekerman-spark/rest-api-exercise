package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ExchangeRate Comment to avoid linter warning
type ExchangeRate struct {
	PurchaseRate float64
	SaleRate     float64
}

// ExchangeRates - global map with purchase/sale rates for different currencies
var ExchangeRates = make(map[string]ExchangeRate)

// SetupRouter - Exported function
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Router accepting currency input parameter as querystring or no querystring parameters at all
	r.GET("/exchangeRates", func(c *gin.Context) {
		// Uncomment code below and comment router above to accept currency input parameter within querystring
		currency := c.Request.URL.Query().Get("currency")
		if len(currency) == 0 {
			fmt.Println("No currency input parameter found -> Exchange rates requested for all currencies")
			c.JSON(http.StatusOK, ExchangeRates)
		} else {
			exchangeRateByCurrencyToResponse(ExchangeRates, currency, c)
		}
	})

	// Router accepting currency input parameter as a resource id URI type
	r.GET("/exchangeRates/:currency", func(c *gin.Context) {
		currency := c.Param("currency")
		exchangeRateByCurrencyToResponse(ExchangeRates, currency, c)
	})

	return r
}

func main() {
	SetupExchangeRates()
	r := SetupRouter()
	r.Run()
}

// SetupExchangeRates - global func
func SetupExchangeRates() {
	ExchangeRates["USD"] = ExchangeRate{1.001, 1.0002}
	ExchangeRates["EUR"] = ExchangeRate{1.003, 1.0004}
	ExchangeRates["JPY"] = ExchangeRate{1.005, 1.0006}
}

func exchangeRateByCurrencyToResponse(ExchangeRates map[string]ExchangeRate, currency string, c *gin.Context) {
	if exchangeRate, ok := ExchangeRates[currency]; ok {
		fmt.Println("Exchange rates requested for", currency)
		c.JSON(http.StatusOK, exchangeRate)
	} else {
		c.String(http.StatusNotFound, "Currency %v not found", currency)
	}
}
