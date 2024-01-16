package pointers_errors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})

}

func TestWallet1(t *testing.T) {

	assertBalance := func(tb testing.TB, wallet Wallet, want Bitcoin) {

		tb.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})

}

func TestWallet2(t *testing.T) {

	//assertError := func(tb testing.TB, err error) {
	//	tb.Helper()
	//	if err == nil {
	//		t.Error("wanted an error but didn't get one")
	//	}
	//}

	//assertError := func(tb testing.TB, got error, want string) {
	//	tb.Helper()
	//	if got == nil {
	//		t.Error("wanted an error but didn't get one")
	//	}
	//
	//	if got.Error() != want {
	//		t.Errorf("got %q, want %q", got, want)
	//	}
	//}

	assertError := func(tb testing.TB, got error, want error) {
		tb.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertBalance := func(tb testing.TB, wallet Wallet, want Bitcoin) {

		tb.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Withdraw", func(t *testing.T) {

		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

	})

}

func TestWallet3(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {

		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})

}

func assertBalance(tb testing.TB, wallet Wallet, want Bitcoin) {

	tb.Helper()
	got := wallet.Balance()
	if got != want {
		tb.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(tb testing.TB, got error) {
	tb.Helper()
	if got != nil {
		tb.Fatal("got an error but didn't want one")
	}
}

func assertError(tb testing.TB, got error, want error) {

	tb.Helper()
	if got == nil {
		tb.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		tb.Errorf("got %s, want %s", got, want)
	}
}
