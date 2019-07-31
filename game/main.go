package main

import (
	"Go-Rocket/src/game/core"
	"Go-Rocket/src/game/layers"
)

func main() {

	testLayer := layers.TestLayer{}
	instanceLayer := layers.InstanceLayer{}

	app := core.App{}

	app.Init()

	app.PushLayer(&testLayer, "test")
	app.PushLayer(&instanceLayer, "instance")
	app.Run()
}
