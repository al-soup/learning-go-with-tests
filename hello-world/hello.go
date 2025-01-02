package main

import "fmt"

func Hello(name, lang string) string {
	var prefix string

	switch lang {
	case "FR":
		prefix = "Bonjour, "
	case "ES":
		prefix = "Hola, "
	default:
		prefix = "Hello, "
	}

	if name == "" {
		name = "World"
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Soup", ""))
}
