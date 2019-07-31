package graphics

import (
	"Go-Rocket/src/game/dependencies/go-gl/gl/v4.1-core/gl"
)

type VBO struct {
	VboID  uint32
	Layout Layout
}

type IBO struct {
	IboID   uint32
	Count   int
	Indices []uint32
}

type Layout struct {
	Elements []Element
	Size     int32
	Stride   int32
}
type Element struct {
	Type    string
	Count   int32
	Size    int32
	Divisor uint32 `default:"0"`
	Index   uint32 `default:"0"`
	Offset  int    `default:"0"`
}

// Vertex - holds data for position and color of vertice
type Vertex struct {
	Vertice  [3]float32
	Fragment [2]float32
}

type VertexInstance struct {
	Offset [3]float32
}

func (vbo *VBO) SetLayout(elements []Element, numVertices int) {

	for v := range elements {
		switch elements[v].Type {
		case "Float":
			elements[v].Size = 4
			elements[v].Count = 1
		case "Float2":
			elements[v].Size = 4 * 2
			elements[v].Count = 2
		case "Float3":
			elements[v].Size = 4 * 3
			elements[v].Count = 3
		case "Float4":
			elements[v].Size = 4 * 4
			elements[v].Count = 4

		case "Integer":
			elements[v].Size = 4
			elements[v].Count = 1
		case "Integer2":
			elements[v].Size = 4 * 2
			elements[v].Count = 2
		case "Integer3":
			elements[v].Size = 4 * 3
			elements[v].Count = 12
		case "Integer4":
			elements[v].Size = 4 * 4
			elements[v].Count = 4
		}
	}

	var offset int

	for v := range elements {
		elements[v].Offset = offset
		offset += int(elements[v].Size)
		vbo.Layout.Stride += elements[v].Size

	}

	vbo.Layout.Size = int32(offset) * int32(numVertices)
	vbo.Layout.Elements = elements

}

func (vbo *VBO) GenVertexBuffer(vertices interface{}, drawMod uint32) {

	gl.GenBuffers(1, &vbo.VboID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.VboID)
	gl.BufferData(gl.ARRAY_BUFFER, int(vbo.Layout.Size), gl.Ptr(vertices), drawMod)

}

func (vbo *VBO) UpdateVertexBuffer(vertices interface{}) {
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.VboID)
	gl.BufferData(gl.ARRAY_BUFFER, int(vbo.Layout.Size), gl.Ptr(vertices), gl.DYNAMIC_DRAW)
}

// GenIndexBuffer - gen index buffer that temporary stores info about order of vertices
func (ibo *IBO) GenIndexBuffer(indices []uint32, count int) {

	gl.GenBuffers(1, &ibo.IboID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo.IboID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, count*4, gl.Ptr(indices), gl.STATIC_DRAW)

	ibo.Count = count
	ibo.Indices = indices
}

func (vbo *VBO) CleanVBO() {
	gl.DeleteBuffers(1, &vbo.VboID)
}
func (ibo *IBO) CleanIBO() {
	gl.DeleteBuffers(1, &ibo.IboID)
}
