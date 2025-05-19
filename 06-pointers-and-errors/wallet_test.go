package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want BitCoin) {
		got := wallet.Balance()

		if got != want {
			// We change the "verb" to to string so the `String()` method of the type is used.
			// t.Errorf("got %d want %d", got, want)
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(BitCoin(10))

		assertBalance(t, wallet, BitCoin(10))

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(10)}

		wallet.Withdraw(5)

		assertBalance(t, wallet, BitCoin(5))
	})
}
