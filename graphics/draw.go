package graphics

import (
	"github.com/go-gl/gl"
	"github.com/SimbaOps/chaotic/graphics/lowlevel"
)

func init() {

}

const (
	DEFAULT_BUF_SIZE = 1000
)

type drawBuffer struct {
	vertexBuffer []float32
	top []float32
	pts int
}

type Sprite struct {
	data []float32
}

type TexCoord struct {
	S float32
	T float32
}

var buf drawBuffer
var vao gl.VertexArray
var vbo gl.Buffer

func Init() {
	vertexBuffer := make([]float32, DEFAULT_BUF_SIZE)
	buf  = drawBuffer{vertexBuffer, vertexBuffer, 0}
	vao = gl.GenVertexArray()
	vao.Bind()
	vbo = lowlevel.CreateVBOxy(vertexBuffer)
	vao.Unbind()
}

func Run() {

}

func DrawSprite(spr Sprite) {
	copy(buf.top, spr.data)
	buf.top = buf.top[len(spr.data):]
	buf.pts += len(spr.data)/2
}

func Flush() {
	vbo.Bind(gl.ARRAY_BUFFER)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, buf.pts * 2 * 4, buf.vertexBuffer)

	vao.Bind()
	gl.DrawArrays(gl.TRIANGLES, 0, buf.pts)
	buf.top = buf.vertexBuffer
	buf.pts = 0
	vao.Unbind()
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
	sprite := Sprite{dat}
	return sprite;
}
