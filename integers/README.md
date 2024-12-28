# Integers

Karena kita mau stick di `TDD (Test Driven Development)`, kita buat test-nya dulu.


## Write the test first

```
package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}
```

## Run the test

Coba run test-nya, terus inspect hasilnya.

```
Go\integers\adder_test.go:6:9: undefined: Add
FAIL	example.com/hello/integers [build failed]
```

Nah hasilnya belum ada method `Add`. Karena itu, kita buat method add.


## Write minimally code

Disini kita cuma nulis kode secara minimal, alias `return 0`

```
func Add(x, y int) int {
	return 0
}
```

## Run the test again

Coba run test-nya lagi. Hasilnya:

```
\Go\integers\adder_test.go:10: expected 4 but got 0
FAIL
```


## Write enough code to make it pass

Ubah kode agar bisa menambahkan `x` dan `y`

```
func Add(x, y int) int {
	return x + y
}
```

Hasil test
```
ok  	example.com/hello/integers	(cached)
```


## Refactor

Gak ada yang bisa di-refactor, so we skip this step.


## Wrapping up

Jadi inti dari bagian ini itu adalah implementasi dari TDD, yaitu

- Tulis test
- Run test
- Run code minimally
- Test again
- Wtite enough code to make it pass
- Refactor
