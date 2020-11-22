package test

import "testing"

func TestHello(t *testing.T) {
	want := "Hi, Tester. Welcome!"
	if got := Hello("Tester"); got != want {
		t.Errorf("Hello(\"Tester\") = %q, want %q", got, want)
	}
}