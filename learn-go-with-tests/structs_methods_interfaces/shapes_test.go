package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if want != got {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {

	t.Run("Rectangles", func(t *testing.T) {

		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if want != got {
			t.Errorf("got %.2f want %.2f", got, want)
		}

	})

	t.Run("Circles", func(t *testing.T) {

		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}

	})

}

func TestAreaShape(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Rectangles", func(t *testing.T) {

		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)

	})

	t.Run("Circles", func(t *testing.T) {

		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)

	})

}
