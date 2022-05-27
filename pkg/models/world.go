package models

const IDEAL_Y = 26.0
const IDEAL_X = 98.0

type NoFlapWorld struct {
    width int
    height int
}

func (n *NoFlapWorld) GetBounds() (int, int) {
    return n.width, n.height
}

func (n *NoFlapWorld) UpdateBounds(width, height int) {
    n.width = width
    n.height = height
}

func (n *NoFlapWorld) ScalingYFactor() float64 {
    return float64(n.height) / IDEAL_Y;
}

func (n *NoFlapWorld) ScalingXFactor() float64 {
    return float64(n.width) / IDEAL_X;
}

func (n *NoFlapWorld) GetFloorY() int {
    return n.height - 2
}
