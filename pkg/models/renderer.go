package models

import "time"

// NOTE: I am just hacking because i don't know how to organize game code
type World interface {
    GetBounds() (int, int)
    GetFloorY() int
    UpdateBounds(width, height int)
    ScalingYFactor() float64
    ScalingXFactor() float64
}

type Renderer interface {
    Render(point *Point, item [][]byte)
}

type Renderable interface {
    Render(renderer Renderer)
}

type Updateable interface {
    Update(t time.Duration)
    UpdateScreen()
}
