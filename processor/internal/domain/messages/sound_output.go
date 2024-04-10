package messages

import (
	"fmt"
	"iotvisual/processor/internal/domain/device"
	"iotvisual/processor/internal/domain/session"
)

// MessageSoundOutput выходное сообщение обработчика.
type MessageSoundOutput struct {
	// ID устройства (датчика).
	Device device.ID `json:"device"`
	// ID сессии.
	SessionUUID session.UUID `json:"session_uuid"`
	// Название мелодии.
	Melody string `json:"melody"`
	// Ожидаемая нота.
	ExpectedNote string `json:"expected_note"`
	// Фактическая нота.
	ActualNote string `json:"actual_note"`
	// Ожидаемая длительность в миллисекундах.
	ExpectedLengthMS int64 `json:"expected_length_ms"`
	// Фактическая длительность в миллисекундах.
	ActualLengthMS int64 `json:"actual_length_ms"`
}

// String реализация интерфейса Stringer.
func (m MessageSoundOutput) String() string {
	s := fmt.Sprintf("\nDevice: %v\nSessionUUID: %v\nMelody: %v\n\tExpectedNote: %v\n\tActualNote: %v\n\tExpectedDuration: %v\n\tActualDuration: %v\n",
		m.Device, m.SessionUUID, m.Melody, m.ExpectedNote, m.ActualNote, m.ExpectedLengthMS, m.ActualLengthMS,
	)
	return s
}
