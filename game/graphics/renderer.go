package graphics

import (
	"Go-Rocket/src/game/dependencies/mathgl/mathgl/mgl32"

	"Go-Rocket/src/game/dependencies/go-gl/gl/v4.1-core/gl"
)

type SceneData struct {
	ViewProjectionMatrix mgl32.Mat4
}

var sceneData SceneData

func BeginScene(camera *Camera) {

	sceneData.ViewProjectionMatrix = *camera.ProjectionViewMatrix
}

func Submit(shader *Shader, vao *VAO, drawMod uint32) {

	shader.Bind()
	shader.UploadMat4("ortho", sceneData.ViewProjectionMatrix)
	vao.BindVertexArrayObject()
	gl.DrawElements(drawMod, int32(vao.Ibo.Count), gl.UNSIGNED_INT, gl.Ptr(vao.Ibo.Indices))
	vao.UnBind()
}

func SubmitInstances(shader *Shader, vao *VAO, numInstances int32, drawMod uint32) {

	gl.PointSize(50)
	shader.Bind()
	shader.UploadMat4("ortho", sceneData.ViewProjectionMatrix)
	vao.BindVertexArrayObject()

	gl.DrawElementsInstanced(drawMod, int32(vao.Ibo.Count), gl.UNSIGNED_INT, gl.Ptr(vao.Ibo.Indices), numInstances)
	vao.UnBind()
}

func EndScene() {

}
