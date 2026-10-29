package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("couldn't create debug.log file")
	}
	defer f.Close()

	model := newModel()
	program := tea.NewProgram(model, tea.WithWindowSize(1600, 1600))

	if _, err := program.Run(); err != nil {
		log.Fatalf("error during runtime: %s", err.Error())
	}
}
