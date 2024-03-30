package note

import "fmt"

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

// GetNote возвращает ноту с учетом нот с одинаковым набором частот.
func GetNote(n Note) string {
	switch n {
	case NoteCd, NoteDb:
		return fmt.Sprintf("%s/%s", NoteCd.String(), NoteDb.String())
	case NoteDd, NoteEb:
		return fmt.Sprintf("%s/%s", NoteDd.String(), NoteEb.String())
	case NoteFd, NoteGb:
		return fmt.Sprintf("%s/%s", NoteFd.String(), NoteGb.String())
	case NoteGd, NoteAb:
		return fmt.Sprintf("%s/%s", NoteGd.String(), NoteAb.String())
	case NoteAd, NoteBb:
		return fmt.Sprintf("%s/%s", NoteAd.String(), NoteBb.String())
	default:
		return n.String()
	}
}
