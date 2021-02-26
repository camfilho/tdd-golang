package loop

import "testing"

func TestRepeat(t *testing.T)  {
	expected := "aaaaa"

	if result := Repeat("a"); result != expected {
		(*t).Errorf("expected %q but got %q", expected, result)
	}
}
