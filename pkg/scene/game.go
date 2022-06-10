package scene

import (
	"time"

	"github.com/theprimeagen/the-game/pkg/models"
)

type GameScene struct {
    context   *models.Context
}

func NewGameScene() {
}

func (g *GameScene) InitializeScene(term *models.Terminal, eventer *models.GameEventer) {
	w, h := term.GetFixedBounds()

    context := &models.Context{}
	bird := models.CreateBird(eventer)
	screen := models.NewScreen2(context, w, h)
	pipes := models.NewPipes(context)
	debug := models.NewDebug(context)

    context.Hydrate(
        screen, bird, term, pipes, eventer, debug,
    )

    g.context = context;
}

func (m *GameScene) Update(delta time.Duration) {
    m.context.Pipes.Update(delta)
    m.context.Bird.Update(delta)
}

func (m *GameScene) Render() string {
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
        m.context.Events.AddEvent(models.GameOverEvent)
    }

    output.RenderAt(&models.Point{
        X: 0,
        Y: float64(offset),
    }, m.context.Screen);

    str := output.String()

	m.context.Screen.Clear()

	return str
}

func (g* GameScene) HandleKeyPress(key string) {
}

func (g* GameScene) WindowResize(width, height int) {
}

