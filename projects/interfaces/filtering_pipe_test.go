package main

import (
	"bytes"
	"testing"
)

// TestFilteringPipe tests the FilteringPipe
func TestFilteringPipe(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "works for empty string", input: "", expected: ""},
		{name: "works for string with no numbers", input: "hello world", expected: "hello world"},
		{name: "works for string with 1 numbers", input: "hello 1 world", expected: "hello  world"},
		{name: "works for string with numbers", input: "hello 123 world", expected: "hello  world"},
		{name: "works for string with numbers and special characters", input: "hello 123 world !@#$%^&*()_+", expected: "hello  world !@#$%^&*()_+"},
		{name: "works for string with only numbers", input: "1234567890", expected: ""},
	}

	//	Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bytes.Buffer{}
			fp := NewFilteringPipe(&b)

			if _, err := fp.Write([]byte(tt.input)); err != nil {
				t.Errorf("got an error but did not want one")
			}

			got := b.String()
			if got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}
