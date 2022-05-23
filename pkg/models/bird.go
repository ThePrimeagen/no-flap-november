package models

import (
	"time"
)

type Bird struct {
    Pos *Point
    Vel *Vector2D
    Acc *Vector2D
}

const BIRD_GRAVITY_Y = 20.8;

func CreateBird() *Bird {
    return &Bird {
        Pos: NewPoint2D(0, 0),
        Vel: NewVector2D(0, 0),
        Acc: NewVector2D(0, 0),
    }
}

// Note: I am no physicist dammit, I am a scientist
func (b *Bird) Update(t time.Duration) {
    delta := float64(t.Milliseconds()) / 1000.0;
    b.Vel.Y += BIRD_GRAVITY_Y * delta
    b.Pos.Y += b.Vel.Y * delta
}

func min(one float64, two int) float64 {
    two_f := float64(two)
    if two_f > one {
        return one
    }
    return two_f
}

func (b *Bird) Render(renderer Renderer) {
    bird := make([][]byte, 1)
    bird[0] = []byte{'@'}

    _, h := renderer.GetBounds()
    b.Pos.Y = min(b.Pos.Y, h)

    renderer.Render(b.Pos, bird)
}

