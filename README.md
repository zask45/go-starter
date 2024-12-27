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
