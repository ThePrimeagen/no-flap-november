package models

const fixedX = 98.0
const fixedY = 33.0

type Terminal struct {
    width int
    height int
}

func NewTerminal() *Terminal {
    return &Terminal{
        width: fixedX,
        height: fixedY,
    }
}

func (t *Terminal) UpdateBounds(width, height int) {
    t.width = width;
    t.height = height;
}

func (t *Terminal) GetBounds() (int, int) {
	return t.width, t.height
}

func (t *Terminal) GetFixedBounds() (int, int) {
	return fixedX, fixedY
}

func (t *Terminal) ScalingYFactor(reduceW int) float64 {
	return float64(t.width - reduceW) / float64(fixedX)
}
func (t *Terminal) ScalingXFactor(reduceH int) float64 {
	return float64(t.height - reduceH) / float64(fixedY)
}

