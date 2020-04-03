package chinese

import "testing"

func TestFormatLayoutEx(t *testing.T) {
	got := FormatLayoutEx(true, true, true)
	want := "06年01月02日"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := FormatLayoutEx(true, true, false)
	want2 := "06年01月02日"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := FormatLayoutEx(true, false, true)
	want3 := "2006年01月02日"
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	got4 := FormatLayoutEx(true, false, false)
	want4 := "2006年01月02日"
	if got4 != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	got5 := FormatLayoutEx(false, true, true)
	want5 := "06年01月02日 15时04分"
	if got5 != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}

	got6 := FormatLayoutEx(false, true, false)
	want6 := "06年01月02日 15时04分05秒"
	if got6 != want6 {
		t.Errorf("got %q; want %q", got6, want6)
	}

	got7 := FormatLayoutEx(false, false, true)
	want7 := "2006年01月02日 15时04分"
	if got7 != want7 {
		t.Errorf("got %q; want %q", got7, want7)
	}

	got8 := FormatLayoutEx(false, false, false)
	want8 := "2006年01月02日 15时04分05秒"
	if got8 != want8 {
		t.Errorf("got %q; want %q", got8, want8)
	}
}

func TestParseLayoutEx(t *testing.T) {
	got := ParseLayoutEx(true, true, true)
	want := "06年1月2日"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := ParseLayoutEx(true, true, false)
	want2 := "06年1月2日"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := ParseLayoutEx(true, false, true)
	want3 := "2006年1月2日"
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	got4 := ParseLayoutEx(true, false, false)
	want4 := "2006年1月2日"
	if got4 != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	got5 := ParseLayoutEx(false, true, true)
	want5 := "06年1月2日 15时4分"
	if got5 != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}

	got6 := ParseLayoutEx(false, true, false)
	want6 := "06年1月2日 15时4分5秒"
	if got6 != want6 {
		t.Errorf("got %q; want %q", got6, want6)
	}

	got7 := ParseLayoutEx(false, false, true)
	want7 := "2006年1月2日 15时4分"
	if got7 != want7 {
		t.Errorf("got %q; want %q", got7, want7)
	}

	got8 := ParseLayoutEx(false, false, false)
	want8 := "2006年1月2日 15时4分5秒"
	if got8 != want8 {
		t.Errorf("got %q; want %q", got8, want8)
	}
}
