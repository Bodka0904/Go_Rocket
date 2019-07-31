package core

import (
	"Go-Rocket/src/game/event"
	"log"

	"Go-Rocket/src/game/dependencies/glfw/v3.2/glfw"
	"Go-Rocket/src/game/dependencies/go-gl/gl/v4.1-core/gl"
)

type Window struct {
	Width  int
	Height int
	Title  string

	GLFWwin     *glfw.Window
	EventHandle func(event *event.Event)
}

// CreateWindow - create window and initialize glfw
func (win *Window) CreateWindow(title string, width int, height int) {

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	window.MakeContextCurrent()

	*win = Window{Width: width, Height: height, Title: title, GLFWwin: window}

	win.GLFWwin.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		win.EventHandle(&event.Event{Event: event.WindowResizedEvent{Xsize: width, Ysize: height}, EventType: event.WindowResized})
		win.Width = width
		win.Height = height
		gl.Viewport(0, 0, int32(width), int32(height))
	})

	win.GLFWwin.SetCloseCallback(func(w *glfw.Window) {
		log.Println("Window Closed")
	})

	win.GLFWwin.SetCharModsCallback(func(w *glfw.Window, char rune, mods glfw.ModifierKey) {
		win.EventHandle(&event.Event{Event: event.CharPressedEvent{Char: char, Mods: int(mods)}, EventType: event.CharPressed})
	})

	win.GLFWwin.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			win.EventHandle(&event.Event{Event: event.KeyPressedEvent{Key: int(key), Mods: int(mods)}, EventType: event.KeyPressed})
		case glfw.Release:
			win.EventHandle(&event.Event{Event: event.KeyReleasedEvent{Key: int(key), Mods: int(mods)}, EventType: event.KeyReleased})
		}
	})

	win.GLFWwin.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		win.EventHandle(&event.Event{Event: event.MousePositionEvent{Xpos: xpos, Ypos: ypos}, EventType: event.MousePosition})
	})

	win.GLFWwin.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			win.EventHandle(&event.Event{Event: event.MouseButtonClickEvent{Button: int(button)}, EventType: event.MouseButtonClick})

		case glfw.Release:
			win.EventHandle(&event.Event{Event: event.MouseButtonReleaseEvent{Button: int(button)}, EventType: event.MouseButtonRelease})
		}

	})
}

// Update - updates window content, and swap buffers , one is current that we see, and one is in background being prepared, should be called at the end of frame
func (win *Window) Update() {
	glfw.PollEvents()
	win.GLFWwin.SwapBuffers()
}

// Clear - clear window , should be called at the start of frame
func (win *Window) Clear(R float32, G float32, B float32, A float32) {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(R, G, B, A)
}
