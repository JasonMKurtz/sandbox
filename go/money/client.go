package main

import (
	"context"
	"fmt"
	"os"

	plaid "github.com/plaid/plaid-go/v20/plaid"
)

type client struct {
	ctx         context.Context
	client      *plaid.APIClient
	publicToken string
	token       string
}

func NewClient() *client {
	c := &client{}
	c.ctx = context.Background()
	return c
}

func (c *client) GetClient() *plaid.APIClient {
	return c.client
}

// environment: plaid.Sandbox, plaid.Production
func (c *client) Init(env plaid.Environment) error {
	// Get your keys from https://dashboard.plaid.com/account/keys
	clientID := os.Getenv("PLAID_CLIENT_ID")
	secret := os.Getenv("PLAID_SECRET")

	cfg := plaid.NewConfiguration()
	cfg.AddDefaultHeader("PLAID-CLIENT-ID", clientID)
	cfg.AddDefaultHeader("PLAID-SECRET", secret)
	cfg.UseEnvironment(env)

	c.client = plaid.NewAPIClient(cfg)

	publicToken, err := c.getPublicToken()
	if err != nil {
		return err
	}
	c.publicToken = publicToken

	token, err := c.getAccessToken()
	if err != nil {
		return err
	}
	c.token = token

	return nil
}

func (c *client) getPublicToken() (string, error) {
	sandboxResp, _, err := c.GetClient().PlaidApi.SandboxPublicTokenCreate(c.ctx).
		SandboxPublicTokenCreateRequest(*plaid.NewSandboxPublicTokenCreateRequest(
			"ins_109508", // Plaid's Sandbox Institution
			[]plaid.Products{plaid.PRODUCTS_TRANSACTIONS},
		)).Execute()

	if err != nil {
		return "", fmt.Errorf("unable to get public token: %v", err)
	}

	return sandboxResp.GetPublicToken(), nil
}

func (c *client) getAccessToken() (string, error) {
	exchangeResp, _, err := c.GetClient().PlaidApi.ItemPublicTokenExchange(c.ctx).
		ItemPublicTokenExchangeRequest(*plaid.NewItemPublicTokenExchangeRequest(c.publicToken)).
		Execute()

	if err != nil {
		return "", fmt.Errorf("failed to exchange public token: %v", err)
	}

	return exchangeResp.GetAccessToken(), nil
}

func (c *client) getAccountNames() (map[string]string, error) {
	resp, _, err := c.GetClient().PlaidApi.AccountsGet(c.ctx).
		AccountsGetRequest(*plaid.NewAccountsGetRequest(c.token)).
		Execute()

	if err != nil {
		return nil, fmt.Errorf("error retrieving accounts: %e", err)
	}
	accounts := resp.GetAccounts()
	accountMap := make(map[string]string, len(accounts))

	for _, account := range accounts {
		accountMap[account.AccountId] = account.Name
	}

	return accountMap, nil
}
