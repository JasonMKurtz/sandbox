package main

import (
	"fmt"
	"log"
	"time"

	"github.com/plaid/plaid-go/v20/plaid"
)

type Manager struct {
	apiclient *client
	accounts  map[string]string
}

func NewManager(env plaid.Environment) (*Manager, error) {
	m := &Manager{}

	m.apiclient = NewClient()
	if err := m.apiclient.Init(env); err != nil {
		return nil, fmt.Errorf("error creating manager: %e", err)
	}

	names, err := m.apiclient.getAccountNames()
	if err != nil {
		return nil, err
	}
	m.accounts = names

	return m, nil
}

func main() {
	m, err := NewManager(plaid.Sandbox)
	if err != nil {
		log.Fatalf("error initializing money manager: %e", err)
	}

	time.Sleep(4 * time.Second)

	transactions, err := m.GetTransactions()
	if err != nil {
		log.Fatalf("error reading transactions: %e", err)
	}

	for _, tx := range transactions {
		fmt.Printf("%s\n", tx)
	}
}
