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

