package reflection

import "reflect"

// Reflection in computing is the ability of a program to examine its own structure, particularly through types; it's a form of metaprogramming
// Through reflection you inspect "anything/unknown" that is passed to your function (like `interface{}`)
// that can not be checked at compile- but at runtime. Plus: flexible. Downside: untyped, slow because of runtime checks. Avoid reflection.

// NOTE this is the goal: golang challenge is to write a function `walk(x interface{}, fn func(string)` which takes a struct `x` and calls `fn` for all strings fields found inside.
// `any` is an alias for `interface{}`

// TODO continue with refactoring this: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reflection#refactor-5
func walk(x any, fn func(input string)) {
	val := getValue(x)

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		// extract underlying value in case of a pointer
		val = val.Elem()
	}

	return val
}
