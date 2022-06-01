package models

import (
	"strings"
)

type Screen2 struct {
	w int
    h int

	data [][]byte
	context *Context
}

func NewScreen2(context *Context, w, h int) *Screen2 {
	data := make([][]byte, h)
	for i := 0; i < h; i++ {
		data[i] = make([]byte, w)
	}

    screen := &Screen2{
        w,
        h,
        data,
        context,
	}

    screen.Clear()

    return screen
}

func (s *Screen2) Clear() {
	for row := 0; row < s.h; row++ {
        for col := 0; col < s.w; col++ {
            s.data[row][col] = ' ';
        }
	}
}

func (s *Screen2) render(pos *Point, data [][]byte) bool {
    collision := false
    for row, line := range data {
        for col, c := range line {
            collision = collision || s.data[int(pos.Y) + row][int(pos.X) + col] != ' '
            s.data[int(pos.Y) + row][int(pos.X) + col] = c
        }
    }
    return collision
}

// flappy
func (s *Screen2) Render(renderable Renderable) bool {
    pos, data := renderable.CreateRender()
    return s.render(pos, data)
}

// pipes
func (s *Screen2) RenderAll(renderable []Renderable) {
    for _, renderable := range renderable {
        s.Render(renderable)
    }
}

func (s *Screen2) RenderAt(offset *Point, renderable Renderable) bool {
    pos, data := renderable.CreateRender()

    if int(offset.Y) + int(pos.Y) + len(data) > len(s.data) {
        return false
    }

    point := &Point{
        X: offset.X + pos.X,
        Y: offset.Y + pos.Y,
    }

    return s.render(point, data)
}

func (s *Screen2) CreateRender() (*Point, [][]byte) {
    return &Point{
        X: 0,
        Y: 0,
    }, s.data
}

func (s *Screen2) String() string {
	screenString := []string{}
	for _, line := range s.data {
		screenString = append(screenString, string(line))
	}

    return strings.Join(screenString, "\n")
}
