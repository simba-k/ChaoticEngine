package graphics

import (
	"github.com/go-gl/gl"
	"github.com/SimbaOps/chaotic/graphics/lowlevel"
)

const (
	DEFAULT_BUF_SIZE = 1000
)

var buf drawBuffer
var vao gl.VertexArray
var vbo gl.Buffer

type drawBuffer struct {
	vertexBuffer []float32
	top []float32
	pts int
}

func Init() {
	vertexBuffer := make([]float32, DEFAULT_BUF_SIZE)
	buf  = drawBuffer{vertexBuffer, vertexBuffer, 0}
	vao = gl.GenVertexArray()
	vao.Bind()
	vbo = lowlevel.CreateVBOxy(vertexBuffer)
	vao.Unbind()
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
