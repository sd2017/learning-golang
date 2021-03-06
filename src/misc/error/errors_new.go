package main

import (
	"errors"
	"fmt"
)

var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}

	return x / y, nil
}

func main() {
	switch z, err := div(10, 0); err {
	case nil:
		fmt.Println(z)
	case ErrDivByZero:
		panic(err)
	}
}
