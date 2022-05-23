package models

import "time"

type Bird struct {
    Pos *Point
    Vel *Vector2D
    Acc *Vector2D
}

const BIRD_GRAVITY_Y = 9.8;

func CreateBird() *Bird {
    return &Bird {
        Pos: NewPoint2D(0, 0),
        Vel: NewVector2D(0, 0),
        Acc: NewVector2D(0, 0),
    }
}

func (b *Bird) Update(t time.Duration) {
    // millis := t.Milliseconds()
    // accel_y := BIRD_GRAVITY_Y * float64(millis) / 1000.0
}

func (b *Bird) Render(renderer Renderer) {
    bird := make([][]byte, 1)
    bird[0] = []byte{'@'}
    renderer.Render(b.Pos, bird)
}

