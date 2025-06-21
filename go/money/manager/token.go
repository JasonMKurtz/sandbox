package manager

import (
	"fmt"

	plaid "github.com/plaid/plaid-go/v20/plaid"
)

type ITokenizer interface {
	GetPublicToken() (string, error)
}

type Tokenizer struct{}

type SandboxTokenizer struct {
	c *Client
}
type ProductionTokenizer struct {
	c *Client
}

func (c *Client) NewTokenizer() ITokenizer {
	var tokenizer ITokenizer
	switch c.env {
	case plaid.Sandbox:
		tokenizer = &SandboxTokenizer{c: c}
	case plaid.Development:
		tokenizer = &ProductionTokenizer{c: c}
	case plaid.Production:
		tokenizer = &ProductionTokenizer{c: c}
	}

	return tokenizer
}

func (c *Client) GetPrivateToken() (string, error) {
	return c.ExchangePrivateTokenFrom(c.publicToken)
}

func (c *Client) ExchangePrivateTokenFrom(token string) (string, error) {
	exchangeResp, _, err := c.client.PlaidApi.ItemPublicTokenExchange(c.ctx).
		ItemPublicTokenExchangeRequest(*plaid.NewItemPublicTokenExchangeRequest(c.publicToken)).
		Execute()

	if err != nil {
		return "", fmt.Errorf("failed to exchange tokens: %v", err)
	}

	return exchangeResp.GetAccessToken(), nil
}

func (c *Client) SetTokens() error {
	publicToken, err := c.NewTokenizer().GetPublicToken()
	if err != nil {
		return fmt.Errorf("couldn't retrieve public token: %s", err)
	}

	switch c.env {
	case plaid.Sandbox:
		c.publicToken = publicToken
	case plaid.Development:
		c.linkToken = publicToken
	case plaid.Production:
		c.linkToken = publicToken
	}

	privateToken, err := c.GetPrivateToken()
	if err != nil {
		return fmt.Errorf("couldn't retrieve private token: %s", err)
	}
	c.token = privateToken

	return nil
}

func (s *SandboxTokenizer) GetPublicToken() (string, error) {
	sandboxResp, _, err := s.c.client.PlaidApi.SandboxPublicTokenCreate(s.c.ctx).
		SandboxPublicTokenCreateRequest(*plaid.NewSandboxPublicTokenCreateRequest(
			"ins_109508", // Plaid's Sandbox Institution
			[]plaid.Products{plaid.PRODUCTS_TRANSACTIONS},
		)).Execute()

	if err != nil {
		return "", fmt.Errorf("unable to get public token: %s", err)
	}

	return sandboxResp.GetPublicToken(), nil
}

func (p *ProductionTokenizer) GetPublicToken() (string, error) {
	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: "jmk-1237", // must be unique per user
	}

	req := plaid.NewLinkTokenCreateRequest(
		"Money Manager",
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
		user,
	)
	req.SetProducts([]plaid.Products{plaid.PRODUCTS_TRANSACTIONS})

	resp, _, err := p.c.GetClient().PlaidApi.LinkTokenCreate(p.c.ctx).LinkTokenCreateRequest(*req).Execute()
	if err != nil {
		return "", fmt.Errorf("unable to create public token: %s", err)
	}

	fmt.Printf("public token: %s\n", resp.GetLinkToken())

	return resp.GetLinkToken(), nil
}
