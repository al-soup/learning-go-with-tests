package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	// Use the writer to send the greeting to the buffer in our test

	// fmt.Printf("Hello, %s", name) uses Fprintf internally like so:
	// `return Fprintln(os.Stdout, a...)`. For testing purposes this
	// dependency can now be passed down with DI
	fmt.Fprintf(writer, "Hello, %s", name)
}

// The Greet function is quite general purpose now. E.g. use it for HTTP calls
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "Horatio")

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
