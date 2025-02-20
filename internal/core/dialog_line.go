package core

type DialogueLine struct {
	Type      DialogueType
	Character *Character // nil for narrator or center text
	Text      string
	Position  Position
	Speed     float64 // symbols per second, 0 for instant
}
