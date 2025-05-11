package tui

type Modifiers int

const (
	ModifierShift Modifiers = 1 << iota
	ModifierAlt
	ModifierControl
	ModifierMeta

	ModifierNone Modifiers = 0
)

func (mk Modifiers) IsShiftPressed() bool {
	return mk&ModifierShift == ModifierShift
}

func (mk Modifiers) IsAltPressed() bool {
	return mk&ModifierAlt == ModifierAlt
}

func (mk Modifiers) IsControlPressed() bool {
	return mk&ModifierControl == ModifierControl
}

func (mk Modifiers) IsMetaPressed() bool {
	return mk&ModifierMeta == ModifierMeta
}
