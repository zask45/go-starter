# Reflection

_Reflection_?

Apa maksudnya?

Jadi gini. Misalnya ada challenge

> Tulis fungsi `walk(x interface{}, fn func(string))` dimana ambil `x` struct dan panggil `fn` untuk ambil semua string yang ada dalem struct.

Kalo kita perhatiin challange di atas, kita diminta masukkin `interface` sama `function`sebagai type-nya kan?

Nah ini yang dimaksud `reflection`.

 > Reflection itu bisa dibilang kemampuan program untuk meriksa strukturnya sendiri, terutama melalui types.


 ## What is interface{}

Kalo kita liat ini,

```
walk(x interface{}, fn func(string))
``` 

ada `interface{}` kan?

Maksudnya ini?

Selama ini kita ngebuat `function` dengan tipe-tipe data kayak `int` `string` atau tipe data sendiri kayak `BankAccount`. Nah kalo tipe data yang dimasukkin kayak gini, `compiler` bisa komplain kalo kita masukkin argumen dengan tipe data yang salah. 

Dari sini, gimana kalo kita mau masukkin nilai dengan tipe data apa aja?

Di sini lah kita pake `interface{}`! Tipe data interface bakal nerima value dengan tipe data apa aja. 

Kok gak pake `any` aja? 
Bisa juga kalo mau pake any. Kenapa? Karena `any` itu sebenernya alias untuk `interface{}`.

## Kenapa gak pake interface dan buat function flexible?

Kita gak bisa pake `interface{}` for everything. Pengecekkan kevalidan tipe data itu salah satu hal yang krusial di pemrograman. Misal kita cuma minta input dalam bentuk `int` masa kita mau accept input-an user dalem bentuk `char`? Gak mungkin kan? Makanya kita gak bisa pake `interface{}` untuk semua kasus.

## Write the test first
 
Sekarang kita coba tulis test-nya dulu kayak biasa.
```
package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}
}
```

Inti test-nya apa?
Intinya kita ngetest berapa banyak isi dari struct `x`. Kalo `0` berarti fungsi `walk`-nya belum kepanggil.

Hasil test

```
undefined: walk
FAIL	example.com/hello/reflection [build failed]
FAIL
```

Nah dikasih tau kalo kita harus define `walk` dulu. 


## Write minimal amount of code to test

```
func walk(x interface{}, fn func(input string)) {

}
```

Hasil test
```
 FAIL: TestWalk (0.00s)
    c:\..\Go\reflection\walk_test.go:18: wrong number of function calls, got 0 want 1
FAIL
FAIL	example.com/hello/reflection	0.410s
```

Kalo sampe `len(got) == 1` berarti function dalem `walk` itu belum kepanggil samsek. Nah sekarang coba benerin `walk` untuk implementasi `fn` di dalem body-nya.

## Write code minimally to test

```
func walk(x interface{}, fn func(input string)) {
	fn("Implement function inside walk")
}
```

Sekarang coba jalanin test-nya

```
ok  	example.com/hello/reflection	(cached)
```

Iya udah `ok` sih, tapi apa bener string yang di-pass udah bener? Harus di cek kan?

## Write the test first

Tambahin `if` untuk ngecek apakah `got` index ke-0 itu nilainya sama kayak `expected`.

```
func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("got %q want %q", got[0], expected)
	}
}
```

Hasil test
```
Go\reflection\walk_test.go:22: got "Implement function inside walk" want "Chris"
FAIL
FAIL	example.com/hello/reflection	0.520s
FAIL
```

## Write minimal amount of code to pass test

Kalo kita liat kode ini di `test`

```
x := struct {
		Name string
	}{expected}

walk(x, func(input string) {
	got = append(got, input)
})
```

Kita bisa liat kalo `x` itu masukkin `expected` sebagai nilai `Name`.

Nah kalo gitu, kita ambil aja dong nilai `x` lewat `walk`.

```
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
```

Kita `reflect` nilai dari `x`. Terus masukkin nilai tersebut sebagai `input` fn.

Karena nilai dari `x` udah di-pass sebagai `input`, harusnya `test` juga udah bisa append `input` ke slice `got` kan?

Coba kita test
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestWalk$ example.com/hello/reflection

