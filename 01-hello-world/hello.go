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

// Return type is 'named return value'. It improves the readability of the code.
// This is treated as a variable that is declared at the top of the function.
// If the variable it not assigned in this case it will default to empty string
// A return statement without arguments returns the named return values. This is known as a "naked" return.
// Only use naked returns in short functions.
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
