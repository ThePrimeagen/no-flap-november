package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type GameEvent = int
const (
    GameOverEvent GameEvent = iota
)

type GameEventer struct {
    messages []tea.Cmd
}

func CreateGameEvent() *GameEventer {
    return &GameEventer {
        messages: []tea.Cmd{},
    }
}

func (g *GameEventer) AddEvent(messageType GameEvent) {
    g.messages = append(g.messages, tea.Cmd(func() tea.Msg {
        return messageType
    }))
}

func (g *GameEventer) GetEvents() []tea.Cmd {
    msg := g.messages

    g.messages = []tea.Cmd{}

    return msg
}
