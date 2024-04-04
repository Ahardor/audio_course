package melody

import (
	"iotvisual/processor/domain/frequency"
	"iotvisual/processor/domain/note"
)

// ID - уникальный идентификатор мелодии.
type ID string

// Melody - мелодия.
type Melody struct {
	ID     ID      `bson:"id"`     // Идентификатор мелодии.
	Sounds []Sound `bson:"sounds"` // Звуки.
}

// Sound - единица разбиения мелодии.
type Sound struct {
	Note       note.Note        `bson:"note"`     // Нота.
	Octave     frequency.Octave `bson:"octave"`   // Октава.
	Serial     int              `bson:"serial"`   // Порядковый номер.
	DurationMS int              `bson:"duration"` // Длительность в мс.
}