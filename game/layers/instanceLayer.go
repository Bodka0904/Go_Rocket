package layers

import (
	"Go-Rocket/src/game/event"
	"Go-Rocket/src/game/graphics"
	"math/rand"

	"Go-Rocket/src/game/dependencies/go-gl/gl/v4.1-core/gl"
)

type InstanceLayer struct {
	testVaoPos [2]float32

	vao    graphics.VAO
	shader graphics.Shader
	camera graphics.Camera

	offsets []graphics.VertexInstance
}

func (layer *InstanceLayer) OnAttach() {

	layer.shader.InitShader(graphics.InstanceVertexShaderSource, graphics.InstanceFragmentShaderSource)
	layer.shader.AddAttrib(0, "position")
	layer.shader.AddAttrib(1, "texCoord")
	layer.shader.AddAttrib(2, "offset")

	layer.shader.LinkShader()
	layer.shader.Bind()
	layer.shader.AddUniform("ortho")

	layer.camera.Init()
	layer.camera.SetProjectionMatrix(0, 1000, 720, 0)
	layer.camera.RecalculateViewMatrix()
	layer.camera.SetPosition(0, 0, 0)
	layer.camera.SetRotation(0)

	layer.testVaoPos = [2]float32{500, 360}
	layer.vao.GenVertexArrayObject()

	var vbo graphics.VBO
	var vbo2 graphics.VBO
	var ibo graphics.IBO

	vertices := []graphics.Vertex{
		graphics.Vertex{Vertice: [3]float32{500, 500, 0}, Fragment: [2]float32{0, 1}},
	}
	indices := []uint32{0}
	ibo.GenIndexBuffer(indices, len(indices))

	vbo.SetLayout([]graphics.Element{graphics.Element{Type: "Float3"}, graphics.Element{Type: "Float3"}}, len(vertices))
	vbo.GenVertexBuffer(vertices, gl.STATIC_DRAW)

	for i := 0; i < 50; i++ {
		layer.offsets = append(layer.offsets, graphics.VertexInstance{Offset: [3]float32{0, float32(i * i), 0}})
	}

	vbo2.SetLayout([]graphics.Element{graphics.Element{Type: "Float3", Index: 2, Divisor: 1}}, len(layer.offsets)*50)
	vbo2.GenVertexBuffer(layer.offsets, gl.STATIC_DRAW)

	layer.vao.AddVertexBuffer(vbo)
	layer.vao.AddVertexBuffer(vbo2)
	layer.vao.AddIndexBuffer(ibo)

	layer.vao.UnBind()
}

func (layer *InstanceLayer) OnEvent(e *event.Event) {

}

func (layer *InstanceLayer) OnUpdate() {

	rand.Seed(15)
	for i := 0; i < 50; i++ {
		layer.offsets[i].Offset[0] += rand.Float32() * 20
	}

	layer.shader.Bind()
	graphics.BeginScene(&layer.camera)
	graphics.SubmitInstances(&layer.shader, &layer.vao, 10, gl.POINTS)
}

func (layer *InstanceLayer) OnDetach() {

}
