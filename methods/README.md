# Structs, methods & interfaces

Kita bakal belajar tentang `structs` `methods` dan `interface`. 

Tapi sebelum masuk ke situ, kita bakal implementasi suatu masalah dulu.

Katakanlah kita butuh program untuk hitung `keliling` area `persegi panjang` dimana panjang dan lebarnya itu berbentuk float. 

Gimana penerapannya?


## Write test first

Sesuai dengan TDD cycle, pertama-tama kita buat test dulu.

```
package area

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 20.0)
	want := 60.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
```

## Run the test

Hasil test
```
methods\area_test.go:6:9: undefined: Perimeter
FAIL	example.com/hello/methods [build failed]
FAIL
```


## Write code minimally

```
func Perimeter(width float64, length float64) float64 {
	return 0
}
```

Jalankan test.
Hasilnya
```
--- FAIL: TestPerimeter (0.00s)
    c:\Users\Keysha\Documents\Go\methods\area_test.go:10: got 0.00 want 60.00
FAIL
FAIL	example.com/hello/methods	0.412s
```

Parameter-nya sudah benar. Tinggal perbaiki isi function `Perimeter`


## Write enough code to make it pass

```
func Perimeter(width float64, length float64) float64 {
	return 2 * width + 2 * length
}
```

Run test. Cek hasilnya
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestPerimeter$ example.com/hello/methods

ok  	example.com/hello/methods	0.324s
```

Sekarang, tambahkan `TestArea`


## Write test

```
func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
```


## Write code

```
func Area(width float64, length float64) float64 {
	return width * length
}
```

Jalankan test.

Hasil test
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestArea$ example.com/hello/methods

ok  	example.com/hello/methods	(cached)
```


## Refactor

Kalo kita liat kode kita, gak ada sesuatu yang nunjukkin kalo kita itu ngitung keliling dan luas `persegi panjang`. 

Ini bisa jadi masalah di masa depan. Mungkin aja developer lain ataupun kita sendiri di masa depan malah masukkin width dan height untuk segitiga.

Makanya itu, kita pake `struct`.

Tambah kode ini di `area.go`

```
type Rectangle struct {
	Width  float64
	Height float64
}
```
