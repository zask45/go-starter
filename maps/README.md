# Maps

Di modul ini kita bakal buat fungsi `Search` dengan bantuan `maps`.


## Write test first

```
package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
```

Struktur map
```
map[key datatype]value datatype{key: value}
```

Hasil test
```
undefined: Search
FAIL	example.com/hello/maps [build failed]
FAIL
```


## Write minimal amount of code to pass the test

```
func Search(dictionary map[string]string, word string) string {
	return ""
}
```

Hasil test

```
--- FAIL: TestSearch (0.00s)
    c:\Users\Keysha\Documents\Go\maps\dictionary_test.go:12: got "" want "this is just a test" given "test"
FAIL
FAIL	example.com/hello/maps	0.368s
FAIL
```

Sekarang revisi kode agar hasil return bisa pass the test.


## Write enough code to pass the test

```
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
```

Hasil test
```
ok  	example.com/hello/maps	0.394s
```


## Refactor

```
package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
```

## Using custom type

Coba pakai custom type `Dictionary` instead of `maps` secara langsung.

```
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertStrings(t, got, want)
}
```

Hasil test

```
undefined: Dictionary
FAIL	example.com/hello/maps [build failed]
FAIL
```

Sekarang buat type `Dictionary` di `dictionary.go`. Masukkan `map[string]string` sebagai datatype-nya.

```
type Dictionary map[string]string
```

Ubah parameter `Search` dari maps jadi `Dictionary`
```
func Search(d Dictionary, word string) string {
	return d[word]
}
```

Run test
```
ok  	example.com/hello/maps	0.326s
```

Sekarang coba implementasi bagaimana kalau kita search word yang gak ada di `Dictionary`?


# Handle error unknown word

## Write the test first

```
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
```

Hasil test
```
Go\maps\dictionary_test.go:9:13: assignment mismatch: 2 variables but dictionary.Search returns 1 value

Go\maps\dictionary_test.go:16:13: assignment mismatch: 2 variables but dictionary.Search returns 1 value

FAIL	example.com/hello/maps [build failed]
FAIL
```


## Write minimal amount code to test

```
package maps

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	return d[word], nil
}
```

Hasil test
```
--- FAIL: TestSearch (0.00s)
    --- FAIL: TestSearch/unknown_word (0.00s)
        c:...\Go\maps\dictionary_test.go:20: expected to get an error
FAIL
FAIL	example.com/hello/maps	0.399s
FAIL
```


## Write enough code to make it pass

Supaya bisa pass the test, kita pake fitur `map lookup` yang return 2 value. Value pertama yang direturn adalah `value dari key yang dicari`. Sedangkan value kedua yang di-return `map lookup` adalah nilai boolean, `true` apabila kata terdapat di maps dan `false` apabila kata tidak terdapat di maps.

```
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.New("could not find the word")
	}

	return definition, nil
}
```


## Refactor

Kita bakal refactor error di dua file, yaitu file `dictionary.go` dan file `test`.

Di `dictionary.go` kita buat variabel error baru bernama `ErrNotFound`

```
var ErrNotFound = errors.New("could not find the word")
```

Lalu ubah fungsi `Search` untuk return `ErrNotFound` apabila `ok == false`

```
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
```

Selanjutnya, refactor error di `file test`.

Ubah testing untuk `unknown word` jadi seperti ini. Intinya kita ingin mengecek apakah error yang kita terima saat kata tidak ada di dictionary itu benar `ErrNotFound`.
```
t.Run("unknown word", func(t *testing.T) {
    _, got := dictionary.Search("unknown")

    if got == nil {
        t.Fatal("expected to get an error")
    }

    assertError(t, got, ErrNotFound)
})
```

Lalu buat fungsi `assertError`. Pastikan bahwa parameter `got` dan `want` bertipe `error`.

```
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
```

Hasil test
```
ok  	example.com/hello/maps	0.456s	coverage: 100.0% of statements
```

Nah, sekarang kita tambah fitur baru yaitu `tambah dictionary`.


# Add word to dictionary

## Write test first
```
func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word")
	}

	assertStrings(t, got, want)
}
```

Hasil test
```
ictionary.Add undefined (type Dictionary has no field or method Add)
FAIL	example.com/hello/maps [build failed]
FAIL
```


## Write minimal amount of code to test

```
func (d Dictionary) Add(key string, value string) {
}
```

Hasil test
```
\Go\maps\dictionary_test.go:34: should find added word
FAIL
FAIL	example.com/hello/maps	0.408s
FAIL
```


## Write enough code to make it pass

```
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
```

Hasil test
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestAdd$ example.com/hello/maps

ok  	example.com/hello/maps	(cached)
```