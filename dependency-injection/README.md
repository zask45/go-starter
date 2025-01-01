# Dependency Injection

Sebenernya materi ini udah pernah lewat pas belajar `android dev`. Tapi jujur kurang paham karena kesannya kompleks banget haha.

Tapi ya, sebenernya pandangan yang nganggep `dependecy injection` ribet itu adalah kesalahpahaman.

Dependency injection itu
- Gak ngebuat design program kita jadi rumit banget
- Gak butuh framework tertentu
- Bisa memfasilitasi testing
- Bisa ngebuat kita nulis fungsi yang bersifat _general purpose_ which is great.

Biar lebih paham soal `dependecy injection` kita bakal ngebuat program sederhana yang bertujuan untuk `Greet` user.

Coba perhatiin fungsi ini
```
func Greet(name string) {
	fmt.Printf("Hello, %s", name)
}
```

Fungsi di atas ini bergantung sama fungsi `fmt.Println` untuk mencetak pesan. Nah, `fmt.Println` ini dependensi karena fungsi `Greet` gak bisa bekerja tanpa `fmt.Println`.

Nah sekarang coba bayangin, gimana caranya kita `test` fungsi ini? Fungsi ini gak return apa-apa, cuma ngelakuin printing kan?

Nah cara yang bisa kita lakuin untuk test fungsi ini itu dengan nge-`inject` dependency dari `printing` ke test.

Kedengeran kompleks banget tapi sebenernya `inject` itu ya `passing`. Jadi yang kita lakuin itu ya `passing dependency dari printing` supaya bisa di-test.

Nah sekarang, coba tulis test-nya dulu!


## Write the test first

```
package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Kita gunain `bytes.Buffer` supaya kita bisa menangkap output dalam memori sebelum dicetak ke `stdout`.

Hasil test

```
Go\dependency-injection\di_test.go:10:2: undefined: Greet
FAIL	example.com/hello/dependency-injection [build failed]
FAIL
```

Sekarang, tulis kode `Greet`.


## Write minimal amount of code to test

```
package dependencyinjection

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Printf("Hello, %s", name)
}
```

Kita masukkan pointer `bytes.Buffer` sebagai parameter. Parameter ini akan digunakan nanti untuk passing nilai dari output ke `test`. 

Hasil test
```
Go\dependency-injection\di_test.go:16: got "" want "Hello, Chris"
FAIL
FAIL	example.com/hello/dependency-injection	0.404s
```

Melihat hasil test berarti parameter-nya sudah benar. Tetapi hasil test menunjukkan bahwa `Greet` langsung cetak greeting ke `stdout` alias ke layar user, bukan ke `Buffer` yang disediain. Sedangkan supaya greeting bisa di pass ke test, maka harus lewat buffer dulu.

Makanya, kita harus ganti `fmt.Println` ke `fmt.Fprintf` supaya bisa pass nilainya ke `buffer`.


## Write enough amount of code to pass test

```
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
```

Hasil test
```
test -timeout 30s -run ^TestGreet$ example.com/hello/dependency-injection

ok  	example.com/hello/dependency-injection	(cached)
```


## Refactor

```
package main

import (
	"bytes"
	"fmt"
	"os"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
```

Gunakan package `main` agar file `go` bisa di run.

Hasil run
```
cannot use os.Stdout (variable of type *os.File) as *bytes.Buffer value in argument to Greet
```

Nah baca error messagenya.

_cannot use os.Stdout as *bytes.Buffer_

Supaya kita bisa menggunakan `os.Stdout`, kita ganti parameter `Greet` ke `io.Writer`. Jangan lupa import juga `io` agar bisa menggunakan _io.Writer_.
```
package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
```

Hasil run 
```
Hello, Elodie
```

Btw kalo kita coba test hasil test-nya masih `ok`

```
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestGreet$ example.com/hello/dependency-injection

ok  	example.com/hello/dependency-injection	0.318s
```

Kenapa ini bisa terjadi?

Greet masih bisa lolos test karena `bytes.Buffer` itu implementasi dari interface `Writer`.

Apa iya? 
Iya! Karena metode `bytes.Buffer` implementasi fungsi `Write`. Jadi secara otomatis metode Buffer dianggap sebagai implementasi `Writer`.
```
func (b *Buffer) Write(p []byte) (n int, err error)
```

Karena `io.Writer` bisa dibilang menaungi `bytes.Buffer` makanya pas parameter diganti test tetap bisa lolos. 

Simplenya, kalo sebelumnya kita gunakan `child` sebagai parameter, kalo kita ganti parameter tersebut jadi `parent`-nya maka harunsya gak ada masalah. 


## More on io.Writer

```
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	//"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Yuta")
}

func main() {
	fmt.Println("http://localhost:5001/")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
	// Greet(os.Stdout, "Elodie")
}
```

Coba run code ini.

Hasilnya

![](../img/{643BB023-96DA-4060-A0D7-A590E9AB11D8}.png)

Kalo kita perhatiin kode ini, bisa dibilang kita passing `http.ResponseWriter` sebagai `io.Writer`. Ini berhasil dilakukan karena `http.ResponseWriter` juga implementasi fungsi `Write` dari interface `io.Writer`. Makanya, proses inject ini berhasil.

Untuk detail kirim data ke server ini bakal dibahas di bab lain. Di sini kita cuma fokus ke bagian inject dependency `Writer`-nya aja.


## Wrapping up

Agar bisa mudah melakukan dependency injection, kita bisa gunain `interface` sebagai parameter. 

Ini mempermudah kita supaya kode bisa di-reuse lebih gampang. Contohnya, kita bisa reuse fungsi `Greet` di dalem `MyGreetHandler` karena Greet meminta interface `io.Writer` sebagai parameter. Kita juga bisa inject `Buffer` sebagai `io.Writer` karena _bytes.Buffer_ merupakan turunan dari _io.Writer_.