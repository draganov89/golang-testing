package math_test

import (
	"testing"

	m "github.com/draganov89/golang-testing/math"
)

func TestAdd(t *testing.T) {

	// Long notation

	// type test struct {
	// 	a        int
	// 	b        int
	// 	expected int
	// }

	// tests := []test{
	// 	{1, 2, 3},
	// 	{0, 0, 0},
	// 	{-8, -88, -96},
	// 	{1, 2, 3},
	// 	{0, -5, -5},
	// 	{100, 100, 200},
	// }

	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-8, -88, -96},
		{-1, 1, 0},
		{0, -5, -5},
		{100, 100, 200},
	}

	for _, test := range tests {
		got := m.Add(test.a, test.b)
		if got != test.expected {
			t.Errorf("Expected %v, got %v\n", test.expected, got)
		}
	}
}

func TestSubtract(t *testing.T) {

	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, -1},
		{0, 0, 0},
		{-8, -88, 80},
		{0, -5, 5},
		{100, 100, 0},
	}

	for _, test := range tests {
		got := m.Subtract(test.a, test.b)
		if got != test.expected {
			t.Errorf("Expected %v, got %v\n", test.expected, got)
		}
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, -1},
		{0, 0, 0},
		{-8, -88, 1},
		{0, -5, 1},
		{99, 100, -1},
	}

	for _, test := range tests {
		got := m.Compare(test.a, test.b)
		if got != test.expected {
			t.Errorf("Expected %v, got %v\n", test.expected, got)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 1, 1, 1, 1}, 5},
		{[]int{0, 0, 0, 0, 0}, 0},
		{[]int{}, 0},
	}

	for _, test := range tests {
		got := m.Sum(test.nums...)
		if got != test.expected {
			t.Errorf("Expected %v, got %v\n", test.expected, got)
		}
	}
}

func TestSumEven(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{1, 1, 1, 1, 1}, 0},
		{[]int{0, 0, 0, 0, 0}, 0},
		{[]int{}, 0},
		{[]int{2, 4, 6, 8}, 20},
	}

	for _, test := range tests {
		got := m.SumEven(test.nums...)
		if got != test.expected {
			t.Errorf("Expected %v, got %v\n", test.expected, got)
		}
	}
}
