package main

import "errors"

var ErrNegativeAge = errors.New("age is negative")

func ParseAge(s string) (int, error) {
	// TODO: convert to number, check on non-negative
	return 0, nil
}

func main() {
	for _, in := range []string{"42", "abc", "-5"} {
		age, err := ParseAge(in)
		// TODO: handle err
		_, _ = age, err
	}
}
