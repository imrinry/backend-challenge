package main

import "testing"

func Test_bigger(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "a มากกว่า b",
			a:        53,
			b:        6,
			expected: 53,
		},
		{
			name:     "b มากกว่า a",
			a:        40,
			b:        53,
			expected: 53,
		},
		{
			name:     "a เท่ากับ b",
			a:        50,
			b:        50,
			expected: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bigger(tt.a, tt.b); got != tt.expected {
				t.Errorf("bigger() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func Test_findSum(t *testing.T) {
	tests := []struct {
		name     string
		pyramid  [][]int
		expected int
	}{
		{
			name: "Example 1",
			pyramid: [][]int{
				{59},
				{73, 41},
				{52, 40, 53},
				{26, 53, 6, 34},
			},
			expected: 237,
		},
		{
			name: "Example 2",
			pyramid: [][]int{
				{5},
				{2, 3},
				{1, 4, 5},
			},
			expected: 13, // 5 -> 3 -> 5 = 13
		},
		{
			name: "Example 3",
			pyramid: [][]int{
				{10},
			},
			expected: 10,
		},
		{
			name: "Example 4",
			pyramid: [][]int{
				{1},
				{2, 3},
			},
			expected: 4, // 1 -> 3 = 4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSum(tt.pyramid); got != tt.expected {
				t.Errorf("findSum() = %v, want %v", got, tt.expected)
			}
		})
	}
}
