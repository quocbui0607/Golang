package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"frontendmasters.com/go/crypto/api"
	"frontendmasters.com/go/crypto/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from a Go program"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	dir, _ := os.Getwd()
	html, err := template.ParseFiles(dir + "./templates/index.tmpl")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	html.Execute(w, data.GetAll())
}

func main() {
	// var wg sync.WaitGroup
	// currencies := []string{"BTC", "ETH", "HIGH"}
	// for _, currency := range currencies {
	// 	wg.Add(1)
	// 	go func(currency string) {
	// 		getCurrencyData(currency)
	// 		wg.Done()
	// 	}(currency)
	// }
	// wg.Wait()
	dir, _ := os.Getwd()

	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	fs := http.FileServer(http.Dir(dir + "./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)
	if err == nil {

	}
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %s \n", rate.Currency, rate.Price)
	}
}
