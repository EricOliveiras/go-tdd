package arrayandslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("should sum numeric collection", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := Sum(numbers)
		expected := 6
	
		if result	!= expected {
			t.Errorf("expected: '%d', result: '%d', data: '%v'", expected, result, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: '%v', result: '%v'", expected, result)
	}
}

func TestSumAllRest(t *testing.T) {
	checkSums := func (t *testing.T, result, expected []int)  {
		t.Helper()
		
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected: '%v', result: '%v'", expected, result)
		}
	}

	t.Run("shoul sum some slices", func(t *testing.T) {
		result := SumAllRest([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
	
		checkSums(t, result, expected)
	})

	t.Run("should sum the empty slice safely", func(t *testing.T) {
		result := SumAllRest([]int{}, []int{0, 9})
		expected := []int{0, 9}
	
		checkSums(t, result, expected)
	})
}