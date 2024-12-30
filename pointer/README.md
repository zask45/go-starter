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