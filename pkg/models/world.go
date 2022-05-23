package models

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

