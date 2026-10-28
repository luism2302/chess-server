package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	model := newModel()
	program := tea.NewProgram(model, tea.WithWindowSize(1600, 1600))

	if _, err := program.Run(); err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
}
