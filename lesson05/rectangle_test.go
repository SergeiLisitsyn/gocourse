package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {

	tests := []struct {
		name    string
		args    Rectangle
		want    float64
		wantErr error
	}{
		{
			name: "if width, or height is <= 0 should return error ",
			args: Rectangle{
				Height: -1.0,
				Width:  0.0,
			},
			want:    0.0,
			wantErr: errors.New("Error is happen! Height and width are neither could be <= 0."),
		},
		{
			name: "if width, or height is <= 0 should return error ",
			args: Rectangle{
				Height: 0.0,
				Width:  -1.0,
			},
			want:    0.0,
			wantErr: errors.New("Error is happen! Height and width are neither could be <= 0."),
		},
		{
			name: "if width, or height is <= 0 should return error ",
			args: Rectangle{
				Height: 0.0,
				Width:  0.0,
			},
			want:    0.0,
			wantErr: errors.New("Error is happen! Height and width are neither could be <= 0."),
		},
		{
			name: "if width and height are = 5 should return 25",
			args: Rectangle{
				Height: 5.0,
				Width:  5.0,
			},
			want:    25.0,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (tt.args).Area()
			assert.Equal(t, err, tt.wantErr)

			if !withTolerane(got, tt.want, tolerance) {
				t.Errorf("Area() = %v, want = %v", got, tt.want)
				assert.Equal(t, got, tt.want)
			}
		})
	}
}
func TestPerimeter(t *testing.T) {

	tests := []struct {
		name    string
		args    Rectangle
		want    float64
		wantErr error
	}{
		{
			name: "if width, or height is <= 0 should return error ",
			args: Rectangle{
				Height: -1.0,
				Width:  0.0,
			},
			want:    0.0,
			wantErr: errors.New("Error is happen! Height and width are neither could be <= 0."),
		},
		{
			name: "if width, or height is <= 0 should return error ",
			args: Rectangle{
				Height: 0.0,
				Width:  -1.0,
			},
			want:    0.0,
			wantErr: errors.New("Error is happen! Height and width are neither could be <= 0."),
		},
		{
			name: "if width, or height is <= 0 should return error ",
			args: Rectangle{
				Height: 0.0,
				Width:  0.0,
			},
			want:    0.0,
			wantErr: errors.New("Error is happen! Height and width are neither could be <= 0."),
		},
		{
			name: "if width and height are = 5 should return 20",
			args: Rectangle{
				Height: 5.0,
				Width:  5.0,
			},
			want:    20.0,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := (tt.args).Perimeter()
			assert.Equal(t, err, tt.wantErr)

			if !withTolerane(got, tt.want, tolerance) {
				t.Errorf("Area() = %v, want = %v", got, tt.want)
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRectangle_String(t *testing.T) {
	tests := []struct {
		name string
		args Rectangle
		want string
	}{
		{
			name: "if width, or height is <= 0 should return error ",
			args: Rectangle{
				Height: 0.0,
				Width:  -1.0,
			},
			want: "\nError.",
		},
		{
			name: "if width, or height is <= 0 should return error ",
			args: Rectangle{
				Height: 0.0,
				Width:  0.0,
			},
			want: "\nError.",
		},
		{
			name: "if width and height are = 5 should return 20",
			args: Rectangle{
				Height: 5.0,
				Width:  5.0,
			},
			want: "\nRectangle with height 5.00 and width 5.00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := (tt.args).String()
			assert.Equal(t, got, tt.want)
		})
	}
}
