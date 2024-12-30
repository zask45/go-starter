package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance().String()
		want := "10 BTC"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Withdraw(Bitcoin(5))

		got := wallet.Balance().String()
		want := "5 BTC"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
