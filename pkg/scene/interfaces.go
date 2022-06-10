package scene

import (
	"time"

	"github.com/theprimeagen/the-game/pkg/models"
)

type State = int

const (
    WelcomeScreen State = iota
	Playing
	GameOver
)

type Scene interface {
    InitializeScene(term *models.Terminal, eventer *models.GameEventer)
    Update(timeElapsed time.Duration)
    Render() string
    HandleKeyPress(key string)
    WindowResize(width, height int)
}

