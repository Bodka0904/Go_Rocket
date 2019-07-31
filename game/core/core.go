package core

import (
	"log"

	"Go-Rocket/src/game/dependencies/glfw/v3.2/glfw"
	"Go-Rocket/src/game/dependencies/go-gl/gl/v4.1-core/gl"
)

// InitGlfw - init glfw lib
func InitGlfw() {
	if err := glfw.Init(); err != nil {
		log.Fatal(err)
	}
	log.Println("GLFW inited")
}

// InitOpenGL - init opengl lib
func InitOpenGL() {
	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	log.Println("OpenGL inited")
}
