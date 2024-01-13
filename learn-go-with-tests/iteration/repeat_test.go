package iteration

import "testing"

func TestRepeat(t *testing.T) {

	repeat := Repeat("a")
	expect := "aaaaa"

	if expect != repeat {
		t.Errorf("expected %q but got %q", expect, repeat)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
