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

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		// `Call` handle possible multiple return values as slice
		for _, res := range valFnResult {
			walkValue(res)
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
