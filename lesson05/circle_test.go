package main

import (
	"errors"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const tolerance float64 = 0.00001

func withTolerane(a, b, tolerance float64) bool {
	if diff := math.Abs(a - b); diff < tolerance {
		//	fmt.Printf("When a=%f and b =%f => Nearly same by tolerance\n", a, b)
		return true
	} else {
		//	fmt.Printf("When a=%f and b=%f => Not same Even by Tolerance\n", a, b)
		return false
	}
}
func TestAreaC(t *testing.T) {

	tests := []struct {
		name    string
		args    Circle
		want    float64
		wantErr error
	}{
		{
			name:    "if radius is <= 0 should return error ",
			args:    Circle{Radius: -1.0},
			want:    0.0,
			wantErr: errors.New("Error is happen! Radius could not be <= 0."),
		},
		{
			name:    "if radius is <= 0 should return error ",
			args:    Circle{Radius: 0.0},
			want:    0.0,
			wantErr: errors.New("Error is happen! Radius could not be <= 0."),
		},
		{
			name:    "if radius = 1 should return Pi",
			args:    Circle{Radius: 1.0},
			want:    math.Pi,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (tt.args).Area()

			if err != nil {
				assert.Equal(t, err, tt.wantErr)

			}
			if !withTolerane(got, tt.want, tolerance) {
				t.Errorf("Area() = %v, want = %v", got, tt.want)
				assert.Equal(t, got, tt.want)
			}
		})
	}
}
func TestPerimeterC(t *testing.T) {

	tests := []struct {
		name    string
		args    Circle
		want    float64
		wantErr error
	}{
		{
			name:    "if radius is <= 0 should return error ",
			args:    Circle{Radius: -1.0},
			want:    0.0,
			wantErr: errors.New("Error is happen! Radius could not be <= 0."),
		},
		{
			name:    "if radius is <= 0 should return error ",
			args:    Circle{Radius: 0.0},
			want:    0.0,
			wantErr: errors.New("Error is happen! Radius could not be <= 0."),
		},
		{
			name:    "if radius = 1 should return Pi",
			args:    Circle{Radius: 1.0},
			want:    2 * math.Pi,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (tt.args).Perimeter()

			if err != nil {
				assert.Equal(t, err, tt.wantErr)

			}
			if !withTolerane(got, tt.want, tolerance) {
				t.Errorf("Area() = %v, want = %v", got, tt.want)
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestCircle_String(t *testing.T) {
	tests := []struct {
		name string
		args Circle
		want string
	}{
		{
			name: "if radius is <= 0 should return error ",
			args: Circle{Radius: -1.0},
			want: "\nError! Circle with Radius = -1 doesn't exist.",
		},
		{
			name: "if radius is <= 0 should return error ",
			args: Circle{Radius: 0.0},
			want: "\nError! Circle with Radius = 0 doesn't exist.",
		},
		{
			name: "if radius = 1 should return Pi",
			args: Circle{Radius: 1.0},
			want: "\nCircle: radius 1.00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := (tt.args).String()

			if got != tt.want {
				assert.Equal(t, got, tt.want)

			}
		})
	}
}
