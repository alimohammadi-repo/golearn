package array

import (
	"reflect"
	"testing"
)

func TestSum1(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum1(numbers)
	want := 15

	if got != want {

		t.Errorf("got %d want %d given, %v", got, want, numbers)

	}

}
func TestSum2(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum2(numbers)
	want := 15

	if got != want {

		t.Errorf("got %d want %d given, %v", got, want, numbers)

	}

}

func TestSum3(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum3(numbers)
		want := 15

		if got != want {

			t.Errorf("got %d want %d given, %v", got, want, numbers)

		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum3(numbers)
		want := 6

		if got != want {

			t.Errorf("got %d want %d given, %v", got, want, numbers)

		}
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
	want := []int{15, 6}

	if !reflect.DeepEqual(got, want) {

		t.Errorf("got %v want %v ", got, want)

	}

}

func TestSumAll1(t *testing.T) {

	got := SumAll1([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
	want := []int{15, 6}

	if !reflect.DeepEqual(got, want) {

		t.Errorf("got %v want %v ", got, want)

	}

}

func TestSumAllTails(t *testing.T) {

	got := SumAllTails([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
	want := []int{14, 5}

	if !reflect.DeepEqual(got, want) {

		t.Errorf("got %v want %v ", got, want)

	}

}

func TestSumAllTails1(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
		want := []int{14, 5}

		if !reflect.DeepEqual(got, want) {

			t.Errorf("got %v want %v ", got, want)

		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails1([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		if !reflect.DeepEqual(got, want) {

			t.Errorf("got %v want %v ", got, want)

		}
	})

}

func TestSumAllTails2(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {

			t.Errorf("got %v want %v ", got, want)

		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
		want := []int{14, 5}
		checkSums(t, got, want)

	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails1([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		checkSums(t, got, want)
	})

}
