package nothing

import "testing"

func TestFormatLayoutEx(t *testing.T) {
	got := FormatLayoutEx(true)
	want := "20060102"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := FormatLayoutEx(false)
	want2 := "20060102150405"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}
}

func TestParseLayoutEx(t *testing.T) {
	got := ParseLayoutEx(true)
	want := "20060102"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := ParseLayoutEx(false)
	want2 := "20060102150405"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}
}
