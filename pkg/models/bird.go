package models

import (
	"time"
)

type Bird struct {
    Pos *Point
    Vel *Vector2D
    Acc *Vector2D
}

const BirdGravityY = 69.8;
const BirdJumpVelocity = -40;

func CreateBird() *Bird {
    return &Bird {
        Pos: NewPoint2D(0, 0),
        Vel: NewVector2D(0, 0),
        Acc: NewVector2D(0, 0),
    }
}

// The Update Note: I am no physicist dammit, I am a scientist
func (b *Bird) Update(t time.Duration) {
    delta := float64(t.Milliseconds()) / 1000.0;
    b.Vel.Y += BirdGravityY * delta + b.Acc.Y
    b.Pos.Y += b.Vel.Y * delta
    b.Acc.Y *= 0.25
}

func (b *Bird) Jump() {
    b.Acc.Y = 0
    b.Vel.Y = BirdJumpVelocity
}

func (b *Bird) Render(renderer Renderer) {
    bird := make([][]byte, 1)
    bird[0] = []byte{'@'}

    renderer.Render(b.Pos, bird)
}


