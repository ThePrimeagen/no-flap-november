package models

import (
	"fmt"
	"time"
)

const MICROSECONDS_TO_X = 45_000
const STARTING_TUBE_SPACING = 25
const MINIMUM_TUBE_SPACING = 3
const REDUCTION_SCALE = 5

type Pipes struct {
	screen *Screen
	world World

	totalPipes int

	elapsedTime     int64
	currentStep     int64
	lastPipeCreated int64

	pipes []*pipe
}

type pipe struct {
    x int
}

func newPipe(startingX int) *pipe {
    return &pipe {
        x: startingX,
    }
}

func NewPipes(world World, screen *Screen) *Pipes {
	return &Pipes{
		screen:          screen,
		world:           world,
		lastPipeCreated: 0,
		elapsedTime:     0,
		currentStep:     0,
		pipes:           []*pipe{},
        totalPipes: 0,
	}
}

func min(one float64, two int) float64 {
    two_f := float64(two)
    if two_f > one {
        return one
    }
    return two_f
}

func maxInt(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}

func (p *Pipes) canCreatePipe() bool {
    if p.lastPipeCreated == 0 {
        return true
    }

    pipeCount := 1
    takenSteps := p.elapsedTime / MICROSECONDS_TO_X

    for {
        currentStepsRequired := maxInt(
            int64(STARTING_TUBE_SPACING - pipeCount / REDUCTION_SCALE),
            MINIMUM_TUBE_SPACING,
        )

        if takenSteps < int64(currentStepsRequired) {
            break;
        }

        pipeCount += 1
        takenSteps -= int64(currentStepsRequired)
    }

    return p.totalPipes < pipeCount
}

func (p *Pipes) Update(delta time.Duration) {
    width, _ := p.world.GetBounds()
	p.elapsedTime += delta.Microseconds()
    if p.canCreatePipe() {
        p.pipes = append(p.pipes, newPipe(width))
        p.lastPipeCreated = p.elapsedTime
        p.totalPipes += 1
    }

    steps := p.elapsedTime / MICROSECONDS_TO_X
    if p.currentStep < steps {
        for _, pipe := range p.pipes {
            pipe.x -= 1
        }
        p.currentStep = steps
    }
}

func (p *Pipes) Render(renderer Renderer) {
    if len(p.pipes) == 0 {
        return
    }

    // so this will always work
    p.screen.AddDebug(fmt.Sprintf("pipes: %+v", p.pipes))
    if p.pipes[0].x < 0 {
        p.pipes = p.pipes[1:]
    }

    height := p.world.GetFloorY()
    startingHeight := height - 5

    pipeAscii := make([][]byte, 5)
    for i := 0; i < 5; i += 1 {
        pipeAscii[i] = []byte{'x'}
    }

    for _, pipe := range p.pipes {
        renderer.Render(&Point{
            X: float64(pipe.x - 1),
            Y: float64(startingHeight),
        }, pipeAscii)
    }
}

func (p *Pipes) UpdateScreen() {
    // TODO: Resizing is going to jack the game
    // So instead lets do some sweeeeeeeeet progressive rendering
}
