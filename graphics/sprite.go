package graphics

import (
	"fmt"
	"math"
)

var deg90 float64 = math.Pi / 4

type Sprite struct {
	data []float32
	matrix []float32
	xVec []float32
	yVec []float32
	loc  []float32
	angle *float64
}

func NewSprite(width, height, centerX, centerY float32) Sprite {
	x1 := -centerX
	y1 := -centerY
	x2 := width - centerX
	y2 := height - centerY
	dat := []float32{x1, y1,
		x2, y1,
		x2, y2,
		x2, y2,
		x1, y2,
		x1, y1}
	mat := []float32{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0}
	ang := new(float64)
	*ang = 0.
	sprite := Sprite{dat,
		mat,
		mat[0:2],
		mat[4:6],
		mat[8:10],
		ang}
	return sprite
}

func (spr *Sprite) MoveForward(units float32) {
	dx := []float32{spr.xVec[0] * units, spr.xVec[1] * units}
	for i := 0; i < len(spr.data); i += 2 {
		spr.data[i] += dx[0]
		spr.data[i+1] += dx[1]
	}
}

func (spr *Sprite) MoveBackwards(units float32) {
	dx := []float32{spr.xVec[0] * -units, spr.xVec[1] * -units}
	for i := 0; i < len(spr.data); i += 2 {
		spr.data[i] += dx[0]
		spr.data[i+1] += dx[1]
	}
}

func (spr *Sprite) RotateRight(rads float64) {
	*spr.angle = math.Mod( *spr.angle + rads, (math.Pi * 2) )
	sin64, cos64 := math.Sincos(*spr.angle)
	sin32 := float32(sin64)
	cos32 := float32(cos64)
	spr.xVec[0] = cos32
	spr.xVec[1] = sin32
	spr.yVec[0] = -sin32
	spr.yVec[1] = cos32
	fmt.Println(spr.xVec, spr.yVec, spr.matrix)
}

func (spr Sprite) RotateLeft(rads float64) {
	*spr.angle = math.Mod( *spr.angle - rads, (math.Pi * 2) )
	sin64, cos64 := math.Sincos(*spr.angle)
	sin32 := float32(sin64)
	cos32 := float32(cos64)
	spr.xVec[0] = cos32
	spr.xVec[1] = sin32
	spr.yVec[0] = -sin32
	spr.yVec[1] = cos32
	fmt.Println(spr.xVec, spr.yVec, spr.matrix)
}
