package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("slice of any size", func(t *testing.T) {
		numbers := []int{5, 4, 3}
		got := Sum(numbers)
		want := 12
		if got != want {
			t.Errorf("got: %d, got: %d", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{0, 9}
		got := SumAllTails(slice1, slice2)
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
