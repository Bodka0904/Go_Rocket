package layers

import (
	"Go-Rocket/src/game/input"
	"log"
	"math"

	"Go-Rocket/src/game/dependencies/glfw/v3.2/glfw"
	"Go-Rocket/src/game/dependencies/go-gl/gl/v4.1-core/gl"

	"Go-Rocket/src/game/event"
	"Go-Rocket/src/game/graphics"
)

type Layer interface {
	OnAttach()
	OnEvent(e *event.Event)
	OnUpdate()
	OnDetach()
}

type TestLayer struct {
	//TEST
	testVaoPos [2]float32
	testVaoRot float32
	vao        graphics.VAO
	shader     graphics.Shader
	camera     graphics.Camera
	transform  graphics.Transform
	texture    *graphics.Texture
}

func (layer *TestLayer) OnAttach() {

	layer.shader.InitShader(graphics.BasicVertexShaderSource, graphics.BasicFragmentShaderSource)
	layer.shader.AddAttrib(0, "position")
	layer.shader.AddAttrib(1, "texCoord")

	layer.shader.LinkShader()
	layer.shader.Bind()
	layer.shader.AddUniform("ortho")
	layer.shader.AddUniform("model")

	layer.camera.Init()
	layer.camera.SetProjectionMatrix(0, 1000, 720, 0)
	layer.camera.RecalculateViewMatrix()
	layer.camera.SetPosition(0, 0, 0)
	layer.camera.SetRotation(0)

	layer.testVaoRot = 0
	layer.testVaoPos = [2]float32{500, 360}
	layer.transform.SetScale(2, 2)

	var err error
	layer.texture, err = graphics.NewTextureFromFile("res/textures/playerShip1_red.png", gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)
	if err != nil {
		log.Fatal(err)
	}
	layer.vao.GenVertexArrayObject()

	var vbo graphics.VBO
	var ibo graphics.IBO

	vertices := []graphics.Vertex{
		graphics.Vertex{Vertice: [3]float32{-20, -20, 0}, Fragment: [2]float32{0, 1}},
		graphics.Vertex{Vertice: [3]float32{20, -20, 0}, Fragment: [2]float32{1, 1}},
		graphics.Vertex{Vertice: [3]float32{20, 20, 0}, Fragment: [2]float32{1, 0}},
		graphics.Vertex{Vertice: [3]float32{-20, 20, 0}, Fragment: [2]float32{0, 0}},
	}

	indices := []uint32{0, 1, 2, 0, 3, 2}
	ibo.GenIndexBuffer(indices, len(indices))

	vbo.SetLayout([]graphics.Element{graphics.Element{Type: "Float3"}, graphics.Element{Type: "Float2"}}, len(vertices))

	vbo.GenVertexBuffer(vertices, gl.STATIC_DRAW)

	layer.vao.AddVertexBuffer(vbo)
	layer.vao.AddIndexBuffer(ibo)

	layer.vao.UnBind()
}

func (layer *TestLayer) OnEvent(e *event.Event) {

	if e.EventType == event.WindowResized {
		eve := e.Event.(event.WindowResizedEvent)
		layer.camera.SetProjectionMatrix(0, float32(eve.Xsize), float32(eve.Ysize), 0)

	}
	//if e.EventType == event.KeyPressed {
	//	eve := e.Event.(event.KeyPressedEvent)
	//	if eve.Key == int(glfw.KeyUp) {
	//
	//	}
	//}
	//
	//if e.EventType == event.KeyReleased {
	//	eve := e.Event.(event.KeyReleasedEvent)
	//	if eve.Key == int(glfw.KeyUp) {
	//
	//	}
	//}
}

func (layer *TestLayer) OnUpdate() {

	if input.IsKeyPressed(glfw.KeyUp) {
		layer.testVaoPos[1] += float32(math.Cos(float64(layer.testVaoRot)))
		layer.testVaoPos[0] -= float32(math.Sin(float64(layer.testVaoRot)))

	} else if input.IsKeyPressed(glfw.KeyDown) {
		layer.testVaoPos[1] += -float32(math.Cos(float64(layer.testVaoRot)))
		layer.testVaoPos[0] -= -float32(math.Sin(float64(layer.testVaoRot)))

	}
	if input.IsKeyPressed(glfw.KeyRight) {
		layer.testVaoRot -= 0.005
	} else if input.IsKeyPressed(glfw.KeyLeft) {
		layer.testVaoRot += 0.005
	}

	layer.shader.Bind()
	layer.transform.SetPosition(layer.testVaoPos[0], layer.testVaoPos[1])
	layer.transform.SetRotation(layer.testVaoRot)
	layer.shader.UploadMat4("model", layer.transform.ModelMatrix)
	layer.texture.Bind(gl.TEXTURE0)

	graphics.BeginScene(&layer.camera)
	graphics.Submit(&layer.shader, &layer.vao, gl.TRIANGLES)

}

func (layer *TestLayer) OnDetach() {
	layer.vao.Clean()
}
