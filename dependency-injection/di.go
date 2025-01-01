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
