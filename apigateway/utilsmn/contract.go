package contract

import "errors"

// TokenContract represents a token smart contract
type TokenContract struct {
	balances map[string]uint
}

// NewTokenContract creates a new token smart contract
func NewTokenContract() *TokenContract {
	return &TokenContract{balances: make(map[string]uint)}
}

// Transfer transfers tokens from one account to another
func (tc *TokenContract) Transfer(sender, receiver string, amount uint) error {
	senderBalance := tc.balances[sender]
	if senderBalance < amount {
		return errors.New("insufficient balance")
	}
	tc.balances[sender] -= amount
	tc.balances[receiver] += amount
	return nil
}

// GetBalance returns the balance of an account
func (tc *TokenContract) GetBalance(account string) uint {
	return tc.balances[account]
}
