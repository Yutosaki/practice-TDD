package main

import (
	"bytes"
	"testing"
)
func TestCountDown(t *testing.T){
	buffer := &bytes.Buffer{}
	CountDown(buffer)

	got := buffer.String()
	want := "3\n2\n1\nGo!\n"

	if got != want {
		t.Errorf("want: %q, but got: %q",want, got)
	}
}