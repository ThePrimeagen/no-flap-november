package game

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/theprimeagen/the-game/pkg/models"
)

type model struct {
	lastUpdate  time.Time
	updateCount int64
	TimeAlive   float64

	updateables []models.Updateable

    context   *models.Context

	state State
}

func InitialModel() *model {
	context := models.Empty()
	term := &models.Terminal{}
	w, h := term.GetFixedBounds()
	gameEvent := models.CreateGameEvent()
	bird := models.CreateBird(gameEvent)
	screen := models.NewScreen2(context, w, h)
	pipes := models.NewPipes(context)
	debug := models.NewDebug(context)

    context.Hydrate(
        screen, bird, term, pipes, gameEvent, debug,
    )

	return &model{
		lastUpdate:  time.Now(),
		updateCount: 0,

		TimeAlive:   0.0,
		updateables: []models.Updateable{pipes, bird},

		state:     Playing,
        context:   context,
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

		events := m.context.Events.GetEvents()

		if len(events) > 0 {
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
			m.context.Bird.Jump()
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.context.Terminal.UpdateBounds(msg.Width, msg.Height)

		// TODO: Flicker?
		m.context.Screen.Clear()
	}

	return m, nil
}

func (m *model) View() string {
	width, height := m.context.Terminal.GetBounds()

	if width == 0 || height == 0 {
		return ""
	}

    // TODO: probably wasteful, but lets start here
    output := models.NewScreen2(m.context, width, height)
    output.Clear()

    // 1 render debug
    // 2 render pipes
    // 3 render bird

    output.Render(m.context.Debug)
    offset := m.context.Debug.LastRenderedHeight

    for _, pipe := range m.context.Pipes.Pipes {
        m.context.Screen.Render(pipe)
    }

    // TODO: collision
    gameEnded := m.context.Screen.Render(m.context.Bird)
    if gameEnded {
        m.context.Events.EmitBirdyCollision()
    }

    output.RenderAt(&models.Point{
        X: 0,
        Y: float64(offset),
    }, m.context.Screen);

    str := output.String()

	m.context.Screen.Clear()

	return str
}
