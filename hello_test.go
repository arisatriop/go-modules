package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Return value should be: '%s' - got: %s", want, got)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Return value shoule be: '%s' - got: %s", want, got)
	}
}
