# Pointers and Errors

Sesuai judul, kita bakal belajar tentang `pointer` dan `error`.

Tapi sebelum masuk ke 2 topik itu, kita bakal bangun `banking system` sederhana dulu.

Apa yang mau kita buat?
`Bitcoin Wallet`!


## Write the test first

```
func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
```

Hasil test
```
undefined: Wallet
FAIL	example.com/hello/pointer [build failed]
FAIL
```

Jadi kita perlu implementasi struct `Wallet` dulu


## Write minimal amount of code to test

```
type Wallet struct {
	
}
```

Jalanin test dan liat hasilnya

```
c:\Users\Keysha\Documents\Go\pointer\wallet_test.go:7:9: wallet.deposit undefined (type Wallet has no field or method Deposit)

c:\Users\Keysha\Documents\Go\pointer\wallet_test.go:9:16: wallet.Balance undefined (type Wallet has no field or method Balance)
FAIL	example.com/hello/pointer [build failed]
FAIL
```

_type Wallet has no field or method Deposit_
_type Wallet has no field or method Balance_

Sesuai error message di atas, kita perlu definisiin method `Deposit` dan `Balance`

```
func (w Wallet) Deposit(amount int) {

}

func (w Wallet) Balance() int {
	return 0
}
```

Hasil test
```
\Go\pointer\wallet_test.go:13: got 0 want 10
FAIL
FAIL	example.com/hello/pointer	0.390s
FAIL
```

Abis ini berarti kita perlu perbaiki kode di method `Deposit` dan `Balance` supaya bisa pass the test.


## Write enough code to make it pass

Tambahkan field `balance` di struct dan perbaiki kode jadi kayak gini

```
type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
```

Coba run test. 
```
--- FAIL: TestWallet (0.00s)
    c:\Users\Keysha\Documents\Go\pointer\wallet_test.go:13: got 0 want 10
FAIL
FAIL	example.com/hello/pointer	0.372s
FAIL
```

Hasil test-nya tetap `got 0 want 10`.


## Kok bisa?

Ini karena pas kita panggil `func (w Wallet) Deposit(amount int)` di file `test` kita tuh ngecopy `w Wallet` ke alokasi memori yang baru. 

Jadi kalo kita mau ngubah nilai wallet one and for all, kita harus ubah nilai tersebut di alokasi memori yang lama (yang original) bukan malah ubah yg copyan-nya. 

Makanya kita pake `pointer`.


# Pointer

Cara supaya kita bisa ngubah nilai di address memori yang lama?

Pake `pointer`!

Gimana pointer tu bekerja? Pointer bakal nunjuk address memori yang lama instead of ngebuat alokasi memori baru buat parameter yang dipanggil.

Nah caranya implementasi pointer tinggal pake tanda bintang `*`. Dengan ini kita ngasih tau kalo kita mau ubah nilai di parameter itu di address memori yang original. Jadi pas metode tsb dipanggil gak akan dibentuk alokasi memori yang baru untuk nampung nilai argumen.

```
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
```


## Refactor

Karena kita mau buat `Bitcoin wallet`, kita implementasiin Bitcoin. 


```
package wallet

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
```

Run test
```
invalid operation: got != want (mismatched types Bitcoin and int)
FAIL	example.com/hello/pointer [build failed]
FAIL
```

Ubah data type di `TestWallet` jadi `Bitcoin`

```
func TestWallet(t *testing.T) {
	...

	wallet.Deposit(Bitcoin(10))

	...

	want := Bitcoin(10)

	...
}
```

Hasil test

```
ok  	example.com/hello/pointer	(cached)
```

Lalu kita bisa buat `Bitcoin` jadi `string` dengan ngebuat method yang bisa melakukan konversi bitcoin ke string

```
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
```

Ubah test jadi  kayak gini
```
func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance().String()
	want := "10 BTC"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
```

Hasil test
```
ok  	example.com/hello/pointer	(cached)
```

Selanjutnya kita bakal tambah method `Withdraw` untuk narik `Bitcoin`.


## Write the test 

Buat test jadi seperti ini. Jangan lupa masukkan `balance` Bitcoin saat test `Withdraw`.

```
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
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(5))

		got := wallet.Balance().String()
		want := "5 BTC"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
```

Hasil test

```
wallet.Withdraw undefined (type Wallet has no field or method Withdraw)
FAIL	example.com/hello/pointer [build failed]
FAIL
```

## Write minimal amount of code to test

func (w *Wallet) Withdraw(amount Bitcoin) Bitcoin {
	return 0
}

Hasil test
```
--- FAIL: TestWallet (0.00s)
    --- FAIL: TestWallet/withdraw (0.00s)
        c:\Users\Keysha\Documents\Go\pointer\wallet_test.go:26: got 0 BTC want 5 BTC
FAIL
FAIL	example.com/hello/pointer	0.383s
FAIL
```


## Write enough code to pass the test

```
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
```

Hasil test

```
ok  	example.com/hello/pointer	0.450s	coverage: 100.0% of statements
```


## Refactor

Yang dilakuin?

Rapihin test. 
Remove duplication pada test.

```
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
```

Sekarang tambah method untuk withdraw unsufficient amount of balance.


## Write test

