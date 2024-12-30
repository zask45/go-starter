package wallet

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance().String()
	want := "10 BTC"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
