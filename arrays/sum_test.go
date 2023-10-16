package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	got := Sum(numbers)
	want := 10
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func TestSumAll(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("with slice values", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{1, 9}, []int{})
		want := []int{3, 10, 0}
		checkSums(t, got, want)
	})
	t.Run("with empty slice", func(t *testing.T) {
		got := SumAll()
		want := []int{0}
		checkSums(t, got, want)
	})
}
