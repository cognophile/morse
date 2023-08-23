package encoder_test

import (
	"testing"

	"cognophile.com/morse/encoder"
)

func TestEncode(t *testing.T) {
	want := ".../---/.../"
	got := encoder.Encode("SOS")

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
