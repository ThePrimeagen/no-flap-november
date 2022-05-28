package models

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const MicrosecondsToX = 40_000
const StartingTubeSpacing = 150
const SpacingStepDown = 50
const MINIMUM_TUBE_SPACING = 12
const REDUCTION_SCALE = 3
const PIPE_DIST_FROM_EDGES = 5
const PIPE_HOLE_SIZE_MILF = 15

type Pipes struct {
	screen *Screen
	term  ITerminal

	totalPipes int

	elapsedTime     int64
	currentStep     int64
	lastPipeCreated int64

	pipes []*pipe
}

type pipe struct {
	x      int
	offset int

	display [][]byte
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func newPipe(startingX, height int) *pipe {

	p := &pipe{
		x:       startingX,
		display: [][]byte{},
	}

	p.sizeUpPipe(height)

	return p
}

func (p *pipe) sizeUpPipe(height int) {

    opening := randInt(PIPE_DIST_FROM_EDGES, height - PIPE_DIST_FROM_EDGES)
	top := opening - PIPE_HOLE_SIZE_MILF/2
	bottom := opening + PIPE_DIST_FROM_EDGES - PIPE_DIST_FROM_EDGES/2

	if top < PIPE_DIST_FROM_EDGES {
		diff := PIPE_DIST_FROM_EDGES - top
		bottom += diff
		top = PIPE_DIST_FROM_EDGES
	} else if bottom > height-PIPE_DIST_FROM_EDGES {
		diff := bottom - (height - PIPE_DIST_FROM_EDGES)
		top -= diff
		bottom = height - PIPE_DIST_FROM_EDGES
	}


    display := make([][]byte, height)
    for i := 0; i < height; i += 1 {
        if i <= top || i >= bottom  {
            display[i] = []byte{'x'}
        } else {
            display[i] = []byte{' '}
        }
    }

    p.display = display
}

func NewPipes(term ITerminal, screen *Screen) *Pipes {
	return &Pipes{
		lastPipeCreated: 0,
		elapsedTime:     0,
		currentStep:     0,
		pipes:           []*pipe{},
		totalPipes:      0,
		term: term,
		screen: screen,
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
	takenSteps := p.elapsedTime /MicrosecondsToX;

	for {
		scaledReduce := pipeCount / REDUCTION_SCALE

		currentStepsRequired := maxInt(
			int64(150*math.Pow(float64(scaledReduce+1), -.71)),
			MINIMUM_TUBE_SPACING,
		)

		if takenSteps < currentStepsRequired {
			break
		}

		pipeCount += 1
		takenSteps -= currentStepsRequired
	}


	return p.totalPipes < pipeCount
}

func (p *Pipes) Update(delta time.Duration) {
	width, height := p.term.GetFixedBounds()
	p.elapsedTime += delta.Microseconds()
	if p.canCreatePipe() {
        pipe := newPipe(width, height)
		p.pipes = append(p.pipes, pipe)
		p.lastPipeCreated = p.elapsedTime
		p.totalPipes += 1
	}

	steps := p.elapsedTime / MicrosecondsToX
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
	if p.pipes[0].x < 0 {
		p.pipes = p.pipes[1:]
	}


	for i, pipe := range p.pipes {
        if i == 0 {
            p.screen.AddDebug(fmt.Sprintf("pipes: %+v", pipe), 0)
        }
		renderer.Render(&Point{
			X: float64(pipe.x - 1),
			Y: float64(0),
		}, pipe.display)
	}
}