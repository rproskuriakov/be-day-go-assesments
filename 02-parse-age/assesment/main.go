package main

import "errors"

var ErrNegativeAge = errors.New("age is negative")

func ParseAge(s string) (int, error) {
	// TODO: convert to number, check on non-negative
}

func main() {
	for _, in := range []string{"42", "abc", "-5"} {
		age, err := ParseAge(in)
		// TODO: handle err
	}
}
