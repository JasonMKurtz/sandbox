package manager

import (
	"context"
	"fmt"
	"os"

	plaid "github.com/plaid/plaid-go/v20/plaid"
)

type Client struct {
	ctx         context.Context
	client      *plaid.APIClient
	linkToken   string
	publicToken string
	token       string
	env         plaid.Environment
}

func (c *Client) GetPublicToken() string {
	return c.publicToken
}

func (c *Client) SetPublicToken(token string) {
	c.publicToken = token
}

func (c *Client) SetPrivateToken(token string) {
	fmt.Printf("setting private token: %s\n", token)
	c.token = token
}

func NewClient(env plaid.Environment) *Client {
	c := &Client{}
	c.ctx = context.Background()
	c.env = env

	clientID := os.Getenv("PLAID_CLIENT_ID")
	secret := os.Getenv("PLAID_SECRET")

	cfg := plaid.NewConfiguration()
	cfg.AddDefaultHeader("PLAID-CLIENT-ID", clientID)
	cfg.AddDefaultHeader("PLAID-SECRET", secret)
	cfg.UseEnvironment(env)

	c.client = plaid.NewAPIClient(cfg)

	return c
}

func (c *Client) GetClient() *plaid.APIClient {
	return c.client
}

// environment: plaid.Sandbox, plaid.Production
func (c *Client) Init(env plaid.Environment) error {
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

func (c *Client) getAccountNames() (map[string]string, error) {
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
