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


## Still some problems


Kalau kita lihat, kita cuma test apakah kita manggil `sleeper.Sleep` 3x. Padahal kita juga harusnya mengkonfirmasi apakah `sleeper.Sleep` dipanggil sebelum menulis line baru seperti ini
```
Print("3")
Sleep
Print("2")
Sleep
Print("1")
Sleep
Print("Go!")
```

Karena itu, kita harus ubah testnya


## Write the test first

```
func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := `3
	2
	1
	Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)

		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string {
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter)
		}
	})

}
```

Pertama, kita test apakah method `Countdown` dapat melakukan print `3 2 1 Go!`. Lalu kita tambahkan juga test untuk mengecek apakah `sleep` dipanggil sebelum mencetak line baru.

Coba run test
```
c:\Users\Keysha\Documents\Go\mocking\mock_test.go:12:22: undefined: SpyCountdownOperations
c:\Users\Keysha\Documents\Go\mocking\mock_test.go:27:23: undefined: SpyCountdownOperations
c:\Users\Keysha\Documents\Go\mocking\mock_test.go:31:4: undefined: write
c:\Users\Keysha\Documents\Go\mocking\mock_test.go:32:4: undefined: sleep
c:\Users\Keysha\Documents\Go\mocking\mock_test.go:33:4: undefined: write
c:\Users\Keysha\Documents\Go\mocking\mock_test.go:34:4: undefined: sleep
c:\Users\Keysha\Documents\Go\mocking\mock_test.go:35:4: undefined: write
c:\Users\Keysha\Documents\Go\mocking\mock_test.go:36:4: undefined: sleep
c:\Users\Keysha\Documents\Go\mocking\mock_test.go:37:4: undefined: write
FAIL	example.com/hello/mocking [build failed]
FAIL
```


## Write minimal amount of code to test

Sekarang, buat `SpyCountdownOperations`

```
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"
```

Jadi intinya `SpyCountdownOperations` ini untuk handle `Sleep` dan `Write`. Karena `SpyCountdownOperations` implement dua fungsi tersebut, maka  bisa dibilang `SpyCountdownOperations` itu bagian dari interface `Sleeper` dan `io.Writer`.

Maka dari itu kalo kita lihat kode ini di `test`

```
spySleepPrinter := &SpyCountdownOperations{}

Countdown(spySleepPrinter, spySleepPrinter)
```

`SpyCountdownOperations` bisa di inject ke dalam `Countdown` baik sebagai argumen pertama maupun kedua karena `SpyCountdownOperations` itu bagian dari interface `Sleeper` dan juga `io.Writer`. Coba cek apa saja parameter dalam `Countdown`

```
func Countdown(writer io.Writer, sleeper Sleeper)
```

Nah sekarang coba jalankan test

Test pertama: `prints 3 to Go!`
```
Go\mocking\mock_test.go:21: got "3\n2\n1\nGo!" want "3\n\t2\n\t1\n\tGo!"
FAIL
FAIL	example.com/hello/mocking	0.384s
```

Test kedua: `sleep before every print`
```
ok  	example.com/hello/mocking	(cached)
```

## Fix test 

Kalo kita lihat, masalah di `test pertama` cuma masalah `tab`. Coba kita perbaiki test pertama karena memang hasil yang diinginkan tidak memakai `tab`. Hapus `tab`

```
t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)

		}
	})
```

Hasil test
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestCountdown$ example.com/hello/mocking

ok  	example.com/hello/mocking	(cached)
```

Hasil test sudah oke nih. Bakal lebih oke lagi kalo kita bisa adjust `sleep time`-nya kan?


# Extending Sleeper to be configurable

Sekarang coba create `ConfigurableSleeper` yang bisa adjust durasi `sleep time`.


## Write the test first

Sebelum buat test, buat struct `ConfigurableSleeper` dulu
```
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}
```

