package graphics

import (
)

const (
	DEFAULT_BUF_SIZE = 1000
)

type drawBuffer struct {
	vertexBuffer []float32
}

type Sprite struct {
	data []float32
}

type SpriteDrawer interface {
	DrawSprite(Sprite)
	Flush()
}

type TexCoord struct {
	S float32
	T float32
}

func NewSprite(widhth, height float32, t1, t2 TexCoord) {

}
