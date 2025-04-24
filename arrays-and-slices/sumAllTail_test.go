package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {

	t.Run("sum tails of slices", func(t *testing.T) {
		got := SumAllTail([]int{1, 2}, []int{3, 4})
		want := []int{2, 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("should accecpt empyt slices", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}
