package chibi_linku

import (
	"testing"
)

func TestBase62Encode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "hello", expected: "h3Avpr9"},
		{"world", "jb5RGRI"},
		{"test", "cijV84"},
	}

	for _, test := range tests {
		result := Base62Encode(test.input)

		if result != test.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", test.input, test.expected, result)
		}
	}
}

func TestBase62Decode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "h3Avpr9", expected: "hello"},
		{"jb5RGRI", "world"},
		{"cijV84", "test"},
	}

	for _, test := range tests {
		result := Base62Decode(test.input)

		if result != test.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", test.input, test.expected, result)
		}
	}
}
