package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        2,
			b:        3,
			expected: 6,
		},
		{
			name:     "negative numbers",
			a:        -2,
			b:        -3,
			expected: 6,
		},
		{
			name:     "mixed positive and negative",
			a:        -2,
			b:        3,
			expected: -6,
		},
		{
			name:     "zero first operand",
			a:        0,
			b:        5,
			expected: 0,
		},
		{
			name:     "zero second operand",
			a:        5,
			b:        0,
			expected: 0,
		},
		{
			name:     "both zero",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "large numbers",
			a:        1000000,
			b:        1000000,
			expected: 1000000000000,
		},
		{
			name:     "identity",
			a:        5,
			b:        1,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}