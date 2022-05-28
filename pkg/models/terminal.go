package models

const idealX = 98.0
const idealY = 35.0

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
	return idealX, idealY
}

func (t *Terminal) ScalingYFactor() float64 {
	return float64(t.width) / float64(idealX)
}
func (t *Terminal) ScalingXFactor() float64 {
	return float64(t.height) / float64(idealY)
}
