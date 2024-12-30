package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, w Wallet, want string) {
		t.Helper()
		got := w.Balance().String()

		if got != want {
			t.Errorf("got %s want %s", got, want)
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
}
