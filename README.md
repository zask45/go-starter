# Run Go

Ini langkah-langkah ngebuat project `Go`.

- Buat dir baru
- Buka terminal lalu ketik `go mod init [nama module]`
- `go mod tidy`
- Buat file go
- Run dengan cara `go run namafile.go`

Dari _go mod init_ bakal terbentuk file `go.mod` yang isinya kayak gini

```
module gitgithub

go 1.23.4
```

## Import another package

Ini kayak proses `npm install`. 
Nah sekarang kita bakal coba install package `rsc.io/quote` . Caranya

```
go get rsc.io/quote
```

Nah nanti `go.mod` bakal berubah isinya jadi kek gini

```
module gitgithub

go 1.23.4

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.3.0 // indirect
)
```

Terus ubah `hello.go` jadi kayak gini
```
package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Konnichiwa")
	fmt.Println(quote.Go())
}
```

## Test

Coba test apakah benar `formatter (fmt)` akan mengembalikan `Konnichiwa`.

Gimana caranya? Kita pake prinsip `Separation of Concern (SoC)`. 

Jadi ya, pas kita ngeprint "Konnichiwa" kita tuh pasti input string kan. Nah proses input string ini tuh bagian dari `domain`, yaitu bagian yang berhubungan langsung dengan logika bisnis. Dimana-nya yang berhubungan dengan logika bisnis? Ya proses input data (string) itu sendiri.

Sedangkan, pas kita print ke layar user itu masuknya ke `side effect`. Side effect ini bisa disebut sebagai efek samping yang terjadi di luar dari domain logika. Contohnya kayak proses `display data` ke layar user, `save data to database`, dan `send http request`.

Nah di kasus ini, kita bakal misahin bagian `Greeting` jadi fungsi tersendiri supaya ada `separation of concern`-nya.

```
func Greeting() string {
    return "Konnichiwa"
}

func main() {
    fmt.Println(Greeting())
}
```

Terus buat file test bernama `hello_test.go`

```
package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting()
	want := "Konnichiwa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Lalu run test

```
go test
```

Nanti hasilnya kayak gini. Kolom kedua itu nama module-nya btw

```
PASS
ok      example.com/hello       0
```