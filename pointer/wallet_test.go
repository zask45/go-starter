package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, w Wallet, want string) {
		t.Helper()
		got := w.Balance().String()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()

		if got == nil {
			t.Fatal("didn't get an error but expecting one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, "10 BTC")
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(5))

		assertBalance(t, wallet, "5 BTC")
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(20))

		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, startingBalance.String())
	})
}
