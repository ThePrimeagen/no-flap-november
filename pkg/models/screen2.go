package models

type Screen2 struct {
	h int
	w int

	data [][]byte
	context *Context
}

func NewScreen2(context *Context, w, h int) *Screen2 {
	data := [][]byte{}
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

func (s *Screen2) Render(pos *Point, data [][]byte) bool {
    collision := false
    for row, line := range data {
        for col, c := range line {
            collision = collision || s.data[int(pos.Y) + row][int(pos.X) + col] != ' '
            s.data[int(pos.Y) + row][int(pos.X) + col] = c
        }
	}
    return collision
}

