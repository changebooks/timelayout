package slash

import "testing"

func TestFormatLayoutEx(t *testing.T) {
	got := FormatLayoutEx(true, true, true)
	want := "06/01/02"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := FormatLayoutEx(true, true, false)
	want2 := "06/01/02"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := FormatLayoutEx(true, false, true)
	want3 := "2006/01/02"
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	got4 := FormatLayoutEx(true, false, false)
	want4 := "2006/01/02"
	if got4 != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	got5 := FormatLayoutEx(false, true, true)
	want5 := "06/01/02 15:04"
	if got5 != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}

	got6 := FormatLayoutEx(false, true, false)
	want6 := "06/01/02 15:04:05"
	if got6 != want6 {
		t.Errorf("got %q; want %q", got6, want6)
	}

	got7 := FormatLayoutEx(false, false, true)
	want7 := "2006/01/02 15:04"
	if got7 != want7 {
		t.Errorf("got %q; want %q", got7, want7)
	}

	got8 := FormatLayoutEx(false, false, false)
	want8 := "2006/01/02 15:04:05"
	if got8 != want8 {
		t.Errorf("got %q; want %q", got8, want8)
	}
}

func TestParseLayoutEx(t *testing.T) {
	got := ParseLayoutEx(true, true, true)
	want := "06/1/2"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := ParseLayoutEx(true, true, false)
	want2 := "06/1/2"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := ParseLayoutEx(true, false, true)
	want3 := "2006/1/2"
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	got4 := ParseLayoutEx(true, false, false)
	want4 := "2006/1/2"
	if got4 != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	got5 := ParseLayoutEx(false, true, true)
	want5 := "06/1/2 15:4"
	if got5 != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}

	got6 := ParseLayoutEx(false, true, false)
	want6 := "06/1/2 15:4:5"
	if got6 != want6 {
		t.Errorf("got %q; want %q", got6, want6)
	}

	got7 := ParseLayoutEx(false, false, true)
	want7 := "2006/1/2 15:4"
	if got7 != want7 {
		t.Errorf("got %q; want %q", got7, want7)
	}

	got8 := ParseLayoutEx(false, false, false)
	want8 := "2006/1/2 15:4:5"
	if got8 != want8 {
		t.Errorf("got %q; want %q", got8, want8)
	}
}
