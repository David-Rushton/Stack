package main

type KeyEvent struct {
	KeyCode   KeyCode
	KeyValue  rune
	Modifiers Modifiers
}

type KeyEventReceiver interface {
	OnKeyPress(keyEvent KeyEvent)
}
