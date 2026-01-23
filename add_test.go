package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "Positive numbers",
			a:        2,
			b:        3,
			expected: 5,
		},
		{
			name:     "Zero values",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "Negative and positive",
			a:        -5,
			b:        10,
			expected: 5,
		},
		{
			name:     "Negative numbers",
			a:        -3,
			b:        -4,
			expected: -7,
		},
		{
			name:     "Large numbers",
			a:        1000000,
			b:        2000000,
			expected: 3000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := Add(tt.a, tt.b)
			assert.Equal(t, tt.expected, result, "Add(%d, %d) should equal %d", tt.a, tt.b, tt.expected)
		})
	}
}