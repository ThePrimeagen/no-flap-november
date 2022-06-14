package scene

import (
	"strings"
	"time"

	"github.com/theprimeagen/the-game/pkg/models"
)

type GameScene struct {
    context   *models.Context
}

func NewGameScene() *GameScene {
    return &GameScene{}
}

func (g *GameScene) InitializeScene(term *models.Terminal, eventer *models.GameEventer) {
    context := &models.Context{
        Terminal: term, // TODO: i don't like this.
    }

    debug := models.NewDebug(context)
    context.Debug = debug // TODO: Look at todo above me.

	bird := models.CreateBird(eventer)
	screen := models.NewScreen2(context, true, 420)
	pipes := models.NewPipes(context)

    screen.Clear()

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

    // 1 render pipes
    // 2 render bird
    // 3 render debug to string + screen to string
    for _, pipe := range m.context.Pipes.Pipes {
        m.context.Screen.Render(pipe)
    }

    gameEnded := m.context.Screen.Render(m.context.Bird)
    if gameEnded {
        m.context.Events.AddEvent(models.GameOverEvent)
    }

    str := strings.Join([]string{m.context.Debug.String(), m.context.Screen.String()}, "\n");

	m.context.Screen.Clear()

	return str
}

func (g* GameScene) HandleKeyPress(key string) {
    // These keys should exit the program.
    if key == "k" {
        g.context.Bird.Jump()
    }
}

func (g* GameScene) WindowResize(width, height int) {
    g.context.Screen.Clear()
}

