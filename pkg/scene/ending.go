package scene

import (
	"time"

	"github.com/theprimeagen/the-game/pkg/models"
)

type EndingScene struct { }

func NewEndingScene() *EndingScene {
    return &EndingScene{ }
}

func (g *EndingScene) InitializeScene(term *models.Terminal, eventer *models.GameEventer) { }

func (g *EndingScene) Update(timeElapsed time.Duration) { }

func (g *EndingScene) Render() string {
    return "You suck at this game"
}

func (g *EndingScene) HandleKeyPress(key string) { }

func (g *EndingScene) WindowResize(width, height int) { }



