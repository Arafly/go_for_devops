package main

import (
	"fmt"
	"testing"
)

func Greet(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name was empty")
	}
	return "Hello " + name, nil
}

func TestGreet(t *testing.T) {
	tests := []struct {
		desc      string // What we are testing
		name      string // The name we will pass
		want      string // What we expect to be returned
		expectErr bool   // Do we expect an error
	}{
		{
			desc:      "Error: name is an empty string",
			expectErr: true,
			// name and want are "", the zero value for string
		},
		{
			desc: "Success",
			name: "John",
			want: "Hello John",
			// expectErr is set to the zero value, false
		},
	}

	// Executes each test.
	for _, test := range tests {
		got, err := Greet(test.name)
		switch {
		// We did not get an error, but expected one
		case err == nil && test.expectErr:
			t.Errorf("TestGreet(%s): got err == nil, want err != nil", test.desc)
			continue
		// We got an error but did not expect one
		case err != nil && !test.expectErr:
			t.Errorf("TestGreet(%s): got err == %s, want err == nil", test.desc, err)
			continue
		// We got an error we expected, so just go to the next test
		case err != nil:
			continue
		}

		// We did not get the result we expected
		if got != test.want {
			t.Errorf("TestGreet(%s): got result %q, want %q", test.desc, got, test.want)
		}
	}
}