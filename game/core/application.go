package core

import (
	"Go-Rocket/src/game/event"
	"Go-Rocket/src/game/input"
	"Go-Rocket/src/game/layers"
	"runtime"
)

type App struct {
	Wnd    Window
	Layers map[string]layers.Layer
}

// Init - Inits application
func (app *App) Init() {
	runtime.LockOSThread()

	InitGlfw()
	app.Wnd.CreateWindow("Game", 1000, 720)

	InitOpenGL()

	input.InitInput(app.Wnd.GLFWwin)

	app.Layers = make(map[string]layers.Layer)
	app.Wnd.EventHandle = app.OnEvent

}

func (app *App) OnEvent(event *event.Event) {

	for v := range app.Layers {
		app.Layers[v].OnEvent(event)

	}

}

// Run - run application
func (app *App) Run() {

	for !app.Wnd.GLFWwin.ShouldClose() {
		app.Wnd.Clear(0, 0, 0, 1)

		for v := range app.Layers {

			app.Layers[v].OnUpdate()
		}

		app.Wnd.Update()

	}
}

func (app *App) PushLayer(newLayer layers.Layer, name string) {

	newLayer.OnAttach()
	app.Layers[name] = newLayer

}

func (app *App) PopLayer(name string) {
	app.Layers[name].OnDetach()
	delete(app.Layers, name)
}
