package models

import (
	"fmt"
	"strings"
	"time"
)

type Screen struct {
    updateCount int
    debugYHeight int
    Screen [][]byte

	startTime  time.Time
    world World
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
        updateCount: 0,
        startTime: time.Now(),
        world: world,
    }
}

func (s *Screen) UpdateScreen() {
    width, height := s.world.GetBounds()
    s.Screen = createScreen(width, height)
}

func (s *Screen) Render(pos *Point, rendered [][]byte) {
    if len(s.Screen) <= int(pos.Y) + len(rendered) {
        return
    }

    for h, row := range rendered {
        for w, b := range row {
            s.Screen[s.debugYHeight + int(pos.Y) + h][int(pos.X) + w] = b
        }
    }
}

func (s *Screen) Update(t time.Duration) {
    s.updateCount += 1

    fps := float64(s.updateCount) / time.Since(s.startTime).Seconds()
    width, height := s.world.GetBounds()

    debugLine := []byte(fmt.Sprintf("(%v, %v): I have updated %v times %v", width, height, s.updateCount, fps))
    if len(debugLine) > width {
        debugLine = debugLine[:width]
    }

    copy(s.Screen[0], debugLine)
}

func (s *Screen) Clear() {
    for _, row := range s.Screen {
        for w := range row {
            row[w] = ' '
        }
    }
}

func (s *Screen) String() string {
    screenString := []string{}
    for _, row := range s.Screen {
        // log.Println(string(row))
        screenString = append(screenString, string(row))
    }

    return strings.Join(screenString, "\n")
}

