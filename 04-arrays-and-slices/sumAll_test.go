package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}

	// reflect.DeepEqual is not type-safe. E.g it will compile if you compare slice and string
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
