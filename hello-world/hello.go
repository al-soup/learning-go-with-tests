package main

import "fmt"

const (
	spanish = "ES"
	french  = "FR"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// Public functions start with a capital letter, private functions start with a lowercase letter
func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

// Return type is 'named return value'. If the variable it not assigned it will default to empty string
func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// named return values don't need to be explicitly returned
	return
}

func main() {
	fmt.Println(Hello("Soup", ""))
}
