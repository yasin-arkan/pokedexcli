package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Howisyourday!",
			expected: []string{"howisyourday!"},
		},
		{
			input:    "Hello,   how  are you!",
			expected: []string{"hello,", "how", "are", "you!"},
		},
		{
			input:    "Im       feeling fine, thankyou",
			expected: []string{"im", "feeling", "fine,", "thankyou"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length does not match")
			fmt.Println(actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Words do not match")
				fmt.Println(word, expectedWord)
			}
		}
	}
}
