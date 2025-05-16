package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(BitCoin(10))

	got := wallet.Balance()

	want := BitCoin(10)

	if got != want {
		// We change the "verb" to to string so the `String()` method of the type is used.
		// t.Errorf("got %d want %d", got, want)
		t.Errorf("got %s want %s", got, want)
	}
}
