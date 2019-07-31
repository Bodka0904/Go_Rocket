package graphics

import (
	"Go-Rocket/src/game/dependencies/mathgl/mathgl/mgl32"
)

type Camera struct {
	ProjectionMatrix     *mgl32.Mat4
	ViewMatrix           *mgl32.Mat4
	ProjectionViewMatrix *mgl32.Mat4

	position mgl32.Vec3
	rotation float32
}

func (camera *Camera) Init() {
	camera.ProjectionViewMatrix = new(mgl32.Mat4)
	camera.ViewMatrix = new(mgl32.Mat4)
	camera.ProjectionMatrix = new(mgl32.Mat4)

}

func (camera *Camera) SetProjectionMatrix(left float32, right float32, top float32, bottom float32) {
	*camera.ProjectionMatrix = mgl32.Ortho(left, right, bottom, top, 0, 1)
	camera.RecalculateViewMatrix()
}

func (camera *Camera) Move(x float32, y float32) {
	camera.position = mgl32.Vec3{camera.position.X() + x, camera.position.Y() + y, camera.position.Z()}
}

func (camera *Camera) SetPosition(x float32, y float32, z float32) {

	camera.position = mgl32.Vec3{x, y, z}
	camera.RecalculateViewMatrix()
}

func (camera *Camera) SetRotation(rotation float32) {
	camera.rotation = rotation
	camera.RecalculateViewMatrix()
}

func (camera *Camera) RecalculateViewMatrix() {
	translation := mgl32.Translate3D(camera.position.X(), camera.position.Y(), camera.position.Z())
	rotation := mgl32.HomogRotate3DZ(camera.rotation)

	*camera.ViewMatrix = mgl32.Mat4(translation.Mul4(rotation))
	*camera.ProjectionViewMatrix = camera.ProjectionMatrix.Mul4(*camera.ViewMatrix)

}
