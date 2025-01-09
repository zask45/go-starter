# Select

Di modul kali ini, kita bakal buat `WebsiteRacer`.

Programnya kayak gimana tuh? Well, simple aja. Kita `GET` 2 url website dan liat mana yang lebih cepet di return.

Sekarang coba write test-nya dulu.

## Write test first

```
func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Hasil test
```
...\Go\select\race_test.go:10:9: undefined: Racer
```


## Write minimal amout of code to test

```
package racer

func Racer(a, b string) string {
	return ""
}
```

Hasil test
```
\Go\select\race_test.go:13: got "" want "http://www.quii.dev"
FAIL
FAIL	example.com/hello/select	0.414s
FAIL
```

## Write enough code to pass the test

```
package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}

	return b
}
```

Intinya ya, kita pake `time.Now` untuk nyatet waktu sekarang. Terus kita `GET` url `a`. Terus kita kurangin waktu setelah selesai GET dengan `startA` dengan pake `time.Since`. Terus lakuin hal yang sama untuk ngukur waktu pas `GET` url `b`.

Hasil test
```
ok  	example.com/hello/select	(cached)
```

## Problems

Test kita pass kan? Emang ada masalah?

Ada! Masalah kode kita ada di `url` yang di-test.

Kita pake url asli untuk test dimana ini gak disaranin untuk proses testing.

Kenapa?

Karena idealnya kita gak boleh bergantung pada sumber lain untuk `testing`. Keyword di `bergantung` ya. 

Bergantung ke sumber lain pas testing itu bisa ngebuat test jadi lama. Ini nyusahin banget kalo ternyata yg di-testnya banyak dan berulang.

Karena itu, kita bisa buat `mock http server` untuk testing. Balik lagi, ini supaya kita gak bergantung sama dependency luar untuk testing.

Makanya, coba ganti `test` jadi kek gini
```
package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Untuk mock http server itu simple sih. Kita buat servernya. Terus untuk ambil URL dari server tinggal pake syntax `ServerName.URL`.

Cara buat `test server`-nya pun cukup simple ya. Tinggal pake `httptest.NewServer` terus masukkin `http.HandlerFunc` sebagai parameter. 

Terus isi parameter `HandlerFunc ` dengan anon function. Parameternya diisi `http.ResponseWriter` dan pointer `*http.Request`. Body function diisi sesuai kebutuhan. Yang slow server kita minta sleep `20 milisecond` supaya lebih slow dari server yang fast.

Kira-kira kayak gini contoh pembuatan mock server
```
slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    time.Sleep(20 * time.Millisecond)
    w.WriteHeader(http.StatusOK)
}))
```

Hasil test
```
unning tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestRacer$ example.com/hello/select

ok  	example.com/hello/select	0.661s
```

Hasilnya udah `ok` ya!

Btw buat `real http server` di `Go` ya mirip-mirip di atas itu. Bedanya kita gak pake `httptest.NewServer` aja.

Testing udah oke, berarti sekarang kita masuk ke tahap selanjutnya `Refactor`!

## Refactor

Kalau kita liat-liat ada pengulangan pas ngukur durasi `GET` url `a` sama `b`. Daripada ada duplikasi kek gitu lebih baik buat fungsi aja kan?

```
func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)

	duration := time.Since(start)

	return duration
}
```

Hasil test
```
C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestRacer$ example.com/hello/select

ok  	example.com/hello/select	(cached)
```

Setelah itu, coba refactor `test` supaya kita bisa langusng panggil function tiap ingin membuat `mock server`
```
package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func makeDelayedServer(delayTime time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delayTime)
		w.WriteHeader(http.StatusOK)
	}))
}
```

Kalo kita liat kode di atas, kita nemu syntax yang asing kan? 

`defer`

Apa tuh?

`defer` itu digunain untuk manggil fungsi lain di akhir fungsi. Kira-kira begini. Dengan pake keyword `defer`, kita bakal manggil `slowServer.Close` dan `fastServer.Close` di akhir fungsi `TestRacer`. Karena fungsi `defer` yang kayak gini, makanya defer sering dipake untuk _close connection_.

Omong-omong ini hasil test-nya
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestRacer$ example.com/hello/select

