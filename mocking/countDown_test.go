package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
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

func TestConfigurableSleeper(t *testing.T){
	sleepTime := 5 * time.Second

	SpyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, SpyTime.Sleep}
	sleeper.Sleep()

	if SpyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, SpyTime.durationSlept)
	}
}
