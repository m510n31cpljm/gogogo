package main

import (
	"errors"
	"fmt"
)

func main() {
	baseErr := errors.New("C")

	c := func() error {
		errC := baseErr

		return errC
	}

	b := func() error {
		errC := c()

		if errC != nil {
			return fmt.Errorf("B failed cause: %w", errC)
		}

		return nil
	}

	a := func() error {
		errB := b()

		if errB != nil {
			return fmt.Errorf("A failed cause: %w", errB)
		}

		return nil
	}

	errA := a()

	if errA != nil {
		fmt.Println(errA)
	}
}
