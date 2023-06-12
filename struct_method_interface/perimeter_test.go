package structmethodinterface

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expected := 40.0

	if expected != result {
		t.Errorf("result: '%.2f', expected: '%.2f'", result, expected)
	}
}

func TestArea(t *testing.T) {
	testsArea := []struct {
		name string
		form Form
		expected float64
	} {
		{name: "Rectangle", form: Rectangle{Width: 12, Height: 6}, expected: 72},
		{name: "Circle", form: Circle{Raio: 10}, expected: 314.1592653589793},
		{name: "Triangle", form: Triangle{Base: 12, Height: 65}, expected: 36},
	}

	for _, tt := range testsArea {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.form.Area()
	
			if result != tt.expected {
				t.Errorf("%#v result: '%.2f', expected: '%.2f'", tt.form, result, tt.expected)
			}
		})
	}
}