# Mocking

Kita bakal buat program untuk menghitung `Countdown` seperti ini
```
3
2
1
Go!
```

Untuk melakukan test pada program seperti ini kita bakal pake `iterative test driven approach`.


## Write the test first

Pertama kita akan test apakah program bisa print `3`.

```
package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Hasil test
```
Go\mocking\mock_test.go:11:2: undefined: Countdown
FAIL	example.com/hello/mocking [build failed]
FAIL
```

Sekarang write fungsi `Countdown`


## Write minimal amount of code to test

```
package main

import (
	"fmt"
	"io"
)

func Countdown(writer io.Writer) {
	fmt.Fprintf(writer, "")
}
```

Hasil test
```
Go\mocking\mock_test.go:17: got "" want "3"
```

## Write enough code to pass the test

Coba ubah `Countdown` jadi seperti ini
```
func Countdown(writer io.Writer) {
	fmt.Fprintf(writer, "3")
}
```

Hasil test
```
ok  	example.com/hello/mocking	0.391s
```


## Refactor

Coba tambahkan fungsi main untuk print
```
package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(writer io.Writer) {
	fmt.Fprintf(writer, "3")
}

func main() {
	Countdown(os.Stdout)
}
```

```
PS C:\...\Documents\Go\mocking> go run mock.go
3
```

Setelah program sudah terbukti bisa print string `3`, kita akan coba apakah program bisa print
```
3
2
1
Go!
```


## Write the test first

```
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
	2
	1
	Go!
	`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Hasil test
```
Go\mocking\mock_test.go:21: got "3" want "3\n\t2\n\t1\n\tGo!\n\t"
FAIL
FAIL	example.com/hello/mocking	0.378s
```

Kalau kita lihat hasil test bis kelihatan kalau ada simbol `\t` yang berarti `tab`.

Kita perbaiki test jadi seperti ini agar tidak ada `tab`.
```
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Kalau test dijalankan, hasilnya
```
Go\mocking\mock_test.go:20: got "3" want "3\n2\n1\nGo!"
````

Karena `\n` itu artinya `new line`, maka hasil test tersebut ekuivalen dengan
```
3
2
1
Go!
```


## Write enough code to pass the test

Gunakan perulangan untuk mencetak `3 2 1`
```
func Countdown(writer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprintf(writer, "Go!")
}
```

Hasil test
```
ok  	example.com/hello/mocking	(cached)
```


## Refactor

Pisahkan `nilai awal countdown` dan `string akhir countdown` ke variabel tersendiri.

```
const finalWord = "Go!"
const countdownStart = 3

func Countdown(writer io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprintf(writer, finalWord)
}
```

Hasil test
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestCountdown$ example.com/hello/mocking

ok  	example.com/hello/mocking	(cached)
```


## Implement time.Sleep

Sekarang buat agar program nunggu 1 detik sebelum print line berikutnya

```
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(writer io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(writer, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
```

Cara test-nya lihat di console apakah benar compiler nunggu satu detik sebelum print line berikutnya.

Hasilnya tidak bisa saya videokan tapi benar bahwa program nunggu 1 detik sebelum print line selanjutnya!


