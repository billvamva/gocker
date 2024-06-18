package pointers

import (
	"testing"
)
func TestWallet(t *testing.T) {

	assertBalance := func (t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		if wallet.balance != want {
			t.Errorf("got %s want %s", wallet.balance, want)
		}
	}

	assertError := func(t testing.TB, got error, want error){
		t.Helper()
		if got == nil {
			t.Fatal("Expected error.")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("deposits", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdrawals" ,func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}