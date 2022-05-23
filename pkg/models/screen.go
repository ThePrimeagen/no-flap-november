package models

import "strings"

type Screen struct {
    Screen [][]byte
}

func createScreen(width, height int) [][]byte {
    screen := make([][]byte, height)
    for i := 0; i < height; i += 1 {
        screen[i] = make([]byte, width)
    }
    return screen
}

func CreateScreen(width, height int) *Screen {
    return &Screen {
        Screen: createScreen(width, height),
    }
}

func (s *Screen) UpdateScreen(width, height int) {
    s.Screen = createScreen(width, height)
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
        screenString = append(screenString, string(row))
    }

    return strings.Join(screenString, "")
}

