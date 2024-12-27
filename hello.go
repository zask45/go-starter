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

	switch language {
	case spanish:
		return helloInSpanish + " " + name
	case french:
		return helloInFrench + " " + name
	}

	return hello + " " + name
}

func main() {
	fmt.Println(Greeting("Yuta-kun", ""))
	fmt.Println(Greeting("Yuta-kun", spanish))
	fmt.Println(Greeting("Yuta-kun", french))
}
