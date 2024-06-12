package structsandinterface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("want: %.2f, but got: %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct{
		shape Shape
		want float64
	}{
		{Rectangle{Width: 10, Height: 10}, 100},
		{Circle{Radius: 10}, 314.1592653589793},
		{Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("want: %g, but got: %g", tt.want, got)
		}
	}
}
