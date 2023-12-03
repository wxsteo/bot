package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hi I'm chat version 0.0.0")

	fmt.Println(greetingAccordingTime())
}

func greetingAccordingTime() string {
	hour := time.Now().Hour()

	switch {
	case hour < 12:
		return "Good Morning!"

	case hour < 18:
		return "Good Afternoon!"

	case hour < 21:
		return "Good Evening!"
	}

	return "Good Night!"
}
