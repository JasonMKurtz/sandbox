package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/plaid/plaid-go/v20/plaid"
)

type Client struct {
	api          *plaid.APIClient
	linkToken    string
	publicToken  string
	privateToken string
	ctx          context.Context
}

func NewClient() *Client {
	return &Client{api: NewAPIClient(), ctx: context.Background()}
}

func NewAPIClient() *plaid.APIClient {
	// Get your keys from https://dashboard.plaid.com/account/keys
	clientID := os.Getenv("PLAID_CLIENT_ID")
	secret := os.Getenv("PLAID_SECRET")

	cfg := plaid.NewConfiguration()
	cfg.AddDefaultHeader("PLAID-CLIENT-ID", clientID)
	cfg.AddDefaultHeader("PLAID-SECRET", secret)
	cfg.UseEnvironment(plaid.Production)

	return plaid.NewAPIClient(cfg)
}

func (c *Client) GetPublicToken() (string, error) {
	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: "jmk-123sasdfsafdsaewrwq7", // must be unique per user
	}

	req := plaid.NewLinkTokenCreateRequest(
		"Money Manager",
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
		user,
	)
	req.SetProducts([]plaid.Products{plaid.PRODUCTS_TRANSACTIONS})

	resp, _, err := c.api.PlaidApi.LinkTokenCreate(c.ctx).LinkTokenCreateRequest(*req).Execute()
	if err != nil {
		return "", fmt.Errorf("unable to create public token: %s", err)
	}

	c.linkToken = resp.GetLinkToken()
	return resp.GetLinkToken(), nil
}

func (c *Client) ExchangeToken(token string) (string, error) {
	exchangeResp, _, err := c.api.PlaidApi.ItemPublicTokenExchange(c.ctx).
		ItemPublicTokenExchangeRequest(*plaid.NewItemPublicTokenExchangeRequest(token)).
		Execute()

	if err != nil {
		return "", fmt.Errorf("failed to exchange tokens: %v", err)
	}

	c.publicToken = exchangeResp.GetAccessToken()
	return exchangeResp.GetAccessToken(), nil
}

type Transaction struct {
	AccountName string

	plaid.Transaction
}

func (c *Client) GetTransactions() ([]Transaction, error) {
	// STEP 3: Fetch transactions
	start := time.Now().AddDate(0, 0, -10).Format("2006-01-02")
	end := time.Now().Format("2006-01-02")

	txReq := plaid.NewTransactionsGetRequest(c.publicToken, start, end)

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
