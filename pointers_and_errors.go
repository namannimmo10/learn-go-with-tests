package main

import "errors"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Returns the number of Bitcoins.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// In Go, errors are values.
var ErrInsufficientFunds = errors.New("insufficient funds.")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
