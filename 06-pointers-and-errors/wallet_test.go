package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(BitCoin(10))

		got := wallet.Balance()

		want := BitCoin(10)

		if got != want {
			// We change the "verb" to to string so the `String()` method of the type is used.
			// t.Errorf("got %d want %d", got, want)
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(10)}

		wallet.Withdraw(5)

		got := wallet.Balance()

		want := BitCoin(5)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})
}
