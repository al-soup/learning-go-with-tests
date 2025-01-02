package main

import "fmt"

func Hello(name string) string {
	const prefix = "Hello, "

	if name == "" {
		name = "World"
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
