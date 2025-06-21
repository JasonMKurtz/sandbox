package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var client *Client

func createTokenReq(w http.ResponseWriter, r *http.Request) {
	token, _ := client.GetPublicToken()
	json.NewEncoder(w).Encode(map[string]string{
		"link_token": token,
	})
}

func exchangeToken(w http.ResponseWriter, r *http.Request) {
	var body struct {
		PublicToken string `json:"public_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := client.ExchangeToken(body.PublicToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("can't decode token: %s", err), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"access_token": token,
	})
	fmt.Printf("exchanged token: %s\n", token)

}

func listTransactions(w http.ResponseWriter, r *http.Request) {
	txs, _ := client.GetTransactions()
	fmt.Printf("transactions: %#v\n", txs)
}

func main() {
	client = NewClient()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/bank-link.html")
	})

	http.HandleFunc("/api/create_link_token", createTokenReq)
	http.HandleFunc("/api/exchange_public_token", exchangeToken)
	http.HandleFunc("/api/transactions", listTransactions)

	log.Println("Serving on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
