package main

import "testing"

func TestcheckSentenceStart(t *testing.T) {

	got := checkSentenceStart("Test.")
	want := true

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
