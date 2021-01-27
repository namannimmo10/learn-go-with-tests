package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit the Bitcoins", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(100))
		assertBalance(t, wallet, Bitcoin(100))
	})

	t.Run("withdraw the Bitcoins", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(50)) // errcheck will catch this.

		assertBalance(t, wallet, Bitcoin(50))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != expected {
		t.Errorf("got %d want %d", got, expected)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error, didn't expect one.")
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if err != want {
		t.Errorf("got %q, expected %q", err, want)
	}
}
