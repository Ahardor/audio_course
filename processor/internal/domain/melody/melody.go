package melody

import (
	"fmt"
	"iotvisual/processor/internal/domain/frequency"
	"iotvisual/processor/internal/domain/note"
	"strings"
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
	DurationMS int64            `bson:"duration"` // Длительность в мс.
}

// String реализация интерфейса Stringer.
func (m Melody) String() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("ID: [%s]", m.ID))
	for _, sound := range m.Sounds {
		s := fmt.Sprintf("\tNote: %v\n\tOctave: %v\n\tSerial: %v\n\tDurationMS: %v\n",
			sound.Note, sound.Octave, sound.Serial, sound.DurationMS,
		)
		b.WriteString(s)
	}
	return b.String()
}
