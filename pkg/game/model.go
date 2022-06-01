package game

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/theprimeagen/the-game/pkg/models"
)

type State = int

const (
	Playing State = iota
	GameOver
)

type model struct {
	lastUpdate  time.Time
	updateCount int64
	TimeAlive   float64

	updateables []models.Updateable
	renderables []models.Renderable

	term      *models.Terminal
	Screen    *models.Screen
	Bird      *models.Bird
	Pipes     *models.Pipes
	GameEvent *models.GameEvent
	state     State
}

func InitialModel() *model {
    context := models.Empty()
	term := &models.Terminal{}
    w, h := term.GetFixedBounds()
	gameEvent := models.CreateGameEvent()
	bird := models.CreateBird(gameEvent)
	screen := models.NewScreen2(context, w, h)
	pipes := models.NewPipes(term, screen)

	return &model{
		lastUpdate:  time.Now(),
		updateCount: 0,

		TimeAlive:   0.0,
		updateables: []models.Updateable{screen, pipes, bird},
		renderables: []models.Renderable{pipes, bird},

		Bird:   bird,
		Screen: screen,
		term:   term,

        GameEvent: gameEvent,
        state: Playing,
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
	case models.GameOverMessage:
        m.state = GameOver

	case frameMsg:

		// TODO: Timing would be great here
        if m.state != Playing {
            return m, animate()
        }

		delta := time.Since(time.Time(msg))
		for _, updateable := range m.updateables {
			updateable.Update(delta)
		}

		events := m.GameEvent.GetEvents()

		if len(events) > 0 {
            log.Fatal("OHH BABE")
			events = append(events, animate())
			return m, tea.Batch(events...)
		}

		// diff := FPS_SECONDS - time.Since(time.Time(msg)).Seconds()
		return m, animate() // slightly not on time updates

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "k":
			m.Bird.Jump()
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.term.UpdateBounds(msg.Width, msg.Height)
		m.Screen.UpdateScreenSize()

		// TODO: Flicker?
		m.Screen.Clear()
	}

	return m, nil
}

func (m *model) View() string {
	width, height := m.term.GetBounds()
	if width == 0 || height == 0 {
		return ""
	}

	for _, renderable := range m.renderables {
		renderable.Render(m.Screen)
	}

	str := m.Screen.String()
	m.Screen.Clear()

	return str
}
