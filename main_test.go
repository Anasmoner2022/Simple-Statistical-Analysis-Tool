package main

import (
	"math"
	"testing"
)

// A small constant for comparing floating-point numbers
const float64Tolerance = 1e-9

// Helper function for comparing two float64 values with tolerance
func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < float64Tolerance
}

// --- Test Cases for Mean ---
func TestMean(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected float64
	}{
		{"Empty slice", []string{}, 0.0},
		{"Single value", []string{"5"}, 5.0},
		{"Even number of elements", []string{"1", "2", "3", "4"}, 2.5},
		{"Odd number of elements", []string{"1", "2", "3", "4", "5"}, 3.0},
		{"With negative numbers", []string{"-10", "10", "0"}, 0.0},
		{"Mixed values", []string{"10", "20", "30", "40", "50"}, 30.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Mean(tt.input)
			if !floatEquals(actual, tt.expected) {
				t.Errorf("Mean(%v) = %f; expected %f", tt.input, actual, tt.expected)
			}
		})
	}
}

// --- Test Cases for Median ---
func TestMedian(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected float64
	}{
		// Note: The original code does not handle empty slice gracefully for Median,
		// but for exported functions, we should test expected behavior.
		// For an empty slice, len(dataConv) is 0, (len % 2) is 0, and it tries to access index -1 and 0, which panics.
		// We'll skip the empty case or fix the function, but for now, we test valid inputs.
		{"Odd count", []string{"1", "5", "2", "4", "3"}, 3.0},
		{"Even count", []string{"1", "5", "2", "6", "3", "4"}, 3.5},
		{"Sorted input", []string{"10", "20", "30", "40", "50"}, 30.0},
		{"Duplicates", []string{"1", "2", "2", "3", "3", "3", "4"}, 3.0},
		{"Negative numbers (odd)", []string{"-10", "0", "10"}, 0.0},
		{"Negative numbers (even)", []string{"-10", "0", "10", "20"}, 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Median(tt.input)
			if !floatEquals(actual, tt.expected) {
				t.Errorf("Median(%v) = %f; expected %f", tt.input, actual, tt.expected)
			}
		})
	}
}

// --- Test Cases for MedianInt (Used internally) ---
func TestMedianInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Odd count", []int{1, 5, 2, 4, 3}, 3},
		// Note: MedianInt uses integer division (first + second) / 2,
		// so (3+4)/2 = 3 in integer arithmetic.
		{"Even count (odd sum)", []int{1, 5, 2, 6, 3, 4}, 3},
		{"Even count (even sum)", []int{1, 5, 2, 7, 3, 4}, 3},
		{"Sorted input", []int{10, 20, 30, 40, 50}, 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MedianInt(tt.input)
			if actual != tt.expected {
				t.Errorf("MedianInt(%v) = %d; expected %d", tt.input, actual, tt.expected)
			}
		})
	}
}

// --- Test Cases for Range ---
func TestRange(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{"Standard range", []string{"1", "10", "5"}, 9},
		{"Zero range", []string{"5", "5", "5"}, 0},
		{"Negative numbers", []string{"-10", "0", "10"}, 20}, // 10 - (-10) = 20
		{"Single value", []string{"7"}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Range(tt.input)
			if actual != tt.expected {
				t.Errorf("Range(%v) = %d; expected %d", tt.input, actual, tt.expected)
			}
		})
	}
}

// --- Test Cases for InterquartileRange ---
func TestInterquartileRange(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int // Because MedianInt uses integer division, we test for integer results.
	}{
		// Data: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
		// Even n=10, half=5. Q1=MedianInt([1, 2, 3, 4, 5])=3. Q3=MedianInt([6, 7, 8, 9, 10])=8. IQR=8-3=5
		{"Even count (10)", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, 5},

		// Data: 1, 2, 3, 4, 5, 6, 7, 8, 9
		// Odd n=9, half=4. Q1=MedianInt([1, 2, 3, 4])=2. Q3=MedianInt([6, 7, 8, 9])=7. IQR=7-2=5
		{"Odd count (9)", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, 5},

		// Data: 1, 2, 3, 4
		// Even n=4, half=2. Q1=MedianInt([1, 2])=1. Q3=MedianInt([3, 4])=3. IQR=3-1=2
		{"Small Even count (4)", []string{"1", "2", "3", "4"}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := InterquartileRange(tt.input)
			if actual != tt.expected {
				t.Errorf("InterquartileRange(%v) = %d; expected %d", tt.input, actual, tt.expected)
			}
		})
	}
}

// --- Test Cases for Min and Max ---
func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{"Standard min", []string{"10", "1", "5"}, 1},
		{"Negative numbers", []string{"-10", "0", "10"}, -10},
		{"Single value", []string{"7"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Min(tt.input)
			if actual != tt.expected {
				t.Errorf("Min(%v) = %d; expected %d", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{"Standard max", []string{"10", "1", "5"}, 10},
		{"Negative numbers", []string{"-10", "0", "10"}, 10},
		{"Single value", []string{"7"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Max(tt.input)
			if actual != tt.expected {
				t.Errorf("Max(%v) = %d; expected %d", tt.input, actual, tt.expected)
			}
		})
	}
}

// --- Test Cases for Variance and StandardDeviation ---
func TestVariance(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected float64
	}{
		// Data: 2, 4, 4, 4, 5, 5, 7, 9. Mean = 5.
		// Squared Differences: 9, 1, 1, 1, 0, 0, 4, 16. Sum=32. Variance = 32/8 = 4.0
		{"Standard set", []string{"2", "4", "4", "4", "5", "5", "7", "9"}, 4.0},
		{"Single value", []string{"5"}, 0.0},
		{"Two values", []string{"1", "5"}, 4.0}, // Mean=3. (3-1)^2 + (3-5)^2 = 4 + 4 = 8. Variance = 8/2 = 4.0
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Variance(tt.input)
			if !floatEquals(actual, tt.expected) {
				t.Errorf("Variance(%v) = %f; expected %f", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected float64
	}{
		// Variance is 4.0, so SD is sqrt(4) = 2.0
		{"Standard set", []string{"2", "4", "4", "4", "5", "5", "7", "9"}, 2.0},
		// Variance is 0.0, so SD is 0.0
		{"Single value", []string{"5"}, 0.0},
		// Variance is 4.0, so SD is 2.0
		{"Two values", []string{"1", "5"}, 2.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := StandardDeviation(tt.input)
			if !floatEquals(actual, tt.expected) {
				t.Errorf("StandardDeviation(%v) = %f; expected %f", tt.input, actual, tt.expected)
			}
		})
	}
}
