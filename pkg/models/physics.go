package models

type Point struct {
    X float64
    Y float64
}

func NewPoint2D(x, y float64) *Point {
    return &Point{ x, y }
}

func AddValues(p1 *Point, x, y float64) Point {
    return Point {
        X: p1.X + x,
        Y: p1.Y + y,
    }
}

func Add(p1, p2 *Point) *Point {
    return &Point {
        X: p1.X + p2.X,
        Y: p1.Y + p2.Y,
    }
}

type Vector2D struct {
    X float64
    Y float64
}

func NewVector2D(x, y float64) *Vector2D {
    return &Vector2D{ x, y, }
}

func (v *Vector2D) Copy() *Vector2D {
    return &Vector2D {
        X: v.X,
        Y: v.Y,
    }
}

func (v *Vector2D) Apply(other *Vector2D, delta float64) {
    v.X *= other.X * delta
    v.Y *= other.Y * delta
}
