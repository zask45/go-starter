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


## Pointers, copies, et al 

Gak kayak `field` biasa dimana perubahan nilai harus diakses lewat pointer, `maps` gak butuh seperti itu.
Ini karena pas kita passing `map` ke function tertentu, kita tu sebenernya passing `pointer` dari `map` itu sendiri. Jadi kita gak perlu ngedefinisiin pointer lagi untuk ngubah value dari map.

By the way, kita sebaiknya juga gak nginisialiasiin nilai `nil` ke map. Ini karena nilai nil bisa ngebuat `runtime panic`.

Kalau memang mau inisialiasi map kosong sebaiknya kayak gini

```
var dictionary = map[string]string{}

// or

var dictionary = make(map[string]string)
```

Bukan malah
```
var m map[string]string 	// BIG NO!
```


## Refactor

Refactor `TestAdd` jadi seperti ini
```
func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "take measures to check the quality, performance, or reliability of something"

	dictionary.Add(word, definition)

	assertDefinition(t, dictionary, word, definition)
}
```
```

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word")
	}

	assertStrings(t, got, definition)
}
```

Hasil test
```
ok  	example.com/hello/maps	(cached)
```


Sekarang, kita buat test untuk kondisi dimana nilai yang ingin kita masukkan sudah ada.


## Write test first

```
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "take measures to check the quality, performance, or reliability of something"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "take measures to check the quality, performance, or reliability of something"

		// definisikan dictionary dengan key "word" dan value "definition"
		dictionary := Dictionary{word: definition}

		// add lagi key "word" dan value "definition" ke dictionary
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}
```

Hasil test
```
Go\maps\dictionary_test.go:32:10: dictionary.Add(word, definition) (no value) used as value

Go\maps\dictionary_test.go:46:10: dictionary.Add(word, definition) (no value) used as value

Go\maps\dictionary_test.go:48:23: undefined: ErrWordExists
FAIL	example.com/hello/maps [build failed]
```

Jika kita baca error message di atas ada dua masalah disini.
`dictionary.Add(word, definition) tidak ada return type`
`undefined ErrWordExists`

Selanjutnya, kita perbaiki masalah di atas.


## Write minimal code to test

```
var (
	ErrNotFound = errors.New("could not find the word")
	ErrWordExists = errors.New("word already exist")
)
```
```
func (d Dictionary) Add(word, definition string) error {
	d[word] = definition
	return nil
}
```

Hasil test
```
-- FAIL: TestAdd (0.00s)
    --- FAIL: TestAdd/existing_word (0.00s)
        c:\Users\Keysha\Documents\Go\maps\dictionary_test.go:48: got error %!q(<nil>) want "word already exist"
FAIL
FAIL	example.com/hello/maps	0.435s
```


## Write enough code to make it pass

```
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
```

Jadi intinya kita panggil method `Search`. 
Kalo `word` belum ada di dictionary, kan pasti Search return error `ErrNotFound`. Nah jika yg di-return `ErrNotFound` maka kita tambahin `word` ke `dictionary`.

Kalo `word` udah ada di dictionary, pasti error yang di-return adalah `nil`. Kalo ini terjadi, kita return error `ErrWordExists`.


Hasil test
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestAdd$ example.com/hello/maps

