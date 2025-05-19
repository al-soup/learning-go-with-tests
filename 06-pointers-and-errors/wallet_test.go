package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(10)}
		err := wallet.Withdraw(BitCoin(5))
		assertNoError(t, err)
		assertBalance(t, wallet, BitCoin(5))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingFund := BitCoin(10)
		wallet := Wallet{startingFund}
		err := wallet.Withdraw(BitCoin(15))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingFund)
	})
}

// Move helpers outside of test body so you can start reading assertions right away
func assertBalance(t testing.TB, wallet Wallet, want BitCoin) {
	got := wallet.Balance()

	if got != want {
		// We change the "verb" to to string so the `String()` method of the type is used.
		// t.Errorf("got %d want %d", got, want)
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		// stops execution here
		t.Fatal("did not receive expected errror")
	}

	// `got.Error()` converts Error to a string
	// It makes sense to check against nil first (above)
	// if got.Error() != want {
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("received unwanted error")
	}
}
