package main

import "math"

// In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// A method is a function with a receiver (it's associated with a specific type)
// Method declarations bind a name to a method and associate it with the receiver's base type
// Unlike regular functions which can be called independently (e.g., Area(rectangle))
// Methods can only be called on instances of their associated type ("things")
// Methods provide a way to add behavior to specific types
// Syntax of a method: func (receiverName ReceiverType) MethodName(args)
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

type Triangle struct {
	Base  float64
	Heigh float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Heigh * 0.5
}

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Height * rectangle.Width
// }
