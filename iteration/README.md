# Iteration

Inget, kita ngembangin `Test Driven Development` jadi write test dulu sebagai langkah pertama.

## Write test

```
package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but go %q", expected, repeated)
	}
}
```

## Run the test

```
undefined: Repeat
FAIL	example.com/hello/iteration [build failed]
FAIL
```

`Repeat` is undefined. Jadi, tulis fungsi `Repeat`.


## Write code minimally

Buat file baru `repeat.go`. Terus isi file seperti ini. Biarkan return string kosong karena kita cuma perlu menulis kode secara minimal pada step ini

```
package iteration

func Repeat(character string) string {
	return ""
}
```

## Run the test again

Hasilnya

```
Go\iteration\repeat_test.go:10: expected "aaaaa" but got ""
FAIL
FAIL	example.com/hello/iteration	0.464s
FAIL
```

## Write enough code to make it pass

```
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}

	return repeated
}
```

Hasil test sudah `ok`!


## Refactor

Apa yang bisa direfactor? 
Kita bisa ngebuat jumlah pengulangan jadi variabel sendiri. Kita kasih nama variabel itu jadi `repeatCount`

```
const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
```

Btw notice kan ya pas kita masukkin nilai ke variabel kita pake simbol `:=` sama `=`. Bedanya apa?

`:=` itu buat deklarasi variabel `tanpa tipe data`. Kita bisa langsung masukkin nilai jadi kayak gini `x := 0`. Dan ini cuma bisa dipake dalem blok `fungsi`. Btw simbol ini namanya **short variabel declaration**.

Sedangkan simbol `=` digunain untuk masukkin nilai dengan atau tanpa tipe data. Bisa digunain di luar blok fungsi. Lebih sering dipake untuk inisiasi nilai ke variabel yang udah ada.

Contoh:
```
var x int
x = 10
```

Beda sama `:=` yang penggunaannya lebih kek gini
```
x := 10
```

Intinya ya, penggunaannya itu biasanya ke gini

- `:=` untuk deklarasi sekaligus inisialisasi nilai
- `=` inisialisasi nilai ke variabel yang udah ada, inisiasi nilai ke variabel di luar scope function