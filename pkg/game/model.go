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
    renderables []models.Renderable

    World  *models.NoFlapWorld
	Screen *models.Screen
	Bird   *models.Bird
}

func InitialModel() *model {
    world := &models.NoFlapWorld{}
    bird := models.CreateBird(world)
    screen := models.CreateScreen(world)

	return &model{
		lastUpdate:  time.Now(),
		updateCount: 0,

		TimeAlive: 0.0,
        updateables: []models.Updateable{screen, bird},
        renderables: []models.Renderable{bird},

        Bird: bird,
        Screen: screen,
        World: world,
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

        // TODO: Timing would be great here

        delta := time.Since(time.Time(msg))
        for _, updateable := range m.updateables {
            updateable.Update(delta)
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

        m.World.UpdateBounds(msg.Width, msg.Height)
        for _, updateable := range m.updateables {
            updateable.UpdateScreen()
        }

        // TODO: Flicker?
        m.Screen.Clear()
	}

	return m, nil
}

func (m *model) View() string {
    m.Bird.Render(m.Screen)
    str := m.Screen.String()
    m.Screen.Clear()

    return str
}
