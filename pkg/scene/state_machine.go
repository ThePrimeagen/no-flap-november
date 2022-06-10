package scene

import "github.com/theprimeagen/the-game/pkg/models"

type GameStateMachine struct {
    scenes map[string]Scene
    state State
}

func NewGameStateMachine() *GameStateMachine {
    return &GameStateMachine{
        scenes: map[string]Scene{
        },
        state: WelcomeScreen,
    }
}

func (g *GameStateMachine) InitializeScene(term *models.Terminal) {
    // todo: ...
}
