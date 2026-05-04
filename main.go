package main

import (
	"fmt"
	"net/http"
	"sync"

	"frontendmasters.com/go/crypto/api"
)

func main() {
	var wg sync.WaitGroup
	currencies := []string{"BTC", "ETH", "HIGH"}
	for _, currency := range currencies {
		wg.Add(1)
		go func(currency string) {
			getCurrencyData(currency)
			wg.Done()
		}(currency)
	}
	wg.Wait()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from a Go program"))
	})

	err := http.ListenAndServe(":3333", nil)
	if err == nil {

	}
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %s \n", rate.Currency, rate.Price)
	}
}
