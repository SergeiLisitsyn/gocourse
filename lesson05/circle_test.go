package main

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreaC(t *testing.T) {

	tests := []struct {
		name    string
		args    Circle
		want    float64
		wantErr bool
	}{
		{
			name: "if radius is <= 0 should return error ",
			args: Circle{
				Radius: 0.0,
			},
			want:    0.0,
			wantErr: true,
		},
		{
			name: "if radius = 1 should return Pi",
			args: Circle{
				Radius: 1.0,
			},
			want:    math.Pi,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (tt.args).Area()

			if (err != nil) != tt.wantErr {
				t.Errorf("Area() error = %v, wantErr %v, got = %v", err, tt.wantErr, got)
				assert.Equal(t, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
		wantErr bool
	}{
		{
			name: "if radius is <= 0 should return error ",
			args: Circle{
				Radius: 0.0,
			},
			want:    0.0,
			wantErr: true,
		},
		{
			name: "if radius = 1 should return 2 * Pi",
			args: Circle{
				Radius: 1.0,
			},
			want:    math.Pi * 2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (tt.args).Perimeter()

			if (err != nil) != tt.wantErr {
				t.Errorf("Perimeter() error = %v, wantErr %v, got = %v", err, tt.wantErr, got)
				assert.Equal(t, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Perimeter() = %v, want = %v", got, tt.want)
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestCircle_String(t *testing.T) {
	c0 := Circle{Radius: 0}
	c5 := Circle{Radius: 5}

	negativeResult := c0.String()
	positiveResult := c5.String()
	// test for radius<= 0
	if negativeResult != "\nError! Circle with Radius = 0 doesn't exist." {
		t.Errorf("String() failed, expected %v, got %v", "Error! Circle with Radius = 0 doesn't exist.", negativeResult)
	}
	// test for radius = 5
	if positiveResult != "\nCircle: radius 5.00" {
		t.Errorf("String() failed, expected %v, got %v", "\nCircle: radius 5.00", positiveResult)
	}
}
