package repository_test

import (
	"Go-Calc/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	calc       = repository.Calculator{}
	calculator = repository.NewCalculator(&calc)
)

func TestCalculator(t *testing.T) {
	t.Run("Test Add", func(t *testing.T) {
		result := calculator.Add(1, 2)
		assert.Equal(t, 3, result)
	})

	t.Run("Test Multiply", func(t *testing.T) {
		result := calculator.Divide(40, 10)
		assert.Equal(t, float64(4), result)
	})

	t.Run("Test Divide", func(t *testing.T) {
		result := calculator.Divide(10, 2)
		assert.Equal(t, float64(5), result)
	})

	t.Run("Test Less", func(t *testing.T) {
		result := calculator.Less(10, 2)
		assert.Equal(t, float64(8), result)
	})

}
