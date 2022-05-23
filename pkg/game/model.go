package game

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/theprimeagen/the-game/pkg/models"
	"github.com/theprimeagen/the-game/pkg/physics"
)

type model struct {
	lastUpdate  time.Time
	updateCount int64
	TimeAlive   float64
	Width       int
	Height      int

	Screen *models.Screen
	Bird   *models.Bird
}

func InitialModel() *model {
	return &model{
		lastUpdate:  time.Now(),
		updateCount: 0,

		TimeAlive: 0.0,
		Width:     0,
		Height:    0,

		Bird:   models.CreateBird(),
		Screen: models.CreateScreen(0, 0),
	}
}

type frameMsg time.Time

func animate() tea.Cmd {
	now := time.Now()
	timeSince := time.Second / FPS
	return tea.Tick(timeSince, func(t time.Time) tea.Msg {
		return frameMsg(now)
	})
}

func (m *model) Init() tea.Cmd {
	return tea.Batch(animate(), tea.EnterAltScreen)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case frameMsg:
		m.updateCount += 1
		// diff := FPS_SECONDS - time.Since(time.Time(msg)).Seconds()
		return m, animate() // slightly not on time updates

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Screen.UpdateScreen(m.Width, m.Height)

        // TODO: Flicker?
        m.Screen.Clear()
	}

	return m, nil
}

func (m *model) View() string {
    return m.Screen.String()
}
