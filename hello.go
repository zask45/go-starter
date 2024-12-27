package main

import (
	"fmt"
)

const helloInJapanese = "Konnichiwa"

func Greeting(name string) string {
	if name == "" {
		return helloInJapanese
	}

	return helloInJapanese + " " + name
}

func main() {
	fmt.Println(Greeting("Yuta-kun"))
}
