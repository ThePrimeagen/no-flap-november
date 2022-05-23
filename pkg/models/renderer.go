package models

type Renderer interface {
    Render(point *Point, item [][]byte)
}

type Renderable interface {
    Render(renderer Renderer)
}


