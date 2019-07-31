package graphics

import (
	"Go-Rocket/src/game/dependencies/mathgl/mathgl/mgl32"
)

type Transform struct {
	position mgl32.Vec3
	rotation mgl32.Vec3
	scale    mgl32.Vec3

	ModelMatrix mgl32.Mat4
}

func (transform *Transform) CalculateModelMatrix() {
	posMatrix := mgl32.Translate3D(transform.position.X(), transform.position.Y(), transform.position.Z())
	rotMatrix := mgl32.HomogRotate3DZ(transform.rotation.Z())
	scaleMatrix := mgl32.Scale3D(transform.scale.X(), transform.scale.Y(), transform.scale.Z())

	transform.ModelMatrix = posMatrix.Mul4(rotMatrix.Mul4(scaleMatrix))
}

func (transform *Transform) SetPosition(x float32, y float32) {
	transform.position = mgl32.Vec3{x, y, 0}

	transform.CalculateModelMatrix()
}

func (transform *Transform) SetRotation(angle float32) {

	transform.rotation = mgl32.Vec3{0, 0, angle}

	transform.CalculateModelMatrix()
}

func (transform *Transform) SetScale(x float32, y float32) {
	transform.scale = mgl32.Vec3{x, y, 1}

	transform.CalculateModelMatrix()
}
