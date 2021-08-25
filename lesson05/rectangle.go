package main

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() (float64, error) {
	if r.Width <= 0 || r.Height <= 0 {
		return 0, errors.New("Error is happen! Height and width are neither could be <= 0.")
	}
	return r.Width * r.Height, nil
}

func (r Rectangle) Perimeter() (float64, error) {
	if r.Width <= 0 || r.Height <= 0 {
		return 0, errors.New("Error is happen! Height and width are neither could be <= 0.")
	}
	return 2 * (r.Width + r.Height), nil
}

func (r Rectangle) String() string {
	if r.Height <= 0 || r.Width <= 0 {
		return "\nError."
	} else {
		return fmt.Sprintf("\nRectangle with height %.2f and width %.2f", r.Height, r.Width)
	}
}
