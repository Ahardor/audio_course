package note

type Note int

const (
	NoteUnknown Note = iota
	NoteC
	NoteCd
	NoteDb
	NoteD
	NoteDd
	NoteEb
	NoteE
	NoteF
	NoteFd
	NoteGb
	NoteG
	NoteGd
	NoteAb
	NoteA
	NoteAd
	NoteBb
	NoteB
)

// String возвращает строку с нотационной записью ноты.
func (n Note) String() string {
	return [...]string{
		"Unknown",
		"C", "C#",
		"Db", "D", "D#",
		"Eb", "E",
		"F", "F#",
		"Gb", "G", "G#",
		"Ab", "A", "A#",
		"Bb", "B",
	}[n]
}
