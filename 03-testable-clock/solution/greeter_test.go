package greeter

import (
	"testing"
	"time"
)

type fixedClock struct{ t time.Time }

func (c fixedClock) Now() time.Time { return c.t }

func TestGreetMorning(t *testing.T) {
	g := Greeter{Clock: fixedClock{time.Date(2026, 1, 1, 9, 0, 0, 0, time.UTC)}}
	if got := g.Greet("Ann"); got != "Good morning, Ann" {
		t.Errorf("got %q", got)
	}
}

func TestGreetAfternoon(t *testing.T) {
	g := Greeter{Clock: fixedClock{time.Date(2026, 1, 1, 15, 0, 0, 0, time.UTC)}}
	if got := g.Greet("Ann"); got != "Hello, Ann" {
		t.Errorf("got %q", got)
	}
}
