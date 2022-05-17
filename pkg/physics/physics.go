package physics

type Vector2D struct {
    x float64
    y float64
}

func NewVector2D(x, y float64) *Vector2D {
    return &Vector2D{ x, y, }
}

func (v *Vector2D) Copy() *Vector2D {
    return &Vector2D {
        x: v.x,
        y: v.y,
    }
}

func (v *Vector2D) Apply(other *Vector2D, delta float64) {
    v.x *= other.x * delta
    v.y *= other.y * delta
}
