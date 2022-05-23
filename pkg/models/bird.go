package models

import "time"

type Bird struct {
    Pos *Vector2D
    Vel *Vector2D
    Acc *Vector2D
}

const BIRD_GRAVITY_Y = 9.8;

func CreateBird() *Bird {
    return &Bird {
        Pos: NewVector2D(0, 0),
        Vel: NewVector2D(0, 0),
        Acc: NewVector2D(0, 0),
    }
}

func (b *Bird) Update(t time.Duration) {
    // millis := t.Milliseconds()
    // accel_y := BIRD_GRAVITY_Y * float64(millis) / 1000.0
}

func (b *Bird) Render(renderStrings [][]byte) {
    if len(renderStrings) < int(b.Pos.Y) {
        return
    }

    renderStrings[int(b.Pos.Y)][int(b.Pos.X)] = '@'
}

