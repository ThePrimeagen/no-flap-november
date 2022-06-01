package models

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type GameEvent struct {
    messages []tea.Cmd
}

func CreateGameEvent() *GameEvent {
    return &GameEvent {
        messages: []tea.Cmd{},
    }
}

type GameOverMessage = int

func (g *GameEvent) emitBirdyCollision() {
    log.Fatal("OHH BABE")
    g.messages = append(g.messages, tea.Cmd(func() tea.Msg {
        return GameOverMessage(1);
    }))
}

func (g *GameEvent) GetEvents() []tea.Cmd {
    msg := g.messages

    g.messages = []tea.Cmd{}

    return msg
}

