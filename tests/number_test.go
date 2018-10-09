package tests

import (
	"math"
	"testing"
)

func TestNumber(t *testing.T) {
	t.Error(float64(88))

	n10 := math.Pow10(2)
	value := math.Trunc((88.999+0.5/n10)*n10) / n10

	t.Error(value, float64(100.00))

}
