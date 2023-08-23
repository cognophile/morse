package decoder_test

import (
	"testing"

	"cognophile.com/morse/decoder"
)

func TestDecode(t *testing.T) {
	want := "SOS"
	got := decoder.Decode(".../---/...//")

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
