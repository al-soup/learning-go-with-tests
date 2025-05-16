package main

import "fmt"

type BitCoin int

// The Stringer interface is defined in the fmt package
type Stringer interface {
	String() string
}

// Creating a method on a type declaration is the same as it is on a struct
func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance BitCoin
}

// In Go, value receivers `(func (w Wallet))` create a copy of the struct, so modifications
// are lost after the method returns. Pointer receivers `(func (w *Wallet))` operate on the
// original struct, allowing permanent modifications. Use value receivers for immutability
// and small structs, and pointer receivers when you need to modify the struct or avoid
// copying large data structures
func (w *Wallet) Balance() BitCoin {
	return w.balance
}

// When you call a function or a method the arguments are copied
func (w *Wallet) Deposit(amount BitCoin) {
	// Using the ampersand operator (&) to get a pointer to a variable and print it with %p
	fmt.Printf("address of balance in Deposit is %p \n", &w)
	w.balance += amount
}
