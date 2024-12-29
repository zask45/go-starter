# Structs, methods & interfaces

Kita bakal belajar tentang `structs` `methods` dan `interface`. 

# Structs

Sebelum masuk ke `structs`, kita bakal implementasi suatu masalah dulu.

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

Tambah kode ini di file kode

```
type Rectangle struct {
	Width  float64
	Height float64
}
```

Lalu refactor test menjadi seperti ini
```
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 20.0}
	got := Perimeter(rectangle)
	want := 60.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := Area(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
```


## Run the test
```
# example.com/hello/methods [example.com/hello/methods.test]
c:\Users\Keysha\Documents\Go\methods\shapes_test.go:7:19: not enough arguments in call to Perimeter
	have (Rectangle)
	want (float64, float64)

c:\Users\Keysha\Documents\Go\methods\shapes_test.go:17:14: not enough arguments in call to Area
	have (Rectangle)
	want (float64, float64)
FAIL	example.com/hello/methods [build failed]
FAIL
```

Nah, dari hasil test-nya keliatan kalo parameter `Perimeter` dan `Area` gak match sama argumen yang dikasih.


## Write enough code to pass the test

Ubah parameter pada fungsi `Perimeter` dan `Area` menjadi `Rectangle`. Kemudian akses `width` dan `height` dari `struct`. 

```
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
```

Hasil test
```
ok  	example.com/hello/methods	0.471s	coverage: 100.0% of statements
```


# Methods
Sebelum masuk ke `methods` kita bakal nambahin penghitungan luas `circle`. 
Pertama-tama, tulis test dulu.


## Write test

```
t.Run("circle", func(t *testing.T) {
    circle := Circle{10}
    got := Area(circle)
    want := 314.1592653589793

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
})
```

## Run the test

Hasil test
```
Go\methods\shapes_test.go:27:13: undefined: Circle
FAIL	example.com/hello/methods [build failed]
FAIL
```

## Write code 

Tambah `struct` circle kayak gini
```
type Circle struct {
	Radius float64
}
```

Coba test lagi
```
Go\methods\shapes_test.go:28:15: cannot use circle (variable of type Circle) as Rectangle value in argument to Area
FAIL	example.com/hello/methods [build failed]
FAIL
```

Lihat error message-nya
_Cannot use circle (variable of type Circle) as Rectangle value ..._

Nah gimana cara penyelesaiannya?

Pake `methods`



## Penjelasan methods

Method itu function tapi ada `receivernya`. Cara pakenya itu kayak gini `circle.Area()`. Mirip-mirip lah sama `extension function` di Kotlin.

Coba implementasi `methods` ini di kode kita.


## Write test


```
func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
```

Hasil run test

```
c:\Users\Keysha\Documents\Go\methods\shapes_test.go:18:20: rectangle.Area undefined (type Rectangle has no field or method Area)
c:\Users\Keysha\Documents\Go\methods\shapes_test.go:28:17: circle.Area undefined (type Circle has no field or method Area)
FAIL	example.com/hello/methods [build failed]
FAIL
```

_rectangle.Area undefined_


## Write minimal amount of code

Pastikan files `shapes.go` memuat kode ini

```
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return 0
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 0
}
```

Hasil test

```
--- FAIL: TestArea (0.00s)
    --- FAIL: TestArea/rectangles (0.00s)
        c:\Users\Keysha\Documents\Go\methods\shapes_test.go:22: got 0.00 want 72.00
    --- FAIL: TestArea/circle (0.00s)
        c:\Users\Keysha\Documents\Go\methods\shapes_test.go:32: got 0.00 want 314.16
FAIL
FAIL	example.com/hello/methods	0.406s
FAIL
```

Berarti implementasinya sudah benar. Tinggal perbaiki isi method.


## Write enough code to make it pass

Tinggal masukkin rumus luas persegi panjang dan lingkaran.

```
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
```
```
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
```

Hasil test:
```
ok  	example.com/hello/methods	0.401s
```