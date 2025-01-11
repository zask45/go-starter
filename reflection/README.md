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