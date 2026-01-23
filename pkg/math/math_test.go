package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        5,
			b:        3,
			expected: 8,
		},
		{
			name:     "negative numbers",
			a:        -5,
			b:        -3,
			expected: -8,
		},
		{
			name:     "mixed positive and negative",
			a:        10,
			b:        -3,
			expected: 7,
		},
		{
			name:     "zero with positive",
			a:        0,
			b:        5,
			expected: 5,
		},
		{
			name:     "zero with negative",
			a:        -5,
			b:        0,
			expected: -5,
		},
		{
			name:     "both zeros",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "large numbers",
			a:        1000000,
			b:        2000000,
			expected: 3000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			assert.Equal(t, tt.expected, result, "Add(%d, %d) should equal %d", tt.a, tt.b, tt.expected)
		})
	}
}

func TestAdd_Commutative(t *testing.T) {
	// Verify that Add is commutative: a + b == b + a
	a, b := 7, 13
	result1 := Add(a, b)
	result2 := Add(b, a)
	require.Equal(t, result1, result2, "Add should be commutative")
}