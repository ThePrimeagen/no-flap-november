package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/theprimeagen/the-game/pkg/game"
)

func main() {
    p := tea.NewProgram(game.InitialModel(), tea.WithAltScreen())
    if err := p.Start(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

