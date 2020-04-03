package chinese

import "testing"

func TestFormat(t *testing.T) {
	got := Format(1136185445, false, false, false)
	want := "2006年01月02日 15时04分05秒"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
