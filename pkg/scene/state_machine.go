package scene

import (
	"time"

	"github.com/theprimeagen/the-game/pkg/models"
)

type GameStateMachine struct {
    scenes map[State]Scene
    state State
}

func NewGameStateMachine() *GameStateMachine {
    return &GameStateMachine{
        scenes: map[State]Scene{
        },
        state: Playing,
    }
}

func (g *GameStateMachine) InitializeScene(term *models.Terminal, eventer *models.GameEventer) {

    g.scenes[Playing] = NewGameScene()
    g.scenes[Playing].InitializeScene(term, eventer)

    g.scenes[GameOver] = NewEndingScene()
    g.scenes[GameOver].InitializeScene(term, eventer)
}

func (g *GameStateMachine) Update(timeElapsed time.Duration) {
    for _, scene := range g.scenes {
        scene.Update(timeElapsed)
    }
}

func (g *GameStateMachine) Render() string {
    return g.scenes[g.state].Render()
}

func (g *GameStateMachine) HandleKeyPress(key string) {
    g.scenes[g.state].HandleKeyPress(key)
}

func (g *GameStateMachine) WindowResize(width, height int) {
    g.scenes[g.state].WindowResize(width, height)
}

func (g *GameStateMachine) HandleStateChange(event models.GameEvent) {
    switch (event) {
    case models.GameOverEvent:
        g.state = GameOver
    }
}

