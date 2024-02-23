package main

import (
	"testing"
)

func TestProccessText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test 1",
			input:    "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			expected: "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			name:     "Test 2",
			input:    "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			expected: "I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			name:     "Test 3",
			input:    "Don not be sad ,because sad backwards is das . And das not good",
			expected: "Don not be sad, because sad backwards is das. And das not good",
		},
		{
			name:     "Test 4",
			input:    "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			expected: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ProccessText(test.input)
			if actual != test.expected {
				t.Errorf("expected:\n%s\n\nactual:\n%s\n", test.expected, actual)
			}
		})
	}
}
