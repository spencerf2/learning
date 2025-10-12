package main

import (
	"errors"
	"fmt"
)

type Calculator struct {
	precision int
}

func NewCalculator(precision int) *Calculator {
	return &Calculator{precision: precision}
}

// TODO: Implement Add(a, b float64) float64
// TODO: Implement Divide(a, b float64) (float64, error) - return error if b is 0
// TODO: Implement Percentage(value, percent float64) float64
