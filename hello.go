package main

import (
	"fmt"
)

const (
	spanish = "spanish"
	french  = "french"

	hello          = "Konnichiwa"
	helloInSpanish = "Hola"
	helloInFrench  = "Bonjour"
)

func Greeting(name string, language string) string {
	if name == "" {
		return hello
	}

	return greetingPrefix(language) + " " + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloInSpanish
	case french:
		prefix = helloInFrench
	default:
		prefix = hello
	}

	return
}

func main() {
	fmt.Println(Greeting("Yuta-kun", ""))
	fmt.Println(Greeting("Yuta-kun", spanish))
	fmt.Println(Greeting("Yuta-kun", french))
}
