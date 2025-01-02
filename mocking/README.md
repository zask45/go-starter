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


# Implement Mocking

Kalau kita run test pasti bakal kerasa kalo `test`-nya itu berjalan selama 3 detik. Ini karena pas melakukan test fungsi `Countdown` dipanggil dan fungsi ini kan butuh 3 detik untuk print seluruh string.

Sekarang bayangin , setiap kali test kita bakal ngehabisin waktu 3 detik untuk nungguin `Countdown` print seluruh isi string. Gak produktif banget kan? 

Makanya, kita bisa pake konsep `mocking`. Maksudnya, kita bisa ganti `time.Sleep` jadi fungsi lain yang gak bener-bener tidur.

Caranya? Pake `dependency injection` untuk melakukan `Spy`. Spy ini teknik buat mantau berapa kali fungsi dipanggil atau mantau argumen apa yang dipanggil tiap pemanggilan.

Bingung? 
Coba perhatiin contoh ini!


## Write the test first

Untuk mock `time.Sleeper` kita buat dulu interface untuk handle Sleeper yang bisa di-inject ke `test`. 

```
type Sleeper interface {
	Sleep()
}
```

Tiap method yang implementasi `Sleep` akan dianggap sebagai bagian dari `Sleeper`.

Sekarang, buat method yang implementasi fungsi `Sleep` sekaligus ngelakuin `spy` ke `Sleeper`.

```
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
```

Karena `spySleeper` ini fungsinya mata-matain berapa kali `Sleeper` dipanggil, maka increment variabel `Calls` tiap kali `spySleeper.Sleep` dipanggil.


Sekarang inject `spySleeper.Sleep` di `Countdown` pada `mock_test.go`
```
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 instead of %d", spySleeper.Calls)
	}
}
```

Coba run test dan liat hasilnya
```
too many arguments in call to Countdown
	have (*bytes.Buffer, *SpySleeper)
	want (io.Writer)
FAIL	example.com/hello/mocking [build failed]
FAIL
```


## Write minimal amount of code to test

Sekarang coba perbaiki parameter pada `Countdown` supaya bisa inject `Sleeper`.

```
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(writer, finalWord)
}
```

Tapi kalo kita run `mock.go`, hasilnya
```
\mock.go:34:12: not enough arguments in call to Countdown
        have (*os.File)
        want (io.Writer, Sleeper)
```

Kalo diliat dari error messagenya, kita diminta perbaiki kode di bawah ini supaya bisa masukkin `Sleeper` sebagai parameter.
```
func main() {
	Countdown(os.Stdout)
}
```

Nah sekarang coba buat `real Sleeper` yang implementasi interface Sleeper di `Countdown`. Karena ini bener-bener untuk implementasi di sisi user bukan cuma untuk di-test, gunain `timeSleep` untuk nge-halt 1 second tiap method ini dipanggil.

```
type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
```
```
func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
```

Hasil run `mock.go` sudah berjalan semestinya.

Sekarang coba run `test`
```
Go\mocking\mock_test.go:25: not enough calls to sleeper, want 3 instead of 0
FAIL
FAIL	example.com/hello/mocking	3.381s
FAIL
```

## Write enough code to make it pass

Sekarang panggil `sleeper.Sleep()` di `Countdown`.
```
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(writer, finalWord)
}
```

Hasil test
```
ok  	example.com/hello/mocking	(cached)
```

Btw test-nya gak selama 3 detik lagi karena kita inject `SpySleeper` sebagai `Sleeper` instead of `DefaultSleeper`.

Bingung?

Jadi intinya, pas `Countdown` dipanggil di file `mock.go`, dia bakal ngelakuin `sleep` selama 1 detik tiap `sleeper.Sleep` dipanggil karena yang di-inject itu `DefaultSleeper`.

Sedangkan, pas kita panggil `Countdown` di test, kita cuma ngitung berapa kali `sleeper.Sleep` dipanggil karena yang diinject itu `SpySleeper`.

Jadi dengan make interface yang sama yaitu `Sleeper`, kita bisa inject dua struct dengan perilaku berbeda ke dalem satu function. Jadi lebih multifungsi kan method-nya?