package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/plaid/plaid-go/v20/plaid"
)

type Client struct {
	api         *plaid.APIClient
	ctx         context.Context
	clientId    string
	secret      string
	accessToken string
}

func NewClient(clientId string, secret string, accessToken string) *Client {
	apiClient := newAPIClient(clientId, secret)

	return &Client{
		api:         apiClient,
		ctx:         context.Background(),
		clientId:    clientId,
		secret:      secret,
		accessToken: accessToken,
	}
}

func newAPIClient(clientId string, secret string) *plaid.APIClient {
	// Get your keys from https://dashboard.plaid.com/account/keys

	cfg := plaid.NewConfiguration()
	cfg.AddDefaultHeader("PLAID-CLIENT-ID", clientId)
	cfg.AddDefaultHeader("PLAID-SECRET", secret)
	cfg.UseEnvironment(plaid.Production)

	return plaid.NewAPIClient(cfg)
}

type Transaction struct {
	AccountName string

	plaid.Transaction
}

func (t Transaction) String() string {
	/*
		A negative amount means INCOME and a positive amount means EXPENSE. WHY???
		There isn't always a counterparty. When there is `t.GetCounterparties()[0].Name` will suffice.
	*/

	var cp string
	if _, ok := t.GetCounterpartiesOk(); ok {
		cp = t.GetCounterparties()[0].Name
	} else {
		cp = ""
	}

	return fmt.Sprintf(
		"Account: %s, Date: %s, Name: %s, Amount: %.2f, Counterparty: %s\n",
		t.AccountName, t.Date, t.Name, t.Amount, cp,
	)
}

func (c *Client) GetTransactions() ([]Transaction, error) {
	// STEP 3: Fetch transactions
	start := time.Now().AddDate(0, 0, -10).Format("2006-01-02")
	end := time.Now().Format("2006-01-02")

	txReq := plaid.NewTransactionsGetRequest(c.accessToken, start, end)

	txResp, _, err := c.api.PlaidApi.TransactionsGet(c.ctx).
		TransactionsGetRequest(*txReq).
		Execute()

	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions: %v", err)
	}

	var txs []Transaction

	for _, tx := range txResp.Transactions {
		t := Transaction{Transaction: tx}
		txs = append(txs, t)
	}

	return txs, nil
}

func main() {
	c := NewClient(
		os.Getenv("CLIENT_ID"),
		os.Getenv("SECRET"),
		os.Getenv("ACCESS_TOKEN"),
	)

	tks, _ := c.GetTransactions()
	fmt.Printf("%s\n", tks[0:5])
}
