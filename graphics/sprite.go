package graphics

type Sprite struct {
	data []float32
	loc []float32
	xVec []float32
	yVec []float32
}

func NewSprite(width, height, centerX, centerY float32) Sprite{
	x1 := -centerX;
	y1 := -centerY;
	x2 := width - centerX;
	y2 := height - centerY;
	dat := []float32{x1, y1,
					x2, y1,
					x2, y2,
					x2, y2,
					x1, y2,
					x1, y1}
	sprite := Sprite{dat, []float32{0, 0}, []float32{1, 0}, []float32{0, 1}}
	return sprite;
}

func (spr Sprite) MoveForward(units float32) {
	dx := []float32{spr.xVec[0] * units, spr.xVec[1] * units}
	for i := 0; i < len(spr.data); i+=2 {
		spr.data[i] += dx[0]
		spr.data[i+1] += dx[1]
	}
}
