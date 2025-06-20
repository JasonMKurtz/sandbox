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
	env         plaid.Environment
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

	c.env = env

	c.client = plaid.NewAPIClient(cfg)
	if err := c.SetTokens(); err != nil {
		return err
	}

	return nil
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
