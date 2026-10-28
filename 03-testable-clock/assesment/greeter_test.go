package greeter

import "testing"

func TestGreetMorning(t *testing.T) {
	// TODO: add a Clock interface and a fixedClock, then make this test
	// pass at any time of day.
	t.Skip("TODO")

	g := Greeter{}
	if got := g.Greet("Ann"); got != "Good morning, Ann" {
		t.Errorf("got %q", got)
	}
}
