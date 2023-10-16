package main

import "testing"

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %f want %f", shape, got, want)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 4.0}
		checkArea(t, rectangle, 40.0)
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
	t.Run("table based test", func(t *testing.T) {
		areaTests := []struct {
			name    string
			shape   Shape
			hasArea float64
		}{
			{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
			{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
			{name: "Triangle", shape: Triangle{Base: 10, Height: 6}, hasArea: 30.0},
		}
		for _, tt := range areaTests {
			checkArea(t, tt.shape, tt.hasArea)
		}
	})
}
