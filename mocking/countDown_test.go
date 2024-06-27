package main

import (
	"bytes"
	"reflect"
	"testing"
)
func TestCountDown(t *testing.T){
	t.Run("print 3 to Go!", func(t *testing.T){
		buffer := &bytes.Buffer{}
		CountDown(buffer, &CountDownOperationsSpy{})

		got := buffer.String()
		want := "3\n2\n1\nGo!\n"

		if got != want {
			t.Errorf("want: %q, but got: %q",want, got)
		}
	})

	t.Run("sleep before every print", func(t *testing.T){
		spySleepPrinter := &CountDownOperationsSpy{}
		CountDown(spySleepPrinter, spySleepPrinter)
		
		want := []string{sleep, write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls){
			t.Errorf("want: %v, but got: %v", want, spySleepPrinter.Calls)
		}
	})
}

