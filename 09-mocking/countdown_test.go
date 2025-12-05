package main

/* Rules of thumb to help you staying with TDD
1. Ask yourself: "Am I testing the behavior or the implementation details"
2. A refactoring should not lead you to having to change any tests.  The outcome stays the same but the code changes in a refactoring.
3. Private functions should not be tested because they describe implementation detail and not public behavior. Test only public behavior.
4. Tests with more than 3 mocks is a red flag. You probably need changes in the design.
5. Use spies with caution because you are coupling the tests to the implementation. Be sure you care about what you spy on.
*/

import (
	"bytes"
	"reflect"
	"testing"
	"time"
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

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
