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

# What is concurrency?

Concurrency tuh apa?
Concurrency tuh bisa dibilang kayak kemampuan processor ngerjain hal lain sambil nunggu proses yang satu-nya selesai.

Kalo diibaratin nih ya, misal kamu punya warung jus yang jual bakso bakar juga. Suatu hari, kamu dapet orderan `jus` dan `bakso bakar`. Nah kamu proses lah jus itu di blender. Daripada kamu bengong nunggu blender selesai, bukannya lebih baik kamu nyiapin arang buat ngebakar bakso?

Nah kira-kira kayak gitu.

Bedanya, di sini pas kita nunggu website buat `respond`, kita bisa minta komputer buat ngeproses request yang baru.