ok  	example.com/hello/select	0.750s
```

Sudah ok ya.

# Kode full

Sebelum masuk ke bagian lain, ini kode full nya.

```
package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)

	duration := time.Since(start)

	return duration
}
```
```
package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func makeDelayedServer(delayTime time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delayTime)
		w.WriteHeader(http.StatusOK)
	}))
}
```

# Synchronising process

Ngapain kita ngecek website mana yang paling cepat kalo kita bisa langsung return website tercepat?

Nah inilah gunanya `Select`.

Basically, kita jalanin kedua website itu di waktu yang bersamaan. Terus kita return yang paling cepet. 

```
package racer

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
```

Keliatan kan ya dari kodenya?

Mekanismenya cukup simple. Kita buat function `ping` untuk `GET` url secara concurrent.

Di `ping` jangan lupa buat `channel` supaya gak terjadi `race condition`. Kita pake `struct kosong` di channel karena kita gak butuh return. Kita cuma perlu untuk request `GET` secara concurrent.

Nah `select` ini fungsinya apa?

`select` ini berguna buat liat `channel` mana yang pertama dibuat. Kalo `ping(a)` ngebuat channel duluan maka return-nya bakal `a`. Sebaliknya kalo `ping(b)` buat `channel` duluan, maka return-nya bakal `b`.

Hasil test
```
ok  	example.com/hello/select	0.709s
```

# Timeouts

Sekarang, kita bakal ngebuat program buat return `error` kalo server butuh waktu lebih dari `10 second` untuk respond request.


## Write the test first

Revisi `TestRacer` jadi kayak gini
```
t.Run("compares speeds of servers, returning fastest url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10 secs", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(20 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}
```

Hasil test
```
Go\select\racer_test.go:36:13: assignment mismatch: 2 variables but Racer returns 1 value
FAIL	example.com/hello/select [build failed]
FAIL
```

Sekarang buat `Racer` untuk return error.

## Write minimal amount of code to test

Revisi `Racer` 
```
func Racer(a, b string) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	}
}
```

Hasil test
```
FAIL: TestRacer (20.00s)
    --- FAIL: TestRacer/returns_an_error_if_a_server_doesn't_respond_within_10_secs (20.00s)
        c:\...\Go\select\racer_test.go:39: expected an error but didn't get one
FAIL
FAIL	example.com/hello/select	20.751s
```

Berdasarkan error message yang diterima berarti parameternya udah bener. Tinggal benerin hasil `return`-nya aja.


## Write enough code to pass the test

```
func Racer(a, b string) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
```

Kita bisa pake `case` untuk send `error` kalo function `Racer` udah jalan lebih dari `10 seconds`. Untuk ngecek apakah fungsinya udah jalan lebih dari 10 detik atau belum, bisa manfaatin fungsi bawaan `time.After`.

Btw notice `time.After` pake `<-` juga? Ini karena sama kayak `ping`, `time.After` ini juga return channel untuk ngirim sinyal sesuai waktu yang ditentuin. Karena return channel, makanya pake `<-`.

Btw, hasil test-nya
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestRacer$/^returns_an_error_if_a_server_doesn't_respond_within_10_secs$ example.com/hello/select

ok  	example.com/hello/select	20.756s
```

Udah `ok`, sih.

Tapi test-nya lama bener. 20 seconds!

Bisa gak test-nya dipercepat?


## Slow test

Bisa!

Test-nya bisa dipercepat.

Caranya?
Inject `time.Duration` sebagai `timeout` ke `Racer`. 

Gak ngerti? 
Intinya mah kita buat `Racer` supaya return error sesuai timeout yang dimasukkin ke parameter. 

Biar gak bingung langsung liat implementasinya aja
```
var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
```

Liat ya? Kita masukkin `10 second` sebagai `timeout`. Jadi setelah 10 second, kita bakal return error.

Jadi karena kita pake timeout sebagai parameter, efeknya kita bisa test apakah bener `Racer` return error sesuai argumen yang dimasukkin sebagai timeout.

Sehingga test kedua bisa direvisi kek gini.
```
t.Run("returns an error if a server doesn't respond withthe specified time ", func(t *testing.T) {
	server := makeDelayedServer(15 * time.Millisecond)

	defer server.Close()

	_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)

	if err == nil {
		t.Errorf("expected an error but didn't get one")
	}
})
```

Kita buat delayed server selama `15 millisecs`. Terus kita cek apakah `ConfigurableRacer` berjalan selama lebih dari `10 milisecs`. Kalo iya, maka hasil test bakal `ok`.

Sekarang liat hasil test-nya
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestRacer$/^returns_an_error_if_a_server_doesn't_respond_within_10_secs$ example.com/hello/select

ok  	example.com/hello/select	0.743s
```

Sip, udah `ok` nih!


## Refactor

Kalo kita liat `test pertama` masih belum check `error` yang di-return `Racer`.

Nah tambahin kode buat check error tersebut.

```
t.Run("compares speeds of servers, returning fastest url", func(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got, err := Racer(slowURL, fastURL)

	if err != nil {
		t.Errorf("didn't expect any error")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
})
```

Hasil test
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestRacer$ example.com/hello/select

ok  	example.com/hello/select	0.703s
```

## Wrapping out

Kesimpulan dari modul ini

- `select` digunain untuk liat channel mana yang pertama terbentuk.
- `httptest` bisa dimanfaatin untuk buat _mock server_. 
- Cara buat mock server: `httptest.NewServer`