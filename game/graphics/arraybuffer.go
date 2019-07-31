package graphics

import (
	"Go-Rocket/src/game/dependencies/go-gl/gl/v4.1-core/gl"
)

// VAO - Vertex Array Object
type VAO struct {
	VaoID uint32
	Vbo   []*VBO
	Ibo   *IBO
	index uint32
}

// GenVertexArrayObject -
func (vertexArray *VAO) GenVertexArrayObject() {
	gl.GenVertexArrays(1, &vertexArray.VaoID)

}

// BindVertexArrayObject - Bind vao for use
func (vertexArray *VAO) BindVertexArrayObject() {
	gl.BindVertexArray(vertexArray.VaoID)
}

// UnBind - unbind vao
func (vertexArray *VAO) UnBind() {
	gl.BindVertexArray(0)
}

// AddVertexBuffer - adds new vertex buffer
func (vertexArray *VAO) AddVertexBuffer(vbo VBO) {

	gl.BindVertexArray(vertexArray.VaoID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.VboID)

	for _, v := range vbo.Layout.Elements {

		gl.EnableVertexAttribArray(vertexArray.index)
		gl.VertexAttribPointer(vertexArray.index, v.Count, gl.FLOAT, false, vbo.Layout.Stride, gl.PtrOffset(v.Offset))
		gl.VertexAttribDivisor(uint32(v.Index), uint32(v.Divisor))
		vertexArray.index++

	}
	vertexArray.Vbo = append(vertexArray.Vbo, &vbo)

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// AddIndexBuffer - Add index buffer with info to vao - now only supporst one ibo
func (vertexArray *VAO) AddIndexBuffer(ibo IBO) {
	gl.BindVertexArray(vertexArray.VaoID)
	gl.BindBuffer(gl.ARRAY_BUFFER, ibo.IboID)

	vertexArray.Ibo = &ibo

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (vertexArray *VAO) Clean() {
	for v := range vertexArray.Vbo {
		vertexArray.Vbo[v].CleanVBO()
	}
	vertexArray.Ibo.CleanIBO()
}
