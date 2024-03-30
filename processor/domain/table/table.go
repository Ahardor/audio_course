package table

import (
	"iotvisual/processor/domain/frequency"
	"iotvisual/processor/domain/note"
)

// noteFrequencies - соответствие ноты и списка частот по октавам.
type noteFrequencies struct {
	// Нота.
	Note note.Note
	// Список частот по октавам.
	Frequencies frequency.List
}

// NoteTable таблица нот с дельтами по октавам.
type NoteTable struct {
	// Ноты с частотами по октавам.
	notes []noteFrequencies
	// Минимальные дельты частот по октавам.
	// Используются в качестве погрешности для определения ноты.
	// Берется половина от минимальной разницы по октаве = (freq(C) - freq(C#/Db))/2 - 1 Hz
	deltas [frequency.OctavesCount]int
}

// Deltas возвращает погрешность по заданной октаве.
func (nt NoteTable) Deltas(octave int) (int, bool) {
	if octave >= 9 || octave < 0 {
		return -1, false
	}
	return nt.deltas[octave], true
}

// Table возвращает таблицу нот с частотами по октавам.
func (nt NoteTable) Table() []noteFrequencies {
	return nt.notes
}

// InitTable создает таблицу нот. Порядок нот имеет значение и идет от низких к высоким.
// Ноты в таблице возрастают по частоте слева направо и сверху вниз.
func InitTable() NoteTable {
	nt := NoteTable{
		notes: []noteFrequencies{
			{Note: note.NoteC, Frequencies: frequency.List{1635, 3270, 6541, 13081, 26163, 52325, 104650, 209300, 418601}},
			{Note: note.NoteCd, Frequencies: frequency.List{1732, 3465, 6930, 13859, 27718, 55437, 110873, 221746, 443492}},
			{Note: note.NoteDb, Frequencies: frequency.List{1732, 3465, 6930, 13859, 27718, 55437, 110873, 221746, 443492}},
			{Note: note.NoteD, Frequencies: frequency.List{1835, 3671, 7342, 14683, 29366, 58733, 117466, 234932, 469863}},
			{Note: note.NoteDd, Frequencies: frequency.List{1945, 3889, 7778, 15556, 31113, 62225, 124451, 248902, 497803}},
			{Note: note.NoteEb, Frequencies: frequency.List{1945, 3889, 7778, 15556, 31113, 62225, 124451, 248902, 497803}},
			{Note: note.NoteE, Frequencies: frequency.List{2060, 4120, 8241, 16481, 32963, 65925, 131851, 263702, 527404}},
			{Note: note.NoteF, Frequencies: frequency.List{2183, 4365, 8731, 17461, 34923, 69846, 139691, 279383, 558765}},
			{Note: note.NoteFd, Frequencies: frequency.List{2312, 4625, 9250, 18500, 36999, 73999, 147998, 295996, 591991}},
			{Note: note.NoteGb, Frequencies: frequency.List{2312, 4625, 9250, 18500, 36999, 73999, 147998, 295996, 591991}},
			{Note: note.NoteG, Frequencies: frequency.List{2450, 4900, 9800, 19600, 39200, 78399, 156798, 313596, 627193}},
			{Note: note.NoteGd, Frequencies: frequency.List{2596, 5191, 10383, 20765, 41530, 83061, 166122, 332244, 664488}},
			{Note: note.NoteAb, Frequencies: frequency.List{2596, 5191, 10383, 20765, 41530, 83061, 166122, 332244, 664488}},
			{Note: note.NoteA, Frequencies: frequency.List{2750, 5500, 11000, 22000, 44000, 88000, 176000, 352000, 704000}},
			{Note: note.NoteAd, Frequencies: frequency.List{2914, 5827, 11654, 23308, 46616, 93233, 186466, 372931, 745862}},
			{Note: note.NoteBb, Frequencies: frequency.List{2914, 5827, 11654, 23308, 46616, 93233, 186466, 372931, 745862}},
			{Note: note.NoteB, Frequencies: frequency.List{3087, 6174, 12347, 24694, 49388, 98777, 197553, 395107, 790213}},
		},
	}

	for oct := 0; oct < frequency.OctavesCount; oct++ {
		nt.deltas[oct] = int(nt.notes[0].Frequencies[oct]-nt.notes[1].Frequencies[oct])/2 - 1
	}
	return nt
}

// FindNote находит ноту по заданной частоте в Hz.
func (nt NoteTable) FindNote(freq float64) note.Note {
	f := frequency.FromFloat(freq)
	for column := 0; column < frequency.OctavesCount; column++ {
		for row, noteEntry := range nt.notes {
			if frequency.IsApproximatelyEqual(f, noteEntry.Frequencies[column], nt.deltas[column]) {
				return noteEntry.Note
			}
			if f < (nt.notes[row+1].Frequencies[column] - frequency.Frequency(nt.deltas[column])) {
				return note.NoteUnknown
			}
		}
	}
	return note.NoteUnknown
}