ok  	example.com/hello/reflection	(cached)
```

_Tapiiii...._

Masih ada masalah, nih!

## Another problem

Coba liat kode ini baik-baik.

```
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
```

Kita cuma ambil index pertama dari `x`! Padahal belum tentu kan nilai di struct itu cuma 1? Terus bisa jadi struct-nya malah kosong. 

Kalo gitu kita harus gimana? 
Benerin dulu `test`-nya!


## Write the test first

```
func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(*testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
```

Bingung?

Coba pahamin pelan-pelan.

Pertama, kita buat `test case` dalem bentuk struct.

```
cases := []struct {
    Name          string
    Input         interface{}
    ExpectedCalls []string
}{
    {
        "struct with one string field",
        struct {
            Name string
        }{"Chris"},
        []string{"Chris"},
    },
}
```

Nilai yang pertama diisi `nama test`.
Yang kedua nilai  `input` yang mau diuji.
Yang ketiga `hasil` yang di-expect.

Terus kita test-apakah `walk` append `x`. Cara ngeceknya dengan masukkin `test.Input` sebagai `x` dan biarin anon function di sebelahnya buat append ke `got`
```
t.Run(test.Name, func(*testing.T) {
    var got []string

    walk(test.Input, func(input string) {
        got = append(got, input)
    })

    if !reflect.DeepEqual(got, test.ExpectedCalls) {
        t.Errorf("got %v, want %v", got, test.ExpectedCalls)
    }
})
```

Hasil test
```
ok  	example.com/hello/reflection	(cached)
```

Sekarang muncul pertanyaan. Gimana kalo input-nya ada 2 string kek gini?
```
struct {
    Name string
    Place string
}
```

## Write the test first

Sekarang coba kita test kalo input-nya ada 2 string

```
func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name  string
				Place string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(*testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
```

Hasil test
```
Go\reflection\walk_test.go:33: got [Chris], want [Chris London]
FAIL
FAIL	example.com/hello/reflection	0.400s
FAIL
```

Kenapa fail? Coba cek kode fungsi `walk` di `walk.go`.

```
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
```

Disini keliatan kalo kita cuma masukkin `field` index ke-`0` sebagai `input`. Coba perbaiki.


## Write enough code to pass test

Benerin `walk` biar bisa ambil semua `Field` di `x`
```
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		fn(val.Field(i).String())
	}
}
```

Hasil test
```
ok  	example.com/hello/reflection	0.397s
```


## Refactor

Inget, ini kan `interface {}`. Berarti bisa masukkin `int` juga dong?
Coba kita test `int` tapi kita buat hasilnya cuma ngeluarin `string`.


## Write the test first

```
func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
	}

	...

}
```

Hasil test
```
Go\reflection\walk_test.go:33: got [Chris <int Value>], want [Chris]
FAIL
FAIL	example.com/hello/reflection	0.403s
FAIL
```

## Write code to pass test

Perbaiki `walk` supaya hanya menginput string.
```
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
```

Hasil test
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestWalk$ example.com/hello/reflection

ok  	example.com/hello/reflection	(cached)
```

## Full code

Sebelum masuk ke bagian selanjutnya, ini kode full-nya
```
package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(*testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
```
```
package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
```

# Nested struct

## Refactor

Skenario selanjutnya yang bakal di-test adalah kalo `struct`-nya nested. Nah gimana kalo kayak gini?


## Write the test first

```
func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"nested",
			struct {
				Name string
				Profile struct {
					Age int
					City string
				}
			}{"Chris", struct {
				Age int
				City string
			}{33, "London"}},
			[]string{"Chris", "London"},
		},
	}

 ...

}
```

Kita jadi nambahin `Profile struct` setelah `name`. Isi `Profile` itu `age: int` dan `city: string`. Expected value-nya juga kita ganti jadi `Chris` dan `London`. Value dari `age` gak dimasukkin sebagai output karena dia bentuknya `int`.

Hasil test
```
Go\reflection\walk_test.go:39: got [Chris], want [Chris London]
FAIL
FAIL	example.com/hello/reflection	0.390s
FAIL
```

Oke, kayaknya fungsi `walk` belum bisa nge-proses `nested struct`. Kita perbaiki!


## Write enough code to pass test

```
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}

		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}
}
```

Intinya kalo hasil pengecekkan nunjukkin `field.Kind() == reflect.Struct` maka kita panggil lagi `walk` dengan masukkin `field` sebagai `x interface{}`.


Hasil test
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestWalk$ example.com/hello/reflection

ok  	example.com/hello/reflection	0.461s
```

## Refactor

Kita bisa rapihin `test` jadi kayak gini

```
package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"nested",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(*testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
```

Yaps, kita bisa buat `struct` di luar fungsi `TestWalk` supaya nested struct-nya bisa lebih rapih.

Terus kita bisa rapihin `if` di fungsi `walk` jadi `switch` kayak gini

```
package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
```

Pastiin hasil refactor `ok`
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestWalk$ example.com/hello/reflection

ok  	example.com/hello/reflection	0.389s
```

# Pointer

Gimana kalo inputnya ada yang berbentuk `Pointer`?

## Write the test first

```
func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
	}

...

}
```

Hasil test
```
--- FAIL: TestWalk (0.00s)
    --- FAIL: TestWalk/pointers_to_things (0.00s)
panic: reflect: call of reflect.Value.NumField on ptr Value [recovered]
	panic: reflect: call of reflect.Value.NumField on ptr Value
```

Nah gimana tuh?

## Write enough code to pass test

Cara nyelesain error di atas itu dengan ekstrak nilai pada `Pointer` pake fungsi `Elem()`. Proses ini dilakuin sebelum masuk ke `switch`. Jadi intinya mah sebelum masuk ke `switch` pastiin dulu bentuk value di Pointernya itu udah diekstrak pake `Elem`.

```
package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
```

```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestWalk$ example.com/hello/reflection

ok  	example.com/hello/reflection	0.398s
```

## Refactor

Kita bisa pisah proses ekstrak value jadi function sendiri
```
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
```
```
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
```

Coba pastiin hasil test masih `ok`
```
ok  	example.com/hello/reflection	0.385s
```

Yups, masih `ok`!

## Full code

Sebelum masuk ke bagian lain, ini kode full nya
```
package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
```

```
package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(*testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
```