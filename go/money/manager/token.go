package main

import (
	"fmt"

	plaid "github.com/plaid/plaid-go/v20/plaid"
)

type ITokenizer interface {
	getPublicToken() (string, error)
	getPrivateToken() (string, error)
}

type Tokenizer struct{}

type SandboxTokenizer struct {
	c *client
}
type ProductionTokenizer struct {
	_ *client
}

func (c *client) SetTokens() error {
	var tokenizer ITokenizer
	switch c.env {
	case plaid.Sandbox:
		tokenizer = &SandboxTokenizer{c: c}
	}

	publicToken, err := tokenizer.getPublicToken()
	if err != nil {
		return fmt.Errorf("couldn't retrieve public token: %e", err)
	}
	c.publicToken = publicToken

	privateToken, err := tokenizer.getPrivateToken()
	if err != nil {
		return fmt.Errorf("couldn't retrieve private token: %e", err)
	}
	c.token = privateToken

	return nil
}

func (s *SandboxTokenizer) getPublicToken() (string, error) {
	sandboxResp, _, err := s.c.client.PlaidApi.SandboxPublicTokenCreate(s.c.ctx).
		SandboxPublicTokenCreateRequest(*plaid.NewSandboxPublicTokenCreateRequest(
			"ins_109508", // Plaid's Sandbox Institution
			[]plaid.Products{plaid.PRODUCTS_TRANSACTIONS},
		)).Execute()

	if err != nil {
		return "", fmt.Errorf("unable to get public token: %v", err)
	}

	return sandboxResp.GetPublicToken(), nil
}

func (s *SandboxTokenizer) getPrivateToken() (string, error) {
	exchangeResp, _, err := s.c.client.PlaidApi.ItemPublicTokenExchange(s.c.ctx).
		ItemPublicTokenExchangeRequest(*plaid.NewItemPublicTokenExchangeRequest(s.c.publicToken)).
		Execute()

	if err != nil {
		return "", fmt.Errorf("failed to exchange public token: %v", err)
	}

	return exchangeResp.GetAccessToken(), nil
}

func (p *ProductionTokenizer) getPublicToken() (string, error) {
	return "", nil
}

func (p *ProductionTokenizer) getPrivateToken() (string, error) {
	return "", nil
}