ok  	example.com/hello/maps	0.383s
```


## Refactor

Buat `error` menjadi `const` dengan tipe `DictionaryErr` seperti di bawah ini.
```
const (
	ErrNotFound   = DictionaryErr("could not find the word")
	ErrWordExists = DictionaryErr("word already exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
```

Di `Go` tipe data custom bisa dijadiin `error` kalau mengimplementasikan interface `error`. Makanya itu, pastikan ada method `Error()` untuk implementasi interface `error`.


Sekarang, kita tambah fitur untuk `Update` definisi dari word.


# Update definition of word

## Write test first

```
func TestUpdate(t *testing.T) {
	word := "test"
	definition := "take measures to check the quality, performance, or reliability of something"

	dictionary := Dictionary{word: definition}

	newDefinition := "measurement to check quality or reability of something"
	dictionary.Update(word, newDefinition)

	assertDefinition(t, dictionary, word, newDefinition)
}
```

Hasil test
```
dictionary.Update undefined (type Dictionary has no field or method Update)
FAIL	example.com/hello/maps [build failed]
FAIL
```


## Write minimal code to test

```
func (d Dictionary) Update(word, newDefinition string) {
}
```

Hasil test
```
got "take measures to check the quality, performance, or reliability of something" want "measurement to check quality or reability of something" given "test"
```

Berarti parameter-nya sudah benar. Tinggal perbaiki kode pada method `Update`.


## Write enough code to pass the test

```
func (d Dictionary) Update(word, newDefinition string) {
	d[word] = newDefinition
}
```

Hasil test
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestUpdate$ example.com/hello/maps

ok  	example.com/hello/maps	0.390s
```

Good, udah pass test!

Masalahnya, kalo kita panggil method `Update` terus masukkin `word` dan `definition` baru, `word` itu bakal ketambah di `dictionary`. Karena ini cuma ubah nilai dari `word` yang udah ada, maka harus ada pengecekan lanjutan kan? Makanya kita harus upgrade kodenya.


## Write the test first

```
func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "take measures to check the quality, performance, or reliability of something"

		dictionary := Dictionary{word: definition}
		newDefinition := "measurement to check quality or reability of something"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "take measures to check the quality, performance, or reliability of something"

		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExists)

	})
}
```

Hasil test
```
# example.com/hello/maps [example.com/hello/maps.test]
Go\maps\dictionary_test.go:61:10: dictionary.Update(word, newDefinition) (no value) used as value

Go\maps\dictionary_test.go:73:10: dictionary.Update(word, definition) (no value) used as value

Go\maps\dictionary_test.go:75:23: undefined: ErrWordDoesNotExists
```


## Write minimal code to test

```
const (
	...
	ErrWordDoesNotExists = DictionaryErr("word does not exist")
)
```
```
func (d Dictionary) Update(word, newDefinition string) error {
	d[word] = newDefinition
	return nil
}
```

Hasil test
```
Go\maps\dictionary_test.go:75: got error %!q(<nil>) want "word does not exist"
FAIL
FAIL	example.com/hello/maps	0.156s
```


## Write enough amount of code to pass test

```
func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}
```

Hasil test
```
TestUpdate$ example.com/hello/maps

ok  	example.com/hello/maps	0.166s
```

Nah udah pass test!

Sekarang, tambah fitur `Delete`.


# Delete word

## Write test first

```
func TestDelete(t *testing.T) {
	word := "test"
	definition := "take measures to check the quality, performance, or reliability of something"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)
}
```

Hasil test
```
Go\maps\dictionary_test.go:85:13: dictionary.Delete undefined (type Dictionary has no field or method Delete)
FAIL	example.com/hello/maps [build failed]
FAIL
```

Saatnya tulis kode untuk method `Delete`.


## Write minimal code to test

```
func (d Dictionary) Delete(word string) {

}
```

Hasil test
```
got error %!q(<nil>) want "could not find the word"
FAIL
FAIL	example.com/hello/maps	0.387s
FAIL
```

Error message menunjukkan bahwa `error` yang didapat dari `Search` itu `nil`. Artinya, `word` masih belum dihapus dari `dictionary`. Memang ini yang diharapkan karena method `Delete` masih kosong. Tapi hasil test ini menandakan bahwa parameter di method `Delete` sudah benar sehingga kita hanya perlu mengisi method `Delete`.

Sekarang fix `Delete` to pass the test!


## Write enough code to pass test

```
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
```

Hasil test
```
TestDelete$ example.com/hello/maps

ok  	example.com/hello/maps	(cached)
```

Sudah pass test!


## Refactor

Gak banyak yang bisa di-refactor, tapi kita bisa tambah implementasi untuk handle delete word yang gak ada di dictionary.

```
func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "take measures to check the quality, performance, or reliability of something"
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)
		assertError(t, err, nil)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("non existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)
		assertError(t, err, ErrWordDoesNotExists)
	})
}
```

Hasilnya, dapat error jika `dictionary.Delete(word) has no value`.

Kita refactor juga method `Delete` agar bisa return `error`. Kita juga akan melakukan pengecekan dengan logika yang mirip dengan yang kita lakukan di method `Update`.
```
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
	case ErrNotFound:
		return ErrWordDoesNotExists
	default:
		return err
	}

	return nil
}
```

Basically kalo `error` yang direturn `Search` itu nil, maka kita hapus `word` dari dictionary. Kalo `Search` malah return `ErrNotFound`, method `Delete` bakal return `ErrDoesNotExists`.

Hasil test
```
^TestDelete$ example.com/hello/maps

ok  	example.com/hello/maps	(cached)
```

Nah, sudah pass test!


# Wrapping up

Yang kita pelajarin dari modul ini
- Buat maps
- Tambah item di maps
- Update item di maps
- Delete item di maps
- Error bertipe `const`

Yang kita lakuin di modul ini basically CRUD(Create, Read, Update, Delete). Tinggal tambah proses writing `dictionary` ke server maka sudah bisa jadi backend api.