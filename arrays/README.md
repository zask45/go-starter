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