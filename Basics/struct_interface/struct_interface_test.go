package struct_interface

import "testing"

func TestPerimeter (t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	// got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f want %.2f", got, want)
	}
}

func TestArea (t *testing.T) {

	t.Run("Test Area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := Area(rectangle)
		// got := Area(12.0, 6.0)
		want := 72.0

		if got != want {
			t.Errorf("Got %.2f want %.2f", got, want)
		}
	})

	t.Run("area of circle", func(t *testing.T) {
		circle := Circle{10}
		got := Area(circle)
		want := 314.1592653589793

		if got != want {
			t.Errorf("Got %.2f want %.2f", got, want)
		}
	})
}