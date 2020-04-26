package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.slow.com"
	fastURL := "http://www.fast.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