```
t.Run("withdraw insufficient funds", func(t *testing.T) {
    startingBalance := Bitcoin(10)
    wallet := Wallet{startingBalance}
    err := wallet.Withdraw(Bitcoin(20))

    assertBalance(t, wallet, startingBalance.String())

    if err == nil {
        t.Error("wanted an error but didn't get one")
    }
})
```

Hasil run test
```
Go\pointer\wallet_test.go:34:10: wallet.Withdraw(Bitcoin(20)) (no value) used as value
FAIL	example.com/hello/pointer [build failed]
FAIL
```

_wallet.Withdraw(Bitcoin(20)) (no value) used as value_

Selanjutnya, kita buat agar method `Withdraw` bisa return error.


## Write minimal amount of code to test

```
func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.balance -= amount
	return nil
}
```

Hasil
```
-- FAIL: TestWallet (0.00s)
    --- FAIL: TestWallet/withdraw_insufficient_funds (0.00s)
        Go\pointer\wallet_test.go:36: got -10 BTC want 10 BTC

        \Go\pointer\wallet_test.go:39: wanted an error but didn't get one
FAIL
FAIL	example.com/hello/pointer	0.395s
FAIL
```


## Write enough code to make it pass

```
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("unsufficient balance")
	}

	w.balance -= amount
	return nil
}
```

Hasil test
```
ok  	example.com/hello/pointer	(cached)
```


## Refactor

```
func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, w Wallet, want string) {
		t.Helper()
		got := w.Balance().String()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
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

		assertBalance(t, wallet, startingBalance.String())
		assertError(t, err)
	})
}
```

Hasil test `ok`.

Sekarang, kita buat supaya `Withdraw` bisa return `error message` yang lebih proper. Btw ini user dianggap dapet error messagenya ya walaupun error message `unsufficient fund`-nya gak muncul di compiler.


## Write the test first

Return `Fatal` untuk stop testing jika ternyata tidak ada error yang di-return `Withdraw`. `Withdraw` memang harus return `error` jika `balance` di wallet tidak cukup untuk ditarik.

```
assertError := func(t testing.TB, got error, want string) {
    t.Helper()

    if got == nil {
        t.Fatal("didn't get an error but expecting one")
    }

    if got.Error() != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
```
```
t.Run("withdraw insufficient funds", func(t *testing.T) {
    startingBalance := Bitcoin(10)
    wallet := Wallet{startingBalance}
    err := wallet.Withdraw(Bitcoin(20))

    assertError(t, err, "cannot withdraw, insufficient funds")
    assertBalance(t, wallet, startingBalance.String())
})
```

Basically ini kita ngecek apakah error-nya ini disebabkan oleh `unsufficient amount of fund`. Kalo error-nya `nil` kita throw error yang ngasih informasi kalo harusnya ada error. 

Hasil test
```
FAIL: TestWallet (0.00s)
--- FAIL: TestWallet/withdraw_insufficient_funds (0.00s)
    c:\Users\Keysha\Documents\Go\pointer\wallet_test.go:48: got "unsufficient balance", want "cannot withdraw, insufficient funds"
```

Sekarang kita ganti error message di `Withdraw` jadi `cannot withdraw, insufficient funds`


## Write enough code to make it pass

```
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}

	w.balance -= amount
	return nil
}
```

Hasil test
```
ok  	example.com/hello/pointer	0.356s
```


## Refactor

Jadiin error-messagenya jadi variabel tersendiri

```
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
```

Karena `ErrInsufficientFunds` jadi variabel tersendiri, kita bisa pake variabel ini juga untuk testing.

```


# Unchecked errors

Walaupun testing lewat compiler itu berguna, bisa aja ada error yang belum dicek.

Untuk handle masalah ini coba install `errcheck`
```
go install github.com/kisielk/errcheck@latest
```

Lalu jalanin kode ini di directory 
```
errcheck .
```

Hasil-nya
```
wallet_test.go:38:18:   wallet.Withdraw(Bitcoin(5))
```

Apa maksudnya?
Maksudnya, pada line 38 kita belum cek error yang di-return sama `Withdraw`.

Makanya kita ubah kode-nya jadi gini
```
t.Run("withdraw", func(t *testing.T) {
    wallet := Wallet{balance: Bitcoin(10)}
    err := wallet.Withdraw(Bitcoin(5))

    assertError(t, err, ErrInsufficientFunds)
    assertBalance(t, wallet, "5 BTC")
})
```

Hasil test
```
--- FAIL: TestWallet (0.00s)
    --- FAIL: TestWallet/withdraw (0.00s)
        c:\Users\Keysha\Documents\Go\pointer\wallet_test.go:40: didn't get an error but expecting one
FAIL
FAIL	example.com/hello/pointer	0.388s
```

Nah, karena memang kita gak expecting any error di test tersebut, kita buat fungsi baru yang namanya `assertNoError` untuk handle `no error`.

```
t.Run("withdraw", func(t *testing.T) {
    wallet := Wallet{balance: Bitcoin(10)}
    err := wallet.Withdraw(Bitcoin(5))

    assertNoError(t, err)
    assertBalance(t, wallet, "5 BTC")
})
```

Hasil test
```
ok  	example.com/hello/pointer	0.389s
```

Coba run `errcheck .`

Gak ada return apapun. 
Berarti sudah oke!


## Wrapping up

- `Pointer` untuk menunjuk memori address suatu variabel. Biasanya digunain untuk mengganti value di object yang ada di argumen.
- `Errors` jangan hanya dicek, tapi harus dihandle secara graceful!