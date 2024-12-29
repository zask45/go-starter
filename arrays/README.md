# Arrays and slices

Inget, TDD. Jadi kita write test lebih dulu

## Write test

```
package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	expected := 15

	if got != expected {
		t.Errorf("got %d expected %d given %v", got, expected, numbers)
	}
}
```

## Run test

Hasil run test

```
c:\Users\Keysha\Documents\Go\arrays\sum_test.go:8:9: undefined: Sum
FAIL	example.com/hello/arrays [build failed]

````

### Write code minimally

Buat file baru `sum.go`. lalu, tulis function `Sum` untuk mengembalikan 0 terlebih dahulu

```
package arrays

func Sum(numbers [5]int) int {
	return 0
}
```

Hasil test
```
 got 0 expected 15 given [1 2 3 4 5]
FAIL
FAIL	example.com/hello/arrays	0.388s
```

## Write enough code to pass the test

Ubah `sum.go` jadi kayak gini

```
package arrays

func Sum(numbers [5]int) int {
	sum := 0

	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}

	return sum
}
```

Hasil test `ok`!


## Refactor

Kita refactor kodenya supaya bisa melakukan perulangan sesuai dengan length array. Ini bisa dilakukan dengan `range`.

Tinggal kosongin index di `for`. Buat variabel untuk tiap element di `range` array. Kita namain variabel tersebut `number`. Lalu tambahkan tiap `number` yang ada di range ke dalam `sum`.

```
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
```

# Slices

Gak mungkin kan kalo kita harus kasih tau length array tiap mau deklarasi array. Makanya, kita bisa pake `slices`.

Caranya mirip sama deklarasi array, cuma kosongin aja lengthnya kayak gini

```
numbers := [5]int{1, 2, 3, 4, 5}
```

Nah sekarang write test first.


## Write test

```
numbers := []int{1, 2, 3}

got := Sum(numbers)
want := 6

if got != want {
    t.Errorf("got %d want %d given, %v", got, want, numbers)
}
```

Sehingga full test-nya jadi kayak gini
```
package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d expected %d given %v", got, expected, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("got %d expected %d given %v", got, expected, numbers)
		}
	})
}
```

## Run the test

Hasil test

```
Go\arrays\sum_test.go:20:14: cannot use numbers (variable of type []int) as [5]int value in argument to Sum
FAIL	example.com/hello/arrays [build failed]
```

Nah berarti kita harus update function `Sum` pada file `sum.go` supaya bisa nerima slice sebagai argumen.


## Write minimal amount of code to pass the test

Objective dari fase ini tu cuma memastikan kalau function `Sum` bisa nerima `slices` sebagai argumen.

Jadi kita cuma update fungsi `Sum` jadi kek gini
```
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
```

Hasil test-nya
```
Go\arrays\sum_test.go:9:14: cannot use numbers (variable of type [5]int) as []int value in argument to Sum
FAIL	example.com/hello/arrays [build failed]
```

Liat error message-nya.

<span style="background-color: #FFFF0065">_cannot use numbers (variable of type [5]int) as []int value in argument to Sum_</span>


Nah ini yang bermasalah adalah test yg pertama. Kita perbaiki.


## Refactor

Ubah test jadi kek gini

```
package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d expected %d given %v", got, expected, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("got %d expected %d given %v", got, expected, numbers)
		}
	})
}
```

Lalu coba jalanin
```
go test -cover
```

Hasilnya
```
PASS
coverage: 100.0% of statements
ok      example.com/hello/arrays0.464s
```


# Sum two slices

## Write test file

```
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
```

## Run the test

Hasil test

```
Go\arrays\sum_test.go:30:9: undefined: SumAll
FAIL	example.com/hello/arrays [build failed]
```

`sumAll` belum didefinisikan.


## Write minimal code

func SumAll(numbersToSum...[]int) []int {
	return nil
}


## Run test again

Hasil test
```
invalid operation: got != want (slice can only be compared to nil)
FAIL	example.com/hello/arrays [build failed]
```

Maksudnya, kita gak bisa ngebandingin `slice 1` dengan `slice lainnya`.

Makanya, kita harus pake cara lain. 
Iterasi elemen satu-satu terus dibandingin? Terlalu panjang!

Makanya, kita pake `reflect.DeepEqual` aja.

```
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
```

Hasilnya
```
got [] want [3 9]
FAIL
FAIL	example.com/hello/arrays	0.208s
```

Selanjutnya kita perbaiki `SumAll` supaya bisa pass the test.


## Write enough code to make it pass

```
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
```

Kurang ngerti kodenya? Gapapa! 
Ini penjelasannya.

`make` itu digunain buat ngebuat slice sebanyak jumlah slice yang ditaruh di parameter. Dalam kasus ini berarti ada dua slice, yaitu `[]int{1, 2}` dan `[]int{0, 9}`.

Sedangkan untuk kode di bawah ini,
```
for i, numbers := range numbersToSum {
    sums[i] = Sum(numbers)
}
```

maksudnya adalah 
> _Untuk tiap elemen `numbers` pada `numbersToSum` (dengan index i), kita akan memanggil fungsi Sum() untuk menjumlahkan tiap elemen numbers._

Saat di-test hasilnya sudah `oke`.

Tapi jujur aja, kode-nya kurang kebaca kan? Makanya kita perlu refactor!


## Refactor

```
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
```

Maksudnya
> _Untuk tiap `numbers` pada `numbersToSum`, append `sums` kemudian isi dengan hasil dari Sum(numbers)_


# Sum Tails

Sekarang, kita cuma akan menjumlahkan semua elemen slice kecuali elemen yang paling depan.


## Write test

```
func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of sum slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
```


## Run test

Hasil 

```
undefined: SumAllTails
```

## Write code minimally

```
func SumAllTails(numbersToSum ...[]int) []int {
	return nil
}
```


## Run the test again

Hasil test

```
--- FAIL: TestSumAllTails (0.00s)
    --- FAIL: TestSumAllTails/make_the_sums_of_sum_slices (0.00s)
        c:\Users\Keysha\Documents\Go\arrays\sum_test.go:47: got [] want [2 9]
    --- FAIL: TestSumAllTails/safely_sum_empty_slices (0.00s)
        c:\Users\Keysha\Documents\Go\arrays\sum_test.go:56: got [] want [0 9]
FAIL
FAIL	example.com/hello/arrays	0.437s
FAIL
```


## Write enough code to pass the test

Intinya ya, kalo `slice`-nya empty hasil SumAllTails == 0. Kalo `slice`-nya ada isi, jumlahkan mulai dari index ke-1 sampai akhir (index pertama == 0).

```
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
```

> Untuk tiap `numbers` pada range `numbersToSum` kita lakukan perulangan. Jika slice `numbers` kosong, maka hasil `sum == 0`. Jika tidak kosong, maka jumlahkan tail.


Hasil test-nya `ok`.


## Refactor

Refactor `sum_test.go` jadi kayak gini

```
package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	expected := 6

	if got != expected {
		t.Errorf("got %d expected %d given %v", got, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of sum slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
```


## Wrapping up

Inti yang udah di pelajarin di modul ini

- Arrays
- Slices: kayak array tapi gak perlu fixed capacity
- `len` untuk cari tahu length dari array atau slice
- `reflect.DeepEqual` untuk ngebandingin slices