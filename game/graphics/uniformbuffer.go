package graphics

import (
	"unsafe"

	"Go-Rocket/src/game/dependencies/go-gl/gl/v4.1-core/gl"
)

type UBO struct {
	UboID uint32
	size  int32
}

func (ubo *UBO) GenUniformBufferObject(data interface{}, size int) {
	gl.GenBuffers(1, &ubo.UboID)
	gl.BindBuffer(gl.UNIFORM_BUFFER, ubo.UboID)
	gl.BufferData(gl.UNIFORM_BUFFER, size, gl.Ptr(data), gl.STREAM_DRAW)

}

func (ubo *UBO) Update(data interface{}) {
	gl.BindBuffer(gl.UNIFORM_BUFFER, ubo.UboID)
	pointer := gl.MapBuffer(gl.UNIFORM_BUFFER, gl.WRITE_ONLY)
	pointer = data.(unsafe.Pointer)
	_ = pointer
	gl.UnmapBuffer(gl.UNIFORM_BUFFER)

}
