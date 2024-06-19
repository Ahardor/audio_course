package cacher

// NoteLoadStatus статус получения ноты из кэша.
type NoteLoadStatus int

const (
	// NoteLoadStatusUnknown неизвестный статус.
	NoteLoadStatusUnknown NoteLoadStatus = iota
	// NoteLoadStatusOK удалось получить ноту.
	NoteLoadStatusOK
	// NoteLoadStatusMelodyDoesNotExist не удалось получить мелодию.
	NoteLoadStatusMelodyDoesNotExist
	// NoteLoadStatusNoteIndexOutOfRange ноты по индексу не существует.
	NoteLoadStatusNoteIndexOutOfRange
)
