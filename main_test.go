package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		size     int
		expected int
	}{
		{0, 0},
		{100000000, 100000000},
	}
	for _, test := range tests {
		result := generateRandomElements(test.size)
		require.Equal(t, test.expected, len(result))
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{42}, 42},
		{[]int{5, 5, 5, 5}, 5},
		{[]int{1, 100, 50, 75, 25}, 100},
		{[]int{1, 1000000, 5000000000, 33544545, 25555}, 5000000000},
	}
	for _, test := range tests {
		result := maximum(test.data)
		require.Equal(t, test.expected, result)
	}
}
