package models

import (
	"fmt"
	"strings"
	"time"
)

type Screen struct {
    updateCount int
    debugYHeight int
    extraDebug[] string
    Screen [][]byte

	startTime  time.Time
    world World
    renderCount int
}

func createScreen(width, height int) [][]byte {
    screen := make([][]byte, height)
    for i := 0; i < height; i += 1 {
        screen[i] = make([]byte, width)
    }
    return screen
}

func CreateScreen(world World) *Screen {
    width, height := world.GetBounds()
    return &Screen {
        Screen: createScreen(width, height),
        debugYHeight: 1,
        extraDebug: []string{},
        updateCount: 0,
        startTime: time.Now(),
        world: world,
        renderCount: 0,
    }
}

func (s *Screen) UpdateScreen() {
    width, height := s.world.GetBounds()
    s.Screen = createScreen(width, height)
}

func (s *Screen) Render(pos *Point, rendered [][]byte) {
    width, height := s.world.GetBounds()
    s.renderCount += 1
    for h, row := range rendered {
        offsetY := s.debugYHeight + len(s.extraDebug) + int(pos.Y) + h
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

            s.Screen[offsetY][offsetX] = b
        }
    }
}

func (s *Screen) debugMsg(msg string) string {
    width, _ := s.world.GetBounds()

    if len(msg) > width {
        msg = msg[:width]
    }
    return msg
}

func (s *Screen) AddDebug(msg string, index int) {
    if len(s.extraDebug) <= index {
        amount := index - len(s.extraDebug)
        for i := 0; i <= amount; i += 1 {
            s.extraDebug = append(s.extraDebug, "")
        }
    }
    s.extraDebug[index] = msg
}

func (s *Screen) debugRender() {
    if len(s.Screen) == 0 {
        return
    }
    width, height := s.world.GetBounds()

    statusLine := s.debugMsg(fmt.Sprintf("(%v, %v)(%v): %v", width, height, s.world.ScalingXFactor(), s.renderCount))

    copy(s.Screen[0], []byte(statusLine))
    for i, line := range s.extraDebug {
        copy(s.Screen[1 + i], []byte(s.debugMsg(line)))
    }

    s.renderCount = 0
}

func (s *Screen) Update(t time.Duration) {
    s.updateCount += 1
}

func (s *Screen) Clear() {
    for _, row := range s.Screen {
        for w := range row {
            row[w] = ' '
        }
    }
}

func (s *Screen) String() string {
    s.debugRender()
    screenString := []string{}
    for _, row := range s.Screen {
        // log.Println(string(row))
        screenString = append(screenString, string(row))
    }

    return strings.Join(screenString, "\n")
}

