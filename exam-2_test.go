package main

import "testing"

func Test_decode(t *testing.T) {
	tests := []struct {
		name     string
		encoded  string
		expected string
	}{
		{
			name:     "Test case 1: LLRR=",
			encoded:  "LLRR=",
			expected: "210122",
		},
		{
			name:     "Test case 2: ==RLL",
			encoded:  "==RLL",
			expected: "000210",
		},
		{
			name:     "Test case 3: =LLRR",
			encoded:  "=LLRR",
			expected: "221012",
		},
		{
			name:     "Test case 4: RRL=R",
			encoded:  "RRL=R",
			expected: "012001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.encoded); got != tt.expected {
				t.Errorf("decode() = %v, want %v", got, tt.expected)
			}
		})
	}
}
