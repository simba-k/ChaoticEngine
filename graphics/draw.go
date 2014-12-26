package graphics

import (
	"github.com/SimbaOps/chaotic/graphics/lowlevel"
	"github.com/go-gl/gl"
)

const (
	DEFAULT_BUF_SIZE = 1200
)

var buf drawBuffer
var vao gl.VertexArray
var vbo gl.Buffer
var ubo gl.Buffer

type drawBuffer struct {
	vertexBuffer []float32
	topvertex    []float32
	matBuffer    []float32
	topmat       []float32
	pts          int
}

func Init(prog lowlevel.Program) {
	vertexBuffer := make([]float32, DEFAULT_BUF_SIZE)
	matBuffer := make([]float32, DEFAULT_BUF_SIZE * 2)
	buf = drawBuffer{vertexBuffer, vertexBuffer, matBuffer, matBuffer, 0}
	vao = gl.GenVertexArray()
	vao.Bind()
	vbo = lowlevel.CreateVBOxy(vertexBuffer)
	vao.Unbind()

	idx := prog.GetUniformBlockIndex("mats")
	prog.UniformBlockBinding(idx, uint(1))
	ubo = lowlevel.CreateUBO(matBuffer)
	prog.LinkUBO("mats", ubo)
}

func DrawSprite(spr Sprite) {
	copy(buf.topvertex, spr.data)
	buf.topvertex = buf.topvertex[len(spr.data):]
	buf.pts += len(spr.data) / 2
	copy(buf.topmat, spr.matrix)
	buf.topmat = buf.topmat[len(spr.matrix):]
}

func Flush() {
	vbo.Bind(gl.ARRAY_BUFFER)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, buf.pts*2*4, buf.vertexBuffer)
	ubo.Bind(gl.UNIFORM_BUFFER)
	gl.BufferSubData(gl.UNIFORM_BUFFER, 0, buf.pts/6*12*4, buf.matBuffer)
	ubo.Unbind(gl.UNIFORM_BUFFER)

	vao.Bind()
	gl.DrawArrays(gl.TRIANGLES, 0, buf.pts)
	buf.topvertex = buf.vertexBuffer
	buf.topmat = buf.matBuffer
	buf.pts = 0
	vao.Unbind()
}
