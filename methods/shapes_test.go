package area

import "testing"

// func TestPerimeter(t *testing.T) {
// 	rectangle := Rectangle{10.0, 20.0}
// 	got := Perimeter(rectangle)
// 	want := 60.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

func TestArea(t *testing.T) {

	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},        // this is `tt``
		{Circle{10}, 314.1592653589793}, // this is also `tt`
		{Triangle{10, 5}, 25.0},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}
