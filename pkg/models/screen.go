package models

import (
	"fmt"
	"strings"
	"time"
)

type Screen struct {
    updateCount int
    screen [][]byte
    render [][]byte
	status string
    debug []string

	startTime  time.Time
    term ITerminal
    renderCount int
}

func createScreen(width, height int) [][]byte {
    screen := make([][]byte, height)
    for i := 0; i < height; i += 1 {
        screen[i] = make([]byte, width)
    }
    return screen
}

func CreateScreen(term ITerminal) *Screen {
    width, height := term.GetFixedBounds()
    w, h := term.GetBounds()
    return &Screen {
        screen: createScreen(width, height),
        render: createScreen(w, h),
        debug: []string{},
        updateCount: 0,
        startTime: time.Now(),
        term: term,
        renderCount: 0,
    }
}

func (s *Screen) CheckForCollisions(pos *Point, rendered [][]byte) bool {
    // TODO: We should figure out a way to make this code to be the same for
    // both rendering and colliding.

    width, height := s.term.GetFixedBounds()
    collision := false

    msg := ""
    outer_loop: for h, row := range rendered {
        offsetY := 1 + len(s.debug) + int(pos.Y) + h
        if offsetY < 0 {
            continue
        }

        if offsetY >= height {
            break;
        }

        for w, b := range row {
            offsetX := int(pos.X) + w
            if offsetX < 0 {
                continue
            }

            if offsetX >= width {
                break;
            }

            msg = fmt.Sprintf("%v(%v,%v(%v), %v,%v(%v)) ", msg, offsetX, offsetY, string(s.screen[offsetX][offsetY]), w, h, string(b))
            if s.screen[offsetX][offsetY] != ' ' && b != ' ' {
                collision = true
                break outer_loop;
            }
        }
    }
    s.AddDebug(msg, 2)

    return collision
}

func (s *Screen) Render(pos *Point, rendered [][]byte) {
    width, height := s.term.GetFixedBounds()
    s.renderCount += 1
    for h, row := range rendered {
        offsetY := 1 + len(s.debug) + int(pos.Y) + h
        if offsetY < 0 {
            continue
        }

        if offsetY >= height {
            break;
        }

        for w, b := range row {
            offsetX := int(pos.X) + w
            if offsetX < 0 {
                continue
            }

            if offsetX >= width {
                break;
            }

            s.screen[offsetY][offsetX] = b
        }
    }
}

func (s *Screen) debugMsg(msg string, bound int) string {
    if len(msg) > bound {
        msg = msg[:bound]
    }
    return msg
}

func (s *Screen) AddDebug(msg string, index int) {
    if len(s.debug) <= index {
        amount := index - len(s.debug)
        for i := 0; i <= amount; i += 1 {
            s.debug = append(s.debug, "")
        }
    }
    s.debug[index] = msg
}

func (s *Screen) debugRender() {
    width, height := s.term.GetBounds()
    statusLine := s.debugMsg(fmt.Sprintf("(%v, %v)(%v): %v", width, height, s.term.ScalingXFactor(), s.renderCount), width)
    copy(s.render[0], []byte(statusLine))

    for i, line := range s.debug {
        copy(s.render[1 + i], []byte(s.debugMsg(line, width)))
    }

    s.renderCount = 0
}

func (s *Screen) Update(t time.Duration) {
    s.updateCount += 1
}

func (s *Screen) UpdateBounds() {
	width, height := s.term.GetBounds()
	s.render = createScreen(width, height)
}

func clear(screen [][]byte) {
    for _, row := range screen {
        for w := range row {
            row[w] = ' '
        }
    }
}

func (s *Screen) UpdateScreenSize() {
	width, height := s.term.GetBounds()
	s.render = createScreen(width, height)
}

func (s *Screen) Clear() {
	clear(s.screen)
}

func (s *Screen) String() string {
	if len(s.render) == 0 {
		return ""
	}

	clear(s.render)
    s.debugRender()

    fw, fh := s.term.GetFixedBounds()
	width, height := s.term.GetBounds()

	offsetX := 0
	offsetY := 0
	if fw < width {
		offsetX = (width - fw) / 2
	}
	if fh < height {
		offsetY = (height - fh) / 2
	}

	debugOffset := len(s.debug) + 1;
	maxY := height - debugOffset

	for i, line := range s.screen {
		if i >= maxY {
			break
		}

		offset := s.render[debugOffset + i + offsetY][offsetX:]
		copy(offset, line)
	}

	screenString := []string{}
	for _, line := range s.render {
		screenString = append(screenString, string(line))
	}

    return strings.Join(screenString, "\n")
}

// 123456789
