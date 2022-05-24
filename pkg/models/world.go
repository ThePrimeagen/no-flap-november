package models

const IDEAL_Y = 35.0

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

func (n *NoFlapWorld) GetFloorY() int {
    return n.height - 2
}
