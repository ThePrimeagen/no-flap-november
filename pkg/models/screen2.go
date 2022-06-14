package models

import (
	"log"
	"math"
	"strings"
)

type Screen2 struct {
	w  int
	h  int
	id int

	data    [][]byte
	context *Context

	renderSize    int
	considerDebug bool
}

func createRenderSurface(w, h int) [][]byte {
	data := make([][]byte, h)
	for i := 0; i < h; i++ {
		data[i] = make([]byte, w)
	}
	return data
}

func NewScreen2(context *Context, considerDebug bool, id int) *Screen2 {
	width, height := context.Terminal.GetFixedBounds()
	screen := &Screen2{
		w:  width,
		h:  height,
		id: id,

		data:          createRenderSurface(width, height),
		context:       context,
		renderSize:    0,
		considerDebug: considerDebug,
	}

	log.Printf("%v creating NewScreen2 %v, %v", id, width, height)

	screen.Clear()

	return screen
}

func (s *Screen2) Clear() {
	size := int(s.getRenderSize())
	width, height := s.getWidthAndHeight()

	if s.renderSize != size {
		log.Printf("%v Screen2#Clear changing width/height %v, %v", s.id, width, height)
		s.data = createRenderSurface(width, height)
		s.renderSize = size
		s.w = width
		s.h = height
	}

	for row := 0; row < s.h; row++ {
		for col := 0; col < s.w; col++ {
			s.data[row][col] = ' '
		}
	}
}

func (s *Screen2) render(pos *Point, data [][]byte) bool {
	collision := false
	for row, line := range data {
		if int(pos.Y)+row >= s.h {
			log.Printf("%v we are about to access row outside of bounds pos(%v, %v) rend(%v, %v) surf(%v, %v) -- %v", s.id, pos.X, pos.Y, len(data), len(data[0]), len(s.data), len(s.data[0]), row)
		}
		for col, c := range line {
			if int(pos.X)+col >= s.w {
				log.Printf("%v we are about to access col outside of bounds pos(%v, %v) rend(%v, %v) surf(%v, %v) -- %v", s.id, pos.X, pos.Y, len(data), len(data[0]), len(s.data), len(s.data[0]), row)
			}
			collision = collision || (s.data[int(pos.Y)+row][int(pos.X)+col] != ' ' &&
                data[row][col] != ' ')

			s.data[int(pos.Y)+row][int(pos.X)+col] = c
		}
	}
	return collision
}

func (s *Screen2) getMinScaling() float64 {
	yReduce := 0

	if s.considerDebug {
		yReduce = s.context.Debug.LineCount()
	}

    scaleX := s.context.Terminal.ScalingXFactor(0)
    scaleY := s.context.Terminal.ScalingXFactor(yReduce)

	log.Printf("%v getMinScaling(%v): math.min(%v, %v)", s.id, yReduce, scaleX, scaleY)
    return math.Min(scaleX, scaleY)
}

func (s *Screen2) getRenderSize() int {
	scale := s.getMinScaling()
    out := 0
	if scale > 4 {
		out = 2
	} else if scale > 2 {
		out =  1
	}
    log.Printf("%v getRenderSize(): %v -> %v", s.id, scale, out)
	return out
}

func (s *Screen2) getWidthAndHeight() (int, int) {
	scale := s.getRenderSize()
	width, height := s.context.Terminal.GetFixedBounds()

	scale_mul := int(math.Pow(2, float64(scale)))
	return width * scale_mul, height * scale_mul
}

// flappy
func (s *Screen2) Render(renderable Renderable) bool {
	pos, data := renderable.CreateRender(s.renderSize)

	scale := float64(s.renderSize) + 1
	pos_scaled := Mul(pos, scale, scale)

	log.Printf("%v Screen2#Render(%v, %v) %+v -> %+v (%v, %v)", s.id, scale, s.renderSize, pos, pos_scaled, len(s.data), len(s.data[0]))
	return s.render(pos_scaled, data)
}

// pipes
func (s *Screen2) RenderAll(renderable []Renderable) {
	for _, renderable := range renderable {
		s.Render(renderable)
	}
}

func (s *Screen2) RenderAt(offset *Point, renderable Renderable) bool {
	pos, data := renderable.CreateRender(s.getRenderSize())

	log.Printf("%v Screen2#RenderAt(%+v,_): %+v, %v > %v", s.id, offset, pos, len(data), len(s.data))
	if int(offset.Y)+int(pos.Y)+len(data) > len(s.data) {
		return false
	}

	point := &Point{
		X: offset.X + pos.X,
		Y: offset.Y + pos.Y,
	}

	return s.render(point, data)
}

func (s *Screen2) CreateRender(size int) (*Point, [][]byte) {
	return &Point{
		X: 0,
		Y: 0,
	}, s.data
}

func (s *Screen2) String() string {
	screenString := []string{}
	log.Print("Printing Screen")
	for _, line := range s.data {
		log.Print(string(line))
		screenString = append(screenString, string(line))
	}

	return strings.Join(screenString, "\n")
}
