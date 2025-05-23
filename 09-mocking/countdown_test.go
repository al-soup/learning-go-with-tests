package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("sleep before every sprint", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}

		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

	t.Run("prints number 3 to 1 and the 'Go!'", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		buffer := &bytes.Buffer{}

		Countdown(buffer, spySleepPrinter)
		want := `3
2
1
Go!`
		got := buffer.String()

		if want != got {
			t.Errorf("want %q, got %q", want, got)
		}

	})
}
