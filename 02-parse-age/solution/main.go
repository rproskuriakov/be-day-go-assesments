package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrNegativeAge = errors.New("age is negative")

func ParseAge(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("age %q: %w", s, err)
	}
	if n < 0 {
		return 0, ErrNegativeAge
	}
	return n, nil
}

func main() {
	for _, in := range []string{"42", "abc", "-5"} {
		age, err := ParseAge(in)
		if errors.Is(err, ErrNegativeAge) {
			fmt.Println("please enter a positive age")
		} else if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Println("age:", age)
		}
	}
}
