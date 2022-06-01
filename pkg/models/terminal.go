package models

const fixedX = 98.0
const fixedY = 35.0

type Terminal struct {
    width int
    height int
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

func (t *Terminal) ScalingYFactor() float64 {
	return float64(t.width) / float64(fixedX)
}
func (t *Terminal) ScalingXFactor() float64 {
	return float64(t.height) / float64(fixedY)
}
