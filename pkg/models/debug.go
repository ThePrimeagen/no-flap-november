package models

import "fmt"

type Debug struct {
	lines              []string
	LastRenderedHeight int
	context            *Context
}

func NewDebug(context *Context) *Debug {
	return &Debug{
		lines:              []string{},
		LastRenderedHeight: 0,

		context: context,
	}
}

func (s *Debug) AddDebug(msg string, index int) {
	if len(s.lines) <= index {
		amount := index - len(s.lines)
		for i := 0; i <= amount; i += 1 {
			s.lines = append(s.lines, "")
		}
	}
	s.lines[index] = msg
}

func (s *Debug) CreateRender() (*Point, [][]byte) {
	width, height := s.context.Terminal.GetBounds()
	screen := make([][]byte, height)
	for i := 0; i < height; i++ {
		screen[i] = make([]byte, width)
	}

	sX := s.context.Terminal.ScalingXFactor()
	sY := s.context.Terminal.ScalingYFactor()

	statusLine := s.debugMsg(fmt.Sprintf("(%v, %v)(%v): %v", width, height, sX, sY), width)

	y := 0
	copy(screen[y], []byte(statusLine))

	y++

	for i, line := range s.lines {
		copy(screen[y+i], []byte(s.debugMsg(line, width)))
	}

    s.LastRenderedHeight = len(screen)

	return &Point{0, 0}, screen
}

func (s *Debug) debugMsg(msg string, bound int) string {
	if len(msg) > bound {
		msg = msg[:bound]
	}
	return msg
}
