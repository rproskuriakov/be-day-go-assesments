package greeter

import "time"

type Greeter struct{}

func (g Greeter) Greet(name string) string {
	if time.Now().Hour() < 12 {
		return "Good morning, " + name
	}
	return "Hello, " + name
}
