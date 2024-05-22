package struct_interface

import (
	"testing"
)

func TestPerimeter (t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	// got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f want %.2f", got, want)
	}
}

// func TestArea (t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("Got %g want %g", got, want)
// 		}
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		// got := Area(rectangle)
// 		// // got := Area(12.0, 6.0)
// 		// want := 72.0

// 		// if got != want {
// 		// 	t.Errorf("Got %.2f want %.2f", got, want)
// 		// }
// 		checkArea(t, rectangle, 72.0)
// 	})

// 	t.Run("area of circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		// got := Area(circle)
// 		// want := 314.1592653589793

// 		// if got != want {
// 		// 	t.Errorf("Got %.2f want %.2f", got, want)
// 		// }
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

// Using table driven tests
func TestArea(t *testing.T){
	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
		// want float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the 't.Run' test name
		t.Run(tt.name, func(t *testing.T){
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}