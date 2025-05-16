package main

import "fmt"

type Wallet struct {
	balance int
}

// In Go, value receivers `(func (w Wallet))` create a copy of the struct, so modifications
// are lost after the method returns. Pointer receivers `(func (w *Wallet))` operate on the
// original struct, allowing permanent modifications. Use value receivers for immutability
// and small structs, and pointer receivers when you need to modify the struct or avoid
// copying large data structures
func (w *Wallet) Balance() int {
	return w.balance
}

// When you call a function or a method the arguments are copied
func (w *Wallet) Deposit(amount int) {
	// Using the ampersand operator (&) to get a pointer to a variable and print it with %p
	fmt.Printf("address of balance in Deposit is %p \n", &w)
	w.balance += amount
}
