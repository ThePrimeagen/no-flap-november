package models

import "time"

type ITerminal interface {
    UpdateBounds(width, height int)
    GetBounds() (int, int)
    GetFixedBounds() (int, int)
    ScalingYFactor() float64
    ScalingXFactor() float64
}

// NOTE: I am just hacking because i don't know how to organize game code
type World interface {
    GetBounds() (int, int)
    GetFloorY() int
}

type Renderer interface {
    Render(Renderable) bool
    RenderAt(*Point, Renderable) bool
    RenderAll([]Renderable)
}

type Renderable interface {
    CreateRender(scale int) (*Point, [][]byte)
}

type Updateable interface {
    Update(t time.Duration)
}
