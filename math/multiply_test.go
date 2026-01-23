package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"math"
)

func TestMultiply(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "Positive numbers",
			a:        2,
			b:        3,
			expected: 6,
		},
		{
			name:     "Zero",
			a:        0,
			b:        5,
			expected: 0,
		},
		{
			name:     "Negative and positive",
			a:        -2,
			b:        3,
			expected: -6,
		},
		{
			name:     "Two negatives",
			a:        -2,
			b:        -3,
			expected: 6,
		},
		{
			name:     "Large numbers",
			a:        1000000,
			b:        1000000,
			expected: 1000000000000,
		},
		{
			name:     "Multiply by zero (negative)",
			a:        -5,
			b:        0,
			expected: 0,
		},
		{
			name:     "MaxInt64",
			a:        math.MaxInt64,
			b:        1,
			expected: math.MaxInt64,
		},
		{
			name:     "MinInt64",
			a:        math.MinInt64,
			b:        1,
			expected: math.MinInt64,
		},
		{
			name:     "Overflow",
			a:        math.MaxInt64,
			b:        2,
			expected: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Multiply(tt.a, tt.b))
		})
	}
}