package models

import (
	"log"
	"time"
)

type Bird struct {
    Pos *Point
    Vel *Vector2D
    Acc *Vector2D
    eventer *GameEventer
}

const BirdGravityY = 69.8;
const BirdJumpVelocity = -40;

func CreateBird(eventer *GameEventer) *Bird {
    return &Bird {
        Pos: NewPoint2D(5, 0),
        Vel: NewVector2D(0, 0),
        Acc: NewVector2D(0, 0),
        eventer: eventer,
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

func (b *Bird) CreateRender(size int) (*Point, [][]byte) {

    bird := make([][]byte, 1 << size)

    switch size {
    case 0:
        bird[0] = []byte{'@'}
    case 1:
        bird[0] = []byte{'/', '\\'}
        bird[1] = []byte{'\\', '/'}
    case 2:
        bird[0] = []byte{' ', ' ', '(', ' ', ' '}
        bird[1] = []byte{' ', '(', '(', '(', ' '}
        bird[2] = []byte{'(', '(', ' ', '(', '('}
    }


    log.Printf("Bird#CreateRender(%v) %+v", size, b.Pos)
    return b.Pos, bird
}


