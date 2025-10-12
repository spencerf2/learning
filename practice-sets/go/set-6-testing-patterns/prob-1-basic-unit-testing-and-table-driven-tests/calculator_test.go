package main

import (
	"math"
	"testing"
)

// TODO: Write TestCalculator_Add that uses table-driven tests with these cases:
// {name: "positive numbers", a: 2.5, b: 3.7, want: 6.2}
// {name: "negative numbers", a: -1.5, b: -2.3, want: -3.8}
// {name: "zero values", a: 0, b: 5.5, want: 5.5}
// {name: "large numbers", a: 1000000.1, b: 2000000.2, want: 3000000.3}

// TODO: Write TestCalculator_Divide that tests:
// - Normal division
// - Division by zero (should return error)
// - Division with precision handling
// Use table-driven tests with struct: {name, a, b, want, wantErr}

// TODO: Write TestCalculator_Percentage that tests:
// - 50% of 100 = 50
// - 25% of 200 = 50
// - 100% of 75 = 75
// - 0% of anything = 0

func TestMain(m *testing.M) {
	// TODO: Add setup/teardown if needed
	// Run tests and exit with result
}

// Helper function for comparing floats
func floatEquals(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

// Answering the questions:
// 1. Why use table-driven tests instead of separate test functions?
//   -
//
// 2. How do you handle floating-point comparison in tests?
//   -
//
// 3. When should you use t.Error() vs t.Fatal()?
//   -
