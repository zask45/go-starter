# Concurrency

Coba liat kasus ini.

Temen satu tim kamu udah nulis fungsi `CheckWebsite` kayak gini.

```
package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
```

Fungsi ini ceritanya bakal ngecek status tiap `url` di array `urls`. Di akhir fungsi bakal return map dengan key `url` dan value `boolean`. Kalo url-nya ngasih good response berarti value-nya `true`, kalo bad response berarti `false`. Proses pengecekkan kalo url-nya oke atau gak dilakuin dengan manggil `wc (WebsiteChecker)`.

Nah temen satu tim kamu juga udah nyiapin test-nya nih.

```
package main

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
```

Intinya kalo urlnya bukan `"waat://furhurterwe.geds` maka url dianggap legit (good response).

Hasil test
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestCheckWebsites$ example.com/hello/concurrency

ok  	example.com/hello/concurrency	0.322s
```

Hasil test-nya udah `ok`.

Terus masalahnya dimana?

Masalahnya, ada ratusan website yang harus dicek di production. Temen-temen setim udah ngeluh kalo proses pengecekannya ini lelet banget.

Jadi sekarang, tugas kamu apa?
Ya, ngecepetin proses pengecekkannya!


## Write a test

Sekarang buat `test` untuk  benchmark kecepatan pemrosesan `CheckWebsite`

```
package main

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
```

Untuk menjalankan benchmark
```
go test -bench=BenchmarkCheckWebsite
```

Hasil test
```
PASS
ok      example.com/hello/concurrency   0.418s   
```

Kalo diliat-liat emang udah cukup cepet sih, tapi apa ada cara supaya lebih cepet?

Ada, manfaatin `concurrency`!
<br><br>

# Concurrency and Goroutine?

Concurrency tuh apa?
Concurrency tuh bisa dibilang kayak kemampuan processor ngerjain hal lain sambil nunggu proses yang satu-nya selesai.

Kalo diibaratin nih ya, misal kamu punya warung jus yang jual bakso bakar juga. Suatu hari, kamu dapet orderan `jus` dan `bakso bakar`. Nah kamu proses lah jus itu di blender. Daripada kamu bengong nunggu blender selesai, bukannya lebih baik kamu nyiapin arang buat ngebakar bakso?

Nah kira-kira kayak gitu.

Bedanya, di sini pas kita nunggu website buat `respond`, kita bisa minta komputer buat ngeproses request yang baru.

Kok bisa? Iya bisa.

Biasanya kode dalem fungsi bakal dibaca dari atas ke bawah. Pas lagi memproses satu line, kita bakal nunggu line tersebut selesai dibaca. Proses ini disebut dengan `blocking`.

Nah kalo kita mau supaya langsung loncat ke line selanjutnya tanpa blocking line di bawahnya, kita harus ngasih tau kalo kita mau lanjutin proses di bawahnya pake 'reader' yang lain. Nah inilah `goroutine`. Semacem thread baru gitu buat jalanin proses yang lain.

Cara pake `goroutine`? Tinggal pake keyword `go`.

```
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func() {
			results[url] = wc(url)
		}()
	}

	return results
}
```

Coba fokus ke `anonymous function` yang di dalem `for`. Itu salah satu contoh penerapan `goroutine`.

Coba kita test
```
FAIL: TestCheckWebsites (0.00s)
    c:\...Go\concurrency\check_website_test.go:28: got map[http://google.com:true waat://furhurterwe.geds:false] want map[http://blog.gypsydave5.com:true http://google.com:true waat://furhurterwe.geds:false]
FAIL
FAIL	example.com/hello/concurrency	0.412s
FAIL
```

Cuma ada 2 data yang masuk ke `results` padahal harusnya ada 3. 

Inilah salah satu efek dari concurrency. Ini mirip-mirip lah sama `thread`. Karena ada yang jalan di `goroutine` makanya kita gak bisa prediksi apa yang bisa terjadi selanjutnya. Bisa aja functionnya udah langsung resturn hasilnya padahal `goroutine` belum selesai melakukan proses. 

Makanya kalo gak di-handle ini bisa berabe, hasilnya bisa gak konsisten. Kadang kosong, kadang gak lengkap, kadang lengkap.
<br><br>

## Handle Goroutine

Well, kalo masalahnya cuma kita harus nunggu semua proses di `goroutine` selesai ya kita tinggal masukkin `time.Sleep` aja kan?

Kita freeze fungsinya selama beberapa seconds pas goroutinenya jalan. Berapa lamanya tinggal kira-kira aja.

Coba `sleep` selama `2 seconds`
```
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func() {
			results[url] = wc(url)
		}()
	}

	time.Sleep(2 * time.Second)

	return results
}
```

Hasil test
```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestCheckWebsites$ example.com/hello/concurrency

ok  	example.com/hello/concurrency	(cached)
```

Ini kita lagi beruntung. Karena ini jalan di `gorutine` bisa aja ada hal-hal yang gak terduga terjadi. Misalnya, `race condition`.

Race condition tu terjadi pas `goroutine` satu sama yang lain berebut sumber. Kalo di kasus ini mereka bisa berebut nulis `maps` nya pada waktu yang bersamaan. 

Nah kalo kayak gini gimana solusinya?

Pake `channels`!

## Channels

Channel tu apa? 

Channel tu bisa dibilang kayak memori supaya gorutine bisa saling sharing. Ini mudahin komunikasi santar proses.

Gimana cara implementasi `channels`?

```
package main

type WebsiteChecker func(url string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
```

Pertama kita bikin `channel`-nya dulu. Karena kita mau berbagi `url` sama `hasil pengecekkan url` antar goroutine, maka kita harus bikin `channel` untuk nampung 	`string` sama `bool`. 


Nah biar gampang, kita buat struct yang bisa nampung `string` sama `bool`. Namain struct tersebut sebagai `result`.

Terus kita tinggal masukkin struct `result` pas buat channel.
```
resultChannel := make(chan result)
```

Abis itu, pas perulangan kita kasih tau kalo kita mau buat `channel` untuk tiap `url` dan hasil dari `wc(url)`.

Terus kita masukkin hasil dari tiap nilai di `channels` ke `maps`
```
for i := 0; i < len(urls); i++ {
	r := <-resultChannel
	results[r.string] = r.bool
}
```

Coba kita test
```
PASS
ok      example.com/hello/concurrency   0.161s  
```

Sekarang coba kita liat hasil `benchmark`
```
ok      example.com/hello/concurrency   0.186s
```

`0.186s`! Hampir 3x lebih cepet dari pas pake `goroutine`!

Btw kayaknya kita belum bahas kenapa `channel` bisa mencegah `race condition` kan?

Ini karena tiap goroutine ngirim hasil ke `channel` bukan langsung ke `map`. Cara ini mastiin hanya ada 1 goroutine yang akses data pada satu waktu sehingga bisa cegah `race condition`.


## Wrapping up

Di modul ini kita belajar tentang 
- `concurrency` yang basically ngebuat kita bisa ngelakuin beberapa proses dalam 1 waktu.
- `goroutine` yang bisa ngebuat kita lakuin concurrency di _Go_.
- `channel` untuk nampung hasil dari proses `goroutine`.


## Make it work, make it right, make it fast

Ada quote dari `Kent Beck` yang sering dipake di `agile development`

> Make it work, make it right, make it fast

Kalo dipikir-pikir yang kita lakuin di modul ini kan ngebuat proses lebih cepet juga. Buktinya hasil benchmark pas kita nerapin `goroutine` bisa hampir 3x lebih cepet dari sebelumnya.
