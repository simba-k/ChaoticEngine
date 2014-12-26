package graphics 
import (
	"math"
)

var deg90 float64 = math.Pi / 4

type Sprite struct {
	data   []float32 //point data
	matrix []float32 //transformation matrix
	xVec   []float32 //X Vector, unit vector pointing in the X direction relative to the sprite
	yVec   []float32 //Y Vector, unit vector pointing in the Y direction relative ot the sprite
	loc    []float32 //x,y location of center of sprite
	angle  *float64  //total angle of rotation of sprite
	relveloc *float32 //linear velocity in direction of xVec
	rotveloc *float64 //rotational velocity
}

func NewSprite(width, height, centerX, centerY float32) Sprite {
	//Set top left point as (x1, y1)
	x1 := -centerX
	y1 := -centerY

	//Set bottom right point as (x1, y1)
	x2 := width - centerX
	y2 := height - centerY

	//set Sprite point data
	dat := []float32{x1, y1,
		x2, y1,
		x2, y2,
		x2, y2,
		x1, y2,
		x1, y1}

	//Initialize mat3 with padding for std140 layout
	mat := []float32{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0}

	//initialize angle as 0
	ang := new(float64)
	*ang = 0.

	relveloc := new(float32)
	*relveloc = 0

	rotveloc := new(float64)
	*rotveloc = 0

	//Create sprite, map xvec, yvec, and loc to their mathmatical correspondants in the matrix
	sprite := Sprite{dat,
		mat,
		mat[0:2],
		mat[4:6],
		mat[8:10],
		ang,
		relveloc,
		[]float32{0,0},
		rotveloc,
		[]float32{0,0},
	}
	return sprite
}

//TODO relveloc should just be one float32
func (spr Sprite) Update() {
	spr.loc[0] += spr.xVec[0] * *spr.relveloc
	spr.loc[1] += spr.xVec[1] * *spr.relveloc
	spr.rotate(*spr.rotveloc)
}

func (spr Sprite) SetRelVeloc(x float32) {
	*spr.relveloc = x
}

func (spr Sprite) SetRot(x float64) {
	*spr.rotveloc = x
}

//testing
func (spr Sprite) Shoot() (rt Sprite){
	rt = NewSprite(.1, .1, .05, .05)
	rt.rotate(*spr.angle)
	rt.SetRelVeloc(.01)
	return
}


//Moves the sprite forward relative to itself
func (spr Sprite) MoveForward(units float32) {
	//xVec is a unit vector in the forward direction
	//moves that many units forward and applies it to the matrix
	spr.loc[0] += spr.xVec[0] * units
	spr.loc[1] += spr.xVec[1] * units
}

//Moves the sprite backwards relative to itself
func (spr Sprite) MoveBackwards(units float32) {
	spr.MoveForward(-units)
}

//Rotates Sprite CounterClockwise
func (spr Sprite) RotateRight(rads float64) {
	spr.rotate(rads)
}

//Rotates Sprite Clockwise
func (spr Sprite) RotateLeft(rads float64) {
	spr.rotate(-rads)
}

func (spr Sprite) rotate(rads float64) {
	//calculate new angle, must be between (-2pi, 2pi)
	*spr.angle = math.Mod(*spr.angle+rads, (math.Pi * 2))
	sin64, cos64 := math.Sincos(*spr.angle)

	//convert to float32
	sin32 := float32(sin64)
	cos32 := float32(cos64)

	//set new xVector and yVector values
	spr.xVec[0] = cos32
	spr.xVec[1] = sin32
	spr.yVec[0] = -sin32
	spr.yVec[1] = cos32
}