Struct ini punya 2 field. Field pertama itu `duration` yang bertipe _time.Duration_. Field kedua itu `sleep`. Field ini bentuknya _function_ yang berparameter _timeDuration_.

```
type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}
```

`ConfigurableSleeper` ini yang bakal handle `Sleep selama xx detik`.

Sekarang, kita buat `SpyTime` supaya pas test kita gak perlu bener-bener `sleep selama 5 detik`. `SpyTime` ini bisa dibilang dibuat untuk diinject sebagai argumen kedua `ConfigurableSleeper`. Argumen kedua `ConfigurableSleeper` ini bentuknya function makanya pastiin ada method `spyTime.Sleep` supaya bisa diinject.

```
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
```

Sekarang coba write test

```
func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v instead", sleepTime, spyTime.durationSlept)
	}
}
```

Hasil run test
```
sleeper.Sleep undefined (type ConfigurableSleeper has no field or method Sleep, but does have field sleep)
FAIL	example.com/hello/mocking [build failed]
FAIL
```

## Write enough code to pass test

Sekarang coba write method `ConfigurableSleeper.Sleep` di `mock.go`

```
func (c *ConfigurableSleeper) Sleep() {

}
```

Hasil run test
```
should have slept for 5s but slept for 0s instead
FAIL
FAIL	example.com/hello/mocking	0.379s
FAIL
```

Berarti konfigurasi method `Sleep` sudah benar. Sekarang perbaiki isi methodnya.

## Write enough code to pass test

```
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
```

Hasil test

```
ok  	example.com/hello/mocking	0.397s
```

Nah, test sudah pass. Sekarang kita tinggal gunakan `ConfigurableSleeper` di `main`.


## Cleanup and Refactor

Ubah `sleeper` di `main` agar menggunakan `ConfigurableSleeper`

```
func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
```


## Isn't mocking evil?

`Mocking` itu gak evil ya. Ya bisa jadi evil kalo terlalu complicated.

Biasanya ya mocking tu jadi over-complicated kalo programmernya gak _listen to the test_ dan gak _respect refactoring stage_.

Jadinya hasilnya ya terlalu complicated untuk di `mock`. Cirinya apa kalo kode terlalu complicated?
- Terlalu banyak dependency to mock
- Test terlalu concern pada _implementation detail_
- Kode terlalu banyak dependency kecil

Kalo udah kayak gitu, ikutin feeling kamu kalo ternyata kode-nya udah terlalu complicated dan coba
- Pisahkan kode jadi bagian-bagian kecil dengan tanggung jawab spesifik
- Uji apa yang kode lakukan, bukan bagaimana kode melakukannya
- Gabungkan beberapa dependency yang sering digunakan bersama pada satu abstraction layer


Contoh penerapan abstraction layer yang buruk
```
type UserService struct {
    userRepository UserRepository
    emailService   EmailService
    logger         Logger
}
```

Contoh penerapa abstraction layer yang lebih baik
```
type NotificationService struct {
    emailService EmailService
    logger       Logger
}

type UserService struct {
    userRepository UserRepository
    notification   NotificationService
}
```

Pada akhirnya supaya bisa melakukan mocking kita perlu
- Sederhanakan kode dengan mengurangi dependency yang terlalu detail
- Pisahkan modul yang terlalu kompleks agar mudah diuji
- Fokus pada hasil yang diharapkan, bukan cara internalnya bekerja


## Wrapping up

Intinya mocking itu ngebuat `objek baru` untuk `niru objek` yang ingin diuji. Contohnya, kita mock `DefaultSleeper` jadi `SpySleeper` supaya bisa diuji.

Yang harus diperhatiin, baik `SpySleeper` dan `DefaultSleeper` harus menjadi bagian dari interface yang sama supaya keduanya bisa bergantian di-`inject` ke satu fungsi.

Perhatiin dua kode ini untuk paham `mocking`

Kode untuk testing
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
````

Kode di `mock.go`
```
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(writer, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
```