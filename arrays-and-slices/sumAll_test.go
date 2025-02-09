package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{1, 2}

	got := SumAll(nums1, nums2)
	want := []int{15, 3}

	// Compare values of slices
	// This function is not type safe and will still compile if you compare slices to a number for example.
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given, %v and %v", got, want, nums1, nums2)
	}
}
