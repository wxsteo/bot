package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ciao sono chat verzione 0.0.0")

	fmt.Println(greetingAccordingTime())

	var command string

	for {

		fmt.Print("Scrivi un commando: ")
		fmt.Scanln(&command)
		fmt.Println("Il tuo commando è:", command)

		if command == "uscire" {
			break
		}

		if command == "ore" {
			tempo := time.Now()
			strtempo := tempo.Format("2006-01-02 15:04")
			fmt.Println(strtempo)
		}

		fmt.Println("commando sconosciuto")
	}
}

func greetingAccordingTime() string {
	hour := time.Now().Hour()

	switch {
	case hour < 12:
		return "Buongiorno!"

	case hour < 18:
		return "Buon Pomeriggio!"

	case hour < 21:
		return "Buona Sera!"
	}

	return "Buona Notte!"
}
