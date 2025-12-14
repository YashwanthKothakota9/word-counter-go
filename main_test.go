package main_test

import (
	"testing"

	counter "github.com/yashwanthkothakota9/counter"
)

func TestCountWords(t *testing.T) {

	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "Empty input",
			input: "",
			wants: 0,
		},
		{
			name:  "Space input",
			input: " ",
			wants: 0,
		},
		{
			name:  "New line",
			input: "one two three\nfour five",
			wants: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := counter.CountWords([]byte(tc.input))
			if result != tc.wants {
				t.Log("expected:", tc.wants, "got:", result)
				t.Fail()
			}
		})
	}
}
