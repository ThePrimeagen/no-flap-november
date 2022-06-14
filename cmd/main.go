package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/theprimeagen/the-game/pkg/game"
)

func main() {
    if f, err := tea.LogToFile("debug.log", "milf"); err != nil {
        fmt.Println("Couldn't open a file for logging:", err)
        os.Exit(1)
    } else {
        defer func() {
            err = f.Close()
            if err != nil {
                log.Fatal(err)
            }
        }()
    }


    log.Printf("--------------------------");
    p := tea.NewProgram(game.InitialModel(), tea.WithAltScreen())
    if err := p.Start(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

