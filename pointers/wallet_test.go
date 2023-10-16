package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))

	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}
		err := wallet.Withdraw(BitCoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, BitCoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		statingBalance := BitCoin(10)
		wallet := Wallet{statingBalance}
		err := wallet.Withdraw(BitCoin(20))
		assertBalance(t, wallet, statingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
func assertBalance(t testing.TB, wallet Wallet, want BitCoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("did not get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
