package event

const (
	None               = 0
	WindowResized      = 1
	MouseButtonClick   = 2
	MouseButtonRelease = 3
	MousePosition      = 4
	KeyPressed         = 5
	KeyReleased        = 6
	CharPressed        = 7
)

type Event struct {
	Event     interface{}
	EventType int
}

type WindowResizedEvent struct {
	Xsize int
	Ysize int
}

type MouseButtonClickEvent struct {
	Button int
}
type MouseButtonReleaseEvent struct {
	Button int
}
type MousePositionEvent struct {
	Xpos float64
	Ypos float64
}
type KeyPressedEvent struct {
	Key  int
	Mods int
}
type KeyReleasedEvent struct {
	Key  int
	Mods int
}
type CharPressedEvent struct {
	Char rune
	Mods int
}
