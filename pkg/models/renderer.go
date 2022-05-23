package models

type Renderer interface {
    Render(point *Point, item [][]byte)
    GetBounds() (int, int)
}

type Renderable interface {
    Render(renderer Renderer)
}


