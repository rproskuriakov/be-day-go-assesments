package greeter

import "time"

type Clock interface {
	Now() time.Time
}

type Greeter struct{ Clock Clock }

func (g Greeter) Greet(name string) string {
	if g.Clock.Now().Hour() < 12 {
		return "Good morning, " + name
	}
	return "Hello, " + name
}
