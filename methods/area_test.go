package area

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 20.0)
	want := 60.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
