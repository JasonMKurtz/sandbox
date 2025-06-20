package main

import (
	"fmt"
	"time"

	plaid "github.com/plaid/plaid-go/v20/plaid"
)

type Transaction struct {
	AccountName string

	plaid.Transaction
}

func defaultString(maybeStr plaid.NullableString) string {
	if maybeStr.IsSet() {
		return *maybeStr.Get()
	}

	return ""
}

func (t Transaction) String() string {
	return fmt.Sprintf(
		"Account: %s, Date: %s, Merchant: %s, Amount: %.2f",
		t.AccountName, t.Date, defaultString(t.MerchantName), t.Amount,
	)
}

func (m *Manager) GetTransactions() ([]Transaction, error) {
	// STEP 3: Fetch transactions
	start := time.Now().AddDate(0, 0, -2).Format("2006-01-02")
	end := time.Now().Format("2006-01-02")

	txReq := plaid.NewTransactionsGetRequest(m.apiclient.token, start, end)

	txResp, _, err := m.apiclient.GetClient().PlaidApi.TransactionsGet(m.apiclient.ctx).
		TransactionsGetRequest(*txReq).
		Execute()

	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions: %v", err)
	}

	var txs []Transaction

	for _, tx := range txResp.Transactions {
		t := Transaction{Transaction: tx}
		accountName, ok := m.accounts[tx.AccountId]
		if ok {
			t.AccountName = accountName
		}
		txs = append(txs, t)
	}

	return txs, nil
}
