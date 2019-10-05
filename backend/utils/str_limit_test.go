package utils

import "testing"

func TestStrLimit(t *testing.T) {
	cases := []struct {
		InputString string
		Limit       int
		Result      string
	}{
		{"Hello man", 6, "Hello ..."},
		{"Long string will be short after this func", 1, "L..."},
		{"short", 20, "short"},
		{"sh", 2, "sh"},
		{"sh", 3, "sh"},
		{"th", 4, "th"},
	}

	for _, c := range cases {
		result := StrLimit(c.InputString, c.Limit)

		if result != c.Result {
			t.Errorf("Result must be %s but got %s", c.Result, result)
		}
	}
}
