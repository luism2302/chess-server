package main

import (
	"flag"
	"log"

	tea "charm.land/bubbletea/v2"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("couldn't create debug.log file")
	}
	defer f.Close()

	fenCode := flag.String("fen", "", "fen string to load")
	flag.Parse()

	model, err := newModel(*fenCode)
	if err != nil {
		log.Fatalf("couldn't initialize model: %s", err.Error())
	}
	program := tea.NewProgram(model, tea.WithWindowSize(1600, 1600))

	if _, err := program.Run(); err != nil {
		log.Fatalf("error during runtime: %s", err.Error())
	}
}
