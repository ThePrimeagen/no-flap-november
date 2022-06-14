package game

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/theprimeagen/the-game/pkg/models"
	"github.com/theprimeagen/the-game/pkg/scene"
)

type model struct {
	lastUpdate  time.Time
	updateCount int64
	timeAlive   float64
    stateMachine *scene.GameStateMachine
    events *models.GameEventer
    terminal *models.Terminal
}

func InitialModel() *model {
	term := models.NewTerminal()
	gameEvent := models.CreateGameEvent()
    sm := scene.NewGameStateMachine()

    sm.InitializeScene(term, gameEvent)

	return &model{
		lastUpdate:  time.Now(),
		updateCount: 0,

		timeAlive:   0.0,
        stateMachine: sm,
        events: gameEvent,
        terminal: term,
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
    case models.GameEvent:
        m.stateMachine.HandleStateChange(msg)

	case frameMsg:
		delta := time.Since(time.Time(msg))
        m.stateMachine.Update(delta)

		events := m.events.GetEvents()

		if len(events) > 0 {
			events = append(events, animate())
			return m, tea.Batch(events...)
		}

        // 3 - rendering items
        // @   x
        //
        // .\
        // @/:

        // 1 . progressive rendering
        // 1.1  we need floor and ceiling
        // 2 . death screen / menu
        // 3 . score

		// diff := FPS_SECONDS - time.Since(time.Time(msg)).Seconds()
		return m, animate() // slightly not on time updates

	// Is it a key press?
	case tea.KeyMsg:
        keyPress := msg.String()
        switch keyPress {
        case "ctrl+c", "q":
            return m, tea.Quit
        default:
            m.stateMachine.HandleKeyPress(keyPress)
            return m, nil
        }

	case tea.WindowSizeMsg:
		m.terminal.UpdateBounds(msg.Width, msg.Height)
		m.stateMachine.WindowResize(msg.Width, msg.Height)
	}

	return m, nil
}

func (m *model) View() string {
    return m.stateMachine.Render()
}
