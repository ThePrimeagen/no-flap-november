package models

import (
	"fmt"
	"strings"
)

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

func (s *Debug) LineCount() int {
    return 1 + len(s.lines)
}

func (s *Debug) String() string {
	width := s.context.Screen.w;
    height := 1 + len(s.lines)

    out := []string{}

	sX := s.context.Terminal.ScalingXFactor(0)
	sY := s.context.Terminal.ScalingYFactor(0)

	out = append(out, s.debugMsg(fmt.Sprintf("(%v, %v)(%v): %v", width, height, sX, sY), width))

	for _, line := range s.lines {
        out = append(out, s.debugMsg(line, width))
	}

	return strings.Join(out, "\n")
}

func (s *Debug) debugMsg(msg string, bound int) string {
	if len(msg) > bound {
		msg = msg[:bound]
	}
	return msg
}
