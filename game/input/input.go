package input

import (
	"Go-Rocket/src/game/dependencies/glfw/v3.2/glfw"
)

type InputInfo struct {
	win *glfw.Window
}

var Input InputInfo

func InitInput(win *glfw.Window) {
	Input.win = win
}

func IsMouseButtonPressed(button glfw.MouseButton) bool {
	action := Input.win.GetMouseButton(button)
	if action == glfw.Press {
		return true
	}

	return false
}

func IsMouseButtonReleased(button glfw.MouseButton) bool {
	action := Input.win.GetMouseButton(button)
	if action == glfw.Release {
		return true
	}

	return false
}

func IsKeyPressed(keycode glfw.Key) bool {
	action := Input.win.GetKey(keycode)
	if action == glfw.Press {
		return true
	}

	return false
}

func IsKeyReleased(keycode int) bool {
	action := Input.win.GetKey(glfw.Key(keycode))
	if action == glfw.Release {
		return true
	}
	return false
}
